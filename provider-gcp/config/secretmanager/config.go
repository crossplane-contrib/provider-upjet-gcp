package secretmanager

import (
	"github.com/upbound/official-providers/provider-gcp/config/common"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_secret_manager_secret", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "replication")
	})

	p.AddResourceConfigurator("google_secret_manager_secret_iam_member", func(r *config.Resource) {
		r.References["secret_id"] = config.Reference{
			Type:      "Secret",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_secret_manager_secret_version", func(r *config.Resource) {
		r.References["secret"] = config.Reference{
			Type: "Secret",
		}
	})
}
