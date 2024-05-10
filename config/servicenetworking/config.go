// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package servicenetworking

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_service_networking_connection", func(r *config.Resource) {
		// Note(donovanmuller): Upjet does not add this reference automatically
		r.References["reserved_peering_ranges"] = config.Reference{
			TerraformName: "google_compute_global_address",
		}
	})
}
