// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package alloydb

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_alloydb_instance", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_alloydb_cluster", func(r *config.Resource) {
		delete(r.TerraformResource.
			Schema["initial_user"].Elem.(*schema.Resource).
			Schema, "password_wo")
		delete(r.TerraformResource.
			Schema["initial_user"].Elem.(*schema.Resource).
			Schema, "password_wo_version")
	})
}
