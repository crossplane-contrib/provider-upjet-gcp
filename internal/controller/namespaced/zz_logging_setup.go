// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	folderbucketconfig "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/logging/folderbucketconfig"
	folderexclusion "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/logging/folderexclusion"
	foldersink "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/logging/foldersink"
	logview "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/logging/logview"
	metric "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/logging/metric"
	projectbucketconfig "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/logging/projectbucketconfig"
	projectexclusion "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/logging/projectexclusion"
	projectsink "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/logging/projectsink"
)

// Setup_logging creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_logging(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		folderbucketconfig.Setup,
		folderexclusion.Setup,
		foldersink.Setup,
		logview.Setup,
		metric.Setup,
		projectbucketconfig.Setup,
		projectexclusion.Setup,
		projectsink.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_logging creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_logging(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		folderbucketconfig.SetupGated,
		folderexclusion.SetupGated,
		foldersink.SetupGated,
		logview.SetupGated,
		metric.SetupGated,
		projectbucketconfig.SetupGated,
		projectexclusion.SetupGated,
		projectsink.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
