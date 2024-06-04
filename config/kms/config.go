// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package kms

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("google_kms_crypto_key_iam_member", func(r *config.Resource) {
		r.References["crypto_key_id"] = config.Reference{
			TerraformName: "google_kms_crypto_key",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_kms_key_ring_iam_member", func(r *config.Resource) {
		r.References["key_ring_id"] = config.Reference{
			TerraformName: "google_kms_key_ring",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_kms_key_ring_import_job", func(r *config.Resource) {
		r.References["key_ring"] = config.Reference{
			TerraformName: "google_kms_key_ring",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		config.MoveToStatus(r.TerraformResource, "public_key")
	})

	p.AddResourceConfigurator("google_kms_secret_ciphertext", func(r *config.Resource) {
		r.TerraformResource.Schema["plaintext"].Sensitive = false
	})
}
