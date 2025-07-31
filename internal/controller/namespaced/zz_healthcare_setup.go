// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	consentstore "github.com/upbound/provider-gcp/internal/controller/namespaced/healthcare/consentstore"
	dataset "github.com/upbound/provider-gcp/internal/controller/namespaced/healthcare/dataset"
	datasetiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/healthcare/datasetiammember"
)

// Setup_healthcare creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_healthcare(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		consentstore.Setup,
		dataset.Setup,
		datasetiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_healthcare creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_healthcare(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		consentstore.SetupGated,
		dataset.SetupGated,
		datasetiammember.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
