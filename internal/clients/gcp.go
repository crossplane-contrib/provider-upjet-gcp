/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/terraform"
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
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

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

		switch pc.Spec.Credentials.Source { //nolint:exhaustive
		case xpv1.CredentialsSourceInjectedIdentity:
			// We don't need to do anything here, as the TF Provider will take care of workloadIdentity etc.
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

		return ps, nil
	}
}
