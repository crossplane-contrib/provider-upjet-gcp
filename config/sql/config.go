// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package sql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/crossplane/upjet/pkg/config"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

	"github.com/upbound/provider-gcp/config/common"
)

// CloudSQL connection detail keys
const (
	CloudSQLSecretServerCACertificateCertKey            = "serverCACertificateCert"
	CloudSQLSecretServerCACertificateCommonNameKey      = "serverCACertificateCommonName"
	CloudSQLSecretServerCACertificateCreateTimeKey      = "serverCACertificateCreateTime"
	CloudSQLSecretServerCACertificateExpirationTimeKey  = "serverCACertificateExpirationTime"
	CloudSQLSecretServerCACertificateSha1FingerprintKey = "serverCACertificateSha1Fingerprint"

	CloudSQLSecretConnectionName = "connectionName"

	PrivateIPKey = "privateIP"
	PublicIPKey  = "publicIP"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("google_sql_database_instance", func(r *config.Resource) {
		// we need to workaround the newly added (with v4.56.0)
		// of `Optional=True` in the native provider
		config.MoveToStatus(r.TerraformResource, "instance_type")
		// NOTE(@tnthornton) most of the connection details that were exported
		// to the connection details secret are marked as non-sensitive for tf.
		// We need to manually construct the secret details for those items.
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"maintenance_version"},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["connection_name"].(string); ok {
				conn[CloudSQLSecretConnectionName] = []byte(a)
			}
			if a, ok := attr["private_ip_address"].(string); ok {
				conn[PrivateIPKey] = []byte(a)
			}
			if a, ok := attr["public_ip_address"].(string); ok {
				conn[PublicIPKey] = []byte(a)
			}
			if a, ok := attr["root_password"].(string); ok {
				conn[xpv1.ResourceCredentialsSecretPasswordKey] = []byte(a)
			}
			// map
			if certSlice, ok := attr["server_ca_cert"].([]interface{}); ok && len(certSlice) > 0 {
				if certattrs, ok := certSlice[0].(map[string]interface{}); ok {
					if a, ok := certattrs["cert"].(string); ok {
						conn[CloudSQLSecretServerCACertificateCertKey] = []byte(a)
					}
					if a, ok := certattrs["common_name"].(string); ok {
						conn[CloudSQLSecretServerCACertificateCommonNameKey] = []byte(a)
					}
					if a, ok := certattrs["create_time"].(string); ok {
						conn[CloudSQLSecretServerCACertificateCreateTimeKey] = []byte(a)
					}
					if a, ok := certattrs["expiration_time"].(string); ok {
						conn[CloudSQLSecretServerCACertificateExpirationTimeKey] = []byte(a)
					}
					if a, ok := certattrs["sha1_fingerprint"].(string); ok {
						conn[CloudSQLSecretServerCACertificateSha1FingerprintKey] = []byte(a)
					}
				}
			}
			return conn, nil
		}
	})
	p.AddResourceConfigurator("google_sql_database", func(r *config.Resource) {
		r.References["instance"] = config.Reference{
			TerraformName: "google_sql_database_instance",
		}
	})
	p.AddResourceConfigurator("google_sql_source_representation_instance", func(r *config.Resource) {
		r.References["instance"] = config.Reference{
			TerraformName: "google_sql_database_instance",
		}
	})
	p.AddResourceConfigurator("google_sql_user", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			id, ok := tfstate["id"].(string)
			if !ok {
				return "", errors.New("cannot get the id field as string in tfstate")
			}
			words := strings.Split(id, "/")
			return words[0], nil
		}
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
			instance, err := common.GetField(parameters, "instance")
			if err != nil {
				return "", err
			}
			host, err := common.GetField(parameters, "host")
			if err != nil {
				host = ""
			}

			return fmt.Sprintf("%s/%s/%s", externalName, host, instance), nil
		}

		r.References["instance"] = config.Reference{
			TerraformName: "google_sql_database_instance",
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["password"].(string); ok {
				conn["password"] = []byte(a)
			}
			return conn, nil
		}
	})
	p.AddResourceConfigurator("google_sql_ssl_cert", func(r *config.Resource) {
		r.References["instance"] = config.Reference{
			TerraformName: "google_sql_database_instance",
		}
	})
}
