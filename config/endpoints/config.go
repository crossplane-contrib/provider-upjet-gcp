package endpoints

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_endpoints_service", func(r *config.Resource) {
		r.TerraformResource.Schema["openapi_config"].Sensitive = true
	})
}
