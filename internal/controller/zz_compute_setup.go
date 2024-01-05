/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	address "github.com/upbound/provider-gcp/internal/controller/compute/address"
	backendservice "github.com/upbound/provider-gcp/internal/controller/compute/backendservice"
	forwardingrule "github.com/upbound/provider-gcp/internal/controller/compute/forwardingrule"
	globaladdress "github.com/upbound/provider-gcp/internal/controller/compute/globaladdress"
	globalforwardingrule "github.com/upbound/provider-gcp/internal/controller/compute/globalforwardingrule"
	healthcheck "github.com/upbound/provider-gcp/internal/controller/compute/healthcheck"
	httphealthcheck "github.com/upbound/provider-gcp/internal/controller/compute/httphealthcheck"
	httpshealthcheck "github.com/upbound/provider-gcp/internal/controller/compute/httpshealthcheck"
	instancegroupmanager "github.com/upbound/provider-gcp/internal/controller/compute/instancegroupmanager"
	instancetemplate "github.com/upbound/provider-gcp/internal/controller/compute/instancetemplate"
	network "github.com/upbound/provider-gcp/internal/controller/compute/network"
	regionbackendservice "github.com/upbound/provider-gcp/internal/controller/compute/regionbackendservice"
	regionhealthcheck "github.com/upbound/provider-gcp/internal/controller/compute/regionhealthcheck"
	regioninstancegroupmanager "github.com/upbound/provider-gcp/internal/controller/compute/regioninstancegroupmanager"
	regionsslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/regionsslcertificate"
	regiontargethttpproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpsproxy"
	regionurlmap "github.com/upbound/provider-gcp/internal/controller/compute/regionurlmap"
<<<<<<< HEAD
	reservation "github.com/upbound/provider-gcp/internal/controller/compute/reservation"
	resourcepolicy "github.com/upbound/provider-gcp/internal/controller/compute/resourcepolicy"
	route "github.com/upbound/provider-gcp/internal/controller/compute/route"
	router "github.com/upbound/provider-gcp/internal/controller/compute/router"
	routerinterface "github.com/upbound/provider-gcp/internal/controller/compute/routerinterface"
	routernat "github.com/upbound/provider-gcp/internal/controller/compute/routernat"
	routerpeer "github.com/upbound/provider-gcp/internal/controller/compute/routerpeer"
	securitypolicy "github.com/upbound/provider-gcp/internal/controller/compute/securitypolicy"
	serviceattachment "github.com/upbound/provider-gcp/internal/controller/compute/serviceattachment"
	sharedvpchostproject "github.com/upbound/provider-gcp/internal/controller/compute/sharedvpchostproject"
	sharedvpcserviceproject "github.com/upbound/provider-gcp/internal/controller/compute/sharedvpcserviceproject"
	snapshot "github.com/upbound/provider-gcp/internal/controller/compute/snapshot"
	snapshotiammember "github.com/upbound/provider-gcp/internal/controller/compute/snapshotiammember"
=======
>>>>>>> a3be7bc6 (Remove unneeded resources)
	sslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/sslcertificate"
	subnetwork "github.com/upbound/provider-gcp/internal/controller/compute/subnetwork"
	targethttpproxy "github.com/upbound/provider-gcp/internal/controller/compute/targethttpproxy"
	targethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/targethttpsproxy"
	targetpool "github.com/upbound/provider-gcp/internal/controller/compute/targetpool"
)

// Setup_compute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		address.Setup,
		backendservice.Setup,
		forwardingrule.Setup,
		globaladdress.Setup,
		globalforwardingrule.Setup,
		healthcheck.Setup,
		httphealthcheck.Setup,
		httpshealthcheck.Setup,
		instancegroupmanager.Setup,
		instancetemplate.Setup,
		network.Setup,
		regionbackendservice.Setup,
		regionhealthcheck.Setup,
		regioninstancegroupmanager.Setup,
		regionsslcertificate.Setup,
		regiontargethttpproxy.Setup,
		regiontargethttpsproxy.Setup,
		regionurlmap.Setup,
<<<<<<< HEAD
		reservation.Setup,
		resourcepolicy.Setup,
		route.Setup,
		router.Setup,
		routerinterface.Setup,
		routernat.Setup,
		routerpeer.Setup,
		securitypolicy.Setup,
		serviceattachment.Setup,
		sharedvpchostproject.Setup,
		sharedvpcserviceproject.Setup,
		snapshot.Setup,
		snapshotiammember.Setup,
=======
>>>>>>> a3be7bc6 (Remove unneeded resources)
		sslcertificate.Setup,
		subnetwork.Setup,
		targethttpproxy.Setup,
		targethttpsproxy.Setup,
		targetpool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
