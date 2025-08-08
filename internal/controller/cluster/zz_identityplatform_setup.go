// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	config "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/config"
	defaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/defaultsupportedidpconfig"
	inboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/inboundsamlconfig"
	oauthidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/oauthidpconfig"
	tenant "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenant"
	tenantdefaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenantdefaultsupportedidpconfig"
	tenantinboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenantinboundsamlconfig"
	tenantoauthidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenantoauthidpconfig"
)

// Setup_identityplatform creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_identityplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		config.Setup,
		defaultsupportedidpconfig.Setup,
		inboundsamlconfig.Setup,
		oauthidpconfig.Setup,
		tenant.Setup,
		tenantdefaultsupportedidpconfig.Setup,
		tenantinboundsamlconfig.Setup,
		tenantoauthidpconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_identityplatform creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_identityplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		config.SetupGated,
		defaultsupportedidpconfig.SetupGated,
		inboundsamlconfig.SetupGated,
		oauthidpconfig.SetupGated,
		tenant.SetupGated,
		tenantdefaultsupportedidpconfig.SetupGated,
		tenantinboundsamlconfig.SetupGated,
		tenantoauthidpconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
