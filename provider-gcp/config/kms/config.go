package kms

import (
	"github.com/upbound/official-providers/provider-gcp/config/common"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_kms_key_ring_iam_member", func(r *config.Resource) {
		// The reference should be manually inferred, but this is not yet activated for GCP
		r.References["key_ring_id"] = config.Reference{
			Type:      "KeyRing",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_kms_key_ring_import_job", func(r *config.Resource) {
		// The reference should be manually inferred, but this is not yet activated for GCP
		r.References["key_ring"] = config.Reference{
			Type:      "KeyRing",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
