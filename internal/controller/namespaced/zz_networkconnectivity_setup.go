// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	hub "github.com/upbound/provider-gcp/internal/controller/namespaced/networkconnectivity/hub"
	serviceconnectionpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/networkconnectivity/serviceconnectionpolicy"
	spoke "github.com/upbound/provider-gcp/internal/controller/namespaced/networkconnectivity/spoke"
)

// Setup_networkconnectivity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkconnectivity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		hub.Setup,
		serviceconnectionpolicy.Setup,
		spoke.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_networkconnectivity creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_networkconnectivity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		hub.SetupGated,
		serviceconnectionpolicy.SetupGated,
		spoke.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
