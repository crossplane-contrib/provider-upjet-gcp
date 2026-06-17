// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connectionprofile "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/datastream/connectionprofile"
	privateconnection "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/datastream/privateconnection"
	stream "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/datastream/stream"
)

// Setup_datastream creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datastream(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connectionprofile.Setup,
		privateconnection.Setup,
		stream.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_datastream creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_datastream(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connectionprofile.SetupGated,
		privateconnection.SetupGated,
		stream.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_datastream registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_datastream(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		connectionprofile.SetupWebhookWithManager,
		privateconnection.SetupWebhookWithManager,
		stream.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
