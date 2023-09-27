/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	tagbinding "github.com/upbound/provider-gcp/internal/controller/tags/tagbinding"
	tagkey "github.com/upbound/provider-gcp/internal/controller/tags/tagkey"
	tagvalue "github.com/upbound/provider-gcp/internal/controller/tags/tagvalue"
)

// Setup_tags creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_tags(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		tagbinding.Setup,
		tagkey.Setup,
		tagvalue.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
