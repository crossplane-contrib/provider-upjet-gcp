// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	quotapreference "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/cloudquotas/quotapreference"
)

// Setup_cloudquotas creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudquotas(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		quotapreference.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudquotas creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudquotas(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		quotapreference.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
