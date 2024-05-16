// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package kms

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"

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

	p.AddResourceConfigurator("google_kms_crypto_key_version", func(r *config.Resource) {
		attestation := r.TerraformResource.Schema["attestation"].Elem.(*schema.Resource)
		attestation.Schema["cert_chains"].MaxItems = 1
		attestation.Schema["external_protection_level_options"].MaxItems = 1

		r.Version = "v1beta2"
		r.PreviousVersions = []string{common.VersionV1Beta1}
		// we would like to set the storage version to v1beta1 to facilitate
		// downgrades.
		r.SetCRDStorageVersion("v1beta1")
		r.ControllerReconcileVersion = "v1beta1"
		r.Conversions = []conversion.Conversion{
			conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, conversion.DefaultPathPrefixes(), r.CRDListConversionPaths()...),
			conversion.NewSingletonListConversion("v1beta1", "v1beta2", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToEmbeddedObject),
			conversion.NewSingletonListConversion("v1beta2", "v1beta1", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToSingletonList)}

		r.AddSingletonListConversion("attestation[*].cert_chains", "attestation[*].certChains")
		r.AddSingletonListConversion("attestation[*].external_protection_level_options", "attestation[*].externalProtectionLevelOptions")
	})
}
