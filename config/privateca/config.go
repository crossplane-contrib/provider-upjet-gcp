// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package privateca

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_privateca_ca_pool_iam_member", func(r *config.Resource) {
		r.References["ca_pool"] = config.Reference{
			TerraformName: "google_privateca_ca_pool",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_privateca_certificate_template_iam_member", func(r *config.Resource) {
		r.References["certificate_template"] = config.Reference{
			TerraformName: "google_privateca_certificate_template",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_privateca_certificate_authority", func(r *config.Resource) {
		r.References["pool"] = config.Reference{
			TerraformName: "google_privateca_ca_pool",
		}
	})

	p.AddResourceConfigurator("google_privateca_certificate", func(r *config.Resource) {
		r.References["pool"] = config.Reference{
			TerraformName: "google_privateca_ca_pool",
		}
		r.TerraformResource.Schema["config"].Elem.(*schema.Resource).
			Schema["public_key"].Elem.(*schema.Resource).
			Schema["key"].Sensitive = true
	})
}
