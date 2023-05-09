/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	function "github.com/upbound/provider-gcp/internal/controller/cloudfunctions2/function"
)

// Setup_cloudfunctions2 creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudfunctions2(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		function.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
