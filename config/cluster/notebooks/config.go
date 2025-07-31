// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package notebooks

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_notebooks_instance_iam_member", func(r *config.Resource) {
		r.References["instance_name"] = config.Reference{
			TerraformName: "google_notebooks_instance",
		}
	})

	p.AddResourceConfigurator("google_notebooks_instance", func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "create_time")
		config.MoveToStatus(r.TerraformResource, "update_time")
	})

	p.AddResourceConfigurator("google_notebooks_runtime_iam_member", func(r *config.Resource) {
		r.References["runtime_name"] = config.Reference{
			TerraformName: "google_notebooks_runtime",
		}
	})
	p.AddResourceConfigurator("google_notebooks_runtime", func(r *config.Resource) {
		r.MetaResource.ArgumentDocs["data_disk.guest_os_features"] = strings.ReplaceAll(r.MetaResource.ArgumentDocs["data_disk.guest_os_features"], "``", "")
	})
}
