// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	upjetmetrics "github.com/crossplane/upjet/v2/pkg/metrics"
	"github.com/google/go-cmp/cmp"
	"github.com/prometheus/client_golang/prometheus/testutil"
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
