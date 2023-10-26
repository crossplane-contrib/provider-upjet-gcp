package oslogin

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_os_login_ssh_public_key", func(r *config.Resource) {
		r.TerraformResource.Schema["key"].Sensitive = true
	})
}
