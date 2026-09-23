// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package certificatemanager

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_certificate_manager_certificate", func(r *config.Resource) {
		// pem_private_key_wo is a write-only Terraform field with no persisted
		// state, which upjet can't represent; pemPrivateKeySecretRef covers the
		// same input. The upstream provider only reads it through the guarded
		// raw-config helper in the expand path and never sets either field on
		// Read, so both can be dropped from the schema.
		selfManaged := r.TerraformResource.Schema["self_managed"].Elem.(*schema.Resource).Schema
		delete(selfManaged, "pem_private_key_wo")
		delete(selfManaged, "pem_private_key_wo_version")
	})
}
