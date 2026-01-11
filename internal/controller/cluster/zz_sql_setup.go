// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	database "github.com/upbound/provider-gcp/v2/internal/controller/cluster/sql/database"
	databaseinstance "github.com/upbound/provider-gcp/v2/internal/controller/cluster/sql/databaseinstance"
	sourcerepresentationinstance "github.com/upbound/provider-gcp/v2/internal/controller/cluster/sql/sourcerepresentationinstance"
	sslcert "github.com/upbound/provider-gcp/v2/internal/controller/cluster/sql/sslcert"
	user "github.com/upbound/provider-gcp/v2/internal/controller/cluster/sql/user"
)

// Setup_sql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		databaseinstance.Setup,
		sourcerepresentationinstance.Setup,
		sslcert.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_sql creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_sql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.SetupGated,
		databaseinstance.SetupGated,
		sourcerepresentationinstance.SetupGated,
		sslcert.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
