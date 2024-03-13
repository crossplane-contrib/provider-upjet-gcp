package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	connection "github.com/upbound/provider-gcp/internal/controller/servicenetworking/connection"
)

// Setup_servicenetworking creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicenetworking(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
