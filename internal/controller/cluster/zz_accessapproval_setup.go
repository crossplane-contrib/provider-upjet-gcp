// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	projectsettings "github.com/upbound/provider-gcp/v3/internal/controller/cluster/accessapproval/projectsettings"
)

// Setup_accessapproval creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_accessapproval(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectsettings.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_accessapproval creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_accessapproval(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectsettings.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_accessapproval registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_accessapproval(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		projectsettings.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
