/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	processor "github.com/upbound/provider-gcp/internal/controller/documentai/processor"
)

// Setup_documentai creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_documentai(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		processor.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
