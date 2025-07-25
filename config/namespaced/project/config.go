// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package project

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_project_default_service_accounts", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of
		// "service_accounts", since it is a map where elements
		// configured as nil, but needs to be String:
		r.TerraformResource.Schema["service_accounts"].Elem = schema.TypeString
	})
}
