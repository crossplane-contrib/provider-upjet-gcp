/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	backup "github.com/upbound/provider-gcp/internal/controller/filestore/backup"
	instance "github.com/upbound/provider-gcp/internal/controller/filestore/instance"
	snapshot "github.com/upbound/provider-gcp/internal/controller/filestore/snapshot"
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
