// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package identityplatform

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_identity_platform_default_supported_idp_config", func(r *config.Resource) {
		r.TerraformResource.Schema["client_id"].Sensitive = true
		r.TerraformResource.Schema["client_secret"].Sensitive = true
	})

	p.AddResourceConfigurator("google_identity_platform_inbound_saml_config", func(r *config.Resource) {
		r.TerraformResource.Schema["idp_config"].Elem.(*schema.Resource).
			Schema["idp_certificates"].Elem.(*schema.Resource).
			Schema["x509_certificate"].Sensitive = true
	})

	p.AddResourceConfigurator("google_identity_platform_oauth_idp_config", func(r *config.Resource) {
		r.TerraformResource.Schema["client_id"].Sensitive = true
		r.TerraformResource.Schema["client_secret"].Sensitive = true
	})

	p.AddResourceConfigurator("google_identity_platform_tenant_default_supported_idp_config", func(r *config.Resource) {
		r.TerraformResource.Schema["client_id"].Sensitive = true
		r.TerraformResource.Schema["client_secret"].Sensitive = true
	})

	p.AddResourceConfigurator("google_identity_platform_tenant_inbound_saml_config", func(r *config.Resource) {
		r.TerraformResource.Schema["idp_config"].Elem.(*schema.Resource).
			Schema["idp_certificates"].Elem.(*schema.Resource).
			Schema["x509_certificate"].Sensitive = true
	})

	p.AddResourceConfigurator("google_identity_platform_tenant_oauth_idp_config", func(r *config.Resource) {
		r.TerraformResource.Schema["client_id"].Sensitive = true
		r.TerraformResource.Schema["client_secret"].Sensitive = true
	})
}
