package logging

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_logging_folder_bucket_config", func(r *config.Resource) {
		r.References["folder"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Folder",
			Extractor: common.ExtractFolderIDFuncPath,
		}
	})
}
