package dataflow

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_dataflow_job", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of "labels", since is a map
		// where elements configured as nil configured as nil, but needs to be
		// String:
		r.TerraformResource.Schema["labels"].Elem = schema.TypeString
		r.TerraformResource.Schema["parameters"].Elem = schema.TypeString
		r.TerraformResource.Schema["transform_name_mapping"].Elem = schema.TypeString

		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil {
				delete(diff.Attributes, "additional_experiments.#")
			}
			return diff, nil
		}
	})
}
