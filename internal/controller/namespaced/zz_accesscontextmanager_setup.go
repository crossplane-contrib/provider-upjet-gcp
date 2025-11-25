// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesslevel "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/accesscontextmanager/accesslevel"
	accesslevelcondition "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/accesscontextmanager/accesslevelcondition"
	accesspolicy "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/accesscontextmanager/accesspolicy"
	accesspolicyiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/accesscontextmanager/accesspolicyiammember"
	serviceperimeter "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/accesscontextmanager/serviceperimeter"
	serviceperimeterresource "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/accesscontextmanager/serviceperimeterresource"
)

// Setup_accesscontextmanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_accesscontextmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslevel.Setup,
		accesslevelcondition.Setup,
		accesspolicy.Setup,
		accesspolicyiammember.Setup,
		serviceperimeter.Setup,
		serviceperimeterresource.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_accesscontextmanager creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_accesscontextmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslevel.SetupGated,
		accesslevelcondition.SetupGated,
		accesspolicy.SetupGated,
		accesspolicyiammember.SetupGated,
		serviceperimeter.SetupGated,
		serviceperimeterresource.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
