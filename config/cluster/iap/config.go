// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package iap

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_iap_web_backend_service_iam_member", func(r *config.Resource) {
		r.References["web_backend_service"] = config.Reference{
			TerraformName: "google_compute_backend_service",
		}
	})

	p.AddResourceConfigurator("google_iap_web_type_app_engine_iam_member", func(r *config.Resource) {
		r.References["app_id"] = config.Reference{
			TerraformName: "google_app_engine_application",
		}
	})
}
