// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	addressgroup "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/addressgroup"
	gatewaysecuritypolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/gatewaysecuritypolicy"
	gatewaysecuritypolicyrule "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/gatewaysecuritypolicyrule"
	tlsinspectionpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/tlsinspectionpolicy"
	urllists "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/urllists"
)

// Setup_networksecurity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networksecurity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addressgroup.Setup,
		gatewaysecuritypolicy.Setup,
		gatewaysecuritypolicyrule.Setup,
		tlsinspectionpolicy.Setup,
		urllists.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_networksecurity creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_networksecurity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addressgroup.SetupGated,
		gatewaysecuritypolicy.SetupGated,
		gatewaysecuritypolicyrule.SetupGated,
		tlsinspectionpolicy.SetupGated,
		urllists.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
