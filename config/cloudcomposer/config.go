package composer

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_composer_environment", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project",
		}
		r.References["node_config.network"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/compute/v1beta1.Network",
		}
		r.References["node_config.subnetwork"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork",
		}
		r.References["node_config.service_account"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount",
		}
		r.References["private_environment_config.cloud_composer_connection_subnetwork"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})
}
