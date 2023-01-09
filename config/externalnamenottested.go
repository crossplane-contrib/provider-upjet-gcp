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

	// accessapproval
	//
	// folders/{{folder_id}}/accessApprovalSettings
	"google_folder_access_approval_settings": config.TemplatedStringAsIdentifier("", "folders/{{ .parameters.folder_id }}/accessApprovalSettings"),
	// organizations/{{organization_id}}/accessApprovalSettings
	"google_organization_access_approval_settings": config.TemplatedStringAsIdentifier("", "organizations/{{ .parameters.organization_id }}/accessApprovalSettings"),
	// projects/{{project_id}}/accessApprovalSettings
	"google_project_access_approval_settings": config.TemplatedStringAsIdentifier("", "projects/{{ .parameters.project_id }}/accessApprovalSettings"),

	// accesscontextmanager
	//
	// {{name}}
	"google_access_context_manager_access_level": config.NameAsIdentifier,
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"google_access_context_manager_access_level_condition": config.IdentifierFromProvider,
	// {{name}}
	"google_access_context_manager_access_policy": config.IdentifierFromProvider,
	// accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin
	"google_access_context_manager_access_policy_iam_binding": config.TemplatedStringAsIdentifier("name", "accessPolicies/{{ .external_name }} {{ .parameters.role }}"),
	// accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin user:jane@example.com
	"google_access_context_manager_access_policy_iam_member": config.TemplatedStringAsIdentifier("name", "accessPolicies/{{ .external_name }} {{ .parameters.role }} {{ .parameters.member }}"),
	// accessPolicies/{{access_policy}}
	"google_access_context_manager_access_policy_iam_policy": config.TemplatedStringAsIdentifier("name", "accessPolicies/{{ .external_name }}"),
	// {{name}}
	"google_access_context_manager_gcp_user_access_binding": config.IdentifierFromProvider,
	// {{name}}
	"google_access_context_manager_service_perimeter": config.NameAsIdentifier,
	// {{perimeter_name}}/{{resource}}
	"google_access_context_manager_service_perimeter_resource": config.TemplatedStringAsIdentifier("resource", "{{ .parameters.perimeter_name }}/{{ .external_name }}"),

	// activedirectory
	//
	// {{name}}: projects/{project}/locations/global/domains/{domainName}
	"google_active_directory_domain": config.TemplatedStringAsIdentifier("domain_name", "projects/{{ .setup.configuration.project }}/locations/global/domains/{{ .external_name }}"),
	// projects/{{project}}/locations/global/domains/{{domain}}/{{target_domain_name}}
	"google_active_directory_domain_trust": config.TemplatedStringAsIdentifier("target_domain_name", "projects/{{ .setup.configuration.project }}/locations/global/domains/{{ .parameters.domain }}/{{ .external_name }}"),

	// apigee
	//
	// {{org_id}}/endpointAttachments/{{endpoint_attachment_id}}
	"google_apigee_endpoint_attachment": config.TemplatedStringAsIdentifier("endpoint_attachment_id", "{{ .parameters.org_id }}/endpointAttachments/{{ .external_name }}"),
	// {{org_id}}/envgroups/{{name}}
	"google_apigee_envgroup": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org_id }}/envgroups/{{ .external_name }}"),
	// {{envgroup_id}}/attachments/{{name}}
	"google_apigee_envgroup_attachment": config.IdentifierFromProvider,
	// {{org_id}}/environments/{{name}}
	"google_apigee_environment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org_id }}/environments/{{ .external_name }}"),
	// {{org_id}}/environments/{{environment}} roles/viewer
	"google_apigee_environment_iam_binding": config.TemplatedStringAsIdentifier("", "{{ .parameters.org_id }}/environments/{{ .parameters.env_id }} {{ .parameters.role }}"),
	// {{org_id}}/environments/{{environment}} roles/viewer user:jane@example.com
	"google_apigee_environment_iam_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.org_id }}/environments/{{ .parameters.env_id }} {{ .parameters.role }} {{ .parameters.member }}"),
	// {{org_id}}/environments/{{environment}}
	"google_apigee_environment_iam_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.org_id }}/environments/{{ .parameters.env_id }}"),
	// {{org_id}}/instances/{{name}}
	"google_apigee_instance": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org_id }}/instances/{{ .external_name }}"),
	// {{instance_id}}/attachments/{{name}}
	"google_apigee_instance_attachment": config.IdentifierFromProvider,
	// organizations/{{name}}
	"google_apigee_organization": config.IdentifierFromProvider,

	// apikeys
	//
	// projects/{{project}}/locations/global/keys/{{name}}
	"google_apikeys_key": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/keys/{{ .external_name }}"),

	// datalossprevention
	//
	// {{parent}}/deidentifyTemplates/{{name}}
	"google_data_loss_prevention_deidentify_template": config.IdentifierFromProvider,
	// {{parent}}/inspectTemplates/{{name}}
	"google_data_loss_prevention_inspect_template": config.IdentifierFromProvider,
	// {{parent}}/jobTriggers/{{name}}
	"google_data_loss_prevention_job_trigger": config.IdentifierFromProvider,
	// {{parent}}/storedInfoTypes/{{name}}
	"google_data_loss_prevention_stored_info_type": config.IdentifierFromProvider,

	// dataproc
	//
	// projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
	"google_dataproc_autoscaling_policy": config.TemplatedStringAsIdentifier("policy_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/autoscalingPolicies/{{ .external_name }}"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"google_dataproc_cluster": config.IdentifierFromProvider,
	// projects/{project}/regions/{region}/clusters/{cluster} roles/editor
	"google_dataproc_cluster_iam_binding": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}regions/{{ .parameters.region }}/clusters/{{ .parameters.cluster }} {{ .parameters.role }}"),
	// projects/{project}/regions/{region}/clusters/{cluster} roles/editor user:jane@example.com
	"google_dataproc_cluster_iam_member": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}regions/{{ .parameters.region }}/clusters/{{ .parameters.cluster }} {{ .parameters.role }} {{ .parameters.member }}"),
	// projects/{project}/regions/{region}/clusters/{cluster}
	"google_dataproc_cluster_iam_policy": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}regions/{{ .parameters.region }}/clusters/{{ .parameters.cluster }}"),
	// No import
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"google_dataproc_job": config.IdentifierFromProvider,
	// projects/{project}/regions/{region}/jobs/{job_id} roles/editor
	"google_dataproc_job_iam_binding": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}regions/{{ .parameters.region }}/jobs/{{ .parameters.job_id }} {{ .parameters.role }}"),
	// projects/{project}/regions/{region}/jobs/{job_id} roles/editor user:jane@example.com
	"google_dataproc_job_iam_member": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}regions/{{ .parameters.region }}/jobs/{{ .parameters.job_id }} {{ .parameters.role }} {{ .parameters.member }}"),
	// projects/{project}/regions/{region}/jobs/{job_id}
	"google_dataproc_job_iam_policy": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}regions/{{ .parameters.region }}/jobs/{{ .parameters.job_id }} {{ .parameters.role }}"),
	// projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}
	"google_dataproc_workflow_template": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}locations/{{ .parameters.location }}/workflowTemplates{{ .external_name }}"),

	// datastore
	//
	// projects/{{project}}/indexes/{{index_id}}
	"google_datastore_index": config.IdentifierFromProvider,

	// deploymentmanager
	//
	// projects/{{project}}/deployments/{{name}}
	"google_deployment_manager_deployment": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/deployments{{ .external_name }}"),

	// dialogflow
	//
	// {{project}}
	"google_dialogflow_agent": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}"),
	// {{name}}
	"google_dialogflow_entity_type": config.IdentifierFromProvider,
	// {{name}}
	"google_dialogflow_fulfillment": config.IdentifierFromProvider,
	// {{name}}
	"google_dialogflow_intent": config.IdentifierFromProvider,
}
