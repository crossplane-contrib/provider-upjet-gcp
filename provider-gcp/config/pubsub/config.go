package pubsub

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_pubsub_lite_reservation", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_pubsub_lite_subscription", func(r *config.Resource) {
		r.References["topic"] = config.Reference{
			Type: "LiteTopic",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = pubsubLiteConnectionDetails
		config.MarkAsRequired(r.TerraformResource, "zone")
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_pubsub_lite_topic", func(r *config.Resource) {
		r.References["reservation_config.throughput_reservation"] = config.Reference{
			Type: "LiteReservation",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = pubsubLiteConnectionDetails
		config.MarkAsRequired(r.TerraformResource, "zone")
		config.MarkAsRequired(r.TerraformResource, "partition_config")
		config.MarkAsRequired(r.TerraformResource, "retention_config")
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_pubsub_schema", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_pubsub_subscription", func(r *config.Resource) {
		r.References["topic"] = config.Reference{
			Type: "Topic",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = pubsubConnectionDetails
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_pubsub_subscription_iam_member", func(r *config.Resource) {
		r.References["subscription"] = config.Reference{
			Type: "Subscription",
		}
	})

	p.AddResourceConfigurator("google_pubsub_topic", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = pubsubConnectionDetails
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_pubsub_topic_iam_member", func(r *config.Resource) {
		r.References["topic"] = config.Reference{
			Type: "Topic",
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
