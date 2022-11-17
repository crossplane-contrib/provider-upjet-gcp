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
	p.AddResourceConfigurator("google_bigquery_table_iam_member", func(r *config.Resource) {
		r.References["dataset_id"] = config.Reference{
			TerraformName: "google_bigquery_dataset",
		}
		r.References["table_id"] = config.Reference{
			TerraformName: "google_bigquery_table",
		}
	})
}
