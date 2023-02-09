package healthcare

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_healthcare_dataset_iam_member", func(r *config.Resource) {
		r.References["dataset_id"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/healthcare/v1beta1.Dataset",
		}
	})
}
