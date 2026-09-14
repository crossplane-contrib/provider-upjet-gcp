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
		// password_wo is a write-only Terraform field with no persisted
		// state, which upjet can't represent. password_wo_version stays
		// in the schema (only its now-dangling RequiredWith is cleared)
		// because the upstream provider's Read unconditionally calls
		// d.Set("password_wo_version", ...): deleting it breaks every
		// read with "Invalid address to set: password_wo_version".
		delete(r.TerraformResource.Schema, "password_wo")
		r.TerraformResource.Schema["password_wo_version"].RequiredWith = nil

		// GCP auto-manages implicit roles for ALLOYDB_IAM_USER (e.g.
		// alloydbiamuser). Late-initializing them into spec.forProvider
		// makes every reconcile resubmit them as an explicit update,
		// which the API rejects: "cannot revoke IAM roles ...; user
		// type cannot be changed".
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"database_roles"},
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
