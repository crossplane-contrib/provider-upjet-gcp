// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package vertexai

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_vertex_ai_index", func(r *config.Resource) {
		r.MarkAsRequired("region")
	})
	p.AddResourceConfigurator("google_vertex_ai_tensorboard", func(r *config.Resource) {
		r.MarkAsRequired("region")
	})
	p.AddResourceConfigurator("google_vertex_ai_reasoning_engine", func(r *config.Resource) {
		r.MarkAsRequired("region")
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"spec.deployment_spec.env",
			},
		}
		r.References["spec.service_account"] = config.Reference{
			TerraformName: "google_service_account",
		}
		r.References["spec.deployment_spec.secret_env.secret_ref.secret"] = config.Reference{
			TerraformName: "google_secret_manager_secret",
		}
		r.References["encryption_spec.kms_key_name"] = config.Reference{
			TerraformName: "google_kms_crypto_key",
		}
		// spec.deployment_spec.psc_interface_config.network_attachment is left a
		// plain string: google_compute_network_attachment is not a managed
		// resource in this provider, so no reference target exists.
	})
	p.AddResourceConfigurator("google_vertex_ai_reasoning_engine_iam_binding", func(r *config.Resource) {
		r.References["reasoning_engine"] = config.Reference{
			TerraformName: "google_vertex_ai_reasoning_engine",
		}
		r.MarkAsRequired("region")
	})
	p.AddResourceConfigurator("google_vertex_ai_reasoning_engine_iam_member", func(r *config.Resource) {
		r.References["reasoning_engine"] = config.Reference{
			TerraformName: "google_vertex_ai_reasoning_engine",
		}
		r.MarkAsRequired("region")
	})
	p.AddResourceConfigurator("google_vertex_ai_reasoning_engine_iam_policy", func(r *config.Resource) {
		r.References["reasoning_engine"] = config.Reference{
			TerraformName: "google_vertex_ai_reasoning_engine",
		}
		r.MarkAsRequired("region")
	})
}
