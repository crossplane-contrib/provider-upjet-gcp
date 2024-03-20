// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	envgroup "github.com/upbound/provider-gcp/internal/controller/apigee/envgroup"
	environment "github.com/upbound/provider-gcp/internal/controller/apigee/environment"
	environmentiammember "github.com/upbound/provider-gcp/internal/controller/apigee/environmentiammember"
	instance "github.com/upbound/provider-gcp/internal/controller/apigee/instance"
	nataddress "github.com/upbound/provider-gcp/internal/controller/apigee/nataddress"
	organization "github.com/upbound/provider-gcp/internal/controller/apigee/organization"
)

// Setup_apigee creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apigee(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		envgroup.Setup,
		environment.Setup,
		environmentiammember.Setup,
		instance.Setup,
		nataddress.Setup,
		organization.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
