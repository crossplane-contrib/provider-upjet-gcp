// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package healthcare

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_healthcare_dataset_iam_member", func(r *config.Resource) {
		r.References["dataset_id"] = config.Reference{
			TerraformName: "google_healthcare_dataset",
		}
	})
}
