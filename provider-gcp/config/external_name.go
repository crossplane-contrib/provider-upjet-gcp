/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-gcp/config/common"
)

var externalNameConfigs = map[string]config.ExternalName{
	// appengine
	//
	// Imported by using the following format: your-project-id
	"google_app_engine_application": config.IdentifierFromProvider,

	// composer
	//
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/environments/{{name}}
	"google_composer_environment": formattedIdentifierUserDefined("projects/%s/locations/%s/environments", "project", "region"),

	// cloudfunctions
	//
	// Imported by using the following format: {{project}}/{{region}}/function-test
	"google_cloudfunctions_function": formattedIdentifierUserDefined("%s/%s/", "project", "region"),
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/functions/{{cloud_function}} roles/viewer user:jane@example.com
	"google_cloudfunctions_function_iam_member": config.IdentifierFromProvider,

	// cloudplatform
	//
	// Folders can be imported using the folder's id, e.g. folders/1234567
	"google_folder": config.IdentifierFromProvider,
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
	"google_project": formattedIdentifierWithResourcePrefix("projects"),
	// Resource with format projects/{{project}}
	// This resource does not support import
	"google_project_default_service_accounts": formattedIdentifierWithResourcePrefix("projects"),
	// Imported by using the following format: your-project-id roles/viewer user:foo@example.com
	"google_project_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: your-project-id foo.googleapis.com
	"google_project_iam_audit_config": config.IdentifierFromProvider,
	// Imported by using the following format: your-project-id/iam.googleapis.com
	"google_project_service": config.IdentifierFromProvider,
	// Imported by using the following format: {{project}}
	"google_project_usage_export_bucket": config.IdentifierFromProvider,
	// Service accounts can be imported using their URI, e.g. projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com
	"google_service_account": googleServiceAccount(),
	// Imported by using the following format: projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31
	"google_service_account_iam_member": config.IdentifierFromProvider,
	// No import
	"google_service_account_key": config.IdentifierFromProvider,
	// Imported by using the following format: services/{service}/projects/{project}/global/networks/{network}/peeredDnsDomains/{name}
	"google_service_networking_peered_dns_domain": formattedIdentifierUserDefined("services/%s/projects/%s/global/networks/%s/peeredDnsDomains", "service", "project", "network"),

	// cloudrun
	//
	// Imported by using the following format: locations/{{location}}/namespaces/{{project}}/domainmappings/{{name}}
	"google_cloud_run_domain_mapping": config.IdentifierFromProvider,
	// Imported by using the following format: locations/{{location}}/namespaces/{{project}}/services/{{name}}
	"google_cloud_run_service": formattedIdentifierUserDefined("locations/%s/namespaces/%s/services", "location", "project"),
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer user:jane@example.com
	"google_cloud_run_service_iam_member": config.IdentifierFromProvider,

	// cloudscheduler
	//
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/jobs/{{name}}
	"google_cloud_scheduler_job": formattedIdentifierUserDefined("projects/%s/locations/%s/jobs", "project", "region"),

	// cloudtasks
	//
	// Imported by using the following format: projects/{{project}}/locations/{{location}}/queues/{{name}}
	"google_cloud_tasks_queue": formattedIdentifierUserDefined("projects/%s/locations/%s/queues", "project", "location"),

	// compute
	//
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instances/{{instance.name}}/{{disk.name}}
	"google_compute_attached_disk": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}
	"google_compute_autoscaler": formattedIdentifierUserDefined("projects/%s/zones/{{zone}}/autoscalers", "project", "zone"),
	// Imported by using the following format: projects/{{project}}/global/backendBuckets/{{name}}
	"google_compute_backend_bucket": formattedIdentifierUserDefined("projects/%s/global/backendBuckets", "project"),
	// Imported by using the following format: This resource does not support import.
	"google_compute_backend_bucket_signed_url_key": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/backendServices/{{name}}
	"google_compute_backend_service": formattedIdentifierUserDefined("projects/%s/global/backendServices", "project"),
	// Imported by using the following format: projects/{{project}}/global/sslCertificates/{{name}}
	"google_compute_managed_ssl_certificate": formattedIdentifierUserDefined("projects/%s/global/sslCertificates", "project"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/subnetworks/{{name}}
	"google_compute_subnetwork": formattedIdentifierUserDefined("projects/%s/regions/%s/subnetworks", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/addresses/{{name}}
	"google_compute_address": formattedIdentifierUserDefined("projects/%s/regions/%s/addresses", "project", "region"),
	// Imported by using the following format: projects/{{project}}/global/firewalls/{{name}}
	"google_compute_firewall": formattedIdentifierUserDefined("projects/%s/global/firewalls", "project"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/routers/{{name}}
	"google_compute_router": formattedIdentifierUserDefined("projects/%s/regions/%s/routers", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}
	"google_compute_router_nat": formattedIdentifierUserDefined("projects/%s/regions/%s/routers/%s", "project", "region", "router"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instances/{{name}}
	"google_compute_instance": formattedIdentifierUserDefined("projects/%s/zones/%s/instances", "project", "zone"),
	// Imported by using the following format: projects/{{project}}/global/networks/{{name}}
	"google_compute_network": formattedIdentifierUserDefined("projects/%s/global/networks", "project"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/disks/{{name}}
	"google_compute_disk": formattedIdentifierUserDefined("projects/%s/zones/%s/disks", "project", "zone"),
	// Imported by using the following format: projects/{{project}}/global/externalVpnGateways/{{name}}
	"google_compute_external_vpn_gateway": formattedIdentifierUserDefined("projects/%s/global/externalVpnGateways", "project"),
	// Imported by using the following format: projects/{{project}}/global/addresses/{{name}}
	"google_compute_global_address": formattedIdentifierUserDefined("projects/%s/global/addresses", "project"),
	// Imported by using the following format: projects/{{project}}/global/networkEndpointGroups/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
	"google_compute_global_network_endpoint_group": formattedIdentifierUserDefined("projects/%s/global/networkEndpointGroups", "project"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}
	"google_compute_ha_vpn_gateway": formattedIdentifierUserDefined("projects/%s/regions/%s/vpnGateways", "project", "region"),
	// Imported by using the following format: projects/{{project}}/global/healthChecks/{{name}}
	"google_compute_health_check": formattedIdentifierUserDefined("projects/%s/global/healthChecks", "project"),
	// Imported by using the following format: projects/{{project}}/global/httpHealthChecks/{{name}}
	"google_compute_http_health_check": formattedIdentifierUserDefined("projects/%s/global/httpHealthChecks", "project"),
	// Imported by using the following format: projects/{{project}}/global/httpsHealthChecks/{{name}}
	"google_compute_https_health_check": formattedIdentifierUserDefined("projects/%s/global/httpsHealthChecks", "project"),
	// Imported by using the following format: projects/{{project}}/global/images/{{name}}
	"google_compute_image": formattedIdentifierUserDefined("projects/%s/global/images", "project"),
	// Imported by using the following format: projects/{{project}/zones/{{zone}}/instanceGroups/{{name}}
	"google_compute_instance_group": formattedIdentifierUserDefined("projects/%s/zones/%s/instanceGroups", "project", "zone"),
	// Imported by using the following format: projects/{{project}}/global/instanceTemplates/{{name}}
	"google_compute_instance_template": config.IdentifierFromProvider,
	// No import
	"google_compute_instance_from_template": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}
	"google_compute_resource_policy": formattedIdentifierUserDefined("projects/%s/regions/%s/resourcePolicies", "project", "region"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/disks/{{disk}}/{{name}}
	"google_compute_disk_resource_policy_attachment": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/disks/{{disk}}
	"google_compute_disk_iam_policy": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer
	"google_compute_disk_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/disks/{{disk}} roles/viewer user:jane@example.com
	"google_compute_disk_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: {{project}}/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
	"google_compute_global_network_endpoint": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/images/{{image}}
	"google_compute_image_iam_policy": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/images/{{image}} roles/compute.imageUser
	"google_compute_image_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/images/{{image}} roles/compute.imageUser user:jane@example.com
	"google_compute_image_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/targetPools/{{name}}
	"google_compute_target_pool": formattedIdentifierUserDefined("projects/%s/regions/%s/targetPools", "project", "region"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}
	"google_compute_instance_group_manager": formattedIdentifierUserDefined("projects/%s/zones/%s/targetPools", "project", "zone"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instances/{{instance}}
	"google_compute_instance_iam_policy": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin
	"google_compute_instance_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin user:jane@example.com
	"google_compute_instance_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}
	"google_compute_interconnect_attachment": formattedIdentifierUserDefined("projects/%s/regions/%s/interconnectAttachments", "project", "region"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}
	"google_compute_network_endpoint_group": formattedIdentifierUserDefined("projects/%s/zones/%s/networkEndpointGroups", "project", "zone"),
	// Imported by using the following format: {{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
	"google_compute_network_endpoint": config.IdentifierFromProvider,
	// Imported by using the following format: project-name/network-name/peering-name
	"google_compute_network_peering": formattedIdentifierWithResourcePrefix("network"),
	// Imported by using the following format: projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}
	"google_compute_network_peering_routes_config": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}
	"google_compute_node_group": formattedIdentifierUserDefined("projects/%s/zones/%s/nodeGroups", "project", "zone"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}
	"google_compute_node_template": formattedIdentifierUserDefined("projects/%s/regions/%s/nodeTemplates", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}
	"google_compute_packet_mirroring": formattedIdentifierUserDefined("projects/%s/regions/%s/packetMirrorings", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}
	"google_compute_forwarding_rule": formattedIdentifierUserDefined("projects/%s/regions/%s/forwardingRules", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/backendServices/{{name}}
	"google_compute_region_backend_service": formattedIdentifierUserDefined("projects/%s/regions/%s/backendServices", "project", "region"),
	// Imported by using the following format: {{name}}
	"google_compute_region_instance_group_manager": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/targetHttpProxies/{{name}}
	"google_compute_region_target_http_proxy": formattedIdentifierUserDefined("projects/%s/regions/%s/targetHttpProxies", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/urlMaps/{{name}}
	"google_compute_region_url_map": formattedIdentifierUserDefined("projects/%s/regions/%s/urlMaps", "project", "region"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/{{name}}
	"google_compute_per_instance_config": config.IdentifierFromProvider,
	// Projects can be imported using the Project ID: your-project-id
	"google_compute_project_default_network_tier": config.IdentifierFromProvider,
	// Projects can be imported using the Project ID: your-project-id
	"google_compute_project_metadata": config.IdentifierFromProvider,
	// Project metadata items can be imported using the key: key
	"google_compute_project_metadata_item": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/autoscalers/{{name}}
	"google_compute_region_autoscaler": formattedIdentifierUserDefined("projects/%s/regions/%s/autoscalers", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/disks/{{name}}
	"google_compute_region_disk": formattedIdentifierUserDefined("projects/%s/regions/%s/disks", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/disks/{{region_disk}} roles/viewer user:jane@example.com
	"google_compute_region_disk_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/disks/{{disk}}/{{name}}
	"google_compute_region_disk_resource_policy_attachment": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/healthChecks/{{name}}
	"google_compute_region_health_check": formattedIdentifierUserDefined("projects/%s/regions/%s/healthChecks", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}
	"google_compute_region_network_endpoint_group": formattedIdentifierUserDefined("projects/%s/regions/%s/networkEndpointGroups", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{region_instance_group_manager}}/{{name}}
	"google_compute_region_per_instance_config": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/sslCertificates/{{name}}
	"google_compute_region_ssl_certificate": formattedIdentifierUserDefined("projects/%s/regions/%s/sslCertificates", "project", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}
	"google_compute_region_target_https_proxy": formattedIdentifierUserDefined("projects/%s/regions/%s/targetHttpsProxies", "project", "region"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/reservations/{{name}}
	"google_compute_reservation": formattedIdentifierUserDefined("projects/%s/zones/%s/reservations", "project", "zone"),
	// Imported by using the following format: projects/{{project}}/global/routes/{{name}}
	"google_compute_route": formattedIdentifierUserDefined("projects/%s/global/routes", "project"),
	// Imported by using the following format: us-central1/router-1/interface-1
	"google_compute_router_interface": config.IdentifierFromProvider,
	// Imported by using the following format: locations/global/firewallPolicies/{{name}}
	"google_compute_firewall_policy": config.IdentifierFromProvider,
	// Imported by using the following format: locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
	"google_compute_firewall_policy_association": config.IdentifierFromProvider,
	// Imported by using the following format: locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
	"google_compute_firewall_policy_rule": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/forwardingRules/{{name}}
	"google_compute_global_forwarding_rule": formattedIdentifierUserDefined("projects/%s/global/forwardingRules", "project"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/{{port}}/{{name}}
	"google_compute_instance_group_named_port": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/global/targetSslProxies/{{name}}
	"google_compute_target_ssl_proxy": formattedIdentifierUserDefined("projects/%s/global/targetSslProxies", "project"),
	// Imported by using the following format: projects/{{project}}/global/sslCertificates/{{name}}
	"google_compute_ssl_certificate": formattedIdentifierUserDefined("projects/%s/global/sslCertificates", "project"),

	// container
	//
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/clusters/my-cluster
	"google_container_cluster": formattedIdentifierUserDefined("projects/%s/locations/%s/clusters", "project", "location"),
	// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
	"google_container_node_pool": containerNodePool(),

	// dns
	//
	// Imported by using the following format: projects/{{project}}/managedZones/{{name}}
	"google_dns_managed_zone": formattedIdentifierUserDefined("projects/%s/managedZones/%s", "project"),
	// Imported by using the following format: projects/{{project}}/policies/{{name}}
	"google_dns_policy": formattedIdentifierUserDefined("projects/%s/policies/%s", "project"),
	// Imported by using the following format: projects/{{project}}/managedZones/{{zone}}/rrsets/{{name}}/{{type}}
	"google_dns_record_set": config.IdentifierFromProvider,

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

	// kms
	//
	// projects/{{project}}/locations/{{location}}/keyRings/{{name}}
	"google_kms_key_ring": formattedIdentifierUserDefined("projects/%s/locations/%s/keyRings", "project", "location"),
	// {{key_ring}}/cryptoKeys/{{name}}
	"google_kms_crypto_key": formattedIdentifierUserDefined("%s/cryptoKeys", "key_ring"),
	// {{name}}
	"google_kms_key_ring_import_job": config.TemplatedStringAsIdentifier("import_job_id", "{{ .parameters.key_ring }}/importJobs/{{ .externalName }}"),
	// terraform import google_kms_key_ring_iam_member.key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com"
	"google_kms_key_ring_iam_member": config.IdentifierFromProvider,
	// terraform import google_kms_crypto_key_iam_member.crypto_key "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com"
	"google_kms_crypto_key_iam_member": config.IdentifierFromProvider,
	// This resource does not support import.
	"google_kms_secret_ciphertext": config.IdentifierFromProvider,

	// monitoring
	//
	// Imported by using the following format: {{name}}
	"google_monitoring_alert_policy": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	"google_monitoring_notification_channel": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}
	"google_monitoring_uptime_check_config": config.IdentifierFromProvider,

	// pubsub
	//
	// Imported by using the following format: projects/{{project}}/topics/{{name}}
	"google_pubsub_topic": formattedIdentifierUserDefined("projects/%s/topics", "project"),
	// Imported by using the following format: projects/{{project}}/topics/{{topic}} roles/viewer user:jane@example.com
	"google_pubsub_topic_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/subscriptions/{{name}}
	"google_pubsub_subscription": formattedIdentifierUserDefined("projects/%s/subscriptions", "project"),
	// Imported by using the following format: projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com
	"google_pubsub_subscription_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/locations/{{zone}}/topics/{{name}}
	"google_pubsub_lite_topic": formattedIdentifierUserDefined("projects/%s/locations/%s/topics", "project", "zone"),
	// Imported by using the following format: projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}
	"google_pubsub_lite_subscription": formattedIdentifierUserDefined("projects/%s/locations/%s/subscriptions", "project", "zone"),
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/reservations/{{name}}
	"google_pubsub_lite_reservation": formattedIdentifierUserDefined("projects/%s/locations/%s/reservations", "project", "region"),
	// Imported by using the following format: projects/{{project}}/schemas/{{name}}
	"google_pubsub_schema": formattedIdentifierUserDefined("projects/%s/schemas", "project"),

	// redis
	//
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/instances/{{name}}
	"google_redis_instance": formattedIdentifierUserDefined("projects/%s/locations/%s/instances", "project", "region"),

	// secretmanager
	//
	// Imported by using the following format: projects/{{project_id}}/secrets/{{secret_id}}
	"google_secret_manager_secret": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/secrets/{{secret_id}} roles/secretmanager.secretAccessor user:jane@example.com
	"google_secret_manager_secret_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: {{name}}/{{name}}
	"google_secret_manager_secret_version": config.IdentifierFromProvider,

	// service networking
	//
	// Imported by using the following format: /projects/{{project}}/global/networks/{{peering-network}}:{{service}}
	"google_service_networking_connection": config.IdentifierFromProvider,

	// sql
	//
	// Imported by using the following format: projects/{{project}}/instances/{{name}}
	"google_sql_database_instance": formattedIdentifierUserDefined("projects/%s/instances", "project"),
	// Imported by using the following format: projects/{{project}}/instances/{{instance}}/databases/{{name}}
	"google_sql_database": formattedIdentifierUserDefined("projects/%s/instances/%s/databases", "project", "instance"),
	// Imported by using the following format: projects/{{project}}/instances/{{name}}
	"google_sql_source_representation_instance": formattedIdentifierUserDefined("projects/%s/instances", "project"),
	// Imported by using the following format: my-project/main-instance/me
	"google_sql_user": formattedIdentifierUserDefined("%s/%s", "project", "instance"),
	// No import
	"google_sql_ssl_cert": config.IdentifierFromProvider,

	// storage
	//
	// Imported by using the following format: tf-test-project/image-store-bucket
	"google_storage_bucket": formattedIdentifierUserDefined("%s", "project"),
	// No import, configures bucket public access
	"google_storage_bucket_access_control": config.IdentifierFromProvider,
	// This resource does not support import.
	"google_storage_bucket_object": config.IdentifierFromProvider,
}

// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
func containerNodePool() config.ExternalName {
	e := config.NameAsIdentifier
	e.GetExternalNameFn = common.GetNameFromFullyQualifiedID
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		project, err := common.GetField(providerConfig, common.KeyProject)
		if err != nil {
			return "", err
		}
		clusterID, err := common.GetField(parameters, "cluster")
		if err != nil {
			return "", err
		}
		if len(strings.Split(clusterID, "/")) != 6 {
			return "", fmt.Errorf("the clusterID is not in expected format")
		}
		location := strings.Split(clusterID, "/")[3]
		cluster := strings.Split(clusterID, "/")[5]
		return fmt.Sprintf("%s/%s/%s/%s", project, location, cluster, externalName), nil
	}
	return e
}

func googleServiceAccount() config.ExternalName {
	e := config.ParameterAsIdentifier("account_id")
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		return common.GetField(tfstate, "account_id")
	}
	e.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		project, err := common.GetField(providerConfig, common.KeyProject)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", project, externalName, project), nil
	}
	return e
}

func formattedIdentifierWithResourcePrefix(resourcePrefix string) config.ExternalName {
	e := config.NameAsIdentifier
	e.GetExternalNameFn = common.GetNameFromFullyQualifiedID
	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]interface{}, _ map[string]interface{}) (string, error) {
		return fmt.Sprintf("%s/%s", resourcePrefix, externalName), nil
	}
	return e
}

func formattedIdentifierUserDefined(format string, paramNames ...string) config.ExternalName {
	e := config.NameAsIdentifier
	e.GetExternalNameFn = common.GetNameFromFullyQualifiedID
	e.GetIDFn = setUserDefinedGetIDFn(format, paramNames...)
	return e
}

func setUserDefinedGetIDFn(format string, paramNames ...string) config.GetIDFn {
	return func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		var params []interface{}
		for _, paramName := range paramNames {
			var param string
			var err error
			switch paramName {
			case common.KeyProject:
				param, err = common.GetField(providerConfig, paramName)
			default:
				param, err = common.GetField(parameters, paramName)
			}
			if err != nil {
				return "", err
			}
			params = append(params, param)
		}
		return strings.Join([]string{fmt.Sprintf(format, params...), externalName}, "/"), nil
	}
}

func externalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := externalNameConfigs[r.Name]; ok {
			r.Version = VersionV1Beta1
			r.ExternalName = e
		}
	}
}
