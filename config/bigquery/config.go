// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package bigquery

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_bigquery_dataset", func(r *config.Resource) {
		delete(r.References, "access.dataset.dataset.project_id")
	})
	p.AddResourceConfigurator("google_bigquery_dataset_access", func(r *config.Resource) {
		delete(r.References, "view.project_id")
		delete(r.References, "dataset.dataset.project_id")
	})
	p.AddResourceConfigurator("google_bigquery_dataset_iam_member", func(r *config.Resource) {
		r.References["dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
	})
	p.AddResourceConfigurator("google_bigquery_dataset_iam_binding", func(r *config.Resource) {
		r.References["dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
	})
	p.AddResourceConfigurator("google_bigquery_table_iam_binding", func(r *config.Resource) {
		r.References["dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
		}
	})
	p.AddResourceConfigurator("google_bigquery_table_iam_policy", func(r *config.Resource) {
		r.References["dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		delete(r.References, "project")
	})
	p.AddResourceConfigurator("google_bigquery_table_iam_member", func(r *config.Resource) {
		r.References["dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
		}
	})
	p.AddResourceConfigurator("google_bigquery_job", func(r *config.Resource) {
		r.References["query.destination_table.dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["extract.source_table.table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["load.destination_table.dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["extract.source_table.dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["query.destination_table.table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["load.destination_table.table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["copy.destination_table.dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["copy.destination_table.table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}

		delete(r.References, "copy.destination_table.project_id")
		delete(r.References, "query.destination_table.project_id")
		delete(r.References, "load.destination_table.project_id")
		delete(r.References, "extract.source_table.project_id")
	})
	p.AddResourceConfigurator("google_bigquery_analytics_hub_data_exchange_iam_member", func(r *config.Resource) {
		r.References["data_exchange_id"] = config.Reference{
			TerraformName: "google_bigquery_analytics_hub_data_exchange",
		}
	})
}
