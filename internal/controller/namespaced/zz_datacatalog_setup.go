// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	entry "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/datacatalog/entry"
	entrygroup "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/datacatalog/entrygroup"
	policytag "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/datacatalog/policytag"
	tag "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/datacatalog/tag"
	tagtemplate "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/datacatalog/tagtemplate"
	taxonomy "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/datacatalog/taxonomy"
)

// Setup_datacatalog creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datacatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		entry.Setup,
		entrygroup.Setup,
		policytag.Setup,
		tag.Setup,
		tagtemplate.Setup,
		taxonomy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_datacatalog creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_datacatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		entry.SetupGated,
		entrygroup.SetupGated,
		policytag.SetupGated,
		tag.SetupGated,
		tagtemplate.SetupGated,
		taxonomy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
