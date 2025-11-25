// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apigee

import (
	"fmt"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-gcp/v2/config/namespaced/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("google_apigee_endpoint_attachment", func(r *config.Resource) {
		r.References["service_attachment"] = config.Reference{
			TerraformName: "google_compute_service_attachment",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_apigee_env_keystore", func(r *config.Resource) {
		r.References["env_id"] = config.Reference{
			TerraformName: "google_apigee_environment",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_apigee_env_references", func(r *config.Resource) {
		r.References["env_id"] = config.Reference{
			TerraformName: "google_apigee_environment",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_apigee_keystores_aliases_key_cert_file", func(r *config.Resource) {
		// org_id requires Organization ID without organization/ prefix, so extract name attribute
		r.References["org_id"] = config.Reference{
			TerraformName: "google_apigee_organization",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("name",true)`,
		}
	})
	p.AddResourceConfigurator("google_apigee_addons_config", func(r *config.Resource) {
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}
			// when all the addons configuration options are disabled,
			// upstream TF provider returns an empty list.
			// when the configuration explicitly specifies all as false,
			// it causes a diff, although they are
			// semantically equivalent. in such situation, ignore the diff.
			isAddonsConfigEmptyState := state.Attributes["addons_config.#"] == "0"
			addonsConfigurationsFields := []string{"advanced_api_ops_config", "api_security_config",
				"connectors_platform_config", "integration_config", "monetization_config"}

			hasActualDiff := false
			for _, field := range addonsConfigurationsFields {
				addonDiff, ok := diff.Attributes[fmt.Sprintf("addons_config.0.%s.0.enabled", field)]
				if ok && !(addonDiff.Old == "" && addonDiff.New == "false") {
					hasActualDiff = true
					break
				}
			}
			if isAddonsConfigEmptyState && !hasActualDiff {
				for k := range diff.Attributes {
					if strings.HasPrefix(k, "addons_config.") {
						delete(diff.Attributes, k)
					}
				}
			}
			return diff, nil
		}
	})
}
