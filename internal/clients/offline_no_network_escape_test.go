// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package clients

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/hashicorp/terraform-provider-google/google/provider"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/upbound/provider-gcp/v3/apis/namespaced"
	storagev1beta2 "github.com/upbound/provider-gcp/v3/apis/namespaced/storage/v1beta2"
	namespacedv1beta1 "github.com/upbound/provider-gcp/v3/apis/namespaced/v1beta1"
)

// TestOfflineTerraformSetupBuilder_NoNetworkEscapeDuringConfigure is the
// end-to-end counterpart to Test_configureNoForkGCPClient_propagatesContextValue:
// where that test only proves our own code threads a context value through,
// this one runs OfflineTerraformSetupBuilder with the real, vendored Google
// provider and asserts the specific outcome
// https://github.com/crossplane-contrib/provider-upjet-gcp/pull/1044#issuecomment-5992112807
// reported - that LoadAndValidate's userinfo lookup reaches
// openidconnect.googleapis.com before the egress guard is installed on the
// provider's own client - no longer happens.
//
// The tripwire is scoped to that one host rather than all of
// http.DefaultTransport on purpose: the GCE metadata probe (metadata.OnGCE)
// also bypasses the guard and also uses http.DefaultTransport, but that gap
// was deliberately left as a documented, accepted limitation rather than
// fixed, and this test should not fail because of it.
func TestOfflineTerraformSetupBuilder_NoNetworkEscapeDuringConfigure(t *testing.T) {
	const guardedHost = "openidconnect.googleapis.com"

	orig := http.DefaultTransport
	var escaped *http.Request
	http.DefaultTransport = roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if r.URL.Host == guardedHost {
			escaped = r
			return nil, fmt.Errorf("blocked by test tripwire: a request reached the real network at %s %s", r.Method, r.URL)
		}
		return orig.RoundTrip(r)
	})
	t.Cleanup(func() { http.DefaultTransport = orig })

	scheme := runtime.NewScheme()
	if err := namespaced.AddToScheme(scheme); err != nil {
		t.Fatalf("cannot build the scheme: %v", err)
	}
	cpc := &namespacedv1beta1.ClusterProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: "default"},
		Spec: namespacedv1beta1.ProviderConfigSpec{
			ProjectID:   "a-project",
			Credentials: namespacedv1beta1.ProviderCredentials{Source: xpv2.CredentialsSourceInjectedIdentity},
		},
	}
	kc := newExactNamespaceClient(t, scheme, cpc)

	bucket := &storagev1beta2.Bucket{
		ObjectMeta: metav1.ObjectMeta{Name: "example", Namespace: "team-a", UID: "test-uid"},
		Spec: storagev1beta2.BucketSpec{
			ManagedResourceSpec: xpv2.ManagedResourceSpec{
				ProviderConfigReference: &xpv2.ProviderConfigReference{
					Kind: namespacedv1beta1.ClusterProviderConfigKind,
					Name: "default",
				},
			},
		},
	}

	setup := OfflineTerraformSetupBuilder(provider.Provider())
	if _, err := setup(context.Background(), kc, bucket); err != nil {
		t.Fatalf("OfflineTerraformSetupBuilder(...)(...): unexpected error: %v", err)
	}
	if escaped != nil {
		t.Errorf("a request reached %s unguarded during Configure: %s %s", guardedHost, escaped.Method, escaped.URL)
	}
}
