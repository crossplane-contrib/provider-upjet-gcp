// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package networkconnectivity

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-gcp/v3/config/cluster/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_network_connectivity_regional_endpoint", func(r *config.Resource) {
		// The API accepts an Address resource URI in the
		// projects/{project}/regions/{region}/addresses/{name} format,
		// which is the Terraform ID of google_compute_address.
		r.References["address"] = config.Reference{
			TerraformName: "google_compute_address",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
}
