// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	agent "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/dialogflowcx/agent"
	entitytype "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/dialogflowcx/entitytype"
	environment "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/dialogflowcx/environment"
	flow "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/dialogflowcx/flow"
	intent "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/dialogflowcx/intent"
	page "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/dialogflowcx/page"
	version "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/dialogflowcx/version"
	webhook "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/dialogflowcx/webhook"
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

// SetupGated_dialogflowcx creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dialogflowcx(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agent.SetupGated,
		entitytype.SetupGated,
		environment.SetupGated,
		flow.SetupGated,
		intent.SetupGated,
		page.SetupGated,
		version.SetupGated,
		webhook.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
