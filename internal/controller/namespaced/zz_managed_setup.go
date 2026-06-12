// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	kafkaacl "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/managed/kafkaacl"
	kafkacluster "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/managed/kafkacluster"
	kafkatopic "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/managed/kafkatopic"
)

// Setup_managed creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_managed(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		kafkaacl.Setup,
		kafkacluster.Setup,
		kafkatopic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_managed creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_managed(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		kafkaacl.SetupGated,
		kafkacluster.SetupGated,
		kafkatopic.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
