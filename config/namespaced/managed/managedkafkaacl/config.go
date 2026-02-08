package managedkafkaacl

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_managed_kafka_acl", func(r *config.Resource) {
		r.ShortGroup = "managedkafkaacl" // group in your CRD (matches your existing ones)
		r.Kind = "KafkaACL"              // desired Kubernetes Kind name

		// These fields are required in the Terraform schema:
		// cluster, location, project, principal, role
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"id",
			},
		}

		// Map Terraform fields to identifiers in Crossplane if needed
		r.UseAsync = true // enables async reconcile for long-running operations

		// Optional: if this resource references a cluster resource
		r.References = map[string]config.Reference{
			"cluster": {
				Type: "github.com/sebembedded/provider-managedkafka/apis/cluster/managedkafkacluster/v1alpha1.KafkaCluster",
			},
		}
	})
}
