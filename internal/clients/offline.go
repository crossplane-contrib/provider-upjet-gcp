// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"net/http"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/diffserver"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transporttpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// offlinePlaceholderAccessToken is handed to the Terraform Google provider as
// a static "access_token". It is the one credential source the provider
// resolves without ever touching Application Default Credentials: no
// metadata server, no gcloud config, no GOOGLE_APPLICATION_CREDENTIALS file.
// The token is never sent anywhere - the egress guard installed below
// refuses every outbound request before it leaves the process.
const offlinePlaceholderAccessToken = "offline-diff-server-placeholder-token" //nolint:gosec // not a credential, see comment above

const errDiffComputationBlockedFmt = "the offline diff server blocked an outbound GCP request (%s %s): computing this diff requires calling GCP, which is not possible offline"

// OfflineTerraformSetupBuilder builds a terraform.SetupFn that configures the
// Google Terraform provider for computing diffs without ever calling GCP. The
// ProviderConfig is resolved exactly as it is for a live reconcile - project,
// its per-resource override and the universe domain are all honoured - but
// credentials are replaced with a static placeholder and the provider's HTTP
// transport is replaced with one that refuses every outbound request.
func OfflineTerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, crClient client.Client, mg resource.Managed) (terraform.Setup, error) {
		_, ps, err := baseConfiguration(ctx, crClient, mg)
		if err != nil {
			return ps, err
		}
		// Do not resolve real credentials offline: a static access token is
		// the only credential source the provider accepts without falling
		// back to Application Default Credentials.
		ps.Configuration[keyAccessToken] = offlinePlaceholderAccessToken

		// deliberately not using the caller context as context used to configure
		// terraform is stored, but the value below still has to reach Configure:
		// the Google provider's own LoadAndValidate fetches the configured
		// identity's userinfo through a throwaway client it builds straight from
		// this context before the provider's real client - the one the egress
		// guard is installed onto below - even exists, so without this, that one
		// call reaches the network unguarded on every single Configure.
		configureCtx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: egressGuardRoundTripper{}})
		//nolint:contextcheck // deliberate: see the comment above
		if err := configureNoForkGCPClient(configureCtx, &ps, *tfProvider); err != nil {
			return ps, errors.Wrap(err, "failed to configure the offline no-fork GCP client")
		}
		// configureNoForkGCPClient already installed its own transport
		// (metrics included); replace it outright rather than chain it - the
		// guard below never calls a next transport, so nothing underneath it
		// ever runs. Offline diffs depend on this guard, so a client it
		// cannot be installed on is an error rather than a silent degradation.
		config, ok := ps.Meta.(*transporttpg.Config)
		if !ok || config.Client == nil {
			return ps, errors.New("cannot install the offline egress guard: the Terraform setup's Meta is not a configured GCP client")
		}
		config.Client.Transport = egressGuardRoundTripper{}
		return ps, nil
	}
}

// egressGuardRoundTripper refuses every HTTP request it sees. It is installed
// as the Google Terraform provider's transport in offline mode so that a
// resource whose CustomizeDiff calls the GCP API - google_storage_object_acl
// is the one case in the current provider - fails loudly instead of reaching
// the network or silently using whatever credentials the host happens to
// have.
type egressGuardRoundTripper struct{}

func (egressGuardRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, diffserver.NewDiffComputationNotSupportedError(
		errors.Errorf(errDiffComputationBlockedFmt, req.Method, req.URL.Host+req.URL.Path),
	)
}
