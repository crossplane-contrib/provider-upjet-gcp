// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	client "github.com/upbound/provider-gcp/internal/controller/containerazure/client"
	cluster "github.com/upbound/provider-gcp/internal/controller/containerazure/cluster"
	nodepool "github.com/upbound/provider-gcp/internal/controller/containerazure/nodepool"
)

// Setup_containerazure creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerazure(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		client.Setup,
		cluster.Setup,
		nodepool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
