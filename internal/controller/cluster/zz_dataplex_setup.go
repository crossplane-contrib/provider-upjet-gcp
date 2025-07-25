// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	aspecttype "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/aspecttype"
	asset "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/asset"
	lake "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/lake"
	lakeiampolicy "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/lakeiampolicy"
	zone "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/zone"
)

// Setup_dataplex creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataplex(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aspecttype.Setup,
		asset.Setup,
		lake.Setup,
		lakeiampolicy.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dataplex creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dataplex(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aspecttype.SetupGated,
		asset.SetupGated,
		lake.SetupGated,
		lakeiampolicy.SetupGated,
		zone.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
