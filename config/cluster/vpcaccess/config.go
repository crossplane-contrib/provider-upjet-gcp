// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package vpcaccess

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_vpc_access_connector", func(r *config.Resource) {
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
		}
		config.MarkAsRequired(r.TerraformResource, "region")
	})
}
