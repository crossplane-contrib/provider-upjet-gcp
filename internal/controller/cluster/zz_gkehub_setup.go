// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	membership "github.com/upbound/provider-gcp/internal/controller/cluster/gkehub/membership"
	membershipiammember "github.com/upbound/provider-gcp/internal/controller/cluster/gkehub/membershipiammember"
)

// Setup_gkehub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_gkehub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		membership.Setup,
		membershipiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_gkehub creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_gkehub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		membership.SetupGated,
		membershipiammember.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
