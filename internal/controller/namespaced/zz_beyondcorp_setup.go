// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appconnection "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/beyondcorp/appconnection"
	appconnector "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/beyondcorp/appconnector"
	appgateway "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/beyondcorp/appgateway"
)

// Setup_beyondcorp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_beyondcorp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appconnection.Setup,
		appconnector.Setup,
		appgateway.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_beyondcorp creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_beyondcorp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appconnection.SetupGated,
		appconnector.SetupGated,
		appgateway.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
