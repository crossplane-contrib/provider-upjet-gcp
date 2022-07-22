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
			Type:              "LiteTopic",
			RefFieldName:      "TopicRef",
			SelectorFieldName: "TopicSelector",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("google_pubsub_lite_topic", func(r *config.Resource) {
		r.References["reservation_config.throughput_reservation"] = config.Reference{
			Type:              "LiteReservation",
			RefFieldName:      "ThroughputReservationRef",
			SelectorFieldName: "ThroughputReservationSelector",
		}
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
			Type:              "Topic",
			RefFieldName:      "TopicRef",
			SelectorFieldName: "TopicSelector",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("google_pubsub_subscription_iam_binding", func(r *config.Resource) {
		r.References["subscription"] = config.Reference{
			Type:              "Subscription",
			RefFieldName:      "SubscriptionRef",
			SelectorFieldName: "SubscriptionSelector",
		}
	})
	p.AddResourceConfigurator("google_pubsub_subscription_iam_member", func(r *config.Resource) {
		r.References["subscription"] = config.Reference{
			Type:              "Subscription",
			RefFieldName:      "SubscriptionRef",
			SelectorFieldName: "SubscriptionSelector",
		}
	})
	p.AddResourceConfigurator("google_pubsub_subscription_iam_policy", func(r *config.Resource) {
		r.References["subscription"] = config.Reference{
			Type:              "Subscription",
			RefFieldName:      "SubscriptionRef",
			SelectorFieldName: "SubscriptionSelector",
		}
	})
	p.AddResourceConfigurator("google_pubsub_topic", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("google_pubsub_topic_iam_binding", func(r *config.Resource) {
		r.References["topic"] = config.Reference{
			Type:              "Topic",
			RefFieldName:      "TopicRef",
			SelectorFieldName: "TopicSelector",
		}
	})
	p.AddResourceConfigurator("google_pubsub_topic_iam_member", func(r *config.Resource) {
		r.References["topic"] = config.Reference{
			Type:              "Topic",
			RefFieldName:      "TopicRef",
			SelectorFieldName: "TopicSelector",
		}
	})
	p.AddResourceConfigurator("google_pubsub_topic_iam_policy", func(r *config.Resource) {
		r.References["topic"] = config.Reference{
			Type:              "Topic",
			RefFieldName:      "TopicRef",
			SelectorFieldName: "TopicSelector",
		}
	})
}
