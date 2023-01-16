package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameNotTestedConfigs contains no-tested configurations for this
// provider.
var ExternalNameNotTestedConfigs = map[string]config.ExternalName{
	// appengine
	//
	// {{project}}
	"google_app_engine_application_url_dispatch_rules": config.TemplatedStringAsIdentifier("", "{{ .setup.configuration.project }}"),
	// apps/{{project}}/domainMappings/{{domain_name}}
	"google_app_engine_domain_mapping": config.TemplatedStringAsIdentifier("domain_name", "apps/{{ .setup.configuration.project }}/domainMappings/{{ .external_name }}"),
	// apps/{{project}}/firewall/ingressRules/{{priority}}
	"google_app_engine_firewall_rule": config.TemplatedStringAsIdentifier("priority", "apps/{{ .setup.configuration.project }}/firewall/ingressRules/{{ .external_name }}"),
	// apps/{{project}}/services/{{service}}/versions/{{version_id}}
	"google_app_engine_flexible_app_version": config.TemplatedStringAsIdentifier("version_id", "apps/{{ .setup.configuration.project }}/services{{ .parameters.service }}/versions/{{ .external_name }}"),
	// apps/{{project}}/services/{{service}}
	"google_app_engine_service_network_settings": config.TemplatedStringAsIdentifier("service", "apps/{{ .setup.configuration.project }}/services/{{ .external_name }}"),
	// apps/{{project}}/services/{{service}}
	"google_app_engine_service_split_traffic": config.TemplatedStringAsIdentifier("service", "apps/{{ .setup.configuration.project }}/services/{{ .external_name }}"),
	// apps/{{project}}/services/{{service}}/versions/{{version_id}}
	"google_app_engine_standard_app_version": config.TemplatedStringAsIdentifier("version_id", "apps/{{ .setup.configuration.project }}/services/{{ .parameters.service }}/versions/{{ .external_name }}"),

	// assuredworkloads
	//
	// organizations/{{organization}}/locations/{{location}}/workloads/{{name}}
	"google_assured_workloads_workload": config.IdentifierFromProvider,

	// bigtable
	//
	// projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}
	"google_bigtable_app_profile": config.TemplatedStringAsIdentifier("app_profile_id", "projects/{{ .setup.configuration.project }}/instances/{{ .parameters.instance }}/appProfiles/{{ .external_name }}"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"google_bigtable_gc_policy": config.IdentifierFromProvider,
	// projects/{{project}}/instances/{{name}}
	"google_bigtable_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/instances/{{ .external_name }}"),
	// projects/{project}/instances/{instance} roles/editor
	"google_bigtable_instance_iam_binding": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/instances/{{ .parameters.instance }} {{ .parameters.role }}"),
	// projects/{project}/instances/{instance} roles/editor user:jane@example.com
	"google_bigtable_instance_iam_member": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/instances/{{ .parameters.instance }} {{ .parameters.role }} {{ .parameters.member }}"),
	// projects/{project}/instances/{instance}
	"google_bigtable_instance_iam_policy": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/instances/{{ .parameters.instance }}"),
	// projects/{{project}}/instances/{{instance_name}}/tables/{{name}}
	"google_bigtable_table": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/instances/{{ .parameters.instance_name }}/tables/{{ .external_name }}"),
	// projects/{project}/tables/{table} roles/editor
	"google_bigtable_table_iam_binding": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/tables/{{ .parameters.table }} {{ .parameters.role }}"),
	// projects/{project}/tables/{table} roles/editor user:jane@example.com
	"google_bigtable_table_iam_member": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/tables/{{ .parameters.table }} {{ .parameters.role }} {{ .parameters.member }}"),
	// projects/{project}/tables/{table}
	"google_bigtable_table_iam_policy": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/instances/{{ .parameters.table }}"),

	// billing
	//
	// your-billing-account-id roles/billing.user
	"google_billing_account_iam_binding": config.TemplatedStringAsIdentifier("", "{{ .parameters.billing_account_id }} {{ .parameters.role }}"),
	// your-billing-account-id roles/billing.user user:jane@example.com
	"google_billing_account_iam_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.billing_account_id }} {{ .parameters.role }} {{ .parameters.member }}"),
	// your-billing-account-id
	"google_billing_account_iam_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.billing_account_id }}"),
	// billingAccounts/{{billing_account}}/budgets/{{name}}
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"google_billing_budget": config.IdentifierFromProvider,

	// binaryauthorization
	//
	// projects/{{project}}/attestors/{{name}}
	"google_binary_authorization_attestor": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/attestors/{{ .external_name }}"),
	// projects/{{project}}/attestors/{{attestor}} roles/viewer
	"google_binary_authorization_attestor_iam_binding": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/attestors/{{ .parameters.attestor }} {{ .parameters.role }}"),
	// projects/{{project}}/attestors/{{attestor}} roles/viewer user:jane@example.com
	"google_binary_authorization_attestor_iam_member": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/attestors/{{ .parameters.attestor }} {{ .parameters.role }} {{ .parameters.member }}"),
	// projects/{{project}}/attestors/{{attestor}}
	"google_binary_authorization_attestor_iam_policy": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}/attestors/{{ .parameters.attestor }}"),
	// projects/{{project}}
	"google_binary_authorization_policy": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}"),

	// certificatemanager
	//
	// projects/{{project}}/locations/global/certificates/{{name}}
	"google_certificate_manager_certificate": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/certificates/{{ .external_name }}"),
	// projects/{{project}}/locations/global/dnsAuthorizations/{{name}}
	"google_certificate_manager_dns_authorization": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/dnsAuthorizations/{{ .external_name }}"),

	// cloudasset
	//
	// folders/{{folder_id}}/feeds/{{name}}
	"google_cloud_asset_folder_feed": config.IdentifierFromProvider,
	// organizations/{{org_id}}/feeds/{{name}}
	"google_cloud_asset_organization_feed": config.IdentifierFromProvider,
	// projects/{{project}}/feeds/{{name}}
	"google_cloud_asset_project_feed": config.IdentifierFromProvider,

	// clouddeploy
	//
	// projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}
	"google_clouddeploy_delivery_pipeline": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/deliveryPipelines/{{ .external_name }}"),
	// projects/{{project}}/locations/{{location}}/targets/{{name}}
	"google_clouddeploy_target": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/targets/{{ .external_name }}"),

	// cloudidentity
	//
	// {{name}}
	"google_cloud_identity_group": config.IdentifierFromProvider,
	// {{name}}
	"google_cloud_identity_group_membership": config.IdentifierFromProvider,

	// datacatalog
	//
	// projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer
	"google_data_catalog_entry_group_iam_binding": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer user:jane@example.com
	"google_data_catalog_entry_group_iam_member": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}}
	"google_data_catalog_entry_group_iam_policy": config.IdentifierFromProvider,
	// {{name}}: projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id} where tag_id is a system-generated identifier
	"google_data_catalog_tag": config.IdentifierFromProvider,
	// {{name}}: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
	"google_data_catalog_tag_template": config.TemplatedStringAsIdentifier("tag_template_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/tagTemplates/{{ .external_name }}"),
	// projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}} roles/viewer
	"google_data_catalog_tag_template_iam_binding": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}} roles/viewer user:jane@example.com
	"google_data_catalog_tag_template_iam_member": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}}
	"google_data_catalog_tag_template_iam_policy": config.IdentifierFromProvider,

	// accessapproval
	//
	// Imported by using the following format: folders/{{folder_id}}/accessApprovalSettings
	"google_folder_access_approval_settings": config.TemplatedStringAsIdentifier("", "folders/{{ .parameters.folder_id }}/accessApprovalSettings"),
	// Imported by using the following format: organizations/{{organization_id}}/accessApprovalSettings
	"google_organization_access_approval_settings": config.TemplatedStringAsIdentifier("", "organizations/{{ .parameters.organization_id }}/accessApprovalSettings"),
	// Imported by using the following format: projects/{{project_id}}/accessApprovalSettings
	"google_project_access_approval_settings": config.TemplatedStringAsIdentifier("", "projects/{{ .parameters.project_id }}/accessApprovalSettings"),

	// accesscontextmanager
	//
	//
	// Imported by using the following format: {{name}}: accessPolicies/{policy_id}/accessLevels/{short_name}
	"google_access_context_manager_access_level": config.IdentifierFromProvider,
	// No import
	"google_access_context_manager_access_level_condition": config.IdentifierFromProvider,
	// Imported by using {{name}}, but name doesn't exist in parameters. Try using IdentifierFromProvider
	"google_access_context_manager_access_policy": config.IdentifierFromProvider,
	// Imported by using the following format: accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin
	"google_access_context_manager_access_policy_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: "accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin user:jane@example.com"
	"google_access_context_manager_access_policy_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: accessPolicies/{{access_policy}}
	"google_access_context_manager_access_policy_iam_policy": config.TemplatedStringAsIdentifier("name", "accessPolicies/{{ .external_name }}"),
	// Imported by using the following format: {{name}}
	"google_access_context_manager_gcp_user_access_binding": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	"google_access_context_manager_service_perimeter": config.IdentifierFromProvider,
	// Imported by using the following format: {{perimeter_name}}/{{resource}}
	"google_access_context_manager_service_perimeter_resource": config.TemplatedStringAsIdentifier("resource", "{{ .parameters.perimeter_name }}/{{ .external_name }}"),

	// activedirectory
	//
	// Imported by using the following format: {{name}}
	"google_active_directory_domain": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}
	"google_active_directory_domain_trust": config.TemplatedStringAsIdentifier("target_domain_name", "projects/{{ .setup.configuration.project }}/locations/global/domains/{{ .parameters.domain }}/{{ .external_name }}"),

	// apigee
	//
	// Imported by using the following format: {{org_id}}/endpointAttachments/{{endpoint_attachment_id}}
	"google_apigee_endpoint_attachment": config.TemplatedStringAsIdentifier("endpoint_attachment_id", "{{ .parameters.org_id }}/endpointAttachments/{{ .external_name }}"),
	// Imported by using the following format: {{org_id}}/envgroups/{{name}}
	"google_apigee_envgroup": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org_id }}/envgroups/{{ .external_name }}"),
	// Imported by using the following format: {{envgroup_id}}/attachments/{{name}}. Name doesn't exist in parameters, try using IdentifierFromProvider
	"google_apigee_envgroup_attachment": config.IdentifierFromProvider,
	// Imported by using the following format: {{org_id}}/environments/{{name}}
	"google_apigee_environment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org_id }}/environments/{{ .external_name }}"),
	// Imported by using the following format: {{org_id}}/environments/{{environment}} roles/viewer
	"google_apigee_environment_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: {{org_id}}/environments/{{environment}} roles/viewer user:jane@example.com
	"google_apigee_environment_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: {{org_id}}/environments/{{environment}}
	"google_apigee_environment_iam_policy": config.IdentifierFromProvider,
	// Imported by using the following format: {{org_id}}/instances/{{name}}
	"google_apigee_instance": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org_id }}/instances/{{ .external_name }}"),
	// Imported by using the following format: {{instance_id}}/attachments/{{name}}. Name doesn't exist in parameters, try using IdentifierFromProvider
	"google_apigee_instance_attachment": config.IdentifierFromProvider,
	// Imported by using the following format: organizations/{{name}}. Name doesn't exist in parameters, try using IdentifierFromProvider
	"google_apigee_organization": config.IdentifierFromProvider,

	// apikeys
	//
	// Imported by using the following format: projects/{{project}}/locations/global/keys/{{name}}
	"google_apikeys_key": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/keys/{{ .external_name }}"),
}
