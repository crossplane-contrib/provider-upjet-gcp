package vpcaccess

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_vpc_access_connector", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
}
