package gameservices

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_game_services_game_server_cluster", func(r *config.Resource) {
		r.References["connection_info.gke_cluster_reference.cluster"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/container/v1beta1.Cluster",
		}
	})
}
