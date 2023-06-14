/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	serviceaccount "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/provider-gcp/internal/controller/cloudplatform/servicenetworkingpeereddnsdomain"
)

// Setup_cloudplatform creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		serviceaccount.Setup,
		serviceaccountiammember.Setup,
		serviceaccountkey.Setup,
		servicenetworkingpeereddnsdomain.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
