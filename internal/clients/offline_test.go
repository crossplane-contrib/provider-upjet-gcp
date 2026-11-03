// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/diffserver"
)

func Test_egressGuardRoundTripper(t *testing.T) {
	type args struct {
		method string
		rawURL string
	}
	cases := map[string]struct {
		args args
	}{
		"storage_object_get": {
			args: args{method: http.MethodGet, rawURL: "https://storage.googleapis.com/storage/v1/b/example-bucket/o/example-object"},
		},
		"compute_post": {
			args: args{method: http.MethodPost, rawURL: "https://compute.googleapis.com/compute/v1/projects/p/zones/z/instances"},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			u, err := url.Parse(tc.args.rawURL)
			if err != nil {
				t.Fatalf("url.Parse(%q): %v", tc.args.rawURL, err)
			}

			resp, err := (egressGuardRoundTripper{}).RoundTrip(&http.Request{Method: tc.args.method, URL: u})

			if resp != nil {
				t.Errorf("RoundTrip response = %v, want nil", resp)
			}
			if err == nil {
				t.Fatalf("RoundTrip err = nil, want a blocked-request error")
			}
			if !diffserver.IsDiffComputationNotSupportedError(err) {
				t.Errorf("IsDiffComputationNotSupportedError(%v) = false, want true", err)
			}
			if !strings.Contains(err.Error(), tc.args.method) || !strings.Contains(err.Error(), u.Host+u.Path) {
				t.Errorf("error %q does not identify the blocked request (%s %s)", err.Error(), tc.args.method, u.Host+u.Path)
			}
		})
	}
}

func Test_offlinePlaceholderAccessToken(t *testing.T) {
	// A non-empty static token is what makes the Google provider pick the
	// access_token credential path instead of falling back to Application
	// Default Credentials - see GetCredentials in the vendored provider.
	if offlinePlaceholderAccessToken == "" {
		t.Error("offlinePlaceholderAccessToken is empty: the provider would fall back to Application Default Credentials")
	}
}
