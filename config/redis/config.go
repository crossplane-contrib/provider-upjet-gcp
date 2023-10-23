package redis

import "github.com/upbound/upjet/pkg/config"

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
}
