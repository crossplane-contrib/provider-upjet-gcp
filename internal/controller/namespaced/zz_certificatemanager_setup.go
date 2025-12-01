// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	certificate "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/certificatemanager/certificate"
	certificatemap "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/certificatemanager/certificatemap"
	certificatemapentry "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/certificatemanager/certificatemapentry"
	dnsauthorization "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/certificatemanager/dnsauthorization"
	trustconfig "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/certificatemanager/trustconfig"
)

// Setup_certificatemanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_certificatemanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.Setup,
		certificatemap.Setup,
		certificatemapentry.Setup,
		dnsauthorization.Setup,
		trustconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_certificatemanager creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_certificatemanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.SetupGated,
		certificatemap.SetupGated,
		certificatemapentry.SetupGated,
		dnsauthorization.SetupGated,
		trustconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
