// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alertpolicy "github.com/upbound/provider-gcp/internal/controller/monitoring/alertpolicy"
	customservice "github.com/upbound/provider-gcp/internal/controller/monitoring/customservice"
	dashboard "github.com/upbound/provider-gcp/internal/controller/monitoring/dashboard"
	group "github.com/upbound/provider-gcp/internal/controller/monitoring/group"
	metricdescriptor "github.com/upbound/provider-gcp/internal/controller/monitoring/metricdescriptor"
	notificationchannel "github.com/upbound/provider-gcp/internal/controller/monitoring/notificationchannel"
	service "github.com/upbound/provider-gcp/internal/controller/monitoring/service"
	slo "github.com/upbound/provider-gcp/internal/controller/monitoring/slo"
	uptimecheckconfig "github.com/upbound/provider-gcp/internal/controller/monitoring/uptimecheckconfig"
)

// Setup_monitoring creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monitoring(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertpolicy.Setup,
		customservice.Setup,
		dashboard.Setup,
		group.Setup,
		metricdescriptor.Setup,
		notificationchannel.Setup,
		service.Setup,
		slo.Setup,
		uptimecheckconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
