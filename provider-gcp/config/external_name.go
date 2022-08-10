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
	// cloudplatform
	//
	// Folders can be imported using the folder's id, e.g. folders/1234567
	"google_folder": config.IdentifierFromProvider,
	// Projects can be imported using the project_id: your-project-id
	// Project-ID has a format as following: projects/{{project}}
	// So, the GetIDFn function implementation for project-id and import method
	// for project resource seems that different.
	"google_project": formattedIdentifierWithResourcePrefix("projects"),
	// Imported by using the following format: your-project-id
	"google_project_iam_policy": config.IdentifierFromProvider,
	// Imported by using the following format: your-project-id roles/viewer
	"google_project_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: your-project-id roles/viewer user:foo@example.com
	"google_project_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: your-project-id foo.googleapis.com
	"google_project_iam_audit_config": config.IdentifierFromProvider,
	// Service accounts can be imported using their URI, e.g. projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com
	"google_service_account": googleServiceAccount(),
	// Imported by using the following format: projects/{your-project-id}/serviceAccounts/{your-service-account-email}
	"google_service_account_iam_policy": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser expires_after_2019_12_31
	"google_service_account_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31
	"google_service_account_iam_member": config.IdentifierFromProvider,
	// No import
	"google_service_account_key": config.IdentifierFromProvider,

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
	// Imported by using the following format: projects/{{project}}/topics/{{topic}}
	"google_pubsub_topic_iam_policy": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/topics/{{topic}} roles/viewer
	"google_pubsub_topic_iam_binding": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/topics/{{topic}} roles/viewer user:jane@example.com
	"google_pubsub_topic_iam_member": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{{project}}/subscriptions/{{name}}
	"google_pubsub_subscription": formattedIdentifierUserDefined("projects/%s/subscriptions", "project"),
	// Imported by using the following format: projects/{your-project-id}/subscriptions/{your-subscription-name}
	"google_pubsub_subscription_iam_policy": config.IdentifierFromProvider,
	// Imported by using the following format: projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor
	"google_pubsub_subscription_iam_binding": config.IdentifierFromProvider,
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
	// "google_secret_manager_secret": googleSecretManagerSecret(),

	// Imported by using the following format: {{name}}/{{name}}
	"google_secret_manager_secret_version": config.IdentifierFromProvider,
	// "google_secret_manager_secret_version": config.TemplatedStringAsIdentifier("name", "{{ .externalName }}/{{ .externalName }}"),

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
}

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
