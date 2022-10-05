package servicenetworking

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_service_networking_connection", func(r *config.Resource) {
		// Note(donovanmuller): Upjet does not add this reference automatically
		r.References["reserved_peering_ranges"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/compute/v1beta1.GlobalAddress",
		}
	})
}
