// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package monitoring

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_monitoring_slo", func(r *config.Resource) {
		r.Path = "sloes"
	})
	p.AddResourceConfigurator("google_monitoring_notification_channel", func(r *config.Resource) {
		delete(r.TerraformResource.
			Schema["sensitive_labels"].Elem.(*schema.Resource).
			Schema, "auth_token_wo")
		delete(r.TerraformResource.
			Schema["sensitive_labels"].Elem.(*schema.Resource).
			Schema, "password_wo")
		delete(r.TerraformResource.
			Schema["sensitive_labels"].Elem.(*schema.Resource).
			Schema, "service_key_wo")
	})
	p.AddResourceConfigurator("google_monitoring_uptime_check_config", func(r *config.Resource) {
		delete(r.TerraformResource.
			Schema["http_check"].Elem.(*schema.Resource).
			Schema["auth_info"].Elem.(*schema.Resource).
			Schema, "password_wo")
	})
}
