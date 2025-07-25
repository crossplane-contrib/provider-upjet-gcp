// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudscheduler

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_cloud_scheduler_job", func(r *config.Resource) {
		r.References["pubsub_target.topic_name"] = config.Reference{
			// Note(donovanmuller): What about support for 'pubsub/v1beta1.LiteTopic' reference?
			TerraformName: "google_pubsub_topic",
		}
		r.MarkAsRequired("region")
	})
}
