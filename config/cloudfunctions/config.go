// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudfunctions

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_cloudfunctions_function", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of
		// "build_environment_variables", since it is a map where elements
		// configured as nil, but needs to be String:
		r.TerraformResource.
			Schema["labels"].Elem = schema.TypeString
		r.TerraformResource.
			Schema["build_environment_variables"].Elem = schema.TypeString
		r.TerraformResource.
			Schema["environment_variables"].Elem = schema.TypeString
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_cloudfunctions_function_iam_binding", func(r *config.Resource) {
		r.References["cloud_function"] = config.Reference{
			TerraformName: "google_cloudfunctions_function",
		}
	})

	p.AddResourceConfigurator("google_cloudfunctions_function_iam_member", func(r *config.Resource) {
		r.References["cloud_function"] = config.Reference{
			TerraformName: "google_cloudfunctions_function",
		}
	})
}
