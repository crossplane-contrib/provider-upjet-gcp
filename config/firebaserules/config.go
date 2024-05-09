// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package firebaserules

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_firebaserules_release", func(r *config.Resource) {
		r.References["ruleset_name"] = config.Reference{
			TerraformName: "google_firebaserules_ruleset",
		}
	})
}
