// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package compute

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-gcp/config/common"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/upbound/provider-gcp/config/compute"
)

var (
	// PathInstanceGroupExtractor is the golang path to InstanceGroupExtractor function
	// in this package.
	PathInstanceGroupExtractor = SelfPackagePath + ".InstanceGroupExtractor()"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint: gocyclo
	// Note(turkenh): We ignore gocyclo in this function since it configures
	//  all resources separately and no complex logic here.

	p.AddResourceConfigurator("google_compute_autoscaler", func(r *config.Resource) {
		r.MarkAsRequired("zone")
	})

	p.AddResourceConfigurator("google_compute_backend_service", func(r *config.Resource) {
		r.References["health_checks"] = config.Reference{
			TerraformName: "google_compute_health_check",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["backend.group"] = config.Reference{
			TerraformName: "google_compute_instance_group_manager",
			Extractor:     PathInstanceGroupExtractor,
		}
	})

	p.AddResourceConfigurator("google_compute_managed_ssl_certificate", func(r *config.Resource) {
		r.Kind = "ManagedSSLCertificate"
	})

	p.AddResourceConfigurator("google_compute_disk", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_compute_disk_iam_member", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			TerraformName: "google_compute_disk",
		}
	})

	p.AddResourceConfigurator("google_compute_shared_vpc_host_project", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "google_project",
			Extractor:     common.ExtractProjectIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_compute_shared_vpc_service_project", func(r *config.Resource) {
		r.References["host_project"] = config.Reference{
			TerraformName: "google_project",
			Extractor:     common.ExtractProjectIDFuncPath,
		}
		r.References["service_project"] = config.Reference{
			TerraformName: "google_project",
			Extractor:     common.ExtractProjectIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_compute_subnetwork", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_address", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_firewall", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.PathSelfLinkExtractor,
		}
	})

	p.AddResourceConfigurator("google_compute_router", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_router_nat", func(r *config.Resource) {
		r.References["router"] = config.Reference{
			TerraformName: "google_compute_router",
		}
		r.References["subnetwork.name"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
		}
		delete(r.References, "region")
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_instance_template", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of
		// "boot_disk.initialize_params.labels", since it is a map where
		// elements configured as nil, defaulting to map[string]string:
		r.TerraformResource.Schema["metadata"].Elem = schema.TypeString
		r.References["network_interface.network"] = config.Reference{
			TerraformName: "google_compute_network",
		}
		r.References["network_interface.subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
		}

		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Destroy {
				return diff, nil
			}
			if cicDiff, ok := diff.Attributes["confidential_instance_config.#"]; ok && cicDiff.Old == "" && cicDiff.New == "" {
				delete(diff.Attributes, "confidential_instance_config.#")
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("google_compute_instance", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of
		// "boot_disk.initialize_params", since "labels" key here is a map where
		// elements configured as nil, defaulting to map[string]string:
		// https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance#nested_initialize_params
		r.TerraformResource.
			Schema["boot_disk"].Elem.(*schema.Resource).
			Schema["initialize_params"].Elem.(*schema.Resource).
			Schema["labels"].Elem = schema.TypeString

		r.References["network_interface.network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["network_interface.subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["boot_disk.initialize_params.image"] = config.Reference{
			TerraformName: "google_compute_image",
		}
		r.MarkAsRequired("zone")
	})

	p.AddResourceConfigurator("google_compute_instance_iam_member", func(r *config.Resource) {
		r.References["instance_name"] = config.Reference{
			TerraformName: "google_compute_instance",
		}
	})

	p.AddResourceConfigurator("google_compute_instance_from_template", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of
		// "boot_disk.initialize_params.labels", since it is a map where
		// elements configured as nil, defaulting to map[string]string:
		// https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance#nested_initialize_params
		r.TerraformResource.
			Schema["boot_disk"].Elem.(*schema.Resource).
			Schema["initialize_params"].Elem.(*schema.Resource).
			Schema["labels"].Elem = schema.TypeString
		r.TerraformResource.Schema["metadata"].Elem = schema.TypeString

		r.References["network_interface.network"] = config.Reference{
			TerraformName: "google_compute_network",
		}
		r.References["network_interface.subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Destroy {
				return diff, nil
			}
			if paramsDiff, ok := diff.Attributes["params.#"]; ok && paramsDiff.Old == "" && paramsDiff.New == "" {
				delete(diff.Attributes, "params.#")
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("google_compute_instance_group", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.MarkAsRequired("zone")
	})

	p.AddResourceConfigurator("google_compute_global_address", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_compute_ha_vpn_gateway", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_instance_from_template", func(r *config.Resource) {
		r.References["source_instance_template"] = config.Reference{
			TerraformName: "google_compute_instance_template",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_compute_instance_group_manager", func(r *config.Resource) {
		r.References["auto_healing_policies.health_check"] = config.Reference{
			TerraformName: "google_compute_health_check",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["version.instance_template"] = config.Reference{
			TerraformName: "google_compute_instance_template",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["target_pools"] = config.Reference{
			TerraformName: "google_compute_target_pool",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.MarkAsRequired("zone")
	})

	p.AddResourceConfigurator("google_compute_interconnect_attachment", func(r *config.Resource) {
		r.References["router"] = config.Reference{
			TerraformName: "google_compute_router",
			Extractor:     common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_network_endpoint_group", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.MarkAsRequired("zone")
	})

	p.AddResourceConfigurator("google_compute_resource_policy", func(r *config.Resource) {
		r.MarkAsRequired("region")
	})

	p.AddResourceConfigurator("google_compute_target_pool", func(r *config.Resource) {
		// Note(donovanmuller): Only legacy google_compute_http_health_check is supported
		// see https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_target_pool#health_checks
		r.References["health_checks"] = config.Reference{
			TerraformName: "google_compute_http_health_check",
		}
		r.MarkAsRequired("region")
	})

	p.AddResourceConfigurator("google_compute_network_endpoint", func(r *config.Resource) {
		delete(r.References, "port")
	})

	p.AddResourceConfigurator("google_compute_node_group", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_compute_node_template", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_packet_mirroring", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_forwarding_rule", func(r *config.Resource) {
		// Note(donovanmuller): See https://github.com/crossplane/upjet/issues/95
		// BackendService is also a valid reference Type
		r.References["backend_service"] = config.Reference{
			TerraformName: "google_compute_region_backend_service",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["ip_address"] = config.Reference{
			TerraformName: "google_compute_address",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["target"] = config.Reference{
			TerraformName: "google_compute_region_target_http_proxy",
			Extractor:     common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_backend_service", func(r *config.Resource) {
		r.References["health_checks"] = config.Reference{
			TerraformName: "google_compute_region_health_check",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["backend.group"] = config.Reference{
			TerraformName: "google_compute_region_instance_group_manager",
			Extractor:     PathInstanceGroupExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_instance_group_manager", func(r *config.Resource) {
		r.References["auto_healing_policies.health_check"] = config.Reference{
			TerraformName: "google_compute_health_check",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["version.instance_template"] = config.Reference{
			TerraformName: "google_compute_instance_template",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["target_pools"] = config.Reference{
			TerraformName: "google_compute_target_pool",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.MarkAsRequired("region")
	})

	p.AddResourceConfigurator("google_compute_region_target_http_proxy", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")

		r.References["url_map"] = config.Reference{
			TerraformName: "google_compute_region_url_map",
			Extractor:     common.PathSelfLinkExtractor,
		}

	})

	p.AddResourceConfigurator("google_compute_region_target_tcp_proxy", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")

		r.References["backend_service"] = config.Reference{
			TerraformName: "google_compute_region_backend_service",
			Extractor:     common.PathSelfLinkExtractor,
		}

	})

	p.AddResourceConfigurator("google_compute_region_url_map", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
	p.AddResourceConfigurator("google_compute_region_autoscaler", func(r *config.Resource) {
		r.MarkAsRequired("region")
	})

	p.AddResourceConfigurator("google_compute_region_disk", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_disk_iam_member", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			TerraformName: "google_compute_region_disk",
		}
	})

	p.AddResourceConfigurator("google_compute_region_disk_resource_policy_attachment", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_health_check", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_ssl_certificate", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_target_https_proxy", func(r *config.Resource) {
		r.References["ssl_certificates"] = config.Reference{
			TerraformName: "google_compute_region_ssl_certificate",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_reservation", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "zone")
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil {
				delete(diff.Attributes, "share_settings.#")
			}
			return diff, nil
		}
	})

	p.AddResourceConfigurator("google_compute_firewall_policy_association", func(r *config.Resource) {
		r.References["ssl_certificates"] = config.Reference{
			TerraformName: "google_compute_region_ssl_certificate",
		}
	})

	p.AddResourceConfigurator("google_compute_instance_group_named_port", func(r *config.Resource) {
		// Note(donovanmuller): see https://github.com/upbound/official-providers/issues/597
		// r.References["group"] = config.Reference{
		// 	Type:      "github.com/upbound/provider-gcp/apis/container/v1beta1.Cluster",
		// 	Extractor: PathClusterInstanceGroupExtractor,
		// }
	})

	p.AddResourceConfigurator("google_compute_target_ssl_proxy", func(r *config.Resource) {
		r.References["ssl_certificates"] = config.Reference{
			TerraformName: "google_compute_ssl_certificate",
		}
	})

	p.AddResourceConfigurator("google_compute_image_iam_member", func(r *config.Resource) {
		r.References["image"] = config.Reference{
			TerraformName: "google_compute_image",
		}
	})

	p.AddResourceConfigurator("google_compute_vpn_tunnel", func(r *config.Resource) {
		r.References["peer_external_gateway"] = config.Reference{
			TerraformName: "google_compute_external_vpn_gateway",
		}
		r.References["router"] = config.Reference{
			TerraformName: "google_compute_router",
		}
		r.References["vpn_gateway"] = config.Reference{
			TerraformName: "google_compute_ha_vpn_gateway",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_target_https_proxy", func(r *config.Resource) {
		r.References["ssl_certificates"] = config.Reference{
			TerraformName: "google_compute_ssl_certificate",
		}
	})

	p.AddResourceConfigurator("google_compute_service_attachment", func(r *config.Resource) {
		r.References["nat_subnets"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
		}
		r.MarkAsRequired("region")
	})

	p.AddResourceConfigurator("google_compute_route", func(r *config.Resource) {
		r.References["next_hop_vpn_tunnel"] = config.Reference{
			TerraformName: "google_compute_vpn_tunnel",
		}
	})

	p.AddResourceConfigurator("google_compute_router_interface", func(r *config.Resource) {
		r.References["router"] = config.Reference{
			TerraformName: "google_compute_router",
		}
		r.References["vpn_tunnel"] = config.Reference{
			TerraformName: "google_compute_vpn_tunnel",
		}
	})

	p.AddResourceConfigurator("google_compute_vpn_gateway", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_target_instance", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_compute_subnetwork_iam_member", func(r *config.Resource) {
		r.References["subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})
	p.AddResourceConfigurator("google_compute_firewall_policy_rule", func(r *config.Resource) {
		r.MetaResource.ArgumentDocs["layer4_configs.ports"] = `An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port.`
	})
	p.AddResourceConfigurator("google_compute_project_metadata_item", func(r *config.Resource) {
		r.MetaResource.ArgumentDocs["id"] = "an identifier for the resource with format `{{key}}`"
	})
}

// InstanceGroupExtractor extracts Instance Group from
// "status.atProvider.instanceGroup"
func InstanceGroupExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.instanceGroup")
		if err != nil {
			return ""
		}
		return r
	}
}

// ClusterInstanceGroupExtractor extracts Instance Group from
// "status.atProvider.nodePool.instanceGroup[0].instanceGroupUrls[0]"
func ClusterInstanceGroupExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.nodePool.instanceGroup[0].instanceGroupUrls[0]")
		if err != nil {
			return ""
		}
		return r
	}
}
