package beyondcorp

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_beyondcorp_app_connection", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
	p.AddResourceConfigurator("google_beyondcorp_app_connector", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
	p.AddResourceConfigurator("google_beyondcorp_app_gateway", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
}
