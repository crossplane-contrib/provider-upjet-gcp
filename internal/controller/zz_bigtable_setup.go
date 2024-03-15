// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	appprofile "github.com/upbound/provider-gcp/internal/controller/bigtable/appprofile"
	garbagecollectionpolicy "github.com/upbound/provider-gcp/internal/controller/bigtable/garbagecollectionpolicy"
	instance "github.com/upbound/provider-gcp/internal/controller/bigtable/instance"
	instanceiambinding "github.com/upbound/provider-gcp/internal/controller/bigtable/instanceiambinding"
	instanceiammember "github.com/upbound/provider-gcp/internal/controller/bigtable/instanceiammember"
	instanceiampolicy "github.com/upbound/provider-gcp/internal/controller/bigtable/instanceiampolicy"
	table "github.com/upbound/provider-gcp/internal/controller/bigtable/table"
	tableiambinding "github.com/upbound/provider-gcp/internal/controller/bigtable/tableiambinding"
	tableiammember "github.com/upbound/provider-gcp/internal/controller/bigtable/tableiammember"
	tableiampolicy "github.com/upbound/provider-gcp/internal/controller/bigtable/tableiampolicy"
)

// Setup_bigtable creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bigtable(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appprofile.Setup,
		garbagecollectionpolicy.Setup,
		instance.Setup,
		instanceiambinding.Setup,
		instanceiammember.Setup,
		instanceiampolicy.Setup,
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
