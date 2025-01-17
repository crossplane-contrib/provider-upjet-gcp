// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dataplex

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_dataplex_aspect_type", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "google_project",
		}
		r.MarkAsRequired("location")
	})

	p.AddResourceConfigurator("google_dataplex_lake_iam_policy", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "google_project",
		}
		r.MarkAsRequired("location")
	})
}
