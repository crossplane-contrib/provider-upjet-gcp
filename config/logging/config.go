// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package logging

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_logging_folder_bucket_config", func(r *config.Resource) {
		r.References["folder"] = config.Reference{
			TerraformName: "google_folder",
			Extractor:     common.ExtractFolderIDFuncPath,
		}
	})
}
