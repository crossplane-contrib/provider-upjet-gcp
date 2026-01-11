// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package datastream

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const group = "datastream"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_datastream_stream", func(r *config.Resource) {
		r.ShortGroup = group

		r.References["source_config.source_connection_profile"] = config.Reference{
			TerraformName: "google_datastream_connection_profile",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}

		r.References["destination_config.destination_connection_profile"] = config.Reference{
			TerraformName: "google_datastream_connection_profile",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		}
	})

}
