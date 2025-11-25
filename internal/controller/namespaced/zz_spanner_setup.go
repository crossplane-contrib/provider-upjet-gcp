// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	database "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/spanner/database"
	databaseiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/spanner/databaseiammember"
	instance "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/spanner/instance"
	instanceiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/spanner/instanceiammember"
)

// Setup_spanner creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_spanner(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		databaseiammember.Setup,
		instance.Setup,
		instanceiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_spanner creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_spanner(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.SetupGated,
		databaseiammember.SetupGated,
		instance.SetupGated,
		instanceiammember.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
