// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package tags

import (
	"github.com/upbound/provider-gcp/config/common"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_tags_tag_binding", func(r *config.Resource) {
		r.UseAsync = true
		r.References["tag_value"] = config.Reference{
			TerraformName: "google_tags_tag_value",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		config.MarkAsRequired(r.TerraformResource, "parent", "tag_value")
	})
	p.AddResourceConfigurator("google_tags_tag_key", func(r *config.Resource) {
		r.UseAsync = true
		config.MarkAsRequired(r.TerraformResource, "parent", "short_name")
	})
	p.AddResourceConfigurator("google_tags_tag_value", func(r *config.Resource) {
		r.UseAsync = true
		r.References["parent"] = config.Reference{
			TerraformName: "google_tags_tag_key",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		config.MarkAsRequired(r.TerraformResource, "parent", "short_name")
	})
}
