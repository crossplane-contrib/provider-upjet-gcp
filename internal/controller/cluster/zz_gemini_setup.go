// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	coderepositoryindex "github.com/upbound/provider-gcp/v2/internal/controller/cluster/gemini/coderepositoryindex"
	codetoolssetting "github.com/upbound/provider-gcp/v2/internal/controller/cluster/gemini/codetoolssetting"
	datasharingwithgooglesetting "github.com/upbound/provider-gcp/v2/internal/controller/cluster/gemini/datasharingwithgooglesetting"
	geminigcpenablementsetting "github.com/upbound/provider-gcp/v2/internal/controller/cluster/gemini/geminigcpenablementsetting"
	loggingsetting "github.com/upbound/provider-gcp/v2/internal/controller/cluster/gemini/loggingsetting"
	releasechannelsetting "github.com/upbound/provider-gcp/v2/internal/controller/cluster/gemini/releasechannelsetting"
	repositorygroup "github.com/upbound/provider-gcp/v2/internal/controller/cluster/gemini/repositorygroup"
)

// Setup_gemini creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_gemini(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		coderepositoryindex.Setup,
		codetoolssetting.Setup,
		datasharingwithgooglesetting.Setup,
		geminigcpenablementsetting.Setup,
		loggingsetting.Setup,
		releasechannelsetting.Setup,
		repositorygroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_gemini creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_gemini(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		coderepositoryindex.SetupGated,
		codetoolssetting.SetupGated,
		datasharingwithgooglesetting.SetupGated,
		geminigcpenablementsetting.SetupGated,
		loggingsetting.SetupGated,
		releasechannelsetting.SetupGated,
		repositorygroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager_gemini registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager_gemini(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		coderepositoryindex.SetupWebhookWithManager,
		codetoolssetting.SetupWebhookWithManager,
		datasharingwithgooglesetting.SetupWebhookWithManager,
		geminigcpenablementsetting.SetupWebhookWithManager,
		loggingsetting.SetupWebhookWithManager,
		releasechannelsetting.SetupWebhookWithManager,
		repositorygroup.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
