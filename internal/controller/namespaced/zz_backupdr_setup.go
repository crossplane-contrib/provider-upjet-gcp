// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backupplan "github.com/upbound/provider-gcp/v3/internal/controller/namespaced/backupdr/backupplan"
	backupplanassociation "github.com/upbound/provider-gcp/v3/internal/controller/namespaced/backupdr/backupplanassociation"
	backupvault "github.com/upbound/provider-gcp/v3/internal/controller/namespaced/backupdr/backupvault"
)

// Setup_backupdr creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_backupdr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupplan.Setup,
		backupplanassociation.Setup,
		backupvault.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_backupdr creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_backupdr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupplan.SetupGated,
		backupplanassociation.SetupGated,
		backupvault.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_backupdr registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_backupdr(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		backupplan.SetupWebhookWithManager,
		backupplanassociation.SetupWebhookWithManager,
		backupvault.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
