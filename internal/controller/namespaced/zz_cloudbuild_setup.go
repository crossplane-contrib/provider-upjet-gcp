// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	trigger "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudbuild/trigger"
	workerpool "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudbuild/workerpool"
)

// Setup_cloudbuild creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudbuild(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		trigger.Setup,
		workerpool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudbuild creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudbuild(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		trigger.SetupGated,
		workerpool.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
