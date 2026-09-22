// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package networksecurity

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-gcp/v3/config/cluster/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_network_security_address_group", func(r *config.Resource) {
		r.MarkAsRequired("parent")
	})

	p.AddResourceConfigurator("google_network_security_backend_authentication_config", func(r *config.Resource) {
		r.References["client_certificate"] = config.Reference{
			TerraformName: "google_certificate_manager_certificate",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["trust_config"] = config.Reference{
			TerraformName: "google_certificate_manager_trust_config",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
}
