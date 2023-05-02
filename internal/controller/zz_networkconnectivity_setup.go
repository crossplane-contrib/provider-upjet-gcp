/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	hub "github.com/upbound/provider-gcp/internal/controller/networkconnectivity/hub"
	spoke "github.com/upbound/provider-gcp/internal/controller/networkconnectivity/spoke"
)

// Setup_networkconnectivity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkconnectivity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		hub.Setup,
		spoke.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
