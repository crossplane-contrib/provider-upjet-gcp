// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package modelarmor

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-gcp/v2/config/cluster/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_model_armor_template", func(r *config.Resource) {
		r.References["filter_config.sdp_settings.advanced_config.inspect_template"] = config.Reference{
			TerraformName: "google_data_loss_prevention_inspect_template",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["filter_config.sdp_settings.advanced_config.deidentify_template"] = config.Reference{
			TerraformName: "google_data_loss_prevention_deidentify_template",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.MetaResource.Description = "A `Template` is a resource of Model Armor that lets you configure how Model Armor screens prompts and responses."
		r.MetaResource.SubCategory = "Model Armor"
	})

	p.AddResourceConfigurator("google_model_armor_floorsetting", func(r *config.Resource) {
		r.Kind = "FloorSetting"
		r.References["filter_config.sdp_settings.advanced_config.inspect_template"] = config.Reference{
			TerraformName: "google_data_loss_prevention_inspect_template",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["filter_config.sdp_settings.advanced_config.deidentify_template"] = config.Reference{
			TerraformName: "google_data_loss_prevention_deidentify_template",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
}
