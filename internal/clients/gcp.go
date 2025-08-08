// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/upbound/provider-gcp/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-gcp/apis/namespaced/v1beta1"
)

const (
	keyProject = "project"

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

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
// NOTE(hasheddan): this function is slightly over our cyclomatic complexity
// goal. Consider refactoring before adding new branches.
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, crClient client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}
		pcSpec, err := resolveProviderConfig(ctx, crClient, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}
		// set provider configuration
		ps.Configuration = map[string]interface{}{
			keyProject: pcSpec.ProjectID,
		}
		// TODO: this will have a performance impact. We need to quantify this.
		p, err := fieldpath.PaveObject(mg, fieldpath.WithMaxFieldPathIndex(1))
		if err != nil {
			return ps, errors.Wrapf(err, errPaveFmt, mg.GetObjectKind().GroupVersionKind().Kind, mg.GetName())
		}
		// TODO: if the managed resource declares its project
		//  in a different parameter, the following will not work.
		resourceProject, err := p.GetString("spec.forProvider.project")
		if err != nil && !fieldpath.IsNotFound(err) {
			return ps, errors.Wrapf(err, errPavedGetValueFmt, mg.GetObjectKind().GroupVersionKind().Kind, mg.GetName())
		}
		if err == nil && resourceProject != "" {
			ps.Configuration[keyProject] = resourceProject
		}

		switch pcSpec.Credentials.Source { //nolint:exhaustive
		case xpv1.CredentialsSourceInjectedIdentity:
			// We don't need to do anything here, as the TF Provider will take care of workloadIdentity etc.
		case impersonateServiceAccount:
			if pcSpec.Credentials.ImpersonateServiceAccountSpec.Name != "" {
				ps.Configuration[keyImpersonateServiceAccount] = pcSpec.Credentials.ImpersonateServiceAccountSpec.Name
			}
		case credentialsSourceAccessToken:
			data, err := resource.CommonCredentialExtractor(ctx, xpv1.CredentialsSourceSecret, crClient, pcSpec.Credentials.CommonCredentialSelectors)
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
		return ps, errors.Wrap(configureNoForkGCPClient(&ps, *tfProvider), "failed to configure the no-fork GCP client")
	}
}

func configureNoForkGCPClient(ps *terraform.Setup, p schema.Provider) error {
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
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(providerTimeout, cancel)

	diag := p.Configure(ctx, &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
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

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
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
