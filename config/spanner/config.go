// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package spanner

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	// This addresses:
	// cannot run refresh: refresh failed: Invalid combination of arguments: "num_nodes": only one of `num_nodes,processing_units` can be specified, but `num_nodes,processing_units` were specified.
	p.AddResourceConfigurator("google_spanner_instance", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"num_nodes", "processing_units"},
		}
	})

	// This configuration does not work, instance cannot be found by using the instaceSelector field
	// The congiguration works only by using instance: name directly
	p.AddResourceConfigurator("google_spanner_instance_iam_member", func(r *config.Resource) {
		r.References["instance"] = config.Reference{
			TerraformName: "google_spanner_instance",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

	// This configuration does not work, instance cannot be found by using the instaceSelector field
	// The congiguration works only by using instance: name directly
	p.AddResourceConfigurator("google_spanner_database_iam_member", func(r *config.Resource) {
		r.References["instance"] = config.Reference{
			TerraformName: "google_spanner_instance",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

	// This configuration does not work, instance cannot be found by using the instaceSelector field
	// The congiguration works only by using database: name directly
	p.AddResourceConfigurator("google_spanner_database_iam_member", func(r *config.Resource) {
		r.References["database"] = config.Reference{
			TerraformName: "google_spanner_database",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})

}
