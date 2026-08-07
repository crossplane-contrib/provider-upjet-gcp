package managedkafkatopic

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_managed_kafka_topic", func(r *config.Resource) {
		r.ShortGroup = "managedkafkatopic"
		r.Kind = "KafkaTopic"
	})
}
