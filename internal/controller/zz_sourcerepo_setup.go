/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	repository "github.com/upbound/provider-gcp/internal/controller/sourcerepo/repository"
	repositoryiammember "github.com/upbound/provider-gcp/internal/controller/sourcerepo/repositoryiammember"
)

// Setup_sourcerepo creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sourcerepo(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		repository.Setup,
		repositoryiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
