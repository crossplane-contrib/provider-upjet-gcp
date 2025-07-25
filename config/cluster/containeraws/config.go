// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package containeraws

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("google_container_aws_cluster", func(r *config.Resource) {
		r.Kind = "Cluster"
	})

	p.AddResourceConfigurator("google_container_aws_node_pool", func(r *config.Resource) {
		r.Kind = "NodePool"

		r.References["cluster"] = config.Reference{
			TerraformName: "google_container_aws_cluster",
		}
	})

}
