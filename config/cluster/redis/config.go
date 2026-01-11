// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package redis

import (
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_redis_instance", redisInstance)
	p.AddResourceConfigurator("google_redis_cluster", redisCluster)
}

func redisInstance(r *config.Resource) {
	config.MarkAsRequired(r.TerraformResource, "region")
	r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
		conn := map[string][]byte{}
		if host, ok := attr["host"].(string); ok {
			conn["host"] = []byte(host)
		}
		if port, ok := attr["port"].(float64); ok {
			conn["port"] = []byte(fmt.Sprintf("%g", port))
		}
		if caCerts, ok := attr["server_ca_certs"].([]any); ok {
			for i, ca := range caCerts {
				if caCerts, ok := ca.(map[string]any); ok && len(caCerts) > 0 {
					if cert, ok := caCerts["cert"].(string); ok {
						key := fmt.Sprintf("server_ca_certs_%d_cert", i)
						conn[key] = []byte(cert)
					}
				}
			}
		}
		return conn, nil
	}
}

func redisCluster(r *config.Resource) {
	r.MarkAsRequired("region")
	r.UseAsync = true
	r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
		conn := map[string][]byte{}
		if discoveryendpoints, ok := attr["discovery_endpoints"].([]any); ok {
			for i, de := range discoveryendpoints {
				if discoveryendpoints, ok := de.(map[string]any); ok && len(discoveryendpoints) > 0 {
					if address, ok := discoveryendpoints["address"].(string); ok {
						key := fmt.Sprintf("discovery_endpoints_%d_address", i)
						conn[key] = []byte(address)
					}
					if port, ok := discoveryendpoints["port"].(float64); ok {
						key := fmt.Sprintf("discovery_endpoints_%d_port", i)
						conn[key] = []byte(fmt.Sprintf("%g", port))
					}
				}
			}
		}
		return conn, nil
	}
}
