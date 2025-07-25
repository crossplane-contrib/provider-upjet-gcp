// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dataflow

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_dataflow_job", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of "labels", since is a map
		// where elements configured as nil configured as nil, but needs to be
		// String:
		r.TerraformResource.Schema["labels"].Elem = schema.TypeString
		r.TerraformResource.Schema["parameters"].Elem = schema.TypeString
		r.TerraformResource.Schema["transform_name_mapping"].Elem = schema.TypeString
		r.MetaResource.Description = "Creates a job in Dataflow according to a provided config file. Dataflow jobs are different from most other Google resources. The Dataflow resource is considered 'existing' while it is in a nonterminal state. If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE', 'CANCELLED'), it will be recreated on the next 'reconcile'. This resource does not support import"
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil {
				delete(diff.Attributes, "additional_experiments.#")
			}
			return diff, nil
		}
	})
}
