package iap

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_iap_web_backend_service_iam_member", func(r *config.Resource) {
		r.References["web_backend_service"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/compute/v1beta1.BackendService",
		}
	})

	p.AddResourceConfigurator("google_iap_web_type_app_engine_iam_member", func(r *config.Resource) {
		r.References["app_id"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/appengine/v1beta1.Application",
		}
	})
}
