// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"testing"
	"time"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	rtfake "github.com/crossplane/crossplane-runtime/v2/pkg/resource/fake"
	"github.com/crossplane/crossplane-runtime/v2/pkg/test"
	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/crossplane/upjet/v2/apis/configuration/v1alpha1"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	clusterv1beta1 "github.com/upbound/provider-gcp/v3/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-gcp/v3/apis/namespaced/v1beta1"
)

const (
	testManagedUID = "test-uid"
	testNamespace  = "team-a"
	testConfigName = "default"
)

func testPolicy() *v1alpha1.ReconciliationPolicy {
	return &v1alpha1.ReconciliationPolicy{
		ExponentialFailureRateLimiter: &v1alpha1.ExponentialFailureRateLimiter{
			BaseDelay: &metav1.Duration{Duration: 5 * time.Second},
			MaxDelay:  &metav1.Duration{Duration: 10 * time.Minute},
		},
	}
}

func modernManaged(kind string) *rtfake.ModernManaged {
	mg := &rtfake.ModernManaged{ObjectMeta: metav1.ObjectMeta{Name: "example", Namespace: testNamespace, UID: testManagedUID}}
	mg.SetProviderConfigReference(&xpv2.ProviderConfigReference{Kind: kind, Name: testConfigName})
	return mg
}

func legacyManaged() *rtfake.LegacyManaged {
	mg := &rtfake.LegacyManaged{ObjectMeta: metav1.ObjectMeta{Name: "example", UID: testManagedUID}}
	mg.SetProviderConfigReference(&xpv2.Reference{Name: testConfigName})
	return mg
}

func newProviderConfigClient(t *testing.T, objs ...client.Object) client.Client {
	t.Helper()
	scheme := runtime.NewScheme()
	if err := clusterv1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("cannot add the cluster-scoped types to the scheme: %v", err)
	}
	if err := namespacedv1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("cannot add the namespaced types to the scheme: %v", err)
	}
	return fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()
}

func trackedUsages(ctx context.Context, t *testing.T, kube client.Client) []string {
	t.Helper()
	var names []string

	legacy := &clusterv1beta1.ProviderConfigUsageList{}
	if err := kube.List(ctx, legacy); err != nil {
		t.Fatalf("cannot list the cluster-scoped ProviderConfigUsages: %v", err)
	}
	for _, u := range legacy.Items {
		names = append(names, "cluster/"+u.GetName())
	}

	modern := &namespacedv1beta1.ProviderConfigUsageList{}
	if err := kube.List(ctx, modern); err != nil {
		t.Fatalf("cannot list the namespaced ProviderConfigUsages: %v", err)
	}
	for _, u := range modern.Items {
		names = append(names, u.GetNamespace()+"/"+u.GetName())
	}
	return names
}

func TestReconciliationPolicy(t *testing.T) {
	type args struct {
		objs []client.Object
		mg   resource.Managed
	}
	type want struct {
		policy *v1alpha1.ReconciliationPolicy
		usages []string
		err    error
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"ClusterProviderConfigWithPolicy": {
			reason: "The policy of a ClusterProviderConfig is returned without tracking its usage.",
			args: args{
				objs: []client.Object{&namespacedv1beta1.ClusterProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: testConfigName},
					Spec:       namespacedv1beta1.ProviderConfigSpec{ReconciliationPolicy: testPolicy()},
				}},
				mg: modernManaged(namespacedv1beta1.ClusterProviderConfigKind),
			},
			want: want{policy: testPolicy()},
		},
		"NamespacedProviderConfigWithPolicy": {
			reason: "The policy of a namespaced ProviderConfig is returned without tracking its usage.",
			args: args{
				objs: []client.Object{&namespacedv1beta1.ProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: testConfigName, Namespace: testNamespace},
					Spec:       namespacedv1beta1.ProviderConfigSpec{ReconciliationPolicy: testPolicy()},
				}},
				mg: modernManaged(namespacedv1beta1.ProviderConfigKind),
			},
			want: want{policy: testPolicy()},
		},
		"LegacyProviderConfigWithPolicy": {
			reason: "The policy of a cluster-scoped legacy ProviderConfig is returned without tracking its usage.",
			args: args{
				objs: []client.Object{&clusterv1beta1.ProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: testConfigName},
					Spec:       clusterv1beta1.ProviderConfigSpec{ReconciliationPolicy: testPolicy()},
				}},
				mg: legacyManaged(),
			},
			want: want{policy: testPolicy()},
		},
		"NoPolicy": {
			reason: "A ProviderConfig without a policy yields no policy and no usage.",
			args: args{
				objs: []client.Object{&namespacedv1beta1.ClusterProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: testConfigName},
				}},
				mg: modernManaged(namespacedv1beta1.ClusterProviderConfigKind),
			},
		},
		"NoProviderConfigReference": {
			reason: "A managed resource without a providerConfigRef is an error.",
			args: args{
				mg: &rtfake.ModernManaged{ObjectMeta: metav1.ObjectMeta{Name: "example", Namespace: testNamespace, UID: testManagedUID}},
			},
			want: want{err: errors.Wrap(errors.New(errNoProviderConfig), errResolveProviderConfig)},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			kube := newProviderConfigClient(t, tc.args.objs...)

			policy, err := ReconciliationPolicy(ctx, kube, tc.args.mg)

			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nReconciliationPolicy(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.policy, policy); diff != "" {
				t.Errorf("\n%s\nReconciliationPolicy(...): -want policy, +got policy:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.usages, trackedUsages(ctx, t, kube)); diff != "" {
				t.Errorf("\n%s\nReconciliationPolicy(...): -want usages, +got usages:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestResolveProviderConfigTracksUsage(t *testing.T) {
	type args struct {
		objs []client.Object
		mg   resource.Managed
	}
	type want struct {
		projectID string
		usages    []string
	}

	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"ModernManaged": {
			reason: "Resolving the ProviderConfig of a namespaced managed resource tracks its usage once.",
			args: args{
				objs: []client.Object{&namespacedv1beta1.ClusterProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: testConfigName},
					Spec:       namespacedv1beta1.ProviderConfigSpec{ProjectID: "a-project"},
				}},
				mg: modernManaged(namespacedv1beta1.ClusterProviderConfigKind),
			},
			want: want{projectID: "a-project", usages: []string{testNamespace + "/" + testManagedUID}},
		},
		"LegacyManaged": {
			reason: "Resolving the ProviderConfig of a cluster-scoped managed resource tracks its usage once.",
			args: args{
				objs: []client.Object{&clusterv1beta1.ProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: testConfigName},
					Spec:       clusterv1beta1.ProviderConfigSpec{ProjectID: "a-project"},
				}},
				mg: legacyManaged(),
			},
			want: want{projectID: "a-project", usages: []string{"cluster/" + testManagedUID}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			kube := newProviderConfigClient(t, tc.args.objs...)

			spec, err := resolveProviderConfig(ctx, kube, tc.args.mg)
			if err != nil {
				t.Fatalf("\n%s\nresolveProviderConfig(...): unexpected error: %v", tc.reason, err)
			}

			if diff := cmp.Diff(tc.want.projectID, spec.ProjectID); diff != "" {
				t.Errorf("\n%s\nresolveProviderConfig(...): -want project, +got project:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.usages, trackedUsages(ctx, t, kube)); diff != "" {
				t.Errorf("\n%s\nresolveProviderConfig(...): -want usages, +got usages:\n%s", tc.reason, diff)
			}
		})
	}
}
