/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/upbound/upjet/pkg/config"
)

var externalNameConfigs = map[string]config.ExternalName{
	// activedirectory
	//
	// Imported by using the following format: {{name}}
	"google_active_directory_domain": config.IdentifierFromProvider,

	// appengine
	//
	// Imported by using the following format: your-project-id
	"google_app_engine_application": config.IdentifierFromProvider,
	// {{project}}
	"google_app_engine_application_url_dispatch_rules": config.TemplatedStringAsIdentifier("", "{{ .setup.configuration.project }}"),
	// Imported by using the following format: your-project-id
	"google_app_engine_service_network_settings": config.IdentifierFromProvider,
	// apps/{{project}}/services/{{service}}/versions/{{version_id}}
	"google_app_engine_standard_app_version": config.TemplatedStringAsIdentifier("version_id", "apps/{{ .setup.configuration.project }}/services/{{ .parameters.service }}/versions/{{ .external_name }}"),
	// apps/{{project}}/firewall/ingressRules/{{priority}}
	"google_app_engine_firewall_rule": config.IdentifierFromProvider,

	// composer
	//
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/environments/{{name}}
	"google_composer_environment": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/environments/{{ .external_name }}"),

	// cloudfunctions
	//
	// Imported by using the following format: {{project}}/{{region}}/function-test
	"google_cloudfunctions_function": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.region }}/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} roles/viewer user:jane@example.com
	"google_cloudfunctions_function_iam_member": config.IdentifierFromProvider,

	// cloudfunctions2
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/functions/{{name}}
	"google_cloudfunctions2_function": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/functions/{{ .external_name }}"),

	// cloudplatform
	//
	// Folders can be imported using the folder's id, e.g. folders/1234567
	"google_folder": config.IdentifierFromProvider,
	// Imported by using the following format: folders/your-folder-id roles/viewer user:foo@example.com
	"google_folder_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: organizations/{{org_id}}/roles/{{role_id}}
	"google_organization_iam_custom_role": config.IdentifierFromProvider,
	// Imported by using the following format: your-orgid roles/viewer user:foo@example.com
	"google_organization_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: your-organization-id foo.googleapis.com
	"google_organization_iam_audit_config": config.IdentifierFromProvider,
	// Projects can be imported using the project_id: your-project-id
	// Project-ID has a format as following: projects/{{project}}
	// So, the GetIDFn function implementation for project-id and import method
	// for project resource seems that different.
	"google_project": config.IdentifierFromProvider,
	// Resource with format projects/{{project}}
	// This resource does not support import
	"google_project_default_service_accounts": TemplatedStringAsIdentifierWithNoName("projects/{{ .external_name }}"),
	// Imported by using the following format: your-project-id roles/viewer user:foo@example.com
	"google_project_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: your-project-id foo.googleapis.com
	"google_project_iam_audit_config": config.IdentifierFromProvider,
	// Imported by using the following format: your-project-id/iam.googleapis.com
	"google_project_service": config.IdentifierFromProvider,
	// Imported by using the following format: {{project}}
	"google_project_usage_export_bucket": config.IdentifierFromProvider,
	// Service accounts can be imported using their URI, e.g. projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com
	"google_service_account": config.TemplatedStringAsIdentifier("account_id", "projects/{{ .setup.configuration.project }}/serviceAccounts/{{ .external_name }}@{{ .setup.configuration.project }}.iam.gserviceaccount.com"),
	// Imported by using the following format: projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31
	"google_service_account_iam_member": config.IdentifierFromProvider,
	// No import
	"google_service_account_key": config.IdentifierFromProvider,
	// Imported by using the following format: services/{service}/projects/{project}/global/networks/{network}/peeredDnsDomains/{name}
	"google_service_networking_peered_dns_domain": config.TemplatedStringAsIdentifier("name", "services/{{ .parameters.service }}/projects/{{ .setup.configuration.project }}/global/networks/{{ .parameters.network }}/peeredDnsDomains/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/roles/{{role_id}}
	"google_project_iam_custom_role": config.TemplatedStringAsIdentifier("role_id", "projects/{{ .setup.configuration.project }}/roles/{{ .external_name }}"),

	// cloudrun
	//
	// Imported by using the following format: locations/{{location}}/namespaces/{{project}}/domainmappings/{{name}}
	"google_cloud_run_domain_mapping": config.IdentifierFromProvider,
	// Imported by using the following format: locations/{{location}}/namespaces/{{project}}/services/{{name}}
	"google_cloud_run_service": config.TemplatedStringAsIdentifier("name", "locations/{{ .parameters.location }}/namespaces/{{ .setup.configuration.project }}/services/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer user:jane@example.com
	"google_cloud_run_service_iam_member": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/locations/{{location}}/jobs/{{name}}
	"google_cloud_run_v2_job": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/jobs/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/services/{{name}}
	"google_cloud_run_v2_service": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/services/{{ .external_name }}"),

	// cloudscheduler
	//
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/jobs/{{name}}
	"google_cloud_scheduler_job": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/jobs/{{ .external_name }}"),

	// cloudtasks
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/queues/{{name}}
	"google_cloud_tasks_queue": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/queues/{{ .external_name }}"),

	// compute
	//
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instances/{{instance.name}}/{{disk.name}}
	"google_compute_attached_disk": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}
	"google_compute_autoscaler": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/zones/{{ .parameters.zone }}/autoscalers/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/backendBuckets/{{name}}
	"google_compute_backend_bucket": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/backendBuckets/{{ .external_name }}"),
	// Imported by using the following format: This resource does not support import.
	"google_compute_backend_bucket_signed_url_key": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/backendServices/{{name}}
	"google_compute_backend_service": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/backendServices/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/sslCertificates/{{name}}
	"google_compute_managed_ssl_certificate": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/sslCertificates/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/subnetworks/{{name}}
	"google_compute_subnetwork": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/subnetworks/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/addresses/{{name}}
	"google_compute_address": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/addresses/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/firewalls/{{name}}
	"google_compute_firewall": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/firewalls/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/routers/{{name}}
	"google_compute_router": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/routers/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}
	"google_compute_router_nat": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/routers/{{ .parameters.router }}/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instances/{{name}}
	"google_compute_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/zones/{{ .parameters.zone }}/instances/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/networks/{{name}}
	"google_compute_network": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/networks/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/disks/{{name}}
	"google_compute_disk": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/zones/{{ .parameters.zone }}/disks/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/externalVpnGateways/{{name}}
	"google_compute_external_vpn_gateway": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/externalVpnGateways/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/addresses/{{name}}
	"google_compute_global_address": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/addresses/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/networkEndpointGroups/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
	"google_compute_global_network_endpoint_group": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/networkEndpointGroups/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}
	"google_compute_ha_vpn_gateway": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/vpnGateways/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/healthChecks/{{name}}
	"google_compute_health_check": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/healthChecks/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/httpHealthChecks/{{name}}
	"google_compute_http_health_check": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/httpHealthChecks/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/httpsHealthChecks/{{name}}
	"google_compute_https_health_check": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/httpsHealthChecks/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/images/{{name}}
	"google_compute_image": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/images/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}/zones/{{zone}}/instanceGroups/{{name}}
	"google_compute_instance_group": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/zones/{{ .parameters.zone }}/instanceGroups/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/instanceTemplates/{{name}}
	"google_compute_instance_template": config.IdentifierFromProvider,
	// No import
	"google_compute_instance_from_template": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}
	"google_compute_resource_policy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/resourcePolicies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/disks/{{disk}}/{{name}}
	"google_compute_disk_resource_policy_attachment": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer user:jane@example.com
	"google_compute_disk_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: {{project}}/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
	"google_compute_global_network_endpoint": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/images/{{image}} roles/compute.imageUser user:jane@example.com
	"google_compute_image_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/targetPools/{{name}}
	"google_compute_target_pool": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/targetPools/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}
	"google_compute_instance_group_manager": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/zones/{{ .parameters.zone }}/targetPools/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin user:jane@example.com
	"google_compute_instance_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}
	"google_compute_interconnect_attachment": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/interconnectAttachments/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}
	"google_compute_network_endpoint_group": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/zones/{{ .parameters.zone }}/networkEndpointGroups/{{ .external_name }}"),
	// Imported by using the following format: {{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
	"google_compute_network_endpoint": config.IdentifierFromProvider,
	// Imported by using the following format: project-name/network-name/peering-name
	"google_compute_network_peering": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.network }}/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}
	"google_compute_network_peering_routes_config": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}
	"google_compute_node_group": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/zones/{{ .parameters.zone }}/nodeGroups/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}
	"google_compute_node_template": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/nodeTemplates/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}
	"google_compute_packet_mirroring": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/packetMirrorings/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}
	"google_compute_forwarding_rule": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/forwardingRules/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/backendServices/{{name}}
	"google_compute_region_backend_service": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/backendServices/{{ .external_name }}"),
	// Imported by using the following format: {{name}}
	"google_compute_region_instance_group_manager": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}
	"google_compute_region_target_http_proxy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/targetHttpProxies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/urlMaps/{{name}}
	"google_compute_region_url_map": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/urlMaps/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/{{name}}
	"google_compute_per_instance_config": config.IdentifierFromProvider,
	// Projects can be imported using the Project ID: your-project-id
	"google_compute_project_default_network_tier": config.IdentifierFromProvider,
	// Projects can be imported using the Project ID: your-project-id
	"google_compute_project_metadata": config.IdentifierFromProvider,
	// Project metadata items can be imported using the key: key
	"google_compute_project_metadata_item": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/autoscalers/{{name}}
	"google_compute_region_autoscaler": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/autoscalers/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/disks/{{name}}
	"google_compute_region_disk": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/disks/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/disks/{{region_disk}} roles/viewer user:jane@example.com
	"google_compute_region_disk_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/disks/{{disk}}/{{name}}
	"google_compute_region_disk_resource_policy_attachment": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/healthChecks/{{name}}
	"google_compute_region_health_check": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/healthChecks/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}
	"google_compute_region_network_endpoint_group": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/networkEndpointGroups/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/{{name}}
	"google_compute_region_per_instance_config": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/sslCertificates/{{name}}
	"google_compute_region_ssl_certificate": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/sslCertificates/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}
	"google_compute_region_target_https_proxy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/targetHttpsProxies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/reservations/{{name}}
	"google_compute_reservation": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/zones/{{ .parameters.zone }}/reservations/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/routes/{{name}}
	"google_compute_route": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/routes/{{ .external_name }}"),
	// Imported by using the following format: us-central1/router-1/interface-1
	"google_compute_router_interface": config.IdentifierFromProvider,
	// Imported by using the following format: locations/global/firewallPolicies/{{name}}
	"google_compute_firewall_policy": config.IdentifierFromProvider,
	// Imported by using the following format: locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
	"google_compute_firewall_policy_association": config.IdentifierFromProvider,
	// Imported by using the following format: locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
	"google_compute_firewall_policy_rule": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/forwardingRules/{{name}}
	"google_compute_global_forwarding_rule": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/forwardingRules/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/{{port}}/{{name}}
	"google_compute_instance_group_named_port": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/targetSslProxies/{{name}}
	"google_compute_target_ssl_proxy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/targetSslProxies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/sslCertificates/{{name}}
	"google_compute_ssl_certificate": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/sslCertificates/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/securityPolicies/{{name}}
	"google_compute_security_policy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/securityPolicies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}
	"google_compute_service_attachment": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/serviceAttachments/{{ .external_name }}"),
	// Note(donovanmuller): Requires organizational level permission 'compute.organizations.enableXpnHost'"
	// Imported by using the following format: host-project-id
	// "google_compute_shared_vpc_host_project": config.IdentifierFromProvider,
	// Note(donovanmuller): The google provider only supports this permission at project or organizational level currently
	// It also requires access to multiple GCP Projects
	// Imported by using the following format: host-project-id/service-project-id-1
	// "google_compute_shared_vpc_service_project": TemplatedStringAsIdentifierWithNoName("{{ .setup.configuration.project }}/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser user:jane@example.com
	"google_compute_subnetwork_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/targetGrpcProxies/{{name}}
	"google_compute_target_grpc_proxy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/targetGrpcProxies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/targetHttpProxies/{{name}}
	"google_compute_target_http_proxy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/targetHttpProxies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/targetHttpsProxies/{{name}}
	"google_compute_target_https_proxy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/targetHttpsProxies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/targetInstances/{{name}}
	"google_compute_target_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/zones/{{ .parameters.zone }}/targetInstances/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/targetTcpProxies/{{name}}
	"google_compute_target_tcp_proxy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/targetTcpProxies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/global/urlMaps/{{name}}
	"google_compute_url_map": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/urlMaps/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}
	"google_compute_vpn_gateway": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/targetVpnGateways/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}
	"google_compute_vpn_tunnel": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/vpnTunnels/{{ .external_name }}"),
	// No import
	"google_compute_backend_service_signed_url_key": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/global/firewallPolicies/{{name}}
	"google_compute_network_firewall_policy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/firewallPolicies/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
	"google_compute_network_firewall_policy_association": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/firewallPolicies/{{ .parameters.firewall_policy }}/associations/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/regions/{{region}}/firewallPolicies/{{name}}
	"google_compute_region_network_firewall_policy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/firewallPolicies/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/regions/{{region}}/firewallPolicies/{{firewall_policy}}/associations/{{name}}
	"google_compute_region_network_firewall_policy_association": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/regions/{{ .parameters.region }}/firewallPolicies/{{ .parameters.firewall_policy }}/associations/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/global/snapshots/{{name}}
	"google_compute_snapshot": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/snapshots/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/global/sslPolicies/{{name}}
	"google_compute_ssl_policy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/global/sslPolicies/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/global/snapshots/{{snapshot}} roles/viewer user:jane@example.com
	"google_compute_snapshot_iam_member": config.IdentifierFromProvider,

	// container
	//
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/clusters/my-cluster
	"google_container_cluster": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/clusters/{{ .external_name }}"),
	// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
	"google_container_node_pool": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.location }}/{{ .parameters.cluster }}/{{ .external_name }}"),

	// containeranalysis
	//
	// Imported by using the following format: projects/{{project}}/notes/{{name}}
	"google_container_analysis_note": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/notes/{{ .external_name }}"),
	// Note(donovanmuller): Requires a reference to a KMS CryptoKey version (a Terraform Datasource)
	// as well as complicated values for serialized_payload's
	// Imported by using the following format: projects/{{project}}/occurrences/{{name}}
	// "google_container_analysis_occurrence": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/occurrences/{{ .external_name }}"),

	// containeraws
	//
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/clusters/my-cluster
	"google_container_aws_cluster": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/clusters/{{ .external_name }}"),
	// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
	"google_container_aws_node_pool": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.location }}/{{ .parameters.cluster }}/{{ .external_name }}"),

	// containerazure
	//
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/clusters/my-cluster
	"google_container_azure_cluster": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/clusters/{{ .external_name }}"),
	// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
	"google_container_azure_node_pool": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.location }}/{{ .parameters.cluster }}/{{ .external_name }}"),
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/azureClients/my-client
	"google_container_azure_client": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/azureClients/{{ .external_name }}"),

	// containerregistry
	//
	// This resource can not be imported. The resource will create a bucket if missing, and do nothing with any bucket it finds
	"google_container_registry": config.IdentifierFromProvider,

	// ids
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/endpoints/{{name}}
	"google_cloud_ids_endpoint": config.IdentifierFromProvider,

	// certificatemanager
	//
	// Imported by using the following projects/{{project}}/locations/global/certificateMaps/{{name}}
	"google_certificate_manager_certificate_map": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/certificateMaps/{{ .external_name }}"),

	// datacatalog
	//
	// Imported by using the following format: {{name}}
	"google_data_catalog_entry": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	"google_data_catalog_entry_group": config.IdentifierFromProvider,
	// {{name}}: projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id} where tag_id is a system-generated identifier
	"google_data_catalog_tag": config.IdentifierFromProvider,
	// {{name}}: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
	"google_data_catalog_tag_template": config.TemplatedStringAsIdentifier("tag_template_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/tagTemplates/{{ .external_name }}"),

	// datastore
	//
	// Imported by using the following format: projects/{{project}}/indexes/{{index_id}}
	"google_datastore_index": config.IdentifierFromProvider,

	// dialogflow
	//
	// More details in https://upboundio.slack.com/archives/C01PK1SMYNN/p1663240546481649
	// All resources must be configured with config.IdentifierFromProvider due to all depending on agent
	"google_dialogflow_cx_agent":       config.IdentifierFromProvider,
	"google_dialogflow_cx_entity_type": config.IdentifierFromProvider,
	"google_dialogflow_cx_environment": config.IdentifierFromProvider,
	"google_dialogflow_cx_flow":        config.IdentifierFromProvider,
	"google_dialogflow_cx_intent":      config.IdentifierFromProvider,
	"google_dialogflow_cx_page":        config.IdentifierFromProvider,
	"google_dialogflow_cx_version":     config.IdentifierFromProvider,
	// Imported by using the following {{parent}}/webhooks/{{name}}
	"google_dialogflow_cx_webhook": config.IdentifierFromProvider,

	// datastream
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}
	"google_datastream_connection_profile": config.TemplatedStringAsIdentifier("connection_profile_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/connectionProfiles/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}
	"google_datastream_private_connection": config.TemplatedStringAsIdentifier("private_connection_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/privateConnections/{{ .external_name }}"),

	// dataplex
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/lakes/{{name}}
	"google_dataplex_lake": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/lakes/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}/assets/{{name}}
	"google_dataplex_asset": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/lakes/{{ .parameters.lake }}/zones/{{ .parameters.dataplex_zone }}/assets/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}
	"google_dataplex_zone": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/lakes/{{ .parameters.lake }}/zones/{{ .external_name }}"),

	// dns
	//
	// Imported by using the following format: projects/{{project}}/managedZones/{{name}}
	"google_dns_managed_zone": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/managedZones/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/policies/{{name}}
	"google_dns_policy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/policies/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/managedZones/{{zone}}/rrsets/{{name}}/{{type}}
	"google_dns_record_set": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/managedZones/{{managed_zone}} roles/viewer user:jane@example.com
	"google_dns_managed_zone_iam_member": config.IdentifierFromProvider,

	// endpoints
	//
	// Note(donovanmuller): Must be owner of the domain name used in serviceName
	// This resource does not support import.
	// "google_endpoints_service": config.IdentifierFromProvider,
	// Imported by using the following format: services/{{service_name}} roles/viewer user:jane@example.com
	// "google_endpoints_service_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: services/{{service_name}}/consumers/{{consumer_project}} roles/servicemanagement.serviceController user:jane@example.com
	// "google_endpoints_service_consumers_iam_member": config.IdentifierFromProvider,

	// essential
	//
	// Imported by using the following format:
	"google_essential_contacts_contact": config.IdentifierFromProvider,

	// eventarc
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/triggers/{{name}}
	"google_eventarc_trigger": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/triggers/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/channels/{{name}}
	"google_eventarc_channel": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/channels/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/googleChannelConfig
	"google_eventarc_google_channel_config": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/googleChannelConfig"),

	// filestore
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/instances/{{name}}
	"google_filestore_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/instances/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/backups/{{name}}
	"google_filestore_backup": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/backups/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/instances/{{instance}}/snapshots/{{name}}
	"google_filestore_snapshot": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/instances/{{ .parameters.instance }}/snapshots/{{ .external_name }}"),

	// firebaserules
	//
	// Imported by using the following format: projects/{{project}}/releases/{{name}}
	"google_firebaserules_release": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/releases/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/rulesets/{{name}}
	"google_firebaserules_ruleset": config.IdentifierFromProvider,

	// firestore
	//
	// Imported by using the following format: {{name}}
	// Note(donovanmuller): This resource creates a Firestore Document on a project that already has Firestore enabled
	// The Cloud Firestore API is not available for Datastore Mode projects
	// "google_firestore_document": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	// Note(donovanmuller): This resource creates a Firestore Document on a project that already has Firestore enabled
	// Requires project level IAM permissions
	// "google_firestore_index": config.IdentifierFromProvider,

	// gameservers
	//
	// Note(donovanmuller): All resourcs return this error: Error 404: Method not found
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}
	// "google_game_services_game_server_cluster": config.TemplatedStringAsIdentifier("cluster_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/realms/{{ .parameters.realm_id }}/gameServerClusters/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}
	// "google_game_services_game_server_config": config.TemplatedStringAsIdentifier("config_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/gameServerDeployments/{{ .parameters.deployment_id }}/configs/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}
	// "google_game_services_game_server_deployment": config.TemplatedStringAsIdentifier("deployment_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/gameServerDeployments/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout
	// "google_game_services_game_server_deployment_rollout": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/realms/{{realm_id}}
	// "google_game_services_realm": config.TemplatedStringAsIdentifier("realm_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/realms/{{ .external_name }}"),

	// gkehub
	//
	// Imported by using the following format: projects/{{project}}/locations/global/memberships/{{membership_id}}
	"google_gke_hub_membership": config.TemplatedStringAsIdentifier("membership_id", "projects/{{ .setup.configuration.project }}/locations/global/memberships/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/memberships/{{membership_id}} roles/viewer user:jane@example.com
	"google_gke_hub_membership_iam_member": config.IdentifierFromProvider,

	// gke
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/backupPlans/{{name}}
	"google_gke_backup_backup_plan": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/backupPlans/{{ .external_name }}"),

	// healthcare
	//
	// Imported by using the following format: {{dataset}}/consentStores/{{name}}
	"google_healthcare_consent_store": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/datasets/{{name}}
	"google_healthcare_dataset": config.IdentifierFromProvider,
	// Imported by using the following format: your-project-id/location-name/dataset-name roles/viewer user:foo@example.com
	"google_healthcare_dataset_iam_member": config.IdentifierFromProvider,

	// iap
	//
	// Imported by using the following format: projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_web_backend_service_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/iap_web roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_web_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_web_type_app_engine_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/iap_web/compute roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_web_type_compute_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_app_engine_service_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_app_engine_version_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/iap_web roles/iap.httpsResourceAccessor user:jane@example.com
	"google_iap_tunnel_iam_member": config.TemplatedStringAsIdentifier("", "/projects/{{ .setup.configuration.project }}/iap_web {{ .parameters.role }} {{ .parameters.member }}"),

	// identityplatform
	//
	// Imported by using the following format: projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}
	"google_identity_platform_default_supported_idp_config": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/inboundSamlConfigs/{{name}}
	"google_identity_platform_inbound_saml_config": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/oauthIdpConfigs/{{name}}
	"google_identity_platform_oauth_idp_config": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/tenants/{{name}}
	"google_identity_platform_tenant": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}
	"google_identity_platform_tenant_default_supported_idp_config": config.IdentifierFromProvider,
	// Imported by using the following: projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}
	"google_identity_platform_tenant_inbound_saml_config": config.IdentifierFromProvider,
	// Imported by using the following: projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}
	"google_identity_platform_tenant_oauth_idp_config": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/config/{{name}}
	"google_identity_platform_project_default_config": config.IdentifierFromProvider,

	// kms
	//
	// projects/{{project}}/locations/{{location}}/keyRings/{{name}}
	"google_kms_key_ring": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/keyRings/{{ .external_name }}"),
	// {{key_ring}}/cryptoKeys/{{name}}
	"google_kms_crypto_key": config.TemplatedStringAsIdentifier("name", "{{ .parameters.key_ring }}/cryptoKeys/{{ .external_name }}"),
	// {{name}}
	"google_kms_key_ring_import_job": config.TemplatedStringAsIdentifier("import_job_id", "{{ .parameters.key_ring }}/importJobs/{{ .external_name }}"),
	// terraform import google_kms_key_ring_iam_member.key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com"
	"google_kms_key_ring_iam_member": config.IdentifierFromProvider,
	// terraform import google_kms_crypto_key_iam_member.crypto_key "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com"
	"google_kms_crypto_key_iam_member": config.IdentifierFromProvider,
	// This resource does not support import.
	"google_kms_secret_ciphertext": config.IdentifierFromProvider,
	// Imported by using the following {{name}}
	"google_kms_crypto_key_version": config.IdentifierFromProvider,

	// monitoring
	//
	// Imported by using the following format: {{name}}
	"google_monitoring_alert_policy": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	"google_monitoring_notification_channel": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	"google_monitoring_uptime_check_config": config.IdentifierFromProvider,
	// Service can be imported using Name
	"google_monitoring_custom_service": config.IdentifierFromProvider,
	// Dashboard can be imported using dashboard_id
	"google_monitoring_dashboard": config.IdentifierFromProvider,
	// Group can be imported using Name
	"google_monitoring_group": config.IdentifierFromProvider,
	// MetricDescriptor can be imported using Name
	"google_monitoring_metric_descriptor": config.IdentifierFromProvider,
	// Slo can be imported using Name
	"google_monitoring_slo": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/services/{{service_id}}
	"google_monitoring_service": config.TemplatedStringAsIdentifier("service_id", "projects/{{ .setup.configuration.project }}/services/{{ .external_name }}"),

	// notebooks
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/environments/{{name}}
	"google_notebooks_environment": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/environments/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/instances/{{name}}
	"google_notebooks_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/instances/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/instances/{{name}}
	"google_notebooks_instance_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/runtimes/{{name}}
	"google_notebooks_runtime": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/runtimes/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/runtimes/{{name}}
	"google_notebooks_runtime_iam_member": config.IdentifierFromProvider,

	// networkmanagement
	//
	// ConnectivityTest can be imported using Name
	"google_network_management_connectivity_test": config.IdentifierFromProvider,

	// network
	//
	// Hub can be imported using Name
	"google_network_connectivity_hub": config.IdentifierFromProvider,
	// Spoke can be imported using {{location}}/{{name}}
	"google_network_connectivity_spoke": config.IdentifierFromProvider,

	// mlengine
	//
	// Model can be imported using Name
	"google_ml_engine_model": config.IdentifierFromProvider,

	// memcache
	//
	// nstance can be imported using Name
	"google_memcache_instance": config.IdentifierFromProvider,

	// osconfig
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/osPolicyAssignments/{{name}}
	"google_os_config_os_policy_assignment": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/osPolicyAssignments/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/patchDeployments/{{name}}
	"google_os_config_patch_deployment": config.TemplatedStringAsIdentifier("patch_deployment_id", "projects/{{ .setup.configuration.project }}/patchDeployments/{{ .external_name }}"),

	// oslogin
	//
	// Imported by using the following format: users/{{user}}/sshPublicKeys/{{fingerprint}}
	"google_os_login_ssh_public_key": config.IdentifierFromProvider,

	// privateca
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/caPools/{{name}}
	"google_privateca_ca_pool": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/caPools/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager user:jane@example.com
	"google_privateca_ca_pool_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificates/{{name}}
	"google_privateca_certificate": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/caPools/{{ .parameters.pool }}/certificates/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificateAuthorities/{{certificate_authority_id}}
	"google_privateca_certificate_authority": config.TemplatedStringAsIdentifier("certificate_authority_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/caPools/{{ .parameters.pool }}/certificateAuthorities/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}
	"google_privateca_certificate_template": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/certificateTemplates/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/certificateTemplates/{{certificate_template}} roles/privateca.templateUser user:jane@example.com
	"google_privateca_certificate_template_iam_member": config.IdentifierFromProvider,

	// pubsub
	//
	// Imported by using the following format: projects/{{project}}/topics/{{name}}
	"google_pubsub_topic": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/topics/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/topics/{{topic}} roles/viewer user:jane@example.com
	"google_pubsub_topic_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/subscriptions/{{name}}
	"google_pubsub_subscription": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/subscriptions/{{ .external_name }}"),
	// Imported by using the following format: projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com
	"google_pubsub_subscription_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/locations/{{zone}}/topics/{{name}}
	"google_pubsub_lite_topic": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.zone }}/topics/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}
	"google_pubsub_lite_subscription": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.zone }}/subscriptions/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/reservations/{{name}}
	"google_pubsub_lite_reservation": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/reservations/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/schemas/{{name}}
	"google_pubsub_schema": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/schemas/{{ .external_name }}"),

	// recaptcha
	//
	// google_recaptcha_enterprise_key.default projects/{{project}}/keys/{{name}}
	// This cannot be tested without elevated org level persmissions: https://github.com/upbound/official-providers/issues/715
	// "google_recaptcha_enterprise_key": config.IdentifierFromProvider,

	// redis
	//
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/instances/{{name}}
	"google_redis_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/instances/{{ .external_name }}"),

	// resource_manager
	//
	// google_resource_manager_lien.default {{parent}}/{{name}}
	// This cannot be tested without elevated org level persmissions: https://github.com/upbound/official-providers/issues/715
	// "google_resource_manager_lien": config.IdentifierFromProvider,

	// secretmanager
	//
	// Imported by using the following format: projects/{{project_id}}/secrets/{{secret_id}}
	"google_secret_manager_secret": config.TemplatedStringAsIdentifier("secret_id", "projects/{{ .setup.configuration.project }}/secrets/{{ .external_name }}"),
	// google_secret_manager_secret_iam_member.editor "projects/{{project}}/secrets/{{secret_id}} roles/secretmanager.secretAccessor user:jane@example.com"
	"google_secret_manager_secret_iam_member": config.IdentifierFromProvider,

	// Imported by using the following format: {{name}}/{{name}}
	"google_secret_manager_secret_version": config.IdentifierFromProvider,

	// service networking
	//
	// Imported by using the following format: /projects/{{project}}/global/networks/{{peering-network}}:{{service}}
	"google_service_networking_connection": config.IdentifierFromProvider,

	// scc
	//
	// google_scc_notification_config.default organizations/{{organization}}/notificationConfigs/{{name}}
	// This cannot be tested without elevated org level persmissions: https://github.com/upbound/official-providers/issues/715
	// "google_scc_notification_config": config.TemplatedStringAsIdentifier("config_id", "{{ .parameters.organization }}/notificationConfigs/{{ .external_name }}"),
	// google_scc_source.default organizations/{{organization}}/sources/{{name}}
	// This cannot be tested without elevated org level persmissions: https://github.com/upbound/official-providers/issues/715
	// "google_scc_source": config.TemplatedStringAsIdentifier("display_name", "{{ .parameters.organization }}/sources/{{ .external_name }}"),

	// sourcerepo
	//
	// google_sourcerepo_repository.default {{name}}
	"google_sourcerepo_repository": config.TemplatedStringAsIdentifier("name", "{{ .terraformProviderConfig.project }}/repos/{{ .external_name }}"),
	// google_sourcerepo_repository_iam_member.editor "projects/{{project}}/repos/{{repository}} roles/viewer user:jane@example.com"
	"google_sourcerepo_repository_iam_member": config.IdentifierFromProvider,

	// spanner
	//
	// google_spanner_database.default {{instance}}/{{name}}
	"google_spanner_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.instance }}/instance/{{ .external_name }}"),
	// google_spanner_instance.default {{project}}/{{name}}
	"google_spanner_instance": config.TemplatedStringAsIdentifier("name", "{{ .terraformProviderConfig.project }}/instance/{{ .external_name }}"),
	// google_spanner_instance_iam_member.instance "project-name/instance-name roles/viewer user:foo@example.com"
	"google_spanner_instance_iam_member": config.IdentifierFromProvider,
	// google_spanner_database_iam_member.database "project-name/instance-name/database-name roles/viewer user:foo@example.com"
	"google_spanner_database_iam_member": config.IdentifierFromProvider,

	// sql
	//
	// Imported by using the following format: projects/{{project}}/instances/{{name}}
	"google_sql_database_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/instances/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/instances/{{instance}}/databases/{{name}}
	"google_sql_database": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/instances/{{ .parameters.instance }}/databases/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/instances/{{name}}
	"google_sql_source_representation_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/instances/{{ .external_name }}"),
	// Imported by using the following format: my-project/main-instance/me
	"google_sql_user": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.instance }}/{{ .external_name }}"),
	// No import
	"google_sql_ssl_cert": config.IdentifierFromProvider,

	// storage
	//
	// Imported by using the following format: tf-test-project/image-store-bucket
	"google_storage_bucket": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .external_name }}"),
	// No import, configures bucket public access
	"google_storage_bucket_access_control": config.IdentifierFromProvider,
	// No import documented.
	"google_storage_bucket_acl": config.IdentifierFromProvider,
	// Imported by using the following format: b/{{bucket}} roles/storage.objectViewer user:jane@example.com
	"google_storage_bucket_iam_member": config.IdentifierFromProvider,
	// This resource does not support import.
	"google_storage_bucket_object": config.IdentifierFromProvider,
	// Imported using the following format: id which should be the format of {{bucket}}/{{entity}}
	"google_storage_default_object_access_control": config.IdentifierFromProvider,
	// No import documented.
	"google_storage_default_object_acl": config.IdentifierFromProvider,
	// Imported by using the following {{bucket}}/{{object}}/{{entity}}
	"google_storage_object_access_control": config.IdentifierFromProvider,
	// No Import
	"google_storage_object_acl": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/agentPools/{{name}}
	"google_storage_transfer_agent_pool": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/agentPools/{{ .external_name }}"),

	// bigquery
	//
	// Imported with the following format: projects/{{project}}/locations/{{location}}/connections/{{connection_id}}
	// However, connection_id is documented to be optional. Thus, we are using config.IdentifierFromProvider
	"google_bigquery_connection": config.IdentifierFromProvider,
	// Imported using ID
	"google_bigquery_data_transfer_config": config.IdentifierFromProvider,
	// Imported with the following format: projects/{{project}}/datasets/{{dataset_id}}
	"google_bigquery_dataset": config.TemplatedStringAsIdentifier("dataset_id", "projects/{{ .parameters.project }}/datasets/{{ .external_name }}"),
	// This resource does not support import
	"google_bigquery_dataset_access": config.IdentifierFromProvider,
	// Binding resource can be imported using the dataset_id and role: "projects/your-project-id/datasets/dataset-id roles/viewer"
	"google_bigquery_dataset_iam_binding": config.TemplatedStringAsIdentifier("", "{{ .parameters.dataset_id }} {{ .parameters.role }}"),
	// <ember resource can be imported using the dataset_id, role, and account: "projects/your-project-id/datasets/dataset-id roles/viewer user:foo@example.com"
	"google_bigquery_dataset_iam_member": config.TemplatedStringAsIdentifier("", "{{ .parameters.dataset_id }} {{ .parameters.role }} {{ .parameters.member }}"),
	// Policy resource can be imported using the dataset_id, role, and account: projects/your-project-id/datasets/dataset-id
	"google_bigquery_dataset_iam_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.dataset_id }}"),
	// TODO: check of the following location or locations
	// TODO: a `project` argument seems to be missing for this resource, to be checked after generation
	// Imported with the following format: projects/{{project}}/jobs/{{job_id}}/location/{{location}}
	"google_bigquery_job": config.IdentifierFromProvider,
	// Imported with the following format: projects/{{project}}/locations/{{location}}/reservations/{{name}}
	"google_bigquery_reservation": config.TemplatedStringAsIdentifier("name", "projects/{{ .parameters.project }}/locations/{{ .parameters.location }}/reservations/{{ .external_name }}"),
	// Imported with the following format: projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}
	"google_bigquery_reservation_assignment": config.TemplatedStringAsIdentifier("", "{{ .parameters.reservation }}/assignments/{{ .external_name }}"),
	// Imported with the following format: projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}
	"google_bigquery_routine": config.TemplatedStringAsIdentifier("routine_id", "{{ .parameters.dataset_id }}/routines/{{ .external_name }}"),
	// BigQuery tables can be imported using the project, dataset_id, and table_id: gcp-project/foo/bar
	"google_bigquery_table": config.TemplatedStringAsIdentifier("table_id", "{{ .parameters.project }}/{{ .parameters.dataset_id }}/{{ .external_name }}"),
	// IAM binding imports use space-delimited identifiers: the resource and the role: "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner"
	"google_bigquery_table_iam_binding": config.IdentifierFromProvider,
	// IAM member imports use space-delimited identifiers: the resource, the role, and the member identity: "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner user:jane@example.com"
	"google_bigquery_table_iam_member": config.TemplatedStringAsIdentifier("", "projects/{{ .parameters.project }}/datasets/{{ .parameters.dataset_id }}/tables/{{ .parameters.table_id }} {{ .parameters.member }}"),
	// IAM policy imports use the identifier of the resource: projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}
	"google_bigquery_table_iam_policy": config.IdentifierFromProvider,

	// dataflow
	//
	// This resource does not support import.
	"google_dataflow_job": config.IdentifierFromProvider,

	// datafusion
	//
	// projects/{{project}}/locations/{{region}}/instances/{{name}}
	"google_data_fusion_instance": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/instances/{{ .external_name }}"),

	// cloudbuild
	//
	// projects/{{project}}/locations/{{location}}/triggers/{{trigger_id}}
	"google_cloudbuild_trigger": config.IdentifierFromProvider,
	// projects/{{project}}/locations/{{location}}/workerPools/{{name}}
	"google_cloudbuild_worker_pool": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/workerPools/{{ .external_name }}"),

	// cloudiot
	//
	// {{registry}}/devices/{{name}}
	"google_cloudiot_device": config.TemplatedStringAsIdentifier("name", "{{ .parameters.registry }}/devices/{{ .external_name }}"),
	// {{project}}/locations/{{region}}/registries/{{name}}
	"google_cloudiot_registry": config.IdentifierFromProvider,

	// bigtable
	//
	// projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}
	"google_bigtable_app_profile": config.TemplatedStringAsIdentifier("app_profile_id", "projects/{{ .setup.configuration.project }}/instances/{{ .parameters.instance }}/appProfiles/{{ .external_name }}"),
	// No import
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

	// apigee
	//
	// Imported by using the following format: {{org_id}}/envgroups/{{name}}
	"google_apigee_envgroup": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org_id }}/envgroups/{{ .external_name }}"),
	// Imported by using the following format: {{org_id}}/environments/{{name}}
	"google_apigee_environment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org_id }}/environments/{{ .external_name }}"),
	// Imported by using the following format: {{org_id}}/instances/{{name}}
	"google_apigee_instance": config.TemplatedStringAsIdentifier("name", "{{ .parameters.org_id }}/instances/{{ .external_name }}"),
	// Imported by using the following format: organizations/{{name}}
	"google_apigee_organization": config.IdentifierFromProvider,
	// Imported by using the following format: {{instance_id}}/natAddresses/{{name}}
	"google_apigee_nat_address": config.TemplatedStringAsIdentifier("name", "{{ .parameters.instance_id }}/natAddresses/{{ .external_name }}"),
	// Imported by using the following format: {{org_id}}/environments/{{environment}} roles/viewer user:jane@example.com
	"google_apigee_environment_iam_member": config.IdentifierFromProvider,

	// binaryauthorization
	//
	// projects/{{project}}/attestors/{{name}}
	"google_binary_authorization_attestor": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/attestors/{{ .external_name }}"),
	// projects/{{project}}
	"google_binary_authorization_policy": config.TemplatedStringAsIdentifier("", "projects/{{ .setup.configuration.project }}"),

	// dataproc
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
	"google_dataproc_autoscaling_policy": config.TemplatedStringAsIdentifier("policy_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/autoscalingPolicies/{{ .external_name }}"),
	// No Import
	"google_dataproc_cluster": config.IdentifierFromProvider,
	// No Import
	"google_dataproc_job": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}
	"google_dataproc_workflow_template": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/workflowTemplates/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/services/{{service_id}}
	"google_dataproc_metastore_service": config.TemplatedStringAsIdentifier("service_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/services/{{ .external_name }}"),

	// iam
	//
	// Imported by using the following projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}
	"google_iam_workload_identity_pool": config.TemplatedStringAsIdentifier("workload_identity_pool_id", "projects/{{ .setup.configuration.project }}/locations/global/workloadIdentityPools/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/providers/{{workload_identity_pool_provider_id}}
	"google_iam_workload_identity_pool_provider": config.TemplatedStringAsIdentifier("workload_identity_pool_provider_id", "projects/{{ .setup.configuration.project }}/locations/global/workloadIdentityPools/{{ .parameters.workload_identity_pool_id }}/providers/{{ .external_name }}"),

	// datalossprevention
	//
	// Imported by using the following format: {{parent}}/deidentifyTemplates/{{name}}
	"google_data_loss_prevention_deidentify_template": config.IdentifierFromProvider,
	// Imported by using the following format: {{parent}}/inspectTemplates/{{name}}
	"google_data_loss_prevention_inspect_template": config.IdentifierFromProvider,
	// Imported by using the following format: {{parent}}/jobTriggers/{{name}}
	"google_data_loss_prevention_job_trigger": config.IdentifierFromProvider,
	// Imported by using the following format: {{parent}}/storedInfoTypes/{{name}}
	"google_data_loss_prevention_stored_info_type": config.IdentifierFromProvider,

	// logging
	//
	// LogView can be imported using any of these accepted formats: {{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}
	"google_logging_log_view": config.TemplatedStringAsIdentifier("name", "{{ .parameters.parent }}/locations/{{ .parameters.location }}/buckets/{{ .parameters.bucket }}/views/{{ .external_name }}"),
	// Metric can be imported using Name
	"google_logging_metric": config.NameAsIdentifier,
	// This resource can be imported using the following format: projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
	"google_logging_project_bucket_config": config.TemplatedStringAsIdentifier("", "projects/{{ .parameters.project }}/locations/{{ .parameters.location }}/buckets/{{ .parameters.bucket_id }}"),
	// Project-level logging exclusions can be imported using their URI
	// projects/my-project/exclusions/my-exclusion
	"google_logging_project_exclusion": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/exclusions/{{ .external_name }}"),
	// Project-level logging sinks can be imported using their URI
	// projects/my-project/sinks/my-sink
	"google_logging_project_sink": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/sinks/{{ .external_name }}"),

	// vertexai
	//
	// No Import
	"google_vertex_ai_dataset": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/locations/{{region}}/featurestores/{{name}}
	"google_vertex_ai_featurestore": config.IdentifierFromProvider,
	// Imported by using the following {{featurestore}}/entityTypes/{{name}}
	"google_vertex_ai_featurestore_entitytype": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/locations/{{region}}/tensorboards/{{name}}
	"google_vertex_ai_tensorboard": config.TemplatedStringAsIdentifier("display_name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/tensorboards/{{ .external_name }}"),

	// documentai
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/processors/{{name}}
	"google_document_ai_processor": config.IdentifierFromProvider,

	// artifactregistry
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
	"google_artifact_registry_repository": config.TemplatedStringAsIdentifier("repository_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/repositories/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/artifactregistry.reader user:jane@example.com
	"google_artifact_registry_repository_iam_member": config.IdentifierFromProvider,

	// beyondcorp
	//
	// Imported by using the following projects/{{project}}/locations/{{region}}/appConnections/{{name}}
	"google_beyondcorp_app_connection": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/locations/{{region}}/appConnectors/{{name}}
	"google_beyondcorp_app_connector": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/appConnectors/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{region}}/appGateways/{{name}}
	"google_beyondcorp_app_gateway": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.region }}/appGateways/{{ .external_name }}"),

	// bigqueryanalyticshub
	//
	// Imported by using the following projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}
	"google_bigquery_analytics_hub_data_exchange": config.IdentifierFromProvider,
	// Imported by using the following projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}
	"google_bigquery_analytics_hub_listing": config.TemplatedStringAsIdentifier("listing_id", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/dataExchanges/{{ .parameters.data_exchange_id }}/listings/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}} roles/viewer user:jane@example.com
	"google_bigquery_analytics_hub_data_exchange_iam_member": config.IdentifierFromProvider,

	// tpu
	//
	// Imported by using the following projects/{{project}}/locations/{{zone}}/nodes/{{name}}
	"google_tpu_node": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.zone }}/nodes/{{ .external_name }}"),

	// workflows
	//
	// No import
	"google_workflows_workflow": config.IdentifierFromProvider,

	// certificatemanager
	//
	// projects/{{project}}/locations/global/certificates/{{name}}
	"google_certificate_manager_certificate": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/certificates/{{ .external_name }}"),
	// projects/{{project}}/locations/global/dnsAuthorizations/{{name}}
	"google_certificate_manager_dns_authorization": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/dnsAuthorizations/{{ .external_name }}"),
	// Imported by using the following projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}
	"google_certificate_manager_certificate_map_entry": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/global/certificateMaps/{{ .parameters.map }}/certificateMapEntries/{{ .external_name }}"),
}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}

func externalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := externalNameConfigs[r.Name]; ok {
			r.Version = VersionV1Beta1
			r.ExternalName = e
		}
	}
}
