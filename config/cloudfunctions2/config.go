package cloudfunctions2

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_cloudfunctions2_function", func(r *config.Resource) {
		r.References["service_config.vpc_connector"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/vpcaccess/v1beta1.Connector",
		}
	})
}
