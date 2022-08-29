package identityplatform

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_identity_platform_default_supported_idp_config", func(r *config.Resource) {
		r.TerraformResource.Schema["client_id"].Sensitive = true
		r.TerraformResource.Schema["client_secret"].Sensitive = true
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_identity_platform_inbound_saml_config", func(r *config.Resource) {
		r.TerraformResource.Schema["idp_config"].Elem.(*schema.Resource).
			Schema["idp_certificates"].Elem.(*schema.Resource).
			Schema["x509_certificate"].Sensitive = true
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_identity_platform_oauth_idp_config", func(r *config.Resource) {
		r.TerraformResource.Schema["client_id"].Sensitive = true
		r.TerraformResource.Schema["client_secret"].Sensitive = true
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_identity_platform_tenant", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_identity_platform_tenant_default_supported_idp_config", func(r *config.Resource) {
		r.TerraformResource.Schema["client_id"].Sensitive = true
		r.TerraformResource.Schema["client_secret"].Sensitive = true
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_identity_platform_tenant_inbound_saml_config", func(r *config.Resource) {
		r.TerraformResource.Schema["idp_config"].Elem.(*schema.Resource).
			Schema["idp_certificates"].Elem.(*schema.Resource).
			Schema["x509_certificate"].Sensitive = true
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_identity_platform_tenant_oauth_idp_config", func(r *config.Resource) {
		r.TerraformResource.Schema["client_id"].Sensitive = true
		r.TerraformResource.Schema["client_secret"].Sensitive = true
		r.UseAsync = true
	})
}
