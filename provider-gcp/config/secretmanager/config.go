package secretmanager

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_secret_manager_secret", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "replication")
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_secret_manager_secret_iam_member", func(r *config.Resource) {
		r.References["secret_id"] = config.Reference{
			Type: "Secret",
		}
	})

	p.AddResourceConfigurator("google_secret_manager_secret_version", func(r *config.Resource) {
		r.UseAsync = true
	})
}
