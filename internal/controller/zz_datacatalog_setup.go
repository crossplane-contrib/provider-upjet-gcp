// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	entry "github.com/upbound/provider-gcp/internal/controller/datacatalog/entry"
	entrygroup "github.com/upbound/provider-gcp/internal/controller/datacatalog/entrygroup"
	tag "github.com/upbound/provider-gcp/internal/controller/datacatalog/tag"
	tagtemplate "github.com/upbound/provider-gcp/internal/controller/datacatalog/tagtemplate"
)

// Setup_datacatalog creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datacatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		entry.Setup,
		entrygroup.Setup,
		tag.Setup,
		tagtemplate.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
