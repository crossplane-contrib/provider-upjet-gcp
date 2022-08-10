package cloudscheduler

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_cloud_scheduler_job", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1.Project",
		}
		r.References["pubsub_target.topic_name"] = config.Reference{
			// Note(donovanmuller): What about support for 'pubsub/v1beta1.LiteTopic' reference?
			Type: "github.com/upbound/official-providers/provider-gcp/apis/pubsub/v1beta1.Topic",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
		r.UseAsync = true
	})
}
