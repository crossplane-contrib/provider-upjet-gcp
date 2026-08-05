// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/appengine/application"
	applicationurldispatchrules "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/appengine/applicationurldispatchrules"
	firewallrule "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/appengine/firewallrule"
	servicenetworksettings "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/appengine/servicenetworksettings"
	standardappversion "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/appengine/standardappversion"
)

// Setup_appengine creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_appengine(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationurldispatchrules.Setup,
		firewallrule.Setup,
		servicenetworksettings.Setup,
		standardappversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_appengine creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_appengine(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.SetupGated,
		applicationurldispatchrules.SetupGated,
		firewallrule.SetupGated,
		servicenetworksettings.SetupGated,
		standardappversion.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_appengine registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_appengine(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		application.SetupWebhookWithManager,
		applicationurldispatchrules.SetupWebhookWithManager,
		firewallrule.SetupWebhookWithManager,
		servicenetworksettings.SetupWebhookWithManager,
		standardappversion.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
