// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backupbackupplan "github.com/upbound/provider-gcp/v2/internal/controller/cluster/gke/backupbackupplan"
)

// Setup_gke creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_gke(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupbackupplan.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_gke creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_gke(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupbackupplan.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
