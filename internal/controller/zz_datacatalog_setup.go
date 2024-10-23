// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	entry "github.com/upbound/provider-gcp/internal/controller/datacatalog/entry"
	entrygroup "github.com/upbound/provider-gcp/internal/controller/datacatalog/entrygroup"
	policytag "github.com/upbound/provider-gcp/internal/controller/datacatalog/policytag"
	policytagiambinding "github.com/upbound/provider-gcp/internal/controller/datacatalog/policytagiambinding"
	policytagiammember "github.com/upbound/provider-gcp/internal/controller/datacatalog/policytagiammember"
	tag "github.com/upbound/provider-gcp/internal/controller/datacatalog/tag"
	tagtemplate "github.com/upbound/provider-gcp/internal/controller/datacatalog/tagtemplate"
	tagtemplateiambinding "github.com/upbound/provider-gcp/internal/controller/datacatalog/tagtemplateiambinding"
	tagtemplateiammember "github.com/upbound/provider-gcp/internal/controller/datacatalog/tagtemplateiammember"
	taxonomy "github.com/upbound/provider-gcp/internal/controller/datacatalog/taxonomy"
	taxonomyiambinding "github.com/upbound/provider-gcp/internal/controller/datacatalog/taxonomyiambinding"
	taxonomyiammember "github.com/upbound/provider-gcp/internal/controller/datacatalog/taxonomyiammember"
)

// Setup_datacatalog creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datacatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		entry.Setup,
		entrygroup.Setup,
		policytag.Setup,
		policytagiambinding.Setup,
		policytagiammember.Setup,
		tag.Setup,
		tagtemplate.Setup,
		tagtemplateiambinding.Setup,
		tagtemplateiammember.Setup,
		taxonomy.Setup,
		taxonomyiambinding.Setup,
		taxonomyiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
