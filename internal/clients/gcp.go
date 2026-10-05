// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/crossplane/upjet/v2/apis/configuration/v1alpha1"
	upjetmetrics "github.com/crossplane/upjet/v2/pkg/metrics"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google/google/fwprovider"
	transporttpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/upbound/provider-gcp/v3/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-gcp/v3/apis/namespaced/v1beta1"
)

const (
	keyProject             = "project"
	keyUserProjectOverride = "user_project_override"
	keyBillingProject      = "billing_project"
	keyUniverseDomain      = "universe_domain"

	credentialsSourceUpbound     = "Upbound"
	keyCredentials               = "credentials"
	credentialsSourceAccessToken = "AccessToken"
	keyAccessToken               = "access_token"

	upboundProviderIdentityTokenFile = "/var/run/secrets/upbound.io/provider/token"
	impersonateServiceAccount        = "ImpersonateServiceAccount"
	keyImpersonateServiceAccount     = "impersonate_service_account"
)

const (
	// error messages
	errNoProviderConfig              = "no providerConfigRef provided"
	errGetProviderConfig             = "cannot get referenced ProviderConfig"
	errTrackUsage                    = "cannot track ProviderConfig usage"
	errExtractKeyCredentials         = "cannot extract JSON key credentials"
	errExtractTokenCredentials       = "cannot extract Access Token credentials"
	errConstructFederatedCredentials = "cannot construct federated identity credentials"
	errMissingFederatedConfiguration = "missing identity federation configuration"
	errPaveFmt                       = "cannot pave the managed resource %s/%s"
	errPavedGetValueFmt              = "cannot get 'spec.forProvider.project' from the managed resource %s/%s"
)

// federatedCredentials is the expected client credential configuration
// structure for federated identity.
type federatedCredentials struct {
	Type                           string               `json:"type"`
	Audience                       string               `json:"audience"`
	SubjectTokenType               string               `json:"subject_token_type"`
	TokenURL                       string               `json:"token_url"`
	CredentialSource               credentialFileSource `json:"credential_source"`
	ServiceAccountImpersonationURL string               `json:"service_account_impersonation_url"`
}

// credentialFileSource is the source of the credential data to be used with
// federated identity.
type credentialFileSource struct {
	File string `json:"file"`
}

// constructFederatedCredentials constructs federated identity credentials with
// the provided identity provider and service account.
func constructFederatedCredentials(providerID, serviceAccount string) ([]byte, error) {
	return json.Marshal(&federatedCredentials{
		Type:                           "external_account",
		Audience:                       fmt.Sprintf("//iam.googleapis.com/%s", providerID),
		SubjectTokenType:               "urn:ietf:params:oauth:token-type:jwt",
		TokenURL:                       "https://sts.googleapis.com/v1/token",
		ServiceAccountImpersonationURL: fmt.Sprintf("https://iamcredentials.googleapis.com/v1/projects/-/serviceAccounts/%s:generateAccessToken", serviceAccount),
		CredentialSource: credentialFileSource{
			File: upboundProviderIdentityTokenFile,
		},
	})
}

// metricsRoundTripper increments the upjet ExternalAPICalls metric for every
// Google API response it observes. It sits at the outermost layer of the GCP
// provider's transport chain, so it counts logical API calls rather than
// individual retry attempts.
type metricsRoundTripper struct {
	base http.RoundTripper
}

func (m *metricsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := m.base.RoundTrip(req)
	if err == nil && resp != nil {
		upjetmetrics.ExternalAPICalls.WithLabelValues(serviceFromURL(req.URL), req.Method).Inc()
	}
	return resp, err
}

// serviceFromURL extracts a short service label from a Google API request URL.
//
// GCP URL patterns:
//   - compute.googleapis.com/compute/v1/...  → "compute"
//   - storage.googleapis.com/...             → "storage"
//   - www.googleapis.com/storage/v1/...      → "storage" (generic host → path fallback)
func serviceFromURL(u *url.URL) string {
	host := u.Hostname()
	if label, _, _ := strings.Cut(host, "."); label != "" && label != "www" {
		return label
	}
	if seg, _, _ := strings.Cut(strings.TrimPrefix(u.Path, "/"), "/"); seg != "" {
		return seg
	}
	return "unknown"
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
// NOTE(hasheddan): this function is slightly over our cyclomatic complexity
// goal. Consider refactoring before adding new branches.
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, crClient client.Client, mg resource.Managed) (terraform.Setup, error) {
		pcSpec, ps, err := baseConfiguration(ctx, crClient, mg)
		if err != nil {
			return ps, err
		}

		switch pcSpec.Credentials.Source { //nolint:exhaustive
		case xpv2.CredentialsSourceInjectedIdentity:
			// We don't need to do anything here, as the TF Provider will take care of workloadIdentity etc.
		case impersonateServiceAccount:
			if pcSpec.Credentials.Name != "" {
				ps.Configuration[keyImpersonateServiceAccount] = pcSpec.Credentials.Name
			}
		case credentialsSourceAccessToken:
			data, err := resource.CommonCredentialExtractor(ctx, xpv2.CredentialsSourceSecret, crClient, pcSpec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, errExtractTokenCredentials)
			}
			ps.Configuration[keyAccessToken] = string(data)
		case credentialsSourceUpbound:
			if pcSpec.Credentials.Upbound == nil || pcSpec.Credentials.Upbound.Federation == nil {
				return ps, errors.Wrap(errors.New(errMissingFederatedConfiguration), errConstructFederatedCredentials)
			}
			data, err := constructFederatedCredentials(pcSpec.Credentials.Upbound.Federation.ProviderID, pcSpec.Credentials.Upbound.Federation.ServiceAccount)
			if err != nil {
				return ps, errors.Wrap(err, errConstructFederatedCredentials)
			}
			ps.Configuration[keyCredentials] = string(data)
		default:
			data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, crClient, pcSpec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, errExtractKeyCredentials)
			}
			ps.Configuration[keyCredentials] = string(data)
		}

		// deliberately not using the caller context as context used to configure terraform is stored
		// nolint:contextcheck
		return ps, errors.Wrap(configureNoForkGCPClient(context.Background(), &ps, *tfProvider), "failed to configure the no-fork GCP client")
	}
}

// baseConfiguration resolves the ProviderConfig referenced by mg and builds
// the Terraform provider configuration shared by every setup function:
// project (with its per-resource override), the ProviderConfig overrides and
// the universe domain. Credentials are deliberately left unset here; each
// caller sets whichever credential source applies to it.
func baseConfiguration(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, terraform.Setup, error) {
	ps := terraform.Setup{}
	pcSpec, err := resolveProviderConfig(ctx, crClient, mg)
	if err != nil {
		return nil, ps, errors.Wrap(err, "cannot resolve provider config")
	}
	// set provider configuration
	ps.Configuration = map[string]interface{}{
		keyProject: pcSpec.ProjectID,
	}
	setProjectOverrides(ps.Configuration, pcSpec)
	setUniverseDomain(ps.Configuration, pcSpec)
	// TODO: this will have a performance impact. We need to quantify this.
	p, err := fieldpath.PaveObject(mg, fieldpath.WithMaxFieldPathIndex(1))
	if err != nil {
		return pcSpec, ps, errors.Wrapf(err, errPaveFmt, mg.GetObjectKind().GroupVersionKind().Kind, mg.GetName())
	}
	// TODO: if the managed resource declares its project
	//  in a different parameter, the following will not work.
	resourceProject, err := p.GetString("spec.forProvider.project")
	if err != nil && !fieldpath.IsNotFound(err) {
		return pcSpec, ps, errors.Wrapf(err, errPavedGetValueFmt, mg.GetObjectKind().GroupVersionKind().Kind, mg.GetName())
	}
	if err == nil && resourceProject != "" {
		ps.Configuration[keyProject] = resourceProject
	}
	return pcSpec, ps, nil
}

// setProjectOverrides populates the user_project_override and billing_project
// provider configuration keys from the resolved ProviderConfig spec. Both keys
// are left unset when the corresponding spec fields are empty so that the
// Terraform provider defaults stay in effect.
func setProjectOverrides(cfg map[string]interface{}, pcSpec *namespacedv1beta1.ProviderConfigSpec) {
	if pcSpec.UserProjectOverride != nil {
		cfg[keyUserProjectOverride] = *pcSpec.UserProjectOverride
	}
	if pcSpec.BillingProject != nil && *pcSpec.BillingProject != "" {
		cfg[keyBillingProject] = *pcSpec.BillingProject
	}
}

// setUniverseDomain populates the universe_domain provider configuration key
// from the resolved ProviderConfig spec. The key is left unset when the spec
// field is empty so that the Terraform provider defaults to googleapis.com.
// The Terraform provider rejects a universe mismatch between its configuration and
// the credentials. So when credentials carry a non-default universe, UniverseDomain
// is required in the ProviderConfigSpec and must match the credentials value.
func setUniverseDomain(cfg map[string]interface{}, pcSpec *namespacedv1beta1.ProviderConfigSpec) {
	if pcSpec.UniverseDomain != nil && *pcSpec.UniverseDomain != "" {
		cfg[keyUniverseDomain] = *pcSpec.UniverseDomain
	}
}

// configureNoForkGCPClient configures the given Terraform provider. ctx roots
// the context it is configured on - deliberately not necessarily the
// caller's live request context, since the provider's own context is stored
// for reuse well past the lifetime of any one call, but a caller can still
// carry a value on ctx, such as the offline egress guard's HTTP client,
// through to Configure.
func configureNoForkGCPClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.

	// Terraform provider stores the context used for its configuration
	// - the context needs to stay active for a longer period of time than terraform plugin sdk operations
	// - the context needs to be eventually cancelled otherwise google provider resources are leaked
	const (
		terraformPluginSDKAsyncTimeout = time.Hour
		gracePeriod                    = 10 * time.Minute
		providerTimeout                = terraformPluginSDKAsyncTimeout + gracePeriod
	)
	ctx, cancel := context.WithCancel(ctx)
	time.AfterFunc(providerTimeout, cancel)

	diag := p.Configure(ctx, &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	// The framework provider must wrap this just-configured provider copy:
	// upstream's FrameworkProvider.Configure ignores the configuration it
	// receives and reads Primary.Meta() instead, so wrapping the shared
	// module-level instance would yield a nil meta and panic, while wrapping
	// the per-reconcile copy also preserves the ProviderConfig isolation
	// described above.
	ps.FrameworkProvider = fwprovider.New(&p)
	if config, ok := ps.Meta.(*transporttpg.Config); ok && config.Client != nil {
		base := config.Client.Transport
		if base == nil {
			base = http.DefaultTransport
		}
		config.Client.Transport = &metricsRoundTripper{base: base}
	}
	return nil
}

func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveV2(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed")
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveV2(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.New("pc is not an Object")
	}

	// Namespace is ignored for a cluster-scoped PC by a real API server's
	// RESTMapper-aware client, but not by the diff server's in-memory one,
	// which keys its store by an exact namespace match and holds a
	// cluster-scoped object under an empty one - so it has to be cleared
	// here instead of relying on the client to do it.
	ns := mg.GetNamespace()
	if configRef.Kind == namespacedv1beta1.ClusterProviderConfigKind {
		ns = ""
	}
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: ns}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	pcu := &namespacedv1beta1.ProviderConfigUsage{}
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		// TODO(erhan)
		return nil, errors.New("unknown")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}

func ReconciliationPolicy(ctx context.Context, client client.Client, mg resource.Managed) (*v1alpha1.ReconciliationPolicy, error) {
	spec, err := resolveProviderConfig(ctx, client, mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve the referenced ProviderConfig")
	}
	return spec.ReconciliationPolicy, nil
}
