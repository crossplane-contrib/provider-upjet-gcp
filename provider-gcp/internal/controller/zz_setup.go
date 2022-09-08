/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/upbound/official-providers/provider-gcp/internal/controller/appengine/application"
	function "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudfunctions/function"
	functioniammember "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudfunctions/functioniammember"
	folder "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/folder"
	organizationiamauditconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/organizationiamauditconfig"
	organizationiamcustomrole "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/organizationiamcustomrole"
	organizationiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/organizationiammember"
	project "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/project"
	projectdefaultserviceaccounts "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectdefaultserviceaccounts"
	projectiamauditconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectiamauditconfig"
	projectiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectiammember"
	projectservice "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectservice"
	projectusageexportbucket "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/projectusageexportbucket"
	serviceaccount "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/servicenetworkingpeereddnsdomain"
	domainmapping "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudrun/domainmapping"
	service "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudrun/service"
	serviceiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudrun/serviceiammember"
	job "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudscheduler/job"
	queue "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudtasks/queue"
	environment "github.com/upbound/official-providers/provider-gcp/internal/controller/composer/environment"
	address "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/address"
	attacheddisk "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/attacheddisk"
	autoscaler "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/autoscaler"
	backendbucket "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/backendbucket"
	backendbucketsignedurlkey "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/backendbucketsignedurlkey"
	backendservice "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/backendservice"
	disk "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/disk"
	diskiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/diskiammember"
	diskresourcepolicyattachment "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/diskresourcepolicyattachment"
	externalvpngateway "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/externalvpngateway"
	firewall "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/firewall"
	firewallpolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/firewallpolicy"
	firewallpolicyassociation "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/firewallpolicyassociation"
	firewallpolicyrule "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/firewallpolicyrule"
	forwardingrule "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/forwardingrule"
	globaladdress "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/globaladdress"
	globalforwardingrule "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/globalforwardingrule"
	globalnetworkendpoint "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/globalnetworkendpoint"
	globalnetworkendpointgroup "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/globalnetworkendpointgroup"
	havpngateway "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/havpngateway"
	healthcheck "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/healthcheck"
	httphealthcheck "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/httphealthcheck"
	httpshealthcheck "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/httpshealthcheck"
	image "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/image"
	imageiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/imageiammember"
	instance "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instance"
	instancefromtemplate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instancefromtemplate"
	instancegroup "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instancegroup"
	instancegroupmanager "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instancegroupmanager"
	instancegroupnamedport "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instancegroupnamedport"
	instanceiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instanceiammember"
	instancetemplate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instancetemplate"
	interconnectattachment "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/interconnectattachment"
	managedsslcertificate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/managedsslcertificate"
	network "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/network"
	networkendpoint "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/networkendpoint"
	networkendpointgroup "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/networkendpointgroup"
	networkpeering "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/networkpeering"
	networkpeeringroutesconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/networkpeeringroutesconfig"
	nodegroup "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/nodegroup"
	nodetemplate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/nodetemplate"
	packetmirroring "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/packetmirroring"
	perinstanceconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/perinstanceconfig"
	projectdefaultnetworktier "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/projectdefaultnetworktier"
	projectmetadata "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/projectmetadata"
	projectmetadataitem "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/projectmetadataitem"
	regionautoscaler "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regionautoscaler"
	regionbackendservice "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regionbackendservice"
	regiondisk "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regiondisk"
	regiondiskiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regiondiskiammember"
	regiondiskresourcepolicyattachment "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regiondiskresourcepolicyattachment"
	regionhealthcheck "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regionhealthcheck"
	regioninstancegroupmanager "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regioninstancegroupmanager"
	regionnetworkendpointgroup "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regionnetworkendpointgroup"
	regionperinstanceconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regionperinstanceconfig"
	regionsslcertificate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regionsslcertificate"
	regiontargethttpproxy "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regiontargethttpsproxy"
	regionurlmap "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/regionurlmap"
	reservation "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/reservation"
	resourcepolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/resourcepolicy"
	route "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/route"
	router "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/router"
	routerinterface "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/routerinterface"
	routernat "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/routernat"
	sslcertificate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/sslcertificate"
	subnetwork "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/subnetwork"
	targetpool "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/targetpool"
	targetsslproxy "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/targetsslproxy"
	cluster "github.com/upbound/official-providers/provider-gcp/internal/controller/container/cluster"
	nodepool "github.com/upbound/official-providers/provider-gcp/internal/controller/container/nodepool"
	registry "github.com/upbound/official-providers/provider-gcp/internal/controller/container/registry"
	managedzone "github.com/upbound/official-providers/provider-gcp/internal/controller/dns/managedzone"
	policy "github.com/upbound/official-providers/provider-gcp/internal/controller/dns/policy"
	recordset "github.com/upbound/official-providers/provider-gcp/internal/controller/dns/recordset"
	webbackendserviceiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/iap/webbackendserviceiammember"
	webiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/iap/webiammember"
	webtypeappengineiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/iap/webtypeappengineiammember"
	webtypecomputeiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/iap/webtypecomputeiammember"
	defaultsupportedidpconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/identityplatform/defaultsupportedidpconfig"
	inboundsamlconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/identityplatform/inboundsamlconfig"
	oauthidpconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/identityplatform/oauthidpconfig"
	tenant "github.com/upbound/official-providers/provider-gcp/internal/controller/identityplatform/tenant"
	tenantdefaultsupportedidpconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/identityplatform/tenantdefaultsupportedidpconfig"
	tenantinboundsamlconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/identityplatform/tenantinboundsamlconfig"
	tenantoauthidpconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/identityplatform/tenantoauthidpconfig"
	cryptokey "github.com/upbound/official-providers/provider-gcp/internal/controller/kms/cryptokey"
	cryptokeyiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/kms/cryptokeyiammember"
	keyring "github.com/upbound/official-providers/provider-gcp/internal/controller/kms/keyring"
	keyringiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/kms/keyringiammember"
	keyringimportjob "github.com/upbound/official-providers/provider-gcp/internal/controller/kms/keyringimportjob"
	secretciphertext "github.com/upbound/official-providers/provider-gcp/internal/controller/kms/secretciphertext"
	alertpolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/monitoring/alertpolicy"
	notificationchannel "github.com/upbound/official-providers/provider-gcp/internal/controller/monitoring/notificationchannel"
	uptimecheckconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/monitoring/uptimecheckconfig"
	environmentnotebooks "github.com/upbound/official-providers/provider-gcp/internal/controller/notebooks/environment"
	instancenotebooks "github.com/upbound/official-providers/provider-gcp/internal/controller/notebooks/instance"
	instanceiammembernotebooks "github.com/upbound/official-providers/provider-gcp/internal/controller/notebooks/instanceiammember"
	runtime "github.com/upbound/official-providers/provider-gcp/internal/controller/notebooks/runtime"
	runtimeiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/notebooks/runtimeiammember"
	ospolicyassignment "github.com/upbound/official-providers/provider-gcp/internal/controller/osconfig/ospolicyassignment"
	patchdeployment "github.com/upbound/official-providers/provider-gcp/internal/controller/osconfig/patchdeployment"
	sshpublickey "github.com/upbound/official-providers/provider-gcp/internal/controller/oslogin/sshpublickey"
	capool "github.com/upbound/official-providers/provider-gcp/internal/controller/privateca/capool"
	capooliammember "github.com/upbound/official-providers/provider-gcp/internal/controller/privateca/capooliammember"
	certificate "github.com/upbound/official-providers/provider-gcp/internal/controller/privateca/certificate"
	certificateauthority "github.com/upbound/official-providers/provider-gcp/internal/controller/privateca/certificateauthority"
	certificatetemplate "github.com/upbound/official-providers/provider-gcp/internal/controller/privateca/certificatetemplate"
	certificatetemplateiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/privateca/certificatetemplateiammember"
	providerconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/providerconfig"
	litereservation "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/litereservation"
	litesubscription "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/litesubscription"
	litetopic "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/litetopic"
	schema "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/schema"
	subscription "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/subscription"
	subscriptioniammember "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/subscriptioniammember"
	topic "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/topic"
	topiciammember "github.com/upbound/official-providers/provider-gcp/internal/controller/pubsub/topiciammember"
	instanceredis "github.com/upbound/official-providers/provider-gcp/internal/controller/redis/instance"
	secret "github.com/upbound/official-providers/provider-gcp/internal/controller/secretmanager/secret"
	secretiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/secretmanager/secretiammember"
	secretversion "github.com/upbound/official-providers/provider-gcp/internal/controller/secretmanager/secretversion"
	connection "github.com/upbound/official-providers/provider-gcp/internal/controller/servicenetworking/connection"
	database "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/database"
	databaseinstance "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/databaseinstance"
	sourcerepresentationinstance "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/sourcerepresentationinstance"
	sslcert "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/sslcert"
	user "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/user"
	bucket "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/bucket"
	bucketaccesscontrol "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/bucketaccesscontrol"
	bucketacl "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/bucketacl"
	bucketiammember "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/bucketiammember"
	bucketobject "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/bucketobject"
	defaultobjectaccesscontrol "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/defaultobjectaccesscontrol"
	defaultobjectacl "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/defaultobjectacl"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		function.Setup,
		functioniammember.Setup,
		folder.Setup,
		organizationiamauditconfig.Setup,
		organizationiamcustomrole.Setup,
		organizationiammember.Setup,
		project.Setup,
		projectdefaultserviceaccounts.Setup,
		projectiamauditconfig.Setup,
		projectiammember.Setup,
		projectservice.Setup,
		projectusageexportbucket.Setup,
		serviceaccount.Setup,
		serviceaccountiammember.Setup,
		serviceaccountkey.Setup,
		servicenetworkingpeereddnsdomain.Setup,
		domainmapping.Setup,
		service.Setup,
		serviceiammember.Setup,
		job.Setup,
		queue.Setup,
		environment.Setup,
		address.Setup,
		attacheddisk.Setup,
		autoscaler.Setup,
		backendbucket.Setup,
		backendbucketsignedurlkey.Setup,
		backendservice.Setup,
		disk.Setup,
		diskiammember.Setup,
		diskresourcepolicyattachment.Setup,
		externalvpngateway.Setup,
		firewall.Setup,
		firewallpolicy.Setup,
		firewallpolicyassociation.Setup,
		firewallpolicyrule.Setup,
		forwardingrule.Setup,
		globaladdress.Setup,
		globalforwardingrule.Setup,
		globalnetworkendpoint.Setup,
		globalnetworkendpointgroup.Setup,
		havpngateway.Setup,
		healthcheck.Setup,
		httphealthcheck.Setup,
		httpshealthcheck.Setup,
		image.Setup,
		imageiammember.Setup,
		instance.Setup,
		instancefromtemplate.Setup,
		instancegroup.Setup,
		instancegroupmanager.Setup,
		instancegroupnamedport.Setup,
		instanceiammember.Setup,
		instancetemplate.Setup,
		interconnectattachment.Setup,
		managedsslcertificate.Setup,
		network.Setup,
		networkendpoint.Setup,
		networkendpointgroup.Setup,
		networkpeering.Setup,
		networkpeeringroutesconfig.Setup,
		nodegroup.Setup,
		nodetemplate.Setup,
		packetmirroring.Setup,
		perinstanceconfig.Setup,
		projectdefaultnetworktier.Setup,
		projectmetadata.Setup,
		projectmetadataitem.Setup,
		regionautoscaler.Setup,
		regionbackendservice.Setup,
		regiondisk.Setup,
		regiondiskiammember.Setup,
		regiondiskresourcepolicyattachment.Setup,
		regionhealthcheck.Setup,
		regioninstancegroupmanager.Setup,
		regionnetworkendpointgroup.Setup,
		regionperinstanceconfig.Setup,
		regionsslcertificate.Setup,
		regiontargethttpproxy.Setup,
		regiontargethttpsproxy.Setup,
		regionurlmap.Setup,
		reservation.Setup,
		resourcepolicy.Setup,
		route.Setup,
		router.Setup,
		routerinterface.Setup,
		routernat.Setup,
		sslcertificate.Setup,
		subnetwork.Setup,
		targetpool.Setup,
		targetsslproxy.Setup,
		cluster.Setup,
		nodepool.Setup,
		registry.Setup,
		managedzone.Setup,
		policy.Setup,
		recordset.Setup,
		webbackendserviceiammember.Setup,
		webiammember.Setup,
		webtypeappengineiammember.Setup,
		webtypecomputeiammember.Setup,
		defaultsupportedidpconfig.Setup,
		inboundsamlconfig.Setup,
		oauthidpconfig.Setup,
		tenant.Setup,
		tenantdefaultsupportedidpconfig.Setup,
		tenantinboundsamlconfig.Setup,
		tenantoauthidpconfig.Setup,
		cryptokey.Setup,
		cryptokeyiammember.Setup,
		keyring.Setup,
		keyringiammember.Setup,
		keyringimportjob.Setup,
		secretciphertext.Setup,
		alertpolicy.Setup,
		notificationchannel.Setup,
		uptimecheckconfig.Setup,
		environmentnotebooks.Setup,
		instancenotebooks.Setup,
		instanceiammembernotebooks.Setup,
		runtime.Setup,
		runtimeiammember.Setup,
		ospolicyassignment.Setup,
		patchdeployment.Setup,
		sshpublickey.Setup,
		capool.Setup,
		capooliammember.Setup,
		certificate.Setup,
		certificateauthority.Setup,
		certificatetemplate.Setup,
		certificatetemplateiammember.Setup,
		providerconfig.Setup,
		litereservation.Setup,
		litesubscription.Setup,
		litetopic.Setup,
		schema.Setup,
		subscription.Setup,
		subscriptioniammember.Setup,
		topic.Setup,
		topiciammember.Setup,
		instanceredis.Setup,
		secret.Setup,
		secretiammember.Setup,
		secretversion.Setup,
		connection.Setup,
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
