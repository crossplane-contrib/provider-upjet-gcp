// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	locationtagbinding "github.com/upbound/provider-gcp/internal/controller/cluster/tags/locationtagbinding"
	tagbinding "github.com/upbound/provider-gcp/internal/controller/cluster/tags/tagbinding"
	tagkey "github.com/upbound/provider-gcp/internal/controller/cluster/tags/tagkey"
	tagvalue "github.com/upbound/provider-gcp/internal/controller/cluster/tags/tagvalue"
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
