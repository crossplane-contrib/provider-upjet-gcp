// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	analyticshubdataexchange "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/analyticshubdataexchange"
	analyticshubdataexchangeiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/analyticshubdataexchangeiammember"
	analyticshublisting "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/analyticshublisting"
	analyticshublistingsubscription "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/analyticshublistingsubscription"
	connection "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/connection"
	dataset "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/dataset"
	datasetaccess "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datasetaccess"
	datasetiambinding "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datasetiambinding"
	datasetiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datasetiammember"
	datasetiampolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datasetiampolicy"
	datatransferconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datatransferconfig"
	job "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/job"
	reservation "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/reservation"
	reservationassignment "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/reservationassignment"
	routine "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/routine"
	table "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/table"
	tableiambinding "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/tableiambinding"
	tableiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/tableiammember"
	tableiampolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/tableiampolicy"
)

// Setup_bigquery creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bigquery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyticshubdataexchange.Setup,
		analyticshubdataexchangeiammember.Setup,
		analyticshublisting.Setup,
		analyticshublistingsubscription.Setup,
		connection.Setup,
		dataset.Setup,
		datasetaccess.Setup,
		datasetiambinding.Setup,
		datasetiammember.Setup,
		datasetiampolicy.Setup,
		datatransferconfig.Setup,
		job.Setup,
		reservation.Setup,
		reservationassignment.Setup,
		routine.Setup,
		table.Setup,
		tableiambinding.Setup,
		tableiammember.Setup,
		tableiampolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_bigquery creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_bigquery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyticshubdataexchange.SetupGated,
		analyticshubdataexchangeiammember.SetupGated,
		analyticshublisting.SetupGated,
		analyticshublistingsubscription.SetupGated,
		connection.SetupGated,
		dataset.SetupGated,
		datasetaccess.SetupGated,
		datasetiambinding.SetupGated,
		datasetiammember.SetupGated,
		datasetiampolicy.SetupGated,
		datatransferconfig.SetupGated,
		job.SetupGated,
		reservation.SetupGated,
		reservationassignment.SetupGated,
		routine.SetupGated,
		table.SetupGated,
		tableiambinding.SetupGated,
		tableiammember.SetupGated,
		tableiampolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
