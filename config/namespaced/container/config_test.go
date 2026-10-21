// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package container

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	gcpcontainer "github.com/hashicorp/terraform-provider-google/google/services/container"
	"github.com/hashicorp/terraform-provider-google/google/transport"
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
		"unset_computed_count_dropped": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"node_config.0.kubelet_config.#": {Old: "", New: "", NewComputed: true},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{}},
		},
		"observed_kubelet_config_not_removed": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"node_config.0.kubelet_config.#": {Old: "1", New: "0"},
					"node_config.0.machine_type":     {Old: "e2-medium", New: "e2-standard-2"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{
				"node_config.0.machine_type": {Old: "e2-medium", New: "e2-standard-2"},
			}},
		},
		"omitted_subtree_preserved_without_hiding_other_pool_updates": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"node_pool.0.node_config.0.kubelet_config.#":                      {Old: "1", New: "0"},
					"node_pool.0.node_config.0.kubelet_config.0.topology_manager.#":   {Old: "1", New: "0"},
					"node_pool.0.node_config.0.kubelet_config.0.cpu_manager_policy":   {Old: "static", New: ""},
					"node_pool.0.node_config.0.kubelet_config.0.cpu_cfs_quota_period": {Old: "100ms", New: ""},
					"node_pool.1.node_config.0.kubelet_config.#":                      {Old: "1", New: "1"},
					"node_pool.1.node_config.0.kubelet_config.0.cpu_cfs_quota":        {Old: "false", New: "true"},
					"node_pool.0.node_config.0.machine_type":                          {Old: "e2-medium", New: "e2-standard-2", RequiresNew: true},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{
				"node_pool.1.node_config.0.kubelet_config.#":               {Old: "1", New: "1"},
				"node_pool.1.node_config.0.kubelet_config.0.cpu_cfs_quota": {Old: "false", New: "true"},
				"node_pool.0.node_config.0.machine_type":                   {Old: "e2-medium", New: "e2-standard-2", RequiresNew: true},
			}},
		},
		"node_pool_nested_count_dropped": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"node_pool.0.node_config.0.kubelet_config.#": {Old: "", New: "0"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{}},
		},
		"configured_kubelet_config_kept": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"node_config.0.kubelet_config.#":                    {Old: "0", New: "1"},
					"node_config.0.kubelet_config.0.cpu_manager_policy": {Old: "", New: "static"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{
				"node_config.0.kubelet_config.#":                    {Old: "0", New: "1"},
				"node_config.0.kubelet_config.0.cpu_manager_policy": {Old: "", New: "static"},
			}},
		},
		"kubelet_config_outside_node_config_kept": {
			args: args{diff: &terraform.InstanceDiff{
				Attributes: map[string]*terraform.ResourceAttrDiff{
					"kubelet_config.#": {Old: "", New: "0"},
				},
			}},
			want: want{attributes: map[string]*terraform.ResourceAttrDiff{
				"kubelet_config.#": {Old: "", New: "0"},
			}},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			dropEmptyKubeletConfigDiff(tc.args.diff)
			var gotAttrs map[string]*terraform.ResourceAttrDiff
			if tc.args.diff != nil {
				gotAttrs = tc.args.diff.Attributes
			}
			if diff := cmp.Diff(tc.want.attributes, gotAttrs); diff != "" {
				t.Errorf("dropEmptyKubeletConfigDiff(...) attributes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestInjectEmptyKubeletConfig(t *testing.T) {
	type args struct {
		tfMap map[string]any
	}
	type want struct {
		tfMap map[string]any
	}
	cases := map[string]struct {
		args args
		want want
	}{
		"no_node_config": {
			args: args{tfMap: map[string]any{"name": "np"}},
			want: want{tfMap: map[string]any{"name": "np"}},
		},
		"kubelet_config_unset": {
			args: args{tfMap: map[string]any{
				"node_config": []any{map[string]any{"machine_type": "e2-medium"}},
			}},
			want: want{tfMap: map[string]any{
				"node_config": []any{map[string]any{"machine_type": "e2-medium", "kubelet_config": []any{}}},
			}},
		},
		"kubelet_config_null": {
			args: args{tfMap: map[string]any{
				"node_config": []any{map[string]any{"kubelet_config": nil}},
			}},
			want: want{tfMap: map[string]any{
				"node_config": []any{map[string]any{"kubelet_config": []any{}}},
			}},
		},
		"empty_kubelet_config_object_kept": {
			args: args{tfMap: map[string]any{
				"node_config": []any{map[string]any{"kubelet_config": []any{map[string]any{}}}},
			}},
			want: want{tfMap: map[string]any{
				"node_config": []any{map[string]any{"kubelet_config": []any{map[string]any{}}}},
			}},
		},
		"configured_kubelet_config_kept": {
			args: args{tfMap: map[string]any{
				"node_config": []any{map[string]any{"kubelet_config": []any{map[string]any{"cpu_manager_policy": "static"}}}},
			}},
			want: want{tfMap: map[string]any{
				"node_config": []any{map[string]any{"kubelet_config": []any{map[string]any{"cpu_manager_policy": "static"}}}},
			}},
		},
		"unexpected_node_config_type_ignored": {
			args: args{tfMap: map[string]any{
				"node_config": map[string]any{"machine_type": "e2-medium"},
			}},
			want: want{tfMap: map[string]any{
				"node_config": map[string]any{"machine_type": "e2-medium"},
			}},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if err := injectEmptyKubeletConfig(nil, tc.args.tfMap); err != nil {
				t.Fatalf("injectEmptyKubeletConfig(...): unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want.tfMap, tc.args.tfMap); diff != "" {
				t.Errorf("injectEmptyKubeletConfig(...) tfMap mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDropEmptyOptionalComputedBlockDiffs(t *testing.T) {
	cases := map[string]struct {
		attrs   map[string]*terraform.ResourceAttrDiff
		dropped []string
		kept    []string
	}{
		"DropsEmptyNodeDrainConfig": {
			attrs:   map[string]*terraform.ResourceAttrDiff{"node_drain_config.#": {Old: "", New: "", NewComputed: true}},
			dropped: []string{"node_drain_config.#"},
		},
		"DropsEmptyPlacementPolicyAutoscalingQueuedProvisioning": {
			attrs: map[string]*terraform.ResourceAttrDiff{
				"placement_policy.#":    {Old: "", New: ""},
				"autoscaling.#":         {Old: "", New: ""},
				"queued_provisioning.#": {Old: "", New: ""},
			},
			dropped: []string{"placement_policy.#", "autoscaling.#", "queued_provisioning.#"},
		},
		"DropsEmptyMaintenancePolicy": {
			attrs:   map[string]*terraform.ResourceAttrDiff{"maintenance_policy.#": {Old: "", New: "", NewComputed: true}},
			dropped: []string{"maintenance_policy.#"},
		},
		"KeepsNonEmptyMaintenancePolicy": {
			attrs: map[string]*terraform.ResourceAttrDiff{"maintenance_policy.#": {Old: "1", New: ""}},
			kept:  []string{"maintenance_policy.#"},
		},
		"KeepsNonEmptyNodeDrainConfig": {
			attrs: map[string]*terraform.ResourceAttrDiff{"node_drain_config.#": {Old: "", New: "1"}},
			kept:  []string{"node_drain_config.#"},
		},
		"IgnoresUnrelatedKeys": {
			attrs: map[string]*terraform.ResourceAttrDiff{"machine_type": {Old: "", New: ""}},
			kept:  []string{"machine_type"},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			diff := &terraform.InstanceDiff{Attributes: tc.attrs}
			dropEmptyOptionalComputedBlockDiffs(diff)
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

// Exercise the real Google schema and CustomizeDiff, not a hand-built diff.
func TestKubeletConfigOmissionPreservesObservedSubtree(t *testing.T) {
	for _, quota := range []bool{false, true} {
		for _, settings := range []string{"defaults", "topology", "custom"} {
			for _, changeLabels := range []bool{false, true} {
				t.Run(fmt.Sprintf("quota=%t/settings=%s/labels=%t", quota, settings, changeLabels), func(t *testing.T) {
					r := gcpcontainer.ResourceContainerNodePool()
					kubelet := map[string]any{
						"cpu_cfs_quota": quota, "max_parallel_image_pulls": 2,
						"insecure_kubelet_readonly_port_enabled": "FALSE",
					}
					if settings == "topology" {
						kubelet["topology_manager"] = []any{map[string]any{"policy": "best-effort", "scope": "container"}}
					}
					if settings == "custom" {
						kubelet["cpu_manager_policy"] = "static"
						kubelet["cpu_cfs_quota_period"] = "100ms"
					}
					observed := omissionTestParameters(kubelet)
					rd := schema.TestResourceDataRaw(t, r.Schema, observed)
					rd.SetId("projects/test-project/locations/us-central1/clusters/test/nodePools/worker")
					state := rd.State()
					desired := omissionTestParameters(nil)
					if changeLabels {
						desired["node_config"].([]any)[0].(map[string]any)["labels"] = map[string]any{"audit": "yes"}
					}
					if err := injectEmptyKubeletConfig(nil, desired); err != nil {
						t.Fatal(err)
					}
					var err error
					state.RawConfig, err = schema.JSONMapToStateValue(desired, r.CoreConfigSchema())
					if err != nil {
						t.Fatal(err)
					}
					state.RawPlan, err = schema.JSONMapToStateValue(observed, r.CoreConfigSchema())
					if err != nil {
						t.Fatal(err)
					}
					diff, err := schema.InternalMap(r.Schema).Diff(t.Context(), state, terraform.NewResourceConfigRaw(desired), r.CustomizeDiff, &transport.Config{Project: "test-project"}, false)
					if err != nil {
						t.Fatal(err)
					}
					dropEmptyKubeletConfigDiff(diff)
					planned, err := schema.InternalMap(r.Schema).Data(state, diff)
					if err != nil {
						t.Fatal(err)
					}
					if planned.HasChange("node_config.0.kubelet_config") {
						t.Error("omission must not enter the Terraform kubelet update handler")
					}
					if delta := cmp.Diff(rd.Get("node_config.0.kubelet_config"), planned.Get("node_config.0.kubelet_config")); delta != "" {
						t.Errorf("omission changed observed kubelet settings (-observed +planned):\n%s", delta)
					}
					if planned.HasChange("node_config.0.labels") != changeLabels {
						t.Error("kubelet suppression must preserve unrelated label updates")
					}
					if diff != nil && (diff.Destroy || diff.RequiresNew()) {
						t.Error("omission must not require replacing the node pool")
					}
				})
			}
		}
	}
}

func omissionTestParameters(kubelet map[string]any) map[string]any {
	node := map[string]any{"machine_type": "e2-medium"}
	if kubelet != nil {
		node["kubelet_config"] = []any{kubelet}
	}
	return map[string]any{
		"name": "worker", "project": "test-project", "location": "us-central1",
		"cluster": "test", "deletion_policy": "DELETE", "node_config": []any{node},
	}
}
