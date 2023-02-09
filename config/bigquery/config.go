// Copyright 2022 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bigquery

import (
	"github.com/upbound/upjet/pkg/config"
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
			Extractor:     "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["load.destination_table.dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["extract.source_table.dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["query.destination_table.table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
			Extractor:     "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["load.destination_table.table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
			Extractor:     "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["copy.destination_table.dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["copy.destination_table.table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
			Extractor:     "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
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
