// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	connectionprofile "github.com/upbound/provider-gcp/internal/controller/datastream/connectionprofile"
	privateconnection "github.com/upbound/provider-gcp/internal/controller/datastream/privateconnection"
)

// Setup_datastream creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datastream(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connectionprofile.Setup,
		privateconnection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
