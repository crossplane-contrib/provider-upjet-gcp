// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bucket "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucket"
	bucketaccesscontrol "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketaccesscontrol"
	bucketacl "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketacl"
	bucketiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketiammember"
	bucketiampolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketiampolicy"
	bucketobject "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketobject"
	defaultobjectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/defaultobjectaccesscontrol"
	defaultobjectacl "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/defaultobjectacl"
	hmackey "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/hmackey"
	managedfolder "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/managedfolder"
	managedfolderiambinding "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/managedfolderiambinding"
	managedfolderiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/managedfolderiammember"
	managedfolderiampolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/managedfolderiampolicy"
	notification "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/notification"
	objectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/objectaccesscontrol"
	objectacl "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/objectacl"
)

// Setup_storage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_storage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		bucketaccesscontrol.Setup,
		bucketacl.Setup,
		bucketiammember.Setup,
		bucketiampolicy.Setup,
		bucketobject.Setup,
		defaultobjectaccesscontrol.Setup,
		defaultobjectacl.Setup,
		hmackey.Setup,
		managedfolder.Setup,
		managedfolderiambinding.Setup,
		managedfolderiammember.Setup,
		managedfolderiampolicy.Setup,
		notification.Setup,
		objectaccesscontrol.Setup,
		objectacl.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_storage creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_storage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.SetupGated,
		bucketaccesscontrol.SetupGated,
		bucketacl.SetupGated,
		bucketiammember.SetupGated,
		bucketiampolicy.SetupGated,
		bucketobject.SetupGated,
		defaultobjectaccesscontrol.SetupGated,
		defaultobjectacl.SetupGated,
		hmackey.SetupGated,
		managedfolder.SetupGated,
		managedfolderiambinding.SetupGated,
		managedfolderiammember.SetupGated,
		managedfolderiampolicy.SetupGated,
		notification.SetupGated,
		objectaccesscontrol.SetupGated,
		objectacl.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
