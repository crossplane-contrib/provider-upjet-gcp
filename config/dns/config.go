package dns

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_dns_managed_zone", func(r *config.Resource) {
		r.References["peering_config.target_network.network_url"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/compute/v1beta1.Network",
			Extractor: common.PathSelfLinkExtractor,
		}
		r.References["private_visibility_config.networks.network_url"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/compute/v1beta1.Network",
			Extractor: common.PathSelfLinkExtractor,
		}
	})
	p.AddResourceConfigurator("google_dns_policy", func(r *config.Resource) {
		r.References["networks.network_url"] = config.Reference{
			Type:              "github.com/upbound/provider-gcp/apis/compute/v1beta1.Network",
			RefFieldName:      "NetworkRef",
			SelectorFieldName: "NetworkSelector",
			Extractor:         common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_dns_record_set", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "managed_zone")
		delete(r.References, "name")
	})
}
