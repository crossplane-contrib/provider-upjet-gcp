package gkehub

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_gke_hub_membership_iam_member", func(r *config.Resource) {
		r.References["membership_id"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/gkehub/v1beta1.Membership",
		}
	})
}
