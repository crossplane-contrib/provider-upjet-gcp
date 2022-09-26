package firebaserules

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_firebaserules_release", func(r *config.Resource) {
		r.References["ruleset_name"] = config.Reference{
			Type: "Ruleset",
		}
	})
}
