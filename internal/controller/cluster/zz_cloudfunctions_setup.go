// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	function "github.com/upbound/provider-gcp/v2/internal/controller/cluster/cloudfunctions/function"
	functioniammember "github.com/upbound/provider-gcp/v2/internal/controller/cluster/cloudfunctions/functioniammember"
)

// Setup_cloudfunctions creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudfunctions(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		function.Setup,
		functioniammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudfunctions creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudfunctions(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		function.SetupGated,
		functioniammember.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
