// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"encoding/json"
	"fmt"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/provider-gcp/apis/v1beta1"
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
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}
		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		// set provider configuration
		ps.Configuration = map[string]interface{}{
			keyProject: pc.Spec.ProjectID,
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

		switch pc.Spec.Credentials.Source { //nolint:exhaustive
		case xpv1.CredentialsSourceInjectedIdentity:
			// We don't need to do anything here, as the TF Provider will take care of workloadIdentity etc.
		case impersonateServiceAccount:
			if pc.Spec.Credentials.ImpersonateServiceAccountSpec.Name != "" {
				ps.Configuration[keyImpersonateServiceAccount] = pc.Spec.Credentials.ImpersonateServiceAccountSpec.Name
			}
		case credentialsSourceAccessToken:
			data, err := resource.CommonCredentialExtractor(ctx, xpv1.CredentialsSourceSecret, client, pc.Spec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, errExtractTokenCredentials)
			}
			ps.Configuration[keyAccessToken] = string(data)
		case credentialsSourceUpbound:
			if pc.Spec.Credentials.Upbound == nil || pc.Spec.Credentials.Upbound.Federation == nil {
				return ps, errors.Wrap(errors.New(errMissingFederatedConfiguration), errConstructFederatedCredentials)
			}
			data, err := constructFederatedCredentials(pc.Spec.Credentials.Upbound.Federation.ProviderID, pc.Spec.Credentials.Upbound.Federation.ServiceAccount)
			if err != nil {
				return ps, errors.Wrap(err, errConstructFederatedCredentials)
			}
			ps.Configuration[keyCredentials] = string(data)
		default:
			data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, errExtractKeyCredentials)
			}
			ps.Configuration[keyCredentials] = string(data)
		}

		return ps, errors.Wrap(configureNoForkGCPClient(ctx, &ps, *tfProvider), "failed to configure the no-fork GCP client")
	}
}

func configureNoForkGCPClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}
