package container

import (
	"encoding/base64"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/upbound/official-providers/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_container_cluster", func(r *config.Resource) {
		r.Kind = "Cluster"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"cluster_ipv4_cidr", "ip_allocation_policy"},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			name, err := common.GetField(attr, "name")
			if err != nil {
				return nil, err
			}
			server, err := common.GetField(attr, "endpoint")
			if err != nil {
				return nil, err
			}
			caData, err := common.GetField(attr, "master_auth[0].cluster_ca_certificate")
			if err != nil {
				return nil, err
			}
			caDataBytes, err := base64.StdEncoding.DecodeString(caData)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot serialize cluster ca data")
			}
			kc := clientcmdapi.Config{
				Kind:       "Config",
				APIVersion: "v1",
				Clusters: map[string]*clientcmdapi.Cluster{
					name: {
						Server:                   server,
						CertificateAuthorityData: caDataBytes,
					},
				},
				AuthInfos: map[string]*clientcmdapi.AuthInfo{
					name: {},
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
		r.UseAsync = true
	})

	p.AddResourceConfigurator("google_container_node_pool", func(r *config.Resource) {
		r.Kind = "NodePool"
		r.References["cluster"] = config.Reference{
			Type:      "Cluster",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.UseAsync = true
	})
}
