package managedkafkacluster

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_managed_kafka_cluster", func(r *config.Resource) {
		r.ShortGroup = "managedkafkacluster"
		r.Kind = "KafkaCluster"
	})
}
