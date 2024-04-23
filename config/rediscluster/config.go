// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package rediscluster

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_redis_cluster", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "psc_configs")
		config.MarkAsRequired(r.TerraformResource, "shard_count")
		config.MarkAsRequired(r.TerraformResource, "name")
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["address"].(string); ok {
				conn["address"] = []byte(a)
			}
			return conn, nil
		}
		r.UseAsync = true
	})
}
