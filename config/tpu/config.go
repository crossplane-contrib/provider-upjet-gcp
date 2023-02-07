package tpu

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_tpu_node", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "zone")
	})
}
