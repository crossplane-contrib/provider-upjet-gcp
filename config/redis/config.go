// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package redis

import (
	"strconv"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_redis_instance", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["host"].(string); ok {
				conn["host"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("google_redis_cluster", func(r *config.Resource) {
		r.MarkAsRequired("region")
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			address, err := common.GetField(attr, "discovery_endpoints[0].address")
			if err != nil {
				return nil, err
			}
			conn["address"] = []byte(address)

			port, err := GetFloat(attr, "discovery_endpoints[0].port")
			if err != nil {
				return nil, err
			}
			conn["port"] = []byte(strconv.FormatFloat(port, 'f', -1, 64))

			return conn, nil
		}
	})
}

// GetFloat value of the supplied field path.
func GetFloat(from map[string]interface{}, path string) (float64, error) {
	v, err := fieldpath.Pave(from).GetValue(path)
	if err != nil {
		return 0, err
	}

	f, ok := v.(float64)
	if !ok {
		return 0, errors.Errorf("%s: not a (float64) number", path)
	}
	return f, nil
}
