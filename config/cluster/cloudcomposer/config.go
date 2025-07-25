// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package composer

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_composer_environment", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "google_project",
		}
		r.References["node_config.network"] = config.Reference{
			TerraformName: "google_compute_network",
		}
		r.References["node_config.subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
		}
		r.References["node_config.service_account"] = config.Reference{
			TerraformName: "google_service_account",
		}
		r.References["private_environment_config.cloud_composer_connection_subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
		}

		r.MarkAsRequired("region")
	})
}
