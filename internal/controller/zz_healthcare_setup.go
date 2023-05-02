/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	consentstore "github.com/upbound/provider-gcp/internal/controller/healthcare/consentstore"
	dataset "github.com/upbound/provider-gcp/internal/controller/healthcare/dataset"
	datasetiammember "github.com/upbound/provider-gcp/internal/controller/healthcare/datasetiammember"
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
