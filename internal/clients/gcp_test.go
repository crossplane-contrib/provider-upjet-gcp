// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	upjetmetrics "github.com/crossplane/upjet/v2/pkg/metrics"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"golang.org/x/oauth2"

	namespacedv1beta1 "github.com/upbound/provider-gcp/v3/apis/namespaced/v1beta1"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func Test_serviceFromURL(t *testing.T) {
	type args struct {
		rawURL string
	}
	type want struct {
		service string
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"service_in_host": {
			args: args{rawURL: "https://compute.googleapis.com/compute/v1/projects/p/zones/z/instances"},
			want: want{service: "compute"},
		},
		"generic_host_path_fallback": {
			args: args{rawURL: "https://www.googleapis.com/storage/v1/b/bucket"},
			want: want{service: "storage"},
		},
		"host_with_port": {
			args: args{rawURL: "https://compute.googleapis.com:443/compute/v1/projects/p"},
			want: want{service: "compute"},
		},
		"generic_host_no_path": {
			args: args{rawURL: "https://www.googleapis.com"},
			want: want{service: "unknown"},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			u, err := url.Parse(tc.args.rawURL)
			if err != nil {
				t.Fatalf("url.Parse(%q): %v", tc.args.rawURL, err)
			}
			got := serviceFromURL(u)
			if diff := cmp.Diff(tc.want.service, got); diff != "" {
				t.Errorf("serviceFromURL(%q) mismatch (-want +got):\n%s", tc.args.rawURL, diff)
			}
		})
	}
}

func Test_setProjectOverrides(t *testing.T) {
	type args struct {
		pcSpec *namespacedv1beta1.ProviderConfigSpec
	}
	type want struct {
		configuration map[string]interface{}
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"unset_fields_leave_configuration_untouched": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{ProjectID: "example-project"}},
			want: want{configuration: map[string]interface{}{keyProject: "example-project"}},
		},
		"override_enabled": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				ProjectID:           "example-project",
				UserProjectOverride: new(true),
			}},
			want: want{configuration: map[string]interface{}{
				keyProject:             "example-project",
				keyUserProjectOverride: true,
			}},
		},
		"override_disabled_explicitly": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				ProjectID:           "example-project",
				UserProjectOverride: new(false),
			}},
			want: want{configuration: map[string]interface{}{
				keyProject:             "example-project",
				keyUserProjectOverride: false,
			}},
		},
		"billing_project_set": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				ProjectID:      "example-project",
				BillingProject: new("billed-project"),
			}},
			want: want{configuration: map[string]interface{}{
				keyProject:        "example-project",
				keyBillingProject: "billed-project",
			}},
		},
		"empty_billing_project_ignored": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				ProjectID:      "example-project",
				BillingProject: new(""),
			}},
			want: want{configuration: map[string]interface{}{keyProject: "example-project"}},
		},
		"override_and_billing_project": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				ProjectID:           "example-project",
				UserProjectOverride: new(true),
				BillingProject:      new("billed-project"),
			}},
			want: want{configuration: map[string]interface{}{
				keyProject:             "example-project",
				keyUserProjectOverride: true,
				keyBillingProject:      "billed-project",
			}},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			cfg := map[string]interface{}{keyProject: tc.args.pcSpec.ProjectID}
			setProjectOverrides(cfg, tc.args.pcSpec)
			if diff := cmp.Diff(tc.want.configuration, cfg); diff != "" {
				t.Errorf("setProjectOverrides(...) configuration mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_setUniverseDomain(t *testing.T) {
	type args struct {
		pcSpec *namespacedv1beta1.ProviderConfigSpec
	}
	type want struct {
		configuration map[string]interface{}
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"unset_field_leaves_configuration_untouched": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{ProjectID: "example-project"}},
			want: want{configuration: map[string]interface{}{keyProject: "example-project"}},
		},
		"universe_domain_set": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				ProjectID:      "example-project",
				UniverseDomain: new("example.partner.com"),
			}},
			want: want{configuration: map[string]interface{}{
				keyProject:        "example-project",
				keyUniverseDomain: "example.partner.com",
			}},
		},
		"default_universe_domain_set": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				ProjectID:      "example-project",
				UniverseDomain: new("googleapis.com"),
			}},
			want: want{configuration: map[string]interface{}{
				keyProject:        "example-project",
				keyUniverseDomain: "googleapis.com",
			}},
		},
		"empty_universe_domain_ignored": {
			args: args{pcSpec: &namespacedv1beta1.ProviderConfigSpec{
				ProjectID:      "example-project",
				UniverseDomain: new(""),
			}},
			want: want{configuration: map[string]interface{}{keyProject: "example-project"}},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			cfg := map[string]interface{}{keyProject: tc.args.pcSpec.ProjectID}
			setUniverseDomain(cfg, tc.args.pcSpec)
			if diff := cmp.Diff(tc.want.configuration, cfg); diff != "" {
				t.Errorf("setUniverseDomain(...) configuration mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_metricsRoundTripper(t *testing.T) {
	errBoom := errors.New("boom")
	type args struct {
		rawURL  string
		method  string
		baseErr error
	}
	type want struct {
		inc float64
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"counts_successful_call": {
			args: args{rawURL: "https://compute.googleapis.com/compute/v1/projects/p", method: http.MethodGet},
			want: want{inc: 1},
		},
		"does_not_count_transport_error": {
			args: args{rawURL: "https://dns.googleapis.com/dns/v1/projects/p", method: http.MethodPost, baseErr: errBoom},
			want: want{inc: 0},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			u, err := url.Parse(tc.args.rawURL)
			if err != nil {
				t.Fatalf("url.Parse(%q): %v", tc.args.rawURL, err)
			}
			counter := upjetmetrics.ExternalAPICalls.WithLabelValues(serviceFromURL(u), tc.args.method)
			before := testutil.ToFloat64(counter)

			var wantResp *http.Response
			if tc.args.baseErr == nil {
				wantResp = &http.Response{StatusCode: http.StatusOK}
			}
			base := roundTripperFunc(func(*http.Request) (*http.Response, error) {
				return wantResp, tc.args.baseErr
			})
			rt := &metricsRoundTripper{base: base}

			resp, err := rt.RoundTrip(&http.Request{Method: tc.args.method, URL: u})
			if !errors.Is(err, tc.args.baseErr) {
				t.Fatalf("RoundTrip err = %v, want %v", err, tc.args.baseErr)
			}
			if resp != wantResp {
				t.Errorf("RoundTrip response not passed through unchanged")
			}
			if diff := cmp.Diff(tc.want.inc, testutil.ToFloat64(counter)-before); diff != "" {
				t.Errorf("counter delta mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// Test_configureNoForkGCPClient_propagatesContextValue reproduces
// https://github.com/crossplane-contrib/provider-upjet-gcp/pull/1044#issuecomment-5992112807:
// the offline egress guard is only installed on the provider's client after
// Configure returns, so a value Configure's own call chain needs - the
// Google provider's LoadAndValidate reads an *http.Client off the context
// under the oauth2.HTTPClient key for its userinfo lookup, before the real
// client exists - has to reach Configure through the context
// configureNoForkGCPClient is given, not through anything set up afterward.
// This does not re-verify the Google provider's own internals (already
// confirmed by reading the pinned source); it verifies the one thing this
// package is responsible for: that a value placed on the context passed in
// survives into the context p.Configure actually receives.
func Test_configureNoForkGCPClient_propagatesContextValue(t *testing.T) {
	want := &http.Client{}
	var got *http.Client

	p := schema.Provider{
		ConfigureContextFunc: func(ctx context.Context, _ *schema.ResourceData) (interface{}, diag.Diagnostics) {
			got, _ = ctx.Value(oauth2.HTTPClient).(*http.Client)
			return struct{}{}, nil
		},
	}

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, want)
	ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}

	if err := configureNoForkGCPClient(ctx, ps, p); err != nil {
		t.Fatalf("configureNoForkGCPClient(...): unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("configureNoForkGCPClient(...): Configure saw HTTP client %v, want the one placed on the given context (%v)", got, want)
	}
}
