// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	locationtagbinding "github.com/upbound/provider-gcp/internal/controller/namespaced/tags/locationtagbinding"
	tagbinding "github.com/upbound/provider-gcp/internal/controller/namespaced/tags/tagbinding"
	tagkey "github.com/upbound/provider-gcp/internal/controller/namespaced/tags/tagkey"
	tagvalue "github.com/upbound/provider-gcp/internal/controller/namespaced/tags/tagvalue"
)

// Setup_tags creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_tags(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		locationtagbinding.Setup,
		tagbinding.Setup,
		tagkey.Setup,
		tagvalue.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_tags creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_tags(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		locationtagbinding.SetupGated,
		tagbinding.SetupGated,
		tagkey.SetupGated,
		tagvalue.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
