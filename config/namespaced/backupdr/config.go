// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package backupdr

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-gcp/v3/config/namespaced/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_backup_dr_backup_vault", func(r *config.Resource) {
		r.References["encryption_config.kms_key_name"] = config.Reference{
			TerraformName: "google_kms_crypto_key",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
}
