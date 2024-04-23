// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package networkconnectivity

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_network_connectivity_service_connection_policy", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "name")
		config.MarkAsRequired(r.TerraformResource, "service_class")
		config.MarkAsRequired(r.TerraformResource, "network")
		config.MarkAsRequired(r.TerraformResource, "location")
		r.References["psc_config.subnets"] = config.Reference{
			TerraformName:     "google_compute_subnetwork",
			SelectorFieldName: "SubnetSelector",
			RefFieldName:      "SubnetRefs",
			Type:              "github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork",
			Extractor:         common.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})
}
