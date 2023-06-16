package accesscontextmanager

import (
	"github.com/upbound/provider-gcp/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_access_context_manager_service_perimeter_resource", func(r *config.Resource) {
		r.References["resource"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project",
			Extractor: common.ExtractProjectNumberFuncPath,
		}
	})
}
