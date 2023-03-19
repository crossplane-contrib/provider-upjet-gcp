package container

import (
	"encoding/base64"
	"net/url"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/upbound/provider-gcp/config/common"
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
				"node_version",
				"enable_autopilot",
				"workload_identity_config",
				"addons_config.network_policy_config",
				"addons_config.gcp_filestore_csi_driver_config",
				"addons_config.dns_cache_config",
				"default_max_pods_per_node",
				"cluster_autoscaling.enabled",
				"cluster_autoscaling.resource_limits",
				"enable_intranode_visibility",
				"network_policy",
				"enable_shielded_nodes",
			},
		}
		config.MoveToStatus(r.TerraformResource, "node_pool")
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
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
				return nil, errors.Wrapf(err, "cannot parse API server endpoint")
			}
			// NOTE(hasheddan): the endpoint returned is just an IP address, and
			// clients will default to http, causing any authentication
			// information to be omitted. We set to https to allow
			// authentication.
			server.Scheme = "https"
			caData, err := common.GetField(attr, "master_auth[0].cluster_ca_certificate")
			if err != nil {
				return nil, err
			}
			caDataBytes, err := base64.StdEncoding.DecodeString(caData)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot serialize cluster ca data")
			}
			clientCertData, err := common.GetField(attr, "master_auth[0].client_certificate")
			if err != nil {
				return nil, err
			}
			clientCertDataBytes, err := base64.StdEncoding.DecodeString(clientCertData)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot serialize cluster client cert data")
			}
			clientKeyData, err := common.GetField(attr, "master_auth[0].client_key")
			if err != nil {
				return nil, err
			}
			clientKeyDataBytes, err := base64.StdEncoding.DecodeString(clientKeyData)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot serialize cluster client key data")
			}
			kc := clientcmdapi.Config{
				Kind:       "Config",
				APIVersion: "v1",
				Clusters: map[string]*clientcmdapi.Cluster{
					name: {
						Server:                   server.String(),
						CertificateAuthorityData: caDataBytes,
					},
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
		r.References["network"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/compute/v1beta1.Network",
			Extractor: common.PathSelfLinkExtractor,
		}
		r.References["subnetwork"] = config.Reference{
			Type:      "github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork",
			Extractor: common.PathSelfLinkExtractor,
		}
		config.MarkAsRequired(r.TerraformResource, "location")
	})

	p.AddResourceConfigurator("google_container_node_pool", func(r *config.Resource) {
		r.Kind = "NodePool"
		r.References["cluster"] = config.Reference{
			Type:      "Cluster",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
