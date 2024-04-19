// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package dialogflowcx

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_dialogflow_cx_version", func(r *config.Resource) {
		r.OverrideFieldNames["NluSettingsInitParameters"] = "VersionNluSettingsInitParameters"
		r.OverrideFieldNames["NluSettingsParameters"] = "VersionNluSettingsParameters"
		r.OverrideFieldNames["NluSettingsObservation"] = "VersionNluSettingsObservation"
	})
}
