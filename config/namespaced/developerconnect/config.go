// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package developerconnect

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const group = "developerconnect"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_developer_connect_account_connector", func(r *config.Resource) {
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("google_developer_connect_connection", func(r *config.Resource) {
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("google_developer_connect_git_repository_link", func(r *config.Resource) {
		r.ShortGroup = group
	})
}
