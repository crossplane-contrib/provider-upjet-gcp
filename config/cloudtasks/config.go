package cloudtasks

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_cloud_tasks_queue", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1.Project",
		}
		config.MarkAsRequired(r.TerraformResource, "location")
	})
}
