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

	// cloudbuild
	//
	// projects/{{project}}/locations/{{location}}/triggers/{{trigger_id}}
	"google_cloudbuild_trigger": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{location}}/workerPools/{{name}}
	"google_cloudbuild_worker_pool": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/workerPools/{{ .external_name }}"),

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

	// cloudiot
	//
	// {{registry}}/devices/{{name}}
	"google_cloudiot_device": config.TemplatedStringAsIdentifier("name", "{{ .parameters.registry }}/devices/{{ .external_name }}"),
	// {{project}}/locations/{{region}}/registries/{{name}}
	"google_cloudiot_registry": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/registries/{{ .external_name }}"),

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

	// datafusion
	//
	// projects/{{project}}/locations/{{region}}/instances/{{name}}
	"google_data_fusion_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/instances/{{ .external_name }}"),
}
