// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	agent "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/agent"
	entitytype "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/entitytype"
	environment "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/environment"
	flow "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/flow"
	intent "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/intent"
	page "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/page"
	version "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/version"
	webhook "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/webhook"
)

// Setup_dialogflowcx creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dialogflowcx(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agent.Setup,
		entitytype.Setup,
		environment.Setup,
		flow.Setup,
		intent.Setup,
		page.Setup,
		version.Setup,
		webhook.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
