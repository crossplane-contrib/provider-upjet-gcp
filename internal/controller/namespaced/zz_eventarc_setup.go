// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	channel "github.com/upbound/provider-gcp/internal/controller/namespaced/eventarc/channel"
	googlechannelconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/eventarc/googlechannelconfig"
	trigger "github.com/upbound/provider-gcp/internal/controller/namespaced/eventarc/trigger"
)

// Setup_eventarc creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eventarc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		channel.Setup,
		googlechannelconfig.Setup,
		trigger.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_eventarc creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_eventarc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		channel.SetupGated,
		googlechannelconfig.SetupGated,
		trigger.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
