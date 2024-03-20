// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	litereservation "github.com/upbound/provider-gcp/internal/controller/pubsub/litereservation"
	litesubscription "github.com/upbound/provider-gcp/internal/controller/pubsub/litesubscription"
	litetopic "github.com/upbound/provider-gcp/internal/controller/pubsub/litetopic"
	schema "github.com/upbound/provider-gcp/internal/controller/pubsub/schema"
	subscription "github.com/upbound/provider-gcp/internal/controller/pubsub/subscription"
	subscriptioniammember "github.com/upbound/provider-gcp/internal/controller/pubsub/subscriptioniammember"
	topic "github.com/upbound/provider-gcp/internal/controller/pubsub/topic"
	topiciammember "github.com/upbound/provider-gcp/internal/controller/pubsub/topiciammember"
)

// Setup_pubsub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_pubsub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		litereservation.Setup,
		litesubscription.Setup,
		litetopic.Setup,
		schema.Setup,
		subscription.Setup,
		subscriptioniammember.Setup,
		topic.Setup,
		topiciammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
