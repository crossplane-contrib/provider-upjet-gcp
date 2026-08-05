// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package storagetransfer

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_storage_transfer_job", func(r *config.Resource) {
		r.References["replication_spec.gcs_data_sink.bucket_name"] = config.Reference{
			TerraformName: "google_storage_bucket",
		}

		r.References["replication_spec.gcs_data_source.bucket_name"] = config.Reference{
			TerraformName: "google_storage_bucket",
		}

		r.References["transfer_spec.gcs_data_sink.bucket_name"] = config.Reference{
			TerraformName: "google_storage_bucket",
		}

		r.References["transfer_spec.gcs_data_source.bucket_name"] = config.Reference{
			TerraformName: "google_storage_bucket",
		}
	})
}
