// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"regexp"
	"strings"

	tjconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/types/name"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// VersionV1Beta1 is used to signify that the resource has been tested and external name configured
const VersionV1Beta1 = "v1beta1"

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

func externalNameConfig() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
	}
}

func groupOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		for k, v := range groupMap {
			ok, err := regexp.MatchString(k, r.Name)
			if err != nil {
				panic(errors.Wrap(err, "cannot match regular expression"))
			}
			if ok {
				r.ShortGroup, r.Kind = v(r.Name)
				return
			}
		}
	}
}

var groupMap = map[string]GroupKindCalculator{
	// Note(turkenh): The following resources are listed under "Cloud Platform"
	// section in Terraform Documentation.
	"google_billing_subaccount$":                   ReplaceGroupWords("cloudplatform", 0),
	"google_folder$":                               ReplaceGroupWords("cloudplatform", 0),
	"google_folder_iam.*":                          ReplaceGroupWords("cloudplatform", 0),
	"google_folder_organization.*":                 ReplaceGroupWords("cloudplatform", 0),
	"google_organization_iam.*":                    ReplaceGroupWords("cloudplatform", 0),
	"google_organization_policy.*":                 ReplaceGroupWords("cloudplatform", 0),
	"google_project$":                              ReplaceGroupWords("cloudplatform", 0),
	"google_project_iam.*":                         ReplaceGroupWords("cloudplatform", 0),
	"google_project_service.*":                     ReplaceGroupWords("cloudplatform", 0),
	"google_project_default_service_accounts$":     ReplaceGroupWords("cloudplatform", 0),
	"google_project_organization_policy$":          ReplaceGroupWords("cloudplatform", 0),
	"google_project_usage_export_bucket$":          ReplaceGroupWords("cloudplatform", 0),
	"google_service_account.*":                     ReplaceGroupWords("cloudplatform", 0),
	"google_service_networking_peered_dns_domain$": ReplaceGroupWords("cloudplatform", 0),

	// Resources in "Access Approval" group.
	// Note(turkenh): The following resources are listed under "Access Approval"
	// section in Terraform Documentation.
	"google_.+_approval_settings$": ReplaceGroupWords("accessapproval", 0),

	"google_access_context_manager.+": ReplaceGroupWords("", 3),
	"google_data_loss_prevention.+":   ReplaceGroupWords("", 3),

	"google_service_networking_connection$": ReplaceGroupWords("", 2),
	"google_active_directory.+":             ReplaceGroupWords("", 2),
	"google_app_engine.+":                   ReplaceGroupWords("", 2),
	"google_assured_workloads.+":            ReplaceGroupWords("", 2),
	"google_binary_authorization.+":         ReplaceGroupWords("", 2),
	"google_container_analysis.+":           ReplaceGroupWords("", 2),
	"google_container_attached.+":           ReplaceGroupWords("", 2),
	"google_container_aws.+":                ReplaceGroupWords("", 2),
	"google_container_azure.+":              ReplaceGroupWords("", 2),
	"google_deployment_manager.+":           ReplaceGroupWords("", 2),
	"google_dialogflow_cx.+":                ReplaceGroupWords("", 2),
	"google_essential_contacts.+":           ReplaceGroupWords("", 2),
	"google_game_services.+":                ReplaceGroupWords("", 2),
	"google_gke_hub.+":                      ReplaceGroupWords("", 2),
	"google_identity_platform.+":            ReplaceGroupWords("", 2),
	"google_ml_engine.+":                    ReplaceGroupWords("", 2),
	"google_model_armor.+":                  ReplaceGroupWords("", 2),
	"google_network_management.+":           ReplaceGroupWords("", 2),
	"google_network_services.+":             ReplaceGroupWords("", 2),
	"google_network_connectivity.+":         ReplaceGroupWords("", 2),
	"google_network_security.+":             ReplaceGroupWords("", 2),
	"google_resource_manager.+":             ReplaceGroupWords("", 2),
	"google_secret_manager.+":               ReplaceGroupWords("", 2),
	"google_storage_transfer.+":             ReplaceGroupWords("", 2),
	"google_org_policy.+":                   ReplaceGroupWords("", 2),
	"google_vertex_ai.+":                    ReplaceGroupWords("", 2),
	"google_vpc_access.+":                   ReplaceGroupWords("", 2),

	"google_cloud_asset.+":     ReplaceGroupWords("", 2),
	"google_cloud_functions.+": ReplaceGroupWords("", 2),
	"google_cloud_identity.+":  ReplaceGroupWords("", 2),
	"google_cloud_tasks.+":     ReplaceGroupWords("", 2),
	"google_cloud_scheduler.+": ReplaceGroupWords("", 2),
	"google_cloud_run.+":       ReplaceGroupWords("", 2),

	"google_data_catalog.+": ReplaceGroupWords("", 2),
	"google_data_flow.+":    ReplaceGroupWords("", 2),
	"google_data_fusion.+":  ReplaceGroupWords("", 2),

	"google_os_config.+": ReplaceGroupWords("", 2),
	"google_os_login.+":  ReplaceGroupWords("", 2),

	"google_certificate_manager.+": ReplaceGroupWords("", 2),

	"google_document_ai.+": ReplaceGroupWords("", 2),

	"google_essential_contacts_contact$": ReplaceGroupWords("", 2),

	"google_cloud_quotas.+$": ReplaceGroupWords("", 2),
}

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		// "google_cloud_run_domain_mapping": "cloudrun" -> (cloudrun, DomainMapping)
		words := strings.Split(strings.TrimPrefix(resource, "google_"), "_")
		if group == "" {
			group = strings.Join(words[:count], "")
		}
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

func defaultVersion() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		r.Version = VersionV1Beta1
	}
}

func descriptionOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		tjconfig.ManipulateEveryField(r.TerraformResource, func(sch *schema.Schema) {
			sch.Description = ""
		})
	}
}

// resourcesWithPreexistingDeletionPolicy is the set of TF resource names that
// had a top-level deletion_policy field before the v7.33.0 provider bump.
// These must be kept to avoid breaking existing CRD APIs.
var resourcesWithPreexistingDeletionPolicy = map[string]struct{}{
	"google_alloydb_cluster":                        {},
	"google_bigtable_gc_policy":                     {},
	"google_billing_subaccount":                     {},
	"google_chronicle_data_table":                   {},
	"google_chronicle_rule":                         {},
	"google_compute_shared_vpc_service_project":     {},
	"google_container_attached_cluster":             {},
	"google_datastream_private_connection":          {},
	"google_firebase_data_connect_service":          {},
	"google_firestore_database":                     {},
	"google_firestore_index":                        {},
	"google_looker_instance":                        {},
	"google_netapp_volume":                          {},
	"google_project":                                {},
	"google_secret_manager_regional_secret_version": {},
	"google_secret_manager_secret_version":          {},
	"google_secure_source_manager_instance":         {},
	"google_secure_source_manager_repository":       {},
	"google_service_networking_connection":          {},
	"google_sql_database":                           {},
	"google_sql_provision_script":                   {},
	"google_sql_user":                               {},
	"google_storage_bucket_object":                  {},
	"google_vertex_ai_reasoning_engine":             {},
}

// deletionPolicyOverride move to status the deletion_policy field for
// resources that gained it in the v7.33.0 provider bump. Resources that
// already exposed this field before the bump are preserved to avoid breaking
// existing CRD APIs.
// Directly removing from the schema causes some runtime issues. So this way
// was preferred.
func deletionPolicyOverride() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		if _, keep := resourcesWithPreexistingDeletionPolicy[r.Name]; keep {
			return
		}
		tjconfig.MoveToStatus(r.TerraformResource, "deletion_policy")
	}
}
