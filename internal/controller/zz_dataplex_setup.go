// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	aspecttype "github.com/upbound/provider-gcp/internal/controller/dataplex/aspecttype"
	asset "github.com/upbound/provider-gcp/internal/controller/dataplex/asset"
	lake "github.com/upbound/provider-gcp/internal/controller/dataplex/lake"
	lakeiampolicy "github.com/upbound/provider-gcp/internal/controller/dataplex/lakeiampolicy"
	zone "github.com/upbound/provider-gcp/internal/controller/dataplex/zone"
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
