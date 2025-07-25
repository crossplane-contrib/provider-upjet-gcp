// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	addonsconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/addonsconfig"
	endpointattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/endpointattachment"
	envgroup "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/envgroup"
	envgroupattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/envgroupattachment"
	environment "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/environment"
	environmentiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/environmentiammember"
	instance "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/instance"
	instanceattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/instanceattachment"
	nataddress "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/nataddress"
	organization "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/organization"
	syncauthorization "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/syncauthorization"
)

// Setup_apigee creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apigee(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addonsconfig.Setup,
		endpointattachment.Setup,
		envgroup.Setup,
		envgroupattachment.Setup,
		environment.Setup,
		environmentiammember.Setup,
		instance.Setup,
		instanceattachment.Setup,
		nataddress.Setup,
		organization.Setup,
		syncauthorization.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_apigee creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apigee(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addonsconfig.SetupGated,
		endpointattachment.SetupGated,
		envgroup.SetupGated,
		envgroupattachment.SetupGated,
		environment.SetupGated,
		environmentiammember.SetupGated,
		instance.SetupGated,
		instanceattachment.SetupGated,
		nataddress.SetupGated,
		organization.SetupGated,
		syncauthorization.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
