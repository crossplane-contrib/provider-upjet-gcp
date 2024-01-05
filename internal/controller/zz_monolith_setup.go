/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

<<<<<<< HEAD
	accesslevel "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevel"
	accesslevelcondition "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevelcondition"
	accesspolicy "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesspolicy"
	accesspolicyiammember "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesspolicyiammember"
	serviceperimeter "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeter"
	serviceperimeterresource "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeterresource"
	domain "github.com/upbound/provider-gcp/internal/controller/activedirectory/domain"
	envgroup "github.com/upbound/provider-gcp/internal/controller/apigee/envgroup"
	environment "github.com/upbound/provider-gcp/internal/controller/apigee/environment"
	environmentiammember "github.com/upbound/provider-gcp/internal/controller/apigee/environmentiammember"
	instance "github.com/upbound/provider-gcp/internal/controller/apigee/instance"
	nataddress "github.com/upbound/provider-gcp/internal/controller/apigee/nataddress"
	organization "github.com/upbound/provider-gcp/internal/controller/apigee/organization"
	application "github.com/upbound/provider-gcp/internal/controller/appengine/application"
	applicationurldispatchrules "github.com/upbound/provider-gcp/internal/controller/appengine/applicationurldispatchrules"
	firewallrule "github.com/upbound/provider-gcp/internal/controller/appengine/firewallrule"
	servicenetworksettings "github.com/upbound/provider-gcp/internal/controller/appengine/servicenetworksettings"
	standardappversion "github.com/upbound/provider-gcp/internal/controller/appengine/standardappversion"
	registryrepository "github.com/upbound/provider-gcp/internal/controller/artifact/registryrepository"
	registryrepositoryiammember "github.com/upbound/provider-gcp/internal/controller/artifact/registryrepositoryiammember"
	appconnection "github.com/upbound/provider-gcp/internal/controller/beyondcorp/appconnection"
	appconnector "github.com/upbound/provider-gcp/internal/controller/beyondcorp/appconnector"
	appgateway "github.com/upbound/provider-gcp/internal/controller/beyondcorp/appgateway"
	analyticshubdataexchange "github.com/upbound/provider-gcp/internal/controller/bigquery/analyticshubdataexchange"
	analyticshubdataexchangeiammember "github.com/upbound/provider-gcp/internal/controller/bigquery/analyticshubdataexchangeiammember"
	analyticshublisting "github.com/upbound/provider-gcp/internal/controller/bigquery/analyticshublisting"
	connection "github.com/upbound/provider-gcp/internal/controller/bigquery/connection"
=======
>>>>>>> a3be7bc6 (Remove unneeded resources)
	dataset "github.com/upbound/provider-gcp/internal/controller/bigquery/dataset"
	table "github.com/upbound/provider-gcp/internal/controller/bigquery/table"
	serviceaccount "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/provider-gcp/internal/controller/cloudplatform/servicenetworkingpeereddnsdomain"
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
	reservationcompute "github.com/upbound/provider-gcp/internal/controller/compute/reservation"
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
	recordset "github.com/upbound/provider-gcp/internal/controller/dns/recordset"
	webbackendserviceiammember "github.com/upbound/provider-gcp/internal/controller/iap/webbackendserviceiammember"
	providerconfig "github.com/upbound/provider-gcp/internal/controller/providerconfig"
	database "github.com/upbound/provider-gcp/internal/controller/sql/database"
	databaseinstance "github.com/upbound/provider-gcp/internal/controller/sql/databaseinstance"
	sourcerepresentationinstance "github.com/upbound/provider-gcp/internal/controller/sql/sourcerepresentationinstance"
	sslcert "github.com/upbound/provider-gcp/internal/controller/sql/sslcert"
	user "github.com/upbound/provider-gcp/internal/controller/sql/user"
	bucket "github.com/upbound/provider-gcp/internal/controller/storage/bucket"
	bucketaccesscontrol "github.com/upbound/provider-gcp/internal/controller/storage/bucketaccesscontrol"
	bucketacl "github.com/upbound/provider-gcp/internal/controller/storage/bucketacl"
	bucketiammember "github.com/upbound/provider-gcp/internal/controller/storage/bucketiammember"
	bucketobject "github.com/upbound/provider-gcp/internal/controller/storage/bucketobject"
	defaultobjectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/storage/defaultobjectaccesscontrol"
	defaultobjectacl "github.com/upbound/provider-gcp/internal/controller/storage/defaultobjectacl"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
<<<<<<< HEAD
		accesslevel.Setup,
		accesslevelcondition.Setup,
		accesspolicy.Setup,
		accesspolicyiammember.Setup,
		serviceperimeter.Setup,
		serviceperimeterresource.Setup,
		domain.Setup,
		envgroup.Setup,
		environment.Setup,
		environmentiammember.Setup,
		instance.Setup,
		nataddress.Setup,
		organization.Setup,
		application.Setup,
		applicationurldispatchrules.Setup,
		firewallrule.Setup,
		servicenetworksettings.Setup,
		standardappversion.Setup,
		registryrepository.Setup,
		registryrepositoryiammember.Setup,
		appconnection.Setup,
		appconnector.Setup,
		appgateway.Setup,
		analyticshubdataexchange.Setup,
		analyticshubdataexchangeiammember.Setup,
		analyticshublisting.Setup,
		connection.Setup,
=======
>>>>>>> a3be7bc6 (Remove unneeded resources)
		dataset.Setup,
		table.Setup,
		serviceaccount.Setup,
		serviceaccountiammember.Setup,
		serviceaccountkey.Setup,
		servicenetworkingpeereddnsdomain.Setup,
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
		reservationcompute.Setup,
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
		recordset.Setup,
		webbackendserviceiammember.Setup,
		providerconfig.Setup,
		database.Setup,
		databaseinstance.Setup,
		sourcerepresentationinstance.Setup,
		sslcert.Setup,
		user.Setup,
		bucket.Setup,
		bucketaccesscontrol.Setup,
		bucketacl.Setup,
		bucketiammember.Setup,
		bucketobject.Setup,
		defaultobjectaccesscontrol.Setup,
		defaultobjectacl.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
