// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	managedzone "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/managedzone"
	managedzoneiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/managedzoneiammember"
	policy "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/policy"
	recordset "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/recordset"
	responsepolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/responsepolicy"
	responsepolicyrule "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/responsepolicyrule"
)

// Setup_dns creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managedzone.Setup,
		managedzoneiammember.Setup,
		policy.Setup,
		recordset.Setup,
		responsepolicy.Setup,
		responsepolicyrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dns creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managedzone.SetupGated,
		managedzoneiammember.SetupGated,
		policy.SetupGated,
		recordset.SetupGated,
		responsepolicy.SetupGated,
		responsepolicyrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
