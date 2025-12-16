// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	instance "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/memorystore/instance"
	instancedesiredusercreatedendpoints "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/memorystore/instancedesiredusercreatedendpoints"
)

// Setup_memorystore creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_memorystore(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instance.Setup,
		instancedesiredusercreatedendpoints.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_memorystore creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_memorystore(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instance.SetupGated,
		instancedesiredusercreatedendpoints.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
