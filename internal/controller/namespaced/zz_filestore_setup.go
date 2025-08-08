// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backup "github.com/upbound/provider-gcp/internal/controller/namespaced/filestore/backup"
	instance "github.com/upbound/provider-gcp/internal/controller/namespaced/filestore/instance"
	snapshot "github.com/upbound/provider-gcp/internal/controller/namespaced/filestore/snapshot"
)

// Setup_filestore creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_filestore(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.Setup,
		instance.Setup,
		snapshot.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_filestore creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_filestore(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backup.SetupGated,
		instance.SetupGated,
		snapshot.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
