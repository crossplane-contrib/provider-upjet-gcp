// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	floorsetting "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/modelarmor/floorsetting"
	template "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/modelarmor/template"
)

// Setup_modelarmor creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_modelarmor(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		floorsetting.Setup,
		template.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_modelarmor creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_modelarmor(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		floorsetting.SetupGated,
		template.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
