// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	bucket "github.com/upbound/provider-gcp/internal/controller/storage/bucket"
	bucketaccesscontrol "github.com/upbound/provider-gcp/internal/controller/storage/bucketaccesscontrol"
	bucketacl "github.com/upbound/provider-gcp/internal/controller/storage/bucketacl"
	bucketiammember "github.com/upbound/provider-gcp/internal/controller/storage/bucketiammember"
	bucketobject "github.com/upbound/provider-gcp/internal/controller/storage/bucketobject"
	defaultobjectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/storage/defaultobjectaccesscontrol"
	defaultobjectacl "github.com/upbound/provider-gcp/internal/controller/storage/defaultobjectacl"
	hmackey "github.com/upbound/provider-gcp/internal/controller/storage/hmackey"
	notification "github.com/upbound/provider-gcp/internal/controller/storage/notification"
	objectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/storage/objectaccesscontrol"
	objectacl "github.com/upbound/provider-gcp/internal/controller/storage/objectacl"
)

// Setup_storage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_storage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bucket.Setup,
		bucketaccesscontrol.Setup,
		bucketacl.Setup,
		bucketiammember.Setup,
		bucketobject.Setup,
		defaultobjectaccesscontrol.Setup,
		defaultobjectacl.Setup,
		hmackey.Setup,
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
