// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudrun

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_cloud_run_domain_mapping", func(r *config.Resource) {
		r.References["metadata.namespace"] = config.Reference{
			TerraformName: "google_project",
		}
		r.References["spec.route_name"] = config.Reference{
			TerraformName: "google_cloud_run_service",
		}
	})
	p.AddResourceConfigurator("google_cloud_run_service", func(r *config.Resource) {
		r.References["metadata.namespace"] = config.Reference{
			TerraformName: "google_project",
		}
		// manually added because during the native bump from 4.22.0->4.48.0
		// crddiff detected the ref fields were removed.
		r.References["template.spec.containers.env.value_from.secret_key_ref.name"] = config.Reference{
			TerraformName: "google_secret_manager_secret",
		}
		r.References["template.spec.volumes.secret.secret_name"] = config.Reference{
			TerraformName: "google_secret_manager_secret",
		}
	})
	p.AddResourceConfigurator("google_cloud_run_service_iam_policy", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "google_project",
		}
		r.References["service"] = config.Reference{
			TerraformName: "google_cloud_run_service",
		}
		config.MarkAsRequired(r.TerraformResource, "location")
	})
	p.AddResourceConfigurator("google_cloud_run_service_iam_binding", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "google_project",
		}
		r.References["service"] = config.Reference{
			TerraformName: "google_cloud_run_service",
		}
		config.MarkAsRequired(r.TerraformResource, "location")
	})
	p.AddResourceConfigurator("google_cloud_run_service_iam_member", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "google_project",
		}
		r.References["service"] = config.Reference{
			TerraformName: "google_cloud_run_service",
		}
		config.MarkAsRequired(r.TerraformResource, "location")
	})
	p.AddResourceConfigurator("google_cloud_run_v2_job", func(r *config.Resource) {
		// This prevents an import cycle not allowed error
		delete(r.References, "template.template.vpc_access.connector")
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff != nil {
				// The Terraform registry docs state that when
				// the `launch_stage` configuration argument
				// is set, the attribute may differ from it.
				// Thus, the provided example manifests
				// ignore changes to this argument using
				// the`ignore_changes` lifecycle meta-argument.
				delete(diff.Attributes, "launch_stage")
			}
			return diff, nil
		}
	})
	p.AddResourceConfigurator("google_cloud_run_v2_service", func(r *config.Resource) {
		// This prevents an import cycle not allowed error
		delete(r.References, "template.vpc_access.connector")
	})
}
