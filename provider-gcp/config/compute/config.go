package compute

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint: gocyclo
	// Note(turkenh): We ignore gocyclo in this function since it configures
	//  all resources separately and no complex logic here.

	p.AddResourceConfigurator("google_compute_managed_ssl_certificate", func(r *config.Resource) {
		r.Kind = "ManagedSSLCertificate"
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_compute_subnetwork", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type: "Network",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_compute_address", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			Type: "Network",
		}
		r.References["subnetwork"] = config.Reference{
			Type: "Subnetwork",
		}
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
	})

	p.AddResourceConfigurator("google_compute_router_nat", func(r *config.Resource) {
		r.References["router"] = config.Reference{
			Type: "Router",
		}
		r.References["subnetwork.name"] = config.Reference{
			Type: "Subnetwork",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_compute_instance_template", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of
		// "boot_disk.initialize_params.labels", since it is a map where
		// elements configured as nil, defaulting to map[string]string:
		r.TerraformResource.Schema["metadata"].Elem = schema.TypeString
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

		r.UseAsync = true
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
			Extractor: common.ExtractResourceIDFuncPath,
		}
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
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_compute_interconnect_attachment", func(r *config.Resource) {
		r.References["router"] = config.Reference{
			Type:      "Router",
			Extractor: common.ExtractResourceIDFuncPath,
		}
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
	})
}
