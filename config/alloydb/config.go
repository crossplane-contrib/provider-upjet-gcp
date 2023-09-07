package alloydb

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_alloydb_instance", func(r *config.Resource) {
		r.UseAsync = true
	})
}
