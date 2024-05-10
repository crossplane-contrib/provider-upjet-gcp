// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package pubsub

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_pubsub_lite_reservation", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})

	p.AddResourceConfigurator("google_pubsub_lite_subscription", func(r *config.Resource) {
		r.References["topic"] = config.Reference{
			TerraformName: "google_pubsub_lite_topic",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = pubsubLiteConnectionDetails
		config.MarkAsRequired(r.TerraformResource, "zone")
	})

	p.AddResourceConfigurator("google_pubsub_lite_topic", func(r *config.Resource) {
		r.References["reservation_config.throughput_reservation"] = config.Reference{
			TerraformName: "google_pubsub_lite_reservation",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = pubsubLiteConnectionDetails
		config.MarkAsRequired(r.TerraformResource, "zone")
		config.MarkAsRequired(r.TerraformResource, "partition_config")
		config.MarkAsRequired(r.TerraformResource, "retention_config")
	})
	p.AddResourceConfigurator("google_pubsub_subscription", func(r *config.Resource) {
		r.References["topic"] = config.Reference{
			TerraformName: "google_pubsub_topic",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = pubsubConnectionDetails
		delete(r.References, "cloud_storage_config.bucket")
	})

	p.AddResourceConfigurator("google_pubsub_subscription_iam_member", func(r *config.Resource) {
		r.References["subscription"] = config.Reference{
			TerraformName: "google_pubsub_subscription",
		}
	})

	p.AddResourceConfigurator("google_pubsub_topic", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = pubsubConnectionDetails
	})

	p.AddResourceConfigurator("google_pubsub_topic_iam_member", func(r *config.Resource) {
		r.References["topic"] = config.Reference{
			TerraformName: "google_pubsub_topic",
		}
	})
}

func pubsubLiteConnectionDetails(attr map[string]interface{}) (map[string][]byte, error) {
	conn := map[string][]byte{}
	if a, ok := attr["id"].(string); ok {
		conn["id"] = []byte(a)
	}
	if a, ok := attr["name"].(string); ok {
		conn["topic"] = []byte(a)
	}
	// Note(donovanmuller): If this is a google_pubsub_lite_subscription
	// resource, then the "topic" attribute will be used to override the
	// "topic" connection key of the google_pubsub_lite_topic subscribed
	// too and the "subscription" attribute to the "name" connection key
	// of the subscription
	if a, ok := attr["topic"].(string); ok {
		conn["topic"] = []byte(a)
		if a, ok := attr["name"].(string); ok {
			conn["subscription"] = []byte(a)
		}
	}
	if a, ok := attr["project"].(string); ok {
		conn["project"] = []byte(a)
	}
	if a, ok := attr["zone"].(string); ok {
		conn["zone"] = []byte(a)
	}
	if a, ok := attr["region"].(string); ok {
		conn["region"] = []byte(a)
	}
	return conn, nil
}

func pubsubConnectionDetails(attr map[string]interface{}) (map[string][]byte, error) {
	conn := map[string][]byte{}
	if a, ok := attr["id"].(string); ok {
		conn["id"] = []byte(a)
	}
	if a, ok := attr["name"].(string); ok {
		conn["topic"] = []byte(a)
	}
	// Note(donovanmuller): If this is a google_pubsub_subscription
	// resource, then the "topic" attribute will be used to override the
	// "topic" connection key of the google_pubsub_topic subscribed
	// too and the "subscription" attribute to the "name" connection key
	// of the subscription
	if a, ok := attr["topic"].(string); ok {
		conn["topic"] = []byte(a)
		if a, ok := attr["name"].(string); ok {
			conn["subscription"] = []byte(a)
		}
	}
	if a, ok := attr["project"].(string); ok {
		conn["project"] = []byte(a)
	}
	return conn, nil
}
