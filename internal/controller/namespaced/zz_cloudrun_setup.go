// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	domainmapping "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/cloudrun/domainmapping"
	service "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/cloudrun/service"
	serviceiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/cloudrun/serviceiammember"
	v2job "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/cloudrun/v2job"
	v2service "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/cloudrun/v2service"
)

// Setup_cloudrun creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudrun(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domainmapping.Setup,
		service.Setup,
		serviceiammember.Setup,
		v2job.Setup,
		v2service.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudrun creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudrun(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		domainmapping.SetupGated,
		service.SetupGated,
		serviceiammember.SetupGated,
		v2job.SetupGated,
		v2service.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_cloudrun registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_cloudrun(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		domainmapping.SetupWebhookWithManager,
		service.SetupWebhookWithManager,
		serviceiammember.SetupWebhookWithManager,
		v2job.SetupWebhookWithManager,
		v2service.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
