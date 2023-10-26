package kms

import (
	"github.com/upbound/provider-gcp/config/common"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("google_kms_crypto_key_iam_member", func(r *config.Resource) {
		r.References["crypto_key_id"] = config.Reference{
			Type:      "CryptoKey",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_kms_key_ring_iam_member", func(r *config.Resource) {
		r.References["key_ring_id"] = config.Reference{
			Type:      "KeyRing",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_kms_key_ring_import_job", func(r *config.Resource) {
		r.References["key_ring"] = config.Reference{
			Type:      "KeyRing",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		config.MoveToStatus(r.TerraformResource, "public_key")
	})

	p.AddResourceConfigurator("google_kms_secret_ciphertext", func(r *config.Resource) {
		r.TerraformResource.Schema["plaintext"].Sensitive = false
	})

}
