package dns

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_dns_managed_zone", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("google_dns_policy", func(r *config.Resource) {
		r.References["networks.network_url"] = config.Reference{
			Type:              "github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Network",
			RefFieldName:      "NetworkRef",
			SelectorFieldName: "NetworkSelector",
			Extractor:         common.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})
}
