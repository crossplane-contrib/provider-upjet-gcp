/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "github.com/upbound/provider-gcp/internal/controller/container/cluster"
	nodepool "github.com/upbound/provider-gcp/internal/controller/container/nodepool"
	registry "github.com/upbound/provider-gcp/internal/controller/container/registry"
)

// Setup_container creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_container(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		nodepool.Setup,
		registry.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
