// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package managedkafka

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-gcp/v3/config/namespaced/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_managed_kafka_cluster", func(r *config.Resource) {
		r.UseAsync = true
		r.References["gcp_config.access_config.network_configs.subnet"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["gcp_config.kms_key"] = config.Reference{
			TerraformName: "google_kms_crypto_key",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["tls_config.trust_config.cas_configs.ca_pool"] = config.Reference{
			TerraformName: "google_privateca_ca_pool",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_managed_kafka_topic", func(r *config.Resource) {
		r.References["cluster"] = config.Reference{
			TerraformName: "google_managed_kafka_cluster",
		}
	})
	p.AddResourceConfigurator("google_managed_kafka_acl", func(r *config.Resource) {
		r.References["cluster"] = config.Reference{
			TerraformName: "google_managed_kafka_cluster",
		}
	})
	p.AddResourceConfigurator("google_managed_kafka_connect_cluster", func(r *config.Resource) {
		r.UseAsync = true
		r.References["kafka_cluster"] = config.Reference{
			TerraformName: "google_managed_kafka_cluster",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["gcp_config.access_config.network_configs.primary_subnet"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
		r.References["gcp_config.access_config.network_configs.additional_subnets"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_managed_kafka_connector", func(r *config.Resource) {
		r.References["connect_cluster"] = config.Reference{
			TerraformName: "google_managed_kafka_connect_cluster",
		}
	})
}
