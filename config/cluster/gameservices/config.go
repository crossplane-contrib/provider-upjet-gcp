// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package gameservices

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_game_services_game_server_cluster", func(r *config.Resource) {
		r.References["connection_info.gke_cluster_reference.cluster"] = config.Reference{
			TerraformName: "google_container_cluster",
		}
	})
}
