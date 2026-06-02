// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package templates

import _ "embed"

// ControllerTemplate is the template for the MR controller setup files.
//
//go:embed controller.go.tmpl
var ControllerTemplate string

// SetupAggregatorTemplate is the template for the controller setup aggregator files.
// It omits the SetupWebhookWithManager aggregator since the controller template
// does not emit a per-resource SetupWebhookWithManager function.
//
//go:embed setup.go.tmpl
var SetupAggregatorTemplate string
