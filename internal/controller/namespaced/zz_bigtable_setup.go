// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appprofile "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/appprofile"
	garbagecollectionpolicy "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/garbagecollectionpolicy"
	instance "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/instance"
	instanceiambinding "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/instanceiambinding"
	instanceiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/instanceiammember"
	instanceiampolicy "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/instanceiampolicy"
	table "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/table"
	tableiambinding "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/tableiambinding"
	tableiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/tableiammember"
	tableiampolicy "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/bigtable/tableiampolicy"
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

// SetupGated_bigtable creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_bigtable(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appprofile.SetupGated,
		garbagecollectionpolicy.SetupGated,
		instance.SetupGated,
		instanceiambinding.SetupGated,
		instanceiammember.SetupGated,
		instanceiampolicy.SetupGated,
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
