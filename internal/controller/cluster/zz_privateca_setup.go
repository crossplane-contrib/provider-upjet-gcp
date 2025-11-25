// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	capool "github.com/upbound/provider-gcp/v2/internal/controller/cluster/privateca/capool"
	capooliammember "github.com/upbound/provider-gcp/v2/internal/controller/cluster/privateca/capooliammember"
	certificate "github.com/upbound/provider-gcp/v2/internal/controller/cluster/privateca/certificate"
	certificateauthority "github.com/upbound/provider-gcp/v2/internal/controller/cluster/privateca/certificateauthority"
	certificatetemplate "github.com/upbound/provider-gcp/v2/internal/controller/cluster/privateca/certificatetemplate"
	certificatetemplateiammember "github.com/upbound/provider-gcp/v2/internal/controller/cluster/privateca/certificatetemplateiammember"
)

// Setup_privateca creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_privateca(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		capool.Setup,
		capooliammember.Setup,
		certificate.Setup,
		certificateauthority.Setup,
		certificatetemplate.Setup,
		certificatetemplateiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_privateca creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_privateca(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		capool.SetupGated,
		capooliammember.SetupGated,
		certificate.SetupGated,
		certificateauthority.SetupGated,
		certificatetemplate.SetupGated,
		certificatetemplateiammember.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
