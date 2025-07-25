// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package accesscontextmanager

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_access_context_manager_service_perimeter", func(r *config.Resource) {
		r.References["status.access_levels"] = config.Reference{
			TerraformName: "google_access_context_manager_access_level",
		}
		r.References["spec.access_levels"] = config.Reference{
			TerraformName: "google_access_context_manager_access_level",
		}
	})
	p.AddResourceConfigurator("google_access_context_manager_service_perimeter_resource", func(r *config.Resource) {
		if s, ok := r.TerraformResource.Schema["resource"]; ok {
			s.Optional = false
			s.Computed = false
		}
		r.MetaResource.Description = "Allows configuring a single GCP resource that should be inside the 'status' block of a service perimeter."
	})
}
