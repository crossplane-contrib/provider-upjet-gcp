// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	deidentifytemplate "github.com/upbound/provider-gcp/internal/controller/datalossprevention/deidentifytemplate"
	inspecttemplate "github.com/upbound/provider-gcp/internal/controller/datalossprevention/inspecttemplate"
	jobtrigger "github.com/upbound/provider-gcp/internal/controller/datalossprevention/jobtrigger"
	storedinfotype "github.com/upbound/provider-gcp/internal/controller/datalossprevention/storedinfotype"
)

// Setup_datalossprevention creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datalossprevention(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		deidentifytemplate.Setup,
		inspecttemplate.Setup,
		jobtrigger.Setup,
		storedinfotype.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
