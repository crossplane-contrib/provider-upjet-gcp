// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package secretmanager

import (
	"github.com/upbound/provider-gcp/config/common"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_secret_manager_secret", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "replication")
	})

	p.AddResourceConfigurator("google_secret_manager_secret_iam_member", func(r *config.Resource) {
		r.References["secret_id"] = config.Reference{
			TerraformName: "google_secret_manager_secret",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_secret_manager_secret_version", func(r *config.Resource) {
		r.References["secret"] = config.Reference{
			TerraformName: "google_secret_manager_secret",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.MetaResource.ArgumentDocs["secret_data"] = `The secret data. Must be no larger than 64KiB.`
	})
}
