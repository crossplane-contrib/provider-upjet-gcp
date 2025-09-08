// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	coderepositoryindex "github.com/upbound/provider-gcp/internal/controller/namespaced/gemini/coderepositoryindex"
	repositorygroup "github.com/upbound/provider-gcp/internal/controller/namespaced/gemini/repositorygroup"
)

// Setup_gemini creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_gemini(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		coderepositoryindex.Setup,
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
		repositorygroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
