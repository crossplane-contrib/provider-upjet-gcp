// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package containerazure

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("google_container_azure_client", func(r *config.Resource) {
		r.Kind = "Client"
	})

	p.AddResourceConfigurator("google_container_azure_cluster", func(r *config.Resource) {
		r.Kind = "Cluster"

		// The client field expects value in the following format:
		// projects/{project-number}/locations/{location}}/azureClients/{name}
		// The current Client resource does not include an ID with the
		// project-number. So we are skipping referencing for now.

		// r.References["client"] = config.Reference{
		// 	Type:      "Client",
		// 	Extractor: common.ExtractResourceIDFuncPath,
		// }

	})

	p.AddResourceConfigurator("google_container_azure_node_pool", func(r *config.Resource) {
		r.Kind = "NodePool"

		r.References["cluster"] = config.Reference{
			TerraformName: "google_container_azure_cluster",
		}
	})

}
