// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	addressgroup "github.com/upbound/provider-gcp/internal/controller/networksecurity/addressgroup"
	gatewaysecuritypolicy "github.com/upbound/provider-gcp/internal/controller/networksecurity/gatewaysecuritypolicy"
	gatewaysecuritypolicyrule "github.com/upbound/provider-gcp/internal/controller/networksecurity/gatewaysecuritypolicyrule"
	tlsinspectionpolicy "github.com/upbound/provider-gcp/internal/controller/networksecurity/tlsinspectionpolicy"
	urllists "github.com/upbound/provider-gcp/internal/controller/networksecurity/urllists"
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
