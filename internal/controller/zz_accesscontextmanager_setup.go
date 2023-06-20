/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accesslevel "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevel"
	accesslevelcondition "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevelcondition"
	accesspolicy "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesspolicy"
)

// Setup_accesscontextmanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_accesscontextmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslevel.Setup,
		accesslevelcondition.Setup,
		accesspolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
