// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	analyticshubdataexchange "github.com/upbound/provider-gcp/internal/controller/bigquery/analyticshubdataexchange"
	analyticshubdataexchangeiammember "github.com/upbound/provider-gcp/internal/controller/bigquery/analyticshubdataexchangeiammember"
	analyticshublisting "github.com/upbound/provider-gcp/internal/controller/bigquery/analyticshublisting"
	connection "github.com/upbound/provider-gcp/internal/controller/bigquery/connection"
	dataset "github.com/upbound/provider-gcp/internal/controller/bigquery/dataset"
	datasetaccess "github.com/upbound/provider-gcp/internal/controller/bigquery/datasetaccess"
	datasetiambinding "github.com/upbound/provider-gcp/internal/controller/bigquery/datasetiambinding"
	datasetiammember "github.com/upbound/provider-gcp/internal/controller/bigquery/datasetiammember"
	datasetiampolicy "github.com/upbound/provider-gcp/internal/controller/bigquery/datasetiampolicy"
	datatransferconfig "github.com/upbound/provider-gcp/internal/controller/bigquery/datatransferconfig"
	job "github.com/upbound/provider-gcp/internal/controller/bigquery/job"
	reservation "github.com/upbound/provider-gcp/internal/controller/bigquery/reservation"
	reservationassignment "github.com/upbound/provider-gcp/internal/controller/bigquery/reservationassignment"
	routine "github.com/upbound/provider-gcp/internal/controller/bigquery/routine"
	table "github.com/upbound/provider-gcp/internal/controller/bigquery/table"
	tableiambinding "github.com/upbound/provider-gcp/internal/controller/bigquery/tableiambinding"
	tableiammember "github.com/upbound/provider-gcp/internal/controller/bigquery/tableiammember"
	tableiampolicy "github.com/upbound/provider-gcp/internal/controller/bigquery/tableiampolicy"
)

// Setup_bigquery creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bigquery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyticshubdataexchange.Setup,
		analyticshubdataexchangeiammember.Setup,
		analyticshublisting.Setup,
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
