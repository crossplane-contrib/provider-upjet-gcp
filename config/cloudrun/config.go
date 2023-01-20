package cloudrun

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_cloud_run_domain_mapping", func(r *config.Resource) {
		r.References["metadata.namespace"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project",
		}
		r.References["spec.route_name"] = config.Reference{
			Type: "Service",
		}
	})
	p.AddResourceConfigurator("google_cloud_run_service", func(r *config.Resource) {
		r.References["metadata.namespace"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project",
		}
		// manually added because during the native bump from 4.22.0->4.48.0
		// crddiff detected the ref fields were removed.
		r.References["template.spec.containers.env.value_from.secret_key_ref.name"] = config.Reference{
			TerraformName: "google_secret_manager_secret",
		}
		r.References["template.spec.volumes.secret.secret_name"] = config.Reference{
			TerraformName: "google_secret_manager_secret",
		}
	})
	p.AddResourceConfigurator("google_cloud_run_service_iam_policy", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project",
		}
		r.References["service"] = config.Reference{
			Type: "Service",
		}
		config.MarkAsRequired(r.TerraformResource, "location")
	})
	p.AddResourceConfigurator("google_cloud_run_service_iam_binding", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project",
		}
		r.References["service"] = config.Reference{
			Type: "Service",
		}
		config.MarkAsRequired(r.TerraformResource, "location")
	})
	p.AddResourceConfigurator("google_cloud_run_service_iam_member", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project",
		}
		r.References["service"] = config.Reference{
			Type: "Service",
		}
		config.MarkAsRequired(r.TerraformResource, "location")
	})
}
