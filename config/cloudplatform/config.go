package cloudplatform

import (
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_folder_iam_member", func(r *config.Resource) {
		r.References["folder"] = config.Reference{
			Type:      "Folder",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_project", func(r *config.Resource) {
		r.TerraformResource.Schema["org_id"].Description =
			"The numeric ID of the organization this project belongs to."
	})
	p.AddResourceConfigurator("google_project_default_service_accounts", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "Project",
		}
	})
	p.AddResourceConfigurator("google_project_iam_member", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "Project",
		}
	})
	p.AddResourceConfigurator("google_project_iam_audit_config", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "Project",
		}
	})
	p.AddResourceConfigurator("google_project_service", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "Project",
		}
	})
	p.AddResourceConfigurator("google_project_usage_export_bucket", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			Type: "Project",
		}
		// Note(donovanmuller): Upjet does not generate this reference automatically
		r.References["bucket_name"] = config.Reference{
			Type: "github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket",
		}
	})
	p.AddResourceConfigurator("google_service_account_key", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of "keepers", since it is a
		// map where elements configured as nil, but needs to be String:
		r.TerraformResource.
			Schema["keepers"].Elem = schema.TypeString
		// Note(donovanmuller): The private_key attribute is already base64 encoded.
		// Therefore, writing it as a connection Secret results in the Secret value being double encoded,
		// so decode it and add to the connection details as private_key
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			privateKeyEncoded, err := common.GetField(attr, "private_key")
			if err != nil {
				return nil, err
			}
			privateKey, err := base64.StdEncoding.DecodeString(privateKeyEncoded)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot decode private_key")
			}
			return map[string][]byte{
				"private_key": privateKey,
			}, nil
		}
		r.References["service_account_id"] = config.Reference{
			Type:      "ServiceAccount",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_service_account", func(r *config.Resource) {
		r.Kind = "ServiceAccount"
	})
	p.AddResourceConfigurator("google_service_account_iam_policy", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type:      "ServiceAccount",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		config.MarkAsRequired(r.TerraformResource, "service_account_id")
	})
	p.AddResourceConfigurator("google_service_account_iam_binding", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type:      "ServiceAccount",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		config.MarkAsRequired(r.TerraformResource, "service_account_id")
	})
	p.AddResourceConfigurator("google_service_account_iam_member", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type:      "ServiceAccount",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		config.MarkAsRequired(r.TerraformResource, "service_account_id")
	})
	p.AddResourceConfigurator("google_service_networking_peered_dns_domain", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "network")
		config.MarkAsRequired(r.TerraformResource, "service")
	})

	p.AddResourceConfigurator("google_storage_hmac_key", func(r *config.Resource) {
		r.References["service_account_email"] = config.Reference{
			Type:      "ServiceAccount",
			Extractor: common.ExtractEmailFuncPath,
		}
	})
}
