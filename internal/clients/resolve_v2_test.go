// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package clients

import (
	"context"
	"testing"

	rtfake "github.com/crossplane/crossplane-runtime/v2/pkg/resource/fake"
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	namespacedv1beta1 "github.com/upbound/provider-gcp/v3/apis/namespaced/v1beta1"
)

// storeKey mirrors the diff server's in-memory client (upjet
// pkg/diffserver/internal.InMemoryClient): it keys its store by GVK plus
// namespace and name, with no awareness that a cluster-scoped kind has no
// namespace to match.
type storeKey struct {
	schema.GroupVersionKind
	types.NamespacedName
}

// exactNamespaceClient requires an exact namespace match on Get, including
// an empty one for a cluster-scoped object. A real API server's
// RESTMapper-aware client ignores a given namespace for a cluster-scoped
// GVK; this one does not, on purpose, so a test against it reproduces the
// exact mechanism the "ClusterProviderConfig not found offline" report
// relies on, rather than risking a generic fake that tolerates the bug by
// accident.
type exactNamespaceClient struct {
	client.Client
	store map[storeKey]client.Object
}

func newExactNamespaceClient(t *testing.T, scheme *runtime.Scheme, objs ...client.Object) *exactNamespaceClient {
	t.Helper()
	store := make(map[storeKey]client.Object, len(objs))
	for _, o := range objs {
		gvks, _, err := scheme.ObjectKinds(o)
		if err != nil || len(gvks) == 0 {
			t.Fatalf("cannot determine the GVK of %T: %v", o, err)
		}
		store[storeKey{GroupVersionKind: gvks[0], NamespacedName: types.NamespacedName{Namespace: o.GetNamespace(), Name: o.GetName()}}] = o
	}
	return &exactNamespaceClient{
		Client: fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build(),
		store:  store,
	}
}

func (c *exactNamespaceClient) Get(_ context.Context, key client.ObjectKey, obj client.Object, _ ...client.GetOption) error {
	gvks, _, err := c.Client.Scheme().ObjectKinds(obj)
	if err != nil || len(gvks) == 0 {
		return err
	}
	stored, ok := c.store[storeKey{GroupVersionKind: gvks[0], NamespacedName: key}]
	if !ok {
		return kerrors.NewNotFound(schema.GroupResource{Group: gvks[0].Group, Resource: gvks[0].Kind}, key.Name)
	}
	return c.Client.Scheme().Convert(stored, obj, nil)
}

// TestResolveV2ClusterProviderConfigAcrossNamespaces reproduces
// https://github.com/crossplane-contrib/provider-upjet-gcp/pull/1044#issuecomment-5992519075:
// a namespaced managed resource referencing a cluster-scoped
// ClusterProviderConfig - the default providerConfigRef for every namespaced
// managed resource when the field is left unset - must resolve it
// regardless of which namespace the managed resource itself is in, since the
// ClusterProviderConfig has none of its own.
func TestResolveV2ClusterProviderConfigAcrossNamespaces(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := namespacedv1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("cannot build the scheme: %v", err)
	}

	cpc := &namespacedv1beta1.ClusterProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: "default"}, // cluster-scoped: no namespace, as the diff server's object store carries it
		Spec: namespacedv1beta1.ProviderConfigSpec{
			ProjectID:   "a-project",
			Credentials: namespacedv1beta1.ProviderCredentials{Source: xpv2.CredentialsSourceInjectedIdentity},
		},
	}
	kc := newExactNamespaceClient(t, scheme, cpc)

	mg := &rtfake.ModernManaged{ObjectMeta: metav1.ObjectMeta{Name: "example", Namespace: "team-a", UID: "test-uid"}}
	mg.SetProviderConfigReference(&xpv2.ProviderConfigReference{
		Kind: namespacedv1beta1.ClusterProviderConfigKind,
		Name: "default",
	})

	spec, err := resolveV2(context.Background(), kc, mg)
	if err != nil {
		t.Fatalf("resolveV2(...): unexpected error resolving a ClusterProviderConfig referenced from a different namespace: %v", err)
	}
	if spec.ProjectID != "a-project" {
		t.Errorf("resolveV2(...): got ProjectID %q, want %q", spec.ProjectID, "a-project")
	}
}

// TestResolveV2NamespacedProviderConfig is the sibling case: a namespaced
// ProviderConfig, resolved from the same namespace it lives in, must keep
// working exactly as before - the fix only changes the lookup namespace for
// a cluster-scoped kind.
func TestResolveV2NamespacedProviderConfig(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := namespacedv1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("cannot build the scheme: %v", err)
	}

	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: "default", Namespace: "team-a"},
		Spec: namespacedv1beta1.ProviderConfigSpec{
			ProjectID:   "team-a-project",
			Credentials: namespacedv1beta1.ProviderCredentials{Source: xpv2.CredentialsSourceInjectedIdentity},
		},
	}
	kc := newExactNamespaceClient(t, scheme, pc)

	mg := &rtfake.ModernManaged{ObjectMeta: metav1.ObjectMeta{Name: "example", Namespace: "team-a", UID: "test-uid"}}
	mg.SetProviderConfigReference(&xpv2.ProviderConfigReference{
		Kind: namespacedv1beta1.ProviderConfigKind,
		Name: "default",
	})

	spec, err := resolveV2(context.Background(), kc, mg)
	if err != nil {
		t.Fatalf("resolveV2(...): unexpected error resolving a namespaced ProviderConfig: %v", err)
	}
	if spec.ProjectID != "team-a-project" {
		t.Errorf("resolveV2(...): got ProjectID %q, want %q", spec.ProjectID, "team-a-project")
	}
}
