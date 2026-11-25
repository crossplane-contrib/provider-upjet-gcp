// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package networkservices

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-gcp/v3/config/cluster/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_network_services_lb_traffic_extension", func(r *config.Resource) {
		r.References["forwarding_rules"] = config.Reference{
			TerraformName: "google_compute_forwarding_rule",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["extension_chains.extensions.service"] = config.Reference{
			TerraformName: "google_compute_region_backend_service",
			Extractor:     common.PathSelfLinkExtractor,
		}
	})
}
