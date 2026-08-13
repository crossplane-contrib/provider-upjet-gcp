// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package container

import (
	"encoding/base64"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/upjet/v2/pkg/config"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/upbound/provider-gcp/v3/config/cluster/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("google_container_cluster", func(r *config.Resource) {
		r.Kind = "Cluster"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"cluster_ipv4_cidr",
				"ip_allocation_policy",
				"node_config",
				"node_version",
				"enable_autopilot",
				"workload_identity_config",
				"addons_config.network_policy_config",
				"addons_config.gcp_filestore_csi_driver_config",
				"addons_config.gcs_fuse_csi_driver_config",
				"addons_config.dns_cache_config",
				"default_max_pods_per_node",
				"cluster_autoscaling.enabled",
				"cluster_autoscaling.resource_limits",
				"enable_intranode_visibility",
				"network_policy",
				"enable_shielded_nodes",
				"logging_config",
				"monitoring_config",
				"logging_service",
				"monitoring_service",
			},
		}
		config.MoveToStatus(r.TerraformResource, "node_pool")
		r.Sensitive.AdditionalConnectionDetailsFn = clusterConnectionDetails
		r.References["network"] = config.Reference{
			TerraformName: "google_compute_network",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.PathSelfLinkExtractor,
		}
		r.References["private_cluster_config.private_endpoint_subnetwork"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.PathSelfLinkExtractor,
		}

		r.MarkAsRequired("location")

		r.TerraformResource.Schema["database_encryption"].Elem.(*schema.Resource).
			Schema["state"].ValidateFunc = validation.StringInSlice([]string{"ENCRYPTED", "ALL_OBJECTS_ENCRYPTION_ENABLED", "DECRYPTED"}, false)
		r.MetaResource.ArgumentDocs["database_encryption.state"] = ""
		r.TerraformResource.Schema["database_encryption"].Elem.(*schema.Resource).
			Schema["state"].Description = `ENCRYPTED, ALL_OBJECTS_ENCRYPTION_ENABLED or DECRYPTED.`
		r.TerraformResource.Schema["database_encryption"].Elem.(*schema.Resource).
			Schema["state"].DiffSuppressFunc = DatabaseEncryptionSuppress

		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}
			delete(diff.Attributes, "autopilot_cluster_policy_config.#")
			return diff, nil
		}
	})

	p.AddResourceConfigurator("google_container_node_pool", func(r *config.Resource) {
		r.Kind = "NodePool"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"version",
				"node_count",
				"initial_node_count",
			},
		}
		r.References["cluster"] = config.Reference{
			TerraformName: "google_container_cluster",
			Extractor:     common.ExtractResourceIDFuncPath,
		}

		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, _ *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Destroy {
				return diff, nil
			}
			if ppDiff, ok := diff.Attributes["placement_policy.#"]; ok && ppDiff.Old == "" && ppDiff.New == "" {
				delete(diff.Attributes, "placement_policy.#")
			}
			if asDiff, ok := diff.Attributes["autoscaling.#"]; ok && asDiff.Old == "" && asDiff.New == "" {
				delete(diff.Attributes, "autoscaling.#")
			}
			if qpDiff, ok := diff.Attributes["queued_provisioning.#"]; ok && qpDiff.Old == "" && qpDiff.New == "" {
				delete(diff.Attributes, "queued_provisioning.#")
			}
			if incDiff, ok := diff.Attributes["initial_node_count"]; ok && incDiff.Old != "" {
				// Changes to actual node count can alter the value TF calculates for initial_node_count, resulting in
				// errors as initial_node_count cannot be updated. TF docs suggest using lifecycle ignore_changes for this
				// attribute.
				delete(diff.Attributes, "initial_node_count")
			}
			return diff, nil
		}
	})
}

// clusterConnectionDetails builds the kubeconfig published in the connection
// secret of a container.Cluster.
func clusterConnectionDetails(attr map[string]interface{}) (map[string][]byte, error) { //nolint:gocyclo // easier to follow as a unit
	name, err := common.GetField(attr, "name")
	if err != nil {
		return nil, err
	}
	endpoint, err := common.GetField(attr, "endpoint")
	if err != nil {
		return nil, err
	}
	server, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse API server endpoint")
	}
	// NOTE(hasheddan): the endpoint returned is just an IP address, and
	// clients will default to http, causing any authentication
	// information to be omitted. We set to https to allow
	// authentication.
	server.Scheme = "https"

	kcCluster := &clientcmdapi.Cluster{
		Server: server.String(),
	}
	// GKE reports the DNS endpoint in place of an IP address when
	// ipEndpointsConfig is disabled. That endpoint is served by Google Front
	// End, which presents a publicly trusted certificate rather than one
	// signed by masterAuth.clusterCaCertificate, so pinning the cluster CA
	// against it fails TLS validation. Leave those clients on the system
	// trust store instead.
	if !servesDNSEndpoint(attr, endpoint) {
		caData, err := common.GetField(attr, "master_auth[0].cluster_ca_certificate")
		if err != nil {
			return nil, err
		}
		caDataBytes, err := base64.StdEncoding.DecodeString(caData)
		if err != nil {
			return nil, errors.Wrap(err, "cannot serialize cluster ca data")
		}
		kcCluster.CertificateAuthorityData = caDataBytes
	}

	clientCertData, err := common.GetField(attr, "master_auth[0].client_certificate")
	if err != nil {
		return nil, err
	}
	clientCertDataBytes, err := base64.StdEncoding.DecodeString(clientCertData)
	if err != nil {
		return nil, errors.Wrap(err, "cannot serialize cluster client cert data")
	}
	clientKeyData, err := common.GetField(attr, "master_auth[0].client_key")
	if err != nil {
		return nil, err
	}
	clientKeyDataBytes, err := base64.StdEncoding.DecodeString(clientKeyData)
	if err != nil {
		return nil, errors.Wrap(err, "cannot serialize cluster client key data")
	}
	kc := clientcmdapi.Config{
		Kind:       "Config",
		APIVersion: "v1",
		Clusters: map[string]*clientcmdapi.Cluster{
			name: kcCluster,
		},
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			name: {
				ClientCertificateData: clientCertDataBytes,
				ClientKeyData:         clientKeyDataBytes,
			},
		},
		Contexts: map[string]*clientcmdapi.Context{
			name: {
				Cluster:  name,
				AuthInfo: name,
			},
		},
		CurrentContext: name,
	}
	kcb, err := clientcmd.Write(kc)
	if err != nil {
		return nil, errors.Wrap(err, "cannot serialize kubeconfig")
	}
	return map[string][]byte{
		"kubeconfig": kcb,
	}, nil
}

// servesDNSEndpoint reports whether endpoint is the cluster's DNS based
// control plane endpoint rather than an IP address.
func servesDNSEndpoint(attr map[string]interface{}, endpoint string) bool {
	dnsEndpoint, err := common.GetField(attr, "control_plane_endpoints_config[0].dns_endpoint_config[0].endpoint")
	if err != nil || dnsEndpoint == "" {
		return false
	}
	return endpointHost(dnsEndpoint) == endpointHost(endpoint)
}

// endpointHost normalizes a GKE endpoint, which the API reports without a
// scheme, so that two endpoints can be compared.
func endpointHost(endpoint string) string {
	if u, err := url.Parse(endpoint); err == nil && u.Host != "" {
		return u.Host
	}
	return endpoint
}

func DatabaseEncryptionSuppress(k, old, new string, d *schema.ResourceData) bool {
	// The API sometimes returns ALL_OBJECTS_ENCRYPTION_ENABLED when the user sets ENCRYPTED
	// and vice versa (depending on the cluster version and underlying resource storage).
	if old == "ALL_OBJECTS_ENCRYPTION_ENABLED" && new == "ENCRYPTED" {
		return true
	}
	if old == "ENCRYPTED" && new == "ALL_OBJECTS_ENCRYPTION_ENABLED" {
		return true
	}
	return false
}
