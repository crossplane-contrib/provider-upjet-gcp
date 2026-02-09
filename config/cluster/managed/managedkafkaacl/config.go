package managedkafkaacl

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_managed_kafka_acl", func(r *config.Resource) {
		r.ShortGroup = "managedkafkaacl" // group in your CRD (matches your existing ones)
		r.Kind = "KafkaACL"              // desired Kubernetes Kind name

		// These fields are required in the Terraform schema:
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"id",
			},
		}
	})
}
