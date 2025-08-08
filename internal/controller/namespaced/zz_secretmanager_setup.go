// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	secret "github.com/upbound/provider-gcp/internal/controller/namespaced/secretmanager/secret"
	secretiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/secretmanager/secretiammember"
	secretversion "github.com/upbound/provider-gcp/internal/controller/namespaced/secretmanager/secretversion"
)

// Setup_secretmanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_secretmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secret.Setup,
		secretiammember.Setup,
		secretversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_secretmanager creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_secretmanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		secret.SetupGated,
		secretiammember.SetupGated,
		secretversion.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
