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

	p.AddResourceConfigurator("google_alloydb_user", func(r *config.Resource) {
		// password_wo/password_wo_version are Terraform write-only
		// (ephemeral) fields with no persisted state; upjet cannot
		// represent them.
		delete(r.TerraformResource.Schema, "password_wo")
		delete(r.TerraformResource.Schema, "password_wo_version")

		r.References["cluster"] = config.Reference{
			TerraformName: "google_alloydb_cluster",
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["password"].(string); ok {
				conn["password"] = []byte(a)
			}
			return conn, nil
		}
	})
}
