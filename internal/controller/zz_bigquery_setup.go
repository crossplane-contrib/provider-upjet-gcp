/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	dataset "github.com/upbound/provider-gcp/internal/controller/bigquery/dataset"
	table "github.com/upbound/provider-gcp/internal/controller/bigquery/table"
)

// Setup_bigquery creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bigquery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dataset.Setup,
		table.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
