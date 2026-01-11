// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connectaccountconnector "github.com/upbound/provider-gcp/v2/internal/controller/cluster/developerconnect/connectaccountconnector"
	connectconnection "github.com/upbound/provider-gcp/v2/internal/controller/cluster/developerconnect/connectconnection"
	connectgitrepositorylink "github.com/upbound/provider-gcp/v2/internal/controller/cluster/developerconnect/connectgitrepositorylink"
)

// Setup_developerconnect creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_developerconnect(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connectaccountconnector.Setup,
		connectconnection.Setup,
		connectgitrepositorylink.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_developerconnect creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_developerconnect(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connectaccountconnector.SetupGated,
		connectconnection.SetupGated,
		connectgitrepositorylink.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
