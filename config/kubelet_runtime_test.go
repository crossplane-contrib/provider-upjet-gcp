// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	uc "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	fp "github.com/hashicorp/terraform-provider-google/google/fwprovider"
	gp "github.com/hashicorp/terraform-provider-google/google/provider"
	transport "github.com/hashicorp/terraform-provider-google/google/transport"
	pc "github.com/upbound/provider-gcp/v3/config"
)

// Exercise the same startup path as cmd/provider/container/zz_main.go.
// Cluster MoveToStatus mutates TF's shared schemaNodePool map, including
// standalone NodePool children. A bare ResourceContainerNodePool is not faithful.
func TestRuntimeKubeletPreservation(t *testing.T) {
	sdk := gp.Provider()
	fw := fp.New(sdk)
	cluster, err := pc.GetProvider(context.Background(), sdk, fw, false)
	if err != nil {
		t.Fatal(err)
	}
	namespaced, err := pc.GetNamespacedProvider(context.Background(), sdk, fw, false)
	if err != nil {
		t.Fatal(err)
	}
	for scope, p := range map[string]*uc.Provider{"cluster": cluster, "namespaced": namespaced} {
		conf := p.Resources["google_container_node_pool"]
		r := conf.TerraformResource
		for _, quota := range []bool{false, true} {
			for _, settings := range []string{"defaults", "topology", "custom"} {
				for _, desired := range []string{"omit", "omit-label-update", "true"} {
					for _, hook := range []string{"atlassian2", "subtree"} {
						t.Run(fmt.Sprintf("%s/quota=%t/%s/%s/%s", scope, quota, settings, desired, hook), func(t *testing.T) {
							kubelet := map[string]any{"cpu_cfs_quota": quota, "insecure_kubelet_readonly_port_enabled": "FALSE", "max_parallel_image_pulls": 2}
							if settings == "topology" {
								kubelet["topology_manager"] = []any{map[string]any{"policy": "best-effort", "scope": "container"}}
							}
							if settings == "custom" {
								kubelet["cpu_manager_policy"] = "static"
								kubelet["cpu_cfs_quota_period"] = "100ms"
							}
							observed := params(kubelet)
							rd := schema.TestResourceDataRaw(t, r.Schema, observed)
							rd.SetId("projects/test-project/locations/us-central1/clusters/test/nodePools/worker")
							state := rd.State()
							wanted := params(nil)
							if desired == "true" {
								wanted = params(map[string]any{"cpu_cfs_quota": true})
							}
							if desired == "omit-label-update" {
								wanted["node_config"].([]any)[0].(map[string]any)["labels"] = map[string]any{"audit": "yes"}
							}
							if err := conf.TerraformConfigurationInjector(nil, wanted); err != nil {
								t.Fatal(err)
							}
							state.RawConfig, err = schema.JSONMapToStateValue(wanted, r.CoreConfigSchema())
							if err != nil {
								t.Fatal(err)
							}
							state.RawPlan, err = schema.JSONMapToStateValue(observed, r.CoreConfigSchema())
							if err != nil {
								t.Fatal(err)
							}
							cfg := terraform.NewResourceConfigRaw(wanted)
							diff, err := schema.InternalMap(r.Schema).Diff(context.Background(), state, cfg, r.CustomizeDiff, &transport.Config{Project: "test-project"}, false)
							if err != nil {
								t.Fatal(err)
							}
							if hook == "atlassian2" {
								countOnlyHook(diff)
							} else {
								diff, err = conf.TerraformCustomDiff(diff, state, cfg)
								if err != nil {
									t.Fatal(err)
								}
							}
							planned, err := schema.InternalMap(r.Schema).Data(state, diff)
							if err != nil {
								t.Fatal(err)
							}
							expectedChange := desired == "true" && !quota
							if planned.HasChange("node_config.0.kubelet_config") != expectedChange {
								t.Errorf("kubelet change got %v want %v", planned.HasChange("node_config.0.kubelet_config"), expectedChange)
							}
							expected := rd.Get("node_config.0.kubelet_config").([]any)
							if desired == "true" {
								expected[0].(map[string]any)["cpu_cfs_quota"] = true
							}
							if d := cmp.Diff(expected, planned.Get("node_config.0.kubelet_config")); d != "" {
								t.Errorf("kubelet settings changed unexpectedly: %s", d)
							}
							if planned.HasChange("node_config.0.labels") != (desired == "omit-label-update") {
								t.Error("unrelated labels diff changed")
							}
							if diff != nil && (diff.Destroy || diff.RequiresNew()) {
								t.Error("requires replacement")
							}
						})
					}
				}
			}
		}
	}
}

// Cluster root node_config is a separate schema instance from the shared
// standalone NodePool schema. Verify this independent generic case at runtime.
func TestRuntimeClusterRootKubeletPreservation(t *testing.T) {
	sdk := gp.Provider()
	fw := fp.New(sdk)
	cluster, err := pc.GetProvider(context.Background(), sdk, fw, false)
	if err != nil {
		t.Fatal(err)
	}
	namespaced, err := pc.GetNamespacedProvider(context.Background(), sdk, fw, false)
	if err != nil {
		t.Fatal(err)
	}
	for scope, p := range map[string]*uc.Provider{"cluster": cluster, "namespaced": namespaced} {
		conf := p.Resources["google_container_cluster"]
		r := conf.TerraformResource
		for _, quota := range []bool{false, true} {
			for _, settings := range []string{"defaults", "topology", "custom"} {
				for _, labels := range []bool{false, true} {
					t.Run(fmt.Sprintf("%s/quota=%t/%s/labels=%t", scope, quota, settings, labels), func(t *testing.T) {
						kubelet := map[string]any{"cpu_cfs_quota": quota, "insecure_kubelet_readonly_port_enabled": "FALSE", "max_parallel_image_pulls": 2}
						if settings == "topology" {
							kubelet["topology_manager"] = []any{map[string]any{"policy": "best-effort", "scope": "container"}}
						}
						if settings == "custom" {
							kubelet["cpu_manager_policy"] = "static"
							kubelet["cpu_cfs_quota_period"] = "100ms"
						}
						observed := params(kubelet)
						delete(observed, "cluster")
						rd := schema.TestResourceDataRaw(t, r.Schema, observed)
						rd.SetId("projects/test-project/locations/us-central1/clusters/test")
						state := rd.State()
						wanted := params(nil)
						delete(wanted, "cluster")
						if labels {
							wanted["node_config"].([]any)[0].(map[string]any)["labels"] = map[string]any{"audit": "yes"}
						}
						if err := conf.TerraformConfigurationInjector(nil, wanted); err != nil {
							t.Fatal(err)
						}
						state.RawConfig, err = schema.JSONMapToStateValue(wanted, r.CoreConfigSchema())
						if err != nil {
							t.Fatal(err)
						}
						state.RawPlan, err = schema.JSONMapToStateValue(observed, r.CoreConfigSchema())
						if err != nil {
							t.Fatal(err)
						}
						cfg := terraform.NewResourceConfigRaw(wanted)
						diff, err := schema.InternalMap(r.Schema).Diff(context.Background(), state, cfg, r.CustomizeDiff, &transport.Config{Project: "test-project"}, false)
						if err != nil {
							t.Fatal(err)
						}
						diff, err = conf.TerraformCustomDiff(diff, state, cfg)
						if err != nil {
							t.Fatal(err)
						}
						planned, err := schema.InternalMap(r.Schema).Data(state, diff)
						if err != nil {
							t.Fatal(err)
						}
						if planned.HasChange("node_config.0.kubelet_config") {
							t.Error("omission must not enter the kubelet update handler")
						}
						if d := cmp.Diff(rd.Get("node_config.0.kubelet_config"), planned.Get("node_config.0.kubelet_config")); d != "" {
							t.Errorf("kubelet settings changed unexpectedly: %s", d)
						}
						if planned.HasChange("node_config.0.labels") != labels {
							t.Error("unrelated labels diff changed")
						}
					})
				}
			}
		}
	}
}
func params(k map[string]any) map[string]any {
	n := map[string]any{"machine_type": "e2-medium"}
	if k != nil {
		n["kubelet_config"] = []any{k}
	}
	return map[string]any{"name": "worker", "project": "test-project", "location": "us-central1", "cluster": "test", "deletion_policy": "DELETE", "node_config": []any{n}}
}

// Exact atlassian-2 effects relevant to these cases: phantom blocks plus
// initial count and only the omitted kubelet parent count, never its children.
func countOnlyHook(d *terraform.InstanceDiff) {
	if d == nil || d.Destroy {
		return
	}
	for _, k := range []string{"placement_policy.#", "autoscaling.#", "queued_provisioning.#", "node_drain_config.#"} {
		if a, ok := d.Attributes[k]; ok && a.Old == "" && a.New == "" {
			delete(d.Attributes, k)
		}
	}
	if a, ok := d.Attributes["initial_node_count"]; ok && a.Old != "" {
		delete(d.Attributes, "initial_node_count")
	}
	for k, a := range d.Attributes {
		if strings.HasSuffix(k, "kubelet_config.#") && strings.Contains(k, "node_config.") && (a.New == "0" || a.New == "") {
			delete(d.Attributes, k)
		}
	}
}
