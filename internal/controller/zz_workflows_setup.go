/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	workflow "github.com/upbound/provider-gcp/internal/controller/workflows/workflow"
)

// Setup_workflows creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_workflows(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		workflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
