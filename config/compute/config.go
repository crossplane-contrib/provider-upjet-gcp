package compute

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"

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
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_compute_backend_service", func(r *config.Resource) {
		r.References["health_checks"] = config.Reference{
			Type:      "HealthCheck",
			Extractor: common.PathSelfLinkExtractor,
		}
		r.References["backend.group"] = config.Reference{
			Type:      "InstanceGroupManager",
			Extractor: PathInstanceGroupExtractor,
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
			Type: "Disk",
		}
	})

	p.AddResourceConfigurator("google_compute_subnetwork", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type: "Network",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_address", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type: "Network",
		}
		r.References["subnetwork"] = config.Reference{
			Type: "Subnetwork",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_firewall", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type:      "Network",
			Extractor: common.PathSelfLinkExtractor,
		}
	})

	p.AddResourceConfigurator("google_compute_router", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type:      "Network",
			Extractor: common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_router_nat", func(r *config.Resource) {
		r.References["router"] = config.Reference{
			Type: "Router",
		}
		r.References["subnetwork.name"] = config.Reference{
			Type: "Subnetwork",
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
			Type: "Network",
		}
		r.References["network_interface.subnetwork"] = config.Reference{
			Type: "Subnetwork",
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
			Type: "Network",
		}
		r.References["network_interface.subnetwork"] = config.Reference{
			Type: "Subnetwork",
		}
		r.References["boot_disk.initialize_params.image"] = config.Reference{
			Type: "Image",
		}
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_compute_instance_iam_member", func(r *config.Resource) {
		r.References["instance_name"] = config.Reference{
			Type: "Instance",
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
			Type: "Network",
		}
		r.References["network_interface.subnetwork"] = config.Reference{
			Type: "Subnetwork",
		}
	})

	p.AddResourceConfigurator("google_compute_instance_group", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type:      "Network",
			Extractor: common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_compute_global_address", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type:      "Network",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_compute_ha_vpn_gateway", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type:      "Network",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_instance_from_template", func(r *config.Resource) {
		r.References["source_instance_template"] = config.Reference{
			Type:      "InstanceTemplate",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_compute_instance_group_manager", func(r *config.Resource) {
		r.References["auto_healing_policies.health_check"] = config.Reference{
			Type:      "HealthCheck",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["version.instance_template"] = config.Reference{
			Type:      "InstanceTemplate",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["target_pools"] = config.Reference{
			Type:      "TargetPool",
			Extractor: common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_compute_interconnect_attachment", func(r *config.Resource) {
		r.References["router"] = config.Reference{
			Type:      "Router",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_network_endpoint_group", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type:      "Network",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["subnetwork"] = config.Reference{
			Type:      "Subnetwork",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_compute_resource_policy", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_target_pool", func(r *config.Resource) {
		// Note(donovanmuller): Only legacy google_compute_http_health_check is supported
		// see https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_target_pool#health_checks
		r.References["health_checks"] = config.Reference{
			Type: "HTTPHealthCheck",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
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
		// Note(donovanmuller): See https://github.com/upbound/upjet/issues/95
		// BackendService is also a valid reference Type
		r.References["backend_service"] = config.Reference{
			Type:      "RegionBackendService",
			Extractor: common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_backend_service", func(r *config.Resource) {
		r.References["health_checks"] = config.Reference{
			Type:      "RegionHealthCheck",
			Extractor: common.PathSelfLinkExtractor,
		}
		r.References["backend.group"] = config.Reference{
			Type:      "RegionInstanceGroupManager",
			Extractor: PathInstanceGroupExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_instance_group_manager", func(r *config.Resource) {
		r.References["auto_healing_policies.health_check"] = config.Reference{
			Type:      "HealthCheck",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["version.instance_template"] = config.Reference{
			Type:      "InstanceTemplate",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["target_pools"] = config.Reference{
			Type:      "TargetPool",
			Extractor: common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_target_http_proxy", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_url_map", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
	p.AddResourceConfigurator("google_compute_region_autoscaler", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_disk", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_disk_iam_member", func(r *config.Resource) {
		r.References["name"] = config.Reference{
			Type: "RegionDisk",
		}
	})

	p.AddResourceConfigurator("google_compute_region_disk_resource_policy_attachment", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_health_check", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_per_instance_config", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_ssl_certificate", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_region_target_https_proxy", func(r *config.Resource) {
		r.References["ssl_certificates"] = config.Reference{
			Type: "RegionSSLCertificate",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_reservation", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_compute_firewall_policy_association", func(r *config.Resource) {
		r.References["ssl_certificates"] = config.Reference{
			Type: "RegionSSLCertificate",
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
			Type: "SSLCertificate",
		}
	})

	p.AddResourceConfigurator("google_compute_image_iam_member", func(r *config.Resource) {
		r.References["image"] = config.Reference{
			Type: "Image",
		}
	})

	p.AddResourceConfigurator("google_compute_vpn_tunnel", func(r *config.Resource) {
		r.References["peer_external_gateway"] = config.Reference{
			Type: "ExternalVPNGateway",
		}
		r.References["router"] = config.Reference{
			Type: "Router",
		}
		r.References["vpn_gateway"] = config.Reference{
			Type: "HaVPNGateway",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_target_https_proxy", func(r *config.Resource) {
		r.References["ssl_certificates"] = config.Reference{
			Type: "SSLCertificate",
		}
	})

	p.AddResourceConfigurator("google_compute_service_attachment", func(r *config.Resource) {
		r.References["nat_subnets"] = config.Reference{
			Type: "Subnetwork",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_compute_route", func(r *config.Resource) {
		r.References["next_hop_vpn_tunnel"] = config.Reference{
			Type: "VPNTunnel",
		}
	})

	p.AddResourceConfigurator("google_compute_router_interface", func(r *config.Resource) {
		r.References["router"] = config.Reference{
			Type: "Router",
		}
		r.References["vpn_tunnel"] = config.Reference{
			Type: "VPNTunnel",
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
			Type: "Subnetwork",
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
