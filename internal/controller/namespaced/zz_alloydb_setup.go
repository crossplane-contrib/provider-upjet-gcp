// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backup "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/alloydb/backup"
	cluster "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/alloydb/cluster"
	instance "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/alloydb/instance"
)

// Setup_alloydb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_alloydb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.Setup,
		cluster.Setup,
		instance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_alloydb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_alloydb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.SetupGated,
		cluster.SetupGated,
		instance.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
