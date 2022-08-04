package secretmanager

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_secret_manager_secret", func(r *config.Resource) {
		r.UseAsync = true

		if s, ok := r.TerraformResource.Schema["replication"]; ok {
			s.Optional = false
		}
	})

	p.AddResourceConfigurator("google_secret_manager_secret_version", func(r *config.Resource) {
		r.UseAsync = true

		r.References["secret"] = config.Reference{
			Type: "Secret",
		}
	})

}
