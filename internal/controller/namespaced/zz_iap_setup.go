// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appengineserviceiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/iap/appengineserviceiammember"
	appengineversioniammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/iap/appengineversioniammember"
	tunneliammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/iap/tunneliammember"
	webbackendserviceiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/iap/webbackendserviceiammember"
	webiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/iap/webiammember"
	webtypeappengineiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/iap/webtypeappengineiammember"
	webtypecomputeiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/iap/webtypecomputeiammember"
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

// SetupGated_iap creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_iap(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appengineserviceiammember.SetupGated,
		appengineversioniammember.SetupGated,
		tunneliammember.SetupGated,
		webbackendserviceiammember.SetupGated,
		webiammember.SetupGated,
		webtypeappengineiammember.SetupGated,
		webtypecomputeiammember.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
