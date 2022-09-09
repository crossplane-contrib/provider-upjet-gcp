package storage

import (
	"time"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_storage_bucket_object", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of
		// "customer_encryption", since it is struct marked as sensitive.
		// https://github.com/upbound/upjet/issues/100#issuecomment-966892273
		r.TerraformResource.Schema["customer_encryption"].Sensitive = false
	})

	p.AddResourceConfigurator("google_storage_bucket", func(r *config.Resource) {
		// Note(turkenh): Terraform provider added retries to bucket resource
		// for eventual consistency: https://github.com/hashicorp/terraform-provider-google/pull/10287
		// However, this causes refresh to keep retrying until timeout if the
		// bucket does not exist indeed. This causes the initial observe call to
		// hang on until timeout. We configure read timeout to a relatively
		// smaller value as a workaround/solution.
		// Related issue: https://github.com/upbound/official-providers/provider-gcp/issues/12
		r.OperationTimeouts.Read = 1 * time.Minute
	})

	p.AddResourceConfigurator("google_storage_bucket_iam_member", func(r *config.Resource) {
		r.References["bucket"] = config.Reference{
			Type: "Bucket",
		}
	})

	p.AddResourceConfigurator("google_storage_bucket_object", func(r *config.Resource) {
		r.TerraformResource.Schema["content"].Sensitive = false
		r.References["bucket"] = config.Reference{
			Type: "Bucket",
		}
	})
}
