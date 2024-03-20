// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	folder "github.com/upbound/provider-gcp/internal/controller/cloudplatform/folder"
	folderiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/folderiammember"
	organizationiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cloudplatform/organizationiamauditconfig"
	organizationiamcustomrole "github.com/upbound/provider-gcp/internal/controller/cloudplatform/organizationiamcustomrole"
	organizationiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/organizationiammember"
	project "github.com/upbound/provider-gcp/internal/controller/cloudplatform/project"
	projectdefaultserviceaccounts "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectdefaultserviceaccounts"
	projectiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectiamauditconfig"
	projectiamcustomrole "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectiamcustomrole"
	projectiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectiammember"
	projectservice "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectservice"
	projectusageexportbucket "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectusageexportbucket"
	serviceaccount "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/provider-gcp/internal/controller/cloudplatform/servicenetworkingpeereddnsdomain"
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
