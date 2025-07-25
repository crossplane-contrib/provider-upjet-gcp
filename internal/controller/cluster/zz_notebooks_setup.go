// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	environment "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/environment"
	instance "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/instance"
	instanceiammember "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/instanceiammember"
	runtime "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/runtime"
	runtimeiammember "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/runtimeiammember"
)

// Setup_notebooks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_notebooks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		environment.Setup,
		instance.Setup,
		instanceiammember.Setup,
		runtime.Setup,
		runtimeiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_notebooks creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_notebooks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		environment.SetupGated,
		instance.SetupGated,
		instanceiammember.SetupGated,
		runtime.SetupGated,
		runtimeiammember.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
