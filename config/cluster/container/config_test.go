// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package container

import (
	"encoding/base64"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/client-go/tools/clientcmd"
)

func Test_suppressEnableComponentsOrderDiff(t *testing.T) {
	type args struct {
		diff *terraform.InstanceDiff
	}
	type want struct {
		attributes map[string]*terraform.ResourceAttrDiff
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"nil_diff": {
			args: args{diff: nil},
			want: want{attributes: nil},
		},
		"no_enable_components_keys": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"node_count": {Old: "1", New: "2"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{
				"node_count": {Old: "1", New: "2"},
			}},
		},
		"pure_reorder_with_count_key": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"monitoring_config.0.enable_components.#": {Old: "2", New: "2"},
					"monitoring_config.0.enable_components.0": {Old: "SYSTEM_COMPONENTS", New: "APISERVER"},
					"monitoring_config.0.enable_components.1": {Old: "APISERVER", New: "SYSTEM_COMPONENTS"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{}},
		},
		"pure_reorder_without_count_key": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"monitoring_config.0.enable_components.0": {Old: "SYSTEM_COMPONENTS", New: "APISERVER"},
					"monitoring_config.0.enable_components.1": {Old: "APISERVER", New: "SYSTEM_COMPONENTS"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{}},
		},
		"real_change_element_replaced": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"monitoring_config.0.enable_components.#": {Old: "2", New: "2"},
					"monitoring_config.0.enable_components.0": {Old: "APISERVER", New: "WORKLOADS"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{
				"monitoring_config.0.enable_components.#": {Old: "2", New: "2"},
				"monitoring_config.0.enable_components.0": {Old: "APISERVER", New: "WORKLOADS"},
			}},
		},
		"real_change_count_increased": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"monitoring_config.0.enable_components.#": {Old: "1", New: "2"},
					"monitoring_config.0.enable_components.0": {Old: "SYSTEM_COMPONENTS", New: "SYSTEM_COMPONENTS"},
					"monitoring_config.0.enable_components.1": {Old: "", New: "APISERVER"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{
				"monitoring_config.0.enable_components.#": {Old: "1", New: "2"},
				"monitoring_config.0.enable_components.0": {Old: "SYSTEM_COMPONENTS", New: "SYSTEM_COMPONENTS"},
				"monitoring_config.0.enable_components.1": {Old: "", New: "APISERVER"},
			}},
		},
		"unrelated_keys_preserved_on_reorder": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"node_count": {Old: "1", New: "2"},
					"monitoring_config.0.enable_components.0": {Old: "SYSTEM_COMPONENTS", New: "APISERVER"},
					"monitoring_config.0.enable_components.1": {Old: "APISERVER", New: "SYSTEM_COMPONENTS"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{
				"node_count": {Old: "1", New: "2"},
			}},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			suppressEnableComponentsOrderDiff(tc.args.diff)
			var gotAttrs map[string]*terraform.ResourceAttrDiff
			if tc.args.diff != nil {
				gotAttrs = tc.args.diff.Attributes
			}
			if diff := cmp.Diff(tc.want.attributes, gotAttrs); diff != "" {
				t.Errorf("suppressEnableComponentsOrderDiff(...) attributes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDropEmptyKubeletConfigDiff(t *testing.T) {
	cases := map[string]struct {
		attrs   map[string]*terraform.ResourceAttrDiff
		dropped []string // keys expected to be removed
		kept    []string // keys expected to remain
	}{
		"DropsEmptyNodePoolNested": {
			attrs: map[string]*terraform.ResourceAttrDiff{
				"node_pool.0.node_config.0.kubelet_config.#": {Old: "", New: "0"},
			},
			dropped: []string{"node_pool.0.node_config.0.kubelet_config.#"},
		},
		"DropsEmptyTopLevel": {
			attrs: map[string]*terraform.ResourceAttrDiff{
				"node_config.0.kubelet_config.#": {Old: "0", New: "0"},
			},
			dropped: []string{"node_config.0.kubelet_config.#"},
		},
		"KeepsNonEmptyKubeletConfig": {
			attrs: map[string]*terraform.ResourceAttrDiff{
				"node_config.0.kubelet_config.#": {Old: "0", New: "1"},
			},
			kept: []string{"node_config.0.kubelet_config.#"},
		},
		"IgnoresUnrelatedKubeletConfigOutsideNodeConfig": {
			attrs: map[string]*terraform.ResourceAttrDiff{
				"kubelet_config.#": {Old: "", New: "0"},
			},
			kept: []string{"kubelet_config.#"},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			diff := &terraform.InstanceDiff{Attributes: tc.attrs}
			dropEmptyKubeletConfigDiff(diff)
			for _, k := range tc.dropped {
				if _, ok := diff.Attributes[k]; ok {
					t.Errorf("expected key %q to be dropped, but it remained", k)
				}
			}
			for _, k := range tc.kept {
				if _, ok := diff.Attributes[k]; !ok {
					t.Errorf("expected key %q to be kept, but it was dropped", k)
				}
			}
		})
	}
}

func TestClusterConnectionDetails(t *testing.T) {
	caPEM := []byte("-----BEGIN CERTIFICATE-----\nprivate-cluster-ca\n-----END CERTIFICATE-----")

	masterAuth := []interface{}{map[string]interface{}{
		"cluster_ca_certificate": base64.StdEncoding.EncodeToString(caPEM),
		"client_certificate":     "",
		"client_key":             "",
	}}
	controlPlaneEndpoints := func(dnsEndpoint string) []interface{} {
		return []interface{}{map[string]interface{}{
			"dns_endpoint_config": []interface{}{map[string]interface{}{
				"endpoint": dnsEndpoint,
			}},
		}}
	}

	type args struct {
		attr map[string]interface{}
	}
	type want struct {
		server string
		caData []byte
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"IPEndpointPinsClusterCA": {
			reason: "An IP endpoint is served by the cluster's own CA, so it must stay pinned.",
			args: args{attr: map[string]interface{}{
				"name":                           "example-cluster",
				"endpoint":                       "34.10.20.30",
				"master_auth":                    masterAuth,
				"control_plane_endpoints_config": controlPlaneEndpoints("uid.europe-west8.gke.goog"),
			}},
			want: want{server: "https://34.10.20.30", caData: caPEM},
		},
		"DNSEndpointOmitsClusterCA": {
			reason: "The DNS endpoint is fronted by Google Front End with a publicly trusted certificate, so pinning the private cluster CA would fail TLS validation.",
			args: args{attr: map[string]interface{}{
				"name":                           "example-cluster",
				"endpoint":                       "uid.europe-west8.gke.goog",
				"master_auth":                    masterAuth,
				"control_plane_endpoints_config": controlPlaneEndpoints("uid.europe-west8.gke.goog"),
			}},
			want: want{server: "https://uid.europe-west8.gke.goog", caData: nil},
		},
		"NoControlPlaneEndpointsConfigPinsClusterCA": {
			reason: "A cluster without controlPlaneEndpointsConfig must keep the pre-existing behaviour.",
			args: args{attr: map[string]interface{}{
				"name":        "example-cluster",
				"endpoint":    "34.10.20.30",
				"master_auth": masterAuth,
			}},
			want: want{server: "https://34.10.20.30", caData: caPEM},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := clusterConnectionDetails(tc.args.attr)
			if err != nil {
				t.Fatalf("ClusterConnectionDetails(...): unexpected error: %v\nreason: %s", err, tc.reason)
			}
			kc, err := clientcmd.Load(got["kubeconfig"])
			if err != nil {
				t.Fatalf("cannot load generated kubeconfig: %v", err)
			}
			cluster := kc.Clusters[kc.CurrentContext]
			result := want{server: cluster.Server, caData: cluster.CertificateAuthorityData}
			if diff := cmp.Diff(tc.want, result, cmp.AllowUnexported(want{})); diff != "" {
				t.Errorf("clusterConnectionDetails(...): -want, +got:\n%s\nreason: %s", diff, tc.reason)
			}
		})
	}
}
