// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	folder "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/folder"
	folderiammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/folderiammember"
	organizationiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/organizationiamauditconfig"
	organizationiamcustomrole "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/organizationiamcustomrole"
	organizationiammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/organizationiammember"
	project "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/project"
	projectdefaultserviceaccounts "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectdefaultserviceaccounts"
	projectiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectiamauditconfig"
	projectiamcustomrole "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectiamcustomrole"
	projectiammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectiammember"
	projectservice "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectservice"
	projectusageexportbucket "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectusageexportbucket"
	serviceaccount "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/servicenetworkingpeereddnsdomain"
)

// Setup_cloudplatform creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		folder.Setup,
		folderiammember.Setup,
		organizationiamauditconfig.Setup,
		organizationiamcustomrole.Setup,
		organizationiammember.Setup,
		project.Setup,
		projectdefaultserviceaccounts.Setup,
		projectiamauditconfig.Setup,
		projectiamcustomrole.Setup,
		projectiammember.Setup,
		projectservice.Setup,
		projectusageexportbucket.Setup,
		serviceaccount.Setup,
		serviceaccountiammember.Setup,
		serviceaccountkey.Setup,
		servicenetworkingpeereddnsdomain.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudplatform creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		folder.SetupGated,
		folderiammember.SetupGated,
		organizationiamauditconfig.SetupGated,
		organizationiamcustomrole.SetupGated,
		organizationiammember.SetupGated,
		project.SetupGated,
		projectdefaultserviceaccounts.SetupGated,
		projectiamauditconfig.SetupGated,
		projectiamcustomrole.SetupGated,
		projectiammember.SetupGated,
		projectservice.SetupGated,
		projectusageexportbucket.SetupGated,
		serviceaccount.SetupGated,
		serviceaccountiammember.SetupGated,
		serviceaccountkey.SetupGated,
		servicenetworkingpeereddnsdomain.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
