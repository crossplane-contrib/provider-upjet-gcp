// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	appengineserviceiammember "github.com/upbound/provider-gcp/internal/controller/iap/appengineserviceiammember"
	appengineversioniammember "github.com/upbound/provider-gcp/internal/controller/iap/appengineversioniammember"
	tunneliammember "github.com/upbound/provider-gcp/internal/controller/iap/tunneliammember"
	webbackendserviceiammember "github.com/upbound/provider-gcp/internal/controller/iap/webbackendserviceiammember"
	webiammember "github.com/upbound/provider-gcp/internal/controller/iap/webiammember"
	webtypeappengineiammember "github.com/upbound/provider-gcp/internal/controller/iap/webtypeappengineiammember"
	webtypecomputeiammember "github.com/upbound/provider-gcp/internal/controller/iap/webtypecomputeiammember"
)

// Setup_iap creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_iap(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appengineserviceiammember.Setup,
		appengineversioniammember.Setup,
		tunneliammember.Setup,
		webbackendserviceiammember.Setup,
		webiammember.Setup,
		webtypeappengineiammember.Setup,
		webtypecomputeiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
