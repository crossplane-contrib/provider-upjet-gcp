// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	workflow "github.com/upbound/provider-gcp/v2/internal/controller/cluster/workflows/workflow"
)

// Setup_workflows creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_workflows(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		workflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_workflows creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_workflows(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		workflow.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_workflows registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_workflows(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		workflow.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
