// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudtasks

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_cloud_tasks_queue", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "google_project",
		}
		config.MarkAsRequired(r.TerraformResource, "location")
	})
}
