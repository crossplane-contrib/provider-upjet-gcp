// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	release "github.com/upbound/provider-gcp/internal/controller/firebaserules/release"
	ruleset "github.com/upbound/provider-gcp/internal/controller/firebaserules/ruleset"
)

// Setup_firebaserules creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_firebaserules(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		release.Setup,
		ruleset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
