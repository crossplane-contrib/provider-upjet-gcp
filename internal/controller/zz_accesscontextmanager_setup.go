// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesslevel "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevel"
	accesslevelcondition "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevelcondition"
	accesspolicy "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesspolicy"
	accesspolicyiammember "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesspolicyiammember"
	serviceperimeter "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeter"
	serviceperimeterdryrunegresspolicy "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeterdryrunegresspolicy"
	serviceperimeterdryruningresspolicy "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeterdryruningresspolicy"
	serviceperimeterdryrunresource "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeterdryrunresource"
	serviceperimeteregresspolicy "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeteregresspolicy"
	serviceperimeteringresspolicy "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeteringresspolicy"
	serviceperimeterresource "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeterresource"
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
		serviceperimeterdryrunegresspolicy.Setup,
		serviceperimeterdryruningresspolicy.Setup,
		serviceperimeterdryrunresource.Setup,
		serviceperimeteregresspolicy.Setup,
		serviceperimeteringresspolicy.Setup,
		serviceperimeterresource.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
