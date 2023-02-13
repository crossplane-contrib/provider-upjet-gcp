package containerattached

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("google_container_attached_cluster", func(r *config.Resource) {
		r.Kind = "Cluster"
	})
}
