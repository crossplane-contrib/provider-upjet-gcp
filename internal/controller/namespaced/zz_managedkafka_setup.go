// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	acl "github.com/upbound/provider-gcp/v3/internal/controller/namespaced/managedkafka/acl"
	cluster "github.com/upbound/provider-gcp/v3/internal/controller/namespaced/managedkafka/cluster"
	connectcluster "github.com/upbound/provider-gcp/v3/internal/controller/namespaced/managedkafka/connectcluster"
	connector "github.com/upbound/provider-gcp/v3/internal/controller/namespaced/managedkafka/connector"
	topic "github.com/upbound/provider-gcp/v3/internal/controller/namespaced/managedkafka/topic"
)

// Setup_managedkafka creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_managedkafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.Setup,
		cluster.Setup,
		connectcluster.Setup,
		connector.Setup,
		topic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_managedkafka creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_managedkafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.SetupGated,
		cluster.SetupGated,
		connectcluster.SetupGated,
		connector.SetupGated,
		topic.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_managedkafka registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_managedkafka(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		acl.SetupWebhookWithManager,
		cluster.SetupWebhookWithManager,
		connectcluster.SetupWebhookWithManager,
		connector.SetupWebhookWithManager,
		topic.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
