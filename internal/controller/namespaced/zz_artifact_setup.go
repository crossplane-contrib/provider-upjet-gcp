// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	registryrepository "github.com/upbound/provider-gcp/internal/controller/namespaced/artifact/registryrepository"
	registryrepositoryiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/artifact/registryrepositoryiammember"
)

// Setup_artifact creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_artifact(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		registryrepository.Setup,
		registryrepositoryiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_artifact creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_artifact(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		registryrepository.SetupGated,
		registryrepositoryiammember.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
