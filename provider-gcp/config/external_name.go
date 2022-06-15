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
	"google_project": formattedIdentifierWithResourcePrefix("projects"),
	// Service accounts can be imported using their URI, e.g. projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com
	"google_service_account": googleServiceAccount(),

	// compute
	//
	// Imported by using the following format: projects/{{project}}/global/sslCertificates/{{name}}
	"google_compute_managed_ssl_certificate": formattedIdentifierUserDefined("projects/%s/global/sslCertificates/%s"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/subnetworks/{{name}}
	"google_compute_subnetwork": formattedIdentifierUserDefined("projects/%s/regions/%s/subnetworks/%s", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/addresses/{{name}}
	"google_compute_address": formattedIdentifierUserDefined("projects/%s/regions/%s/addresses/%s", "region"),
	// Imported by using the following format: projects/{{project}}/global/firewalls/{{name}}
	"google_compute_firewall": formattedIdentifierUserDefined("projects/%s/global/firewalls/%s"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/routers/{{name}}
	"google_compute_router": formattedIdentifierUserDefined("projects/%s/regions/%s/routers/%s", "region"),
	// Imported by using the following format: projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}
	"google_compute_router_nat": formattedIdentifierUserDefined("projects/%s/regions/%s/routers/%s/%s", "region", "router"),
	// Imported by using the following format: projects/{{project}}/zones/{{zone}}/instances/{{name}}
	"google_compute_instance": formattedIdentifierUserDefined("projects/%s/zones/%s/instances/%s", "zone"),
	// Imported by using the following format: projects/{{project}}/global/networks/{{name}}
	"google_compute_network": formattedIdentifierUserDefined("projects/%s/global/networks/%s"),

	// container
	//
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/clusters/my-cluster
	"google_container_cluster": formattedIdentifierUserDefined("projects/%s/locations/%s/clusters/%s", "location"),
	// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
	"google_container_node_pool": containerNodePool(),

	// redis
	//
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/instances/{{name}}
	"google_redis_instance": formattedIdentifierUserDefined("projects/%s/locations/%s/instances/%s", "region"),

	// sql
	//
	// Imported by using the following format: projects/{{project}}/instances/{{name}}
	"google_sql_database_instance": formattedIdentifierUserDefined("projects/%s/instances/%s"),
	// Imported by using the following format: projects/{{project}}/instances/{{instance}}/databases/{{name}}
	"google_sql_database": formattedIdentifierUserDefined("projects/%s/instances/%s/databases/%s", "instance"),
	// Imported by using the following format: projects/{{project}}/instances/{{name}}
	"google_sql_source_representation_instance": formattedIdentifierUserDefined("projects/%s/instances/%s/databases/%s", "instance"),
	// Imported by using the following format: my-project/main-instance/me
	"google_sql_user": formattedIdentifierUserDefined("%s/%s/%s", "instance"),
	// No import
	"google_sql_ssl_cert": formattedIdentifierFromProvider("projects/%s/instances/%s/sslCerts/%s", "instance"),

	// storage
	//
	// Imported by using the following format: tf-test-project/image-store-bucket
	"google_storage_bucket": formattedIdentifierUserDefined("%s/%s"),
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
		location := strings.Split(clusterID, "/")[3]
		cluster := strings.Split(clusterID, "/")[5]
		return fmt.Sprintf("%s/%s/%s/%s", project, location, cluster, externalName), nil
	}
	return e
}

func googleServiceAccount() config.ExternalName {
	e := parameterAsExternalName("account_id")
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, err := common.GetField(tfstate, "account_id")
		if err != nil {
			return "", err
		}
		return id, nil
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

func formattedIdentifierFromProvider(format string, paramNames ...string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = common.GetNameFromFullyQualifiedID
	e.GetIDFn = setUserDefinedGetIDFn(format, paramNames...)
	return e
}

func formattedIdentifierUserDefined(format string, paramNames ...string) config.ExternalName {
	e := config.NameAsIdentifier
	e.GetExternalNameFn = common.GetNameFromFullyQualifiedID
	e.GetIDFn = setUserDefinedGetIDFn(format, paramNames...)
	return e
}

func parameterAsExternalName(paramName string) config.ExternalName {
	e := config.NameAsIdentifier
	e.SetIdentifierArgumentFn = func(base map[string]interface{}, externalName string) {
		base[paramName] = externalName
	}
	e.OmittedFields = []string{paramName}
	return e
}

func setUserDefinedGetIDFn(format string, paramNames ...string) config.GetIDFn {
	return func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		project, err := common.GetField(providerConfig, common.KeyProject)
		if err != nil {
			return "", err
		}
		params := []interface{}{project}
		for _, paramName := range paramNames {
			param, err := common.GetField(parameters, paramName)
			if err != nil {
				return "", err
			}
			params = append(params, param)
		}
		params = append(params, externalName)
		return fmt.Sprintf(format, params...), nil
	}
}

func externalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := externalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}
