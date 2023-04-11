/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	function "github.com/upbound/provider-gcp/internal/controller/cloudfunctions/function"
	functioniammember "github.com/upbound/provider-gcp/internal/controller/cloudfunctions/functioniammember"
)

// Setup_cloudfunctions creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudfunctions(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		function.Setup,
		functioniammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
