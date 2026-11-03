// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package beyondcorp

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-gcp/v3/config/cluster/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	// google_beyondcorp_app_connection and google_beyondcorp_app_connector were
	// deprecated upstream in v7.46.0 and their doc examples were removed, so the
	// references that used to be injected from those examples are pinned here to
	// keep the CRD APIs stable.
	p.AddResourceConfigurator("google_beyondcorp_app_connection", func(r *config.Resource) {
		r.MarkAsRequired("region")
		r.References["connectors"] = config.Reference{
			TerraformName: "google_beyondcorp_app_connector",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["gateway.app_gateway"] = config.Reference{
			TerraformName: "google_beyondcorp_app_gateway",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_beyondcorp_app_connector", func(r *config.Resource) {
		r.MarkAsRequired("region")
		r.References["principal_info.service_account.email"] = config.Reference{
			TerraformName: "google_service_account",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("email",true)`,
		}
	})
	p.AddResourceConfigurator("google_beyondcorp_app_gateway", func(r *config.Resource) {
		r.MarkAsRequired("region")
	})
}
