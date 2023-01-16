/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/upbound/provider-gcp/internal/controller/appengine/application"
	connection "github.com/upbound/provider-gcp/internal/controller/bigquery/connection"
	dataset "github.com/upbound/provider-gcp/internal/controller/bigquery/dataset"
	datasetaccess "github.com/upbound/provider-gcp/internal/controller/bigquery/datasetaccess"
	datasetiambinding "github.com/upbound/provider-gcp/internal/controller/bigquery/datasetiambinding"
	datasetiammember "github.com/upbound/provider-gcp/internal/controller/bigquery/datasetiammember"
	datasetiampolicy "github.com/upbound/provider-gcp/internal/controller/bigquery/datasetiampolicy"
	datatransferconfig "github.com/upbound/provider-gcp/internal/controller/bigquery/datatransferconfig"
	job "github.com/upbound/provider-gcp/internal/controller/bigquery/job"
	reservation "github.com/upbound/provider-gcp/internal/controller/bigquery/reservation"
	reservationassignment "github.com/upbound/provider-gcp/internal/controller/bigquery/reservationassignment"
	routine "github.com/upbound/provider-gcp/internal/controller/bigquery/routine"
	table "github.com/upbound/provider-gcp/internal/controller/bigquery/table"
	tableiambinding "github.com/upbound/provider-gcp/internal/controller/bigquery/tableiambinding"
	tableiammember "github.com/upbound/provider-gcp/internal/controller/bigquery/tableiammember"
	tableiampolicy "github.com/upbound/provider-gcp/internal/controller/bigquery/tableiampolicy"
	trigger "github.com/upbound/provider-gcp/internal/controller/cloudbuild/trigger"
	workerpool "github.com/upbound/provider-gcp/internal/controller/cloudbuild/workerpool"
	function "github.com/upbound/provider-gcp/internal/controller/cloudfunctions/function"
	functioniammember "github.com/upbound/provider-gcp/internal/controller/cloudfunctions/functioniammember"
	device "github.com/upbound/provider-gcp/internal/controller/cloudiot/device"
	registry "github.com/upbound/provider-gcp/internal/controller/cloudiot/registry"
	folder "github.com/upbound/provider-gcp/internal/controller/cloudplatform/folder"
	folderiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/folderiammember"
	organizationiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cloudplatform/organizationiamauditconfig"
	organizationiamcustomrole "github.com/upbound/provider-gcp/internal/controller/cloudplatform/organizationiamcustomrole"
	organizationiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/organizationiammember"
	project "github.com/upbound/provider-gcp/internal/controller/cloudplatform/project"
	projectdefaultserviceaccounts "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectdefaultserviceaccounts"
	projectiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectiamauditconfig"
	projectiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectiammember"
	projectservice "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectservice"
	projectusageexportbucket "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectusageexportbucket"
	serviceaccount "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/provider-gcp/internal/controller/cloudplatform/servicenetworkingpeereddnsdomain"
	domainmapping "github.com/upbound/provider-gcp/internal/controller/cloudrun/domainmapping"
	service "github.com/upbound/provider-gcp/internal/controller/cloudrun/service"
	serviceiammember "github.com/upbound/provider-gcp/internal/controller/cloudrun/serviceiammember"
	jobcloudscheduler "github.com/upbound/provider-gcp/internal/controller/cloudscheduler/job"
	queue "github.com/upbound/provider-gcp/internal/controller/cloudtasks/queue"
	environment "github.com/upbound/provider-gcp/internal/controller/composer/environment"
	address "github.com/upbound/provider-gcp/internal/controller/compute/address"
	attacheddisk "github.com/upbound/provider-gcp/internal/controller/compute/attacheddisk"
	autoscaler "github.com/upbound/provider-gcp/internal/controller/compute/autoscaler"
	backendbucket "github.com/upbound/provider-gcp/internal/controller/compute/backendbucket"
	backendbucketsignedurlkey "github.com/upbound/provider-gcp/internal/controller/compute/backendbucketsignedurlkey"
	backendservice "github.com/upbound/provider-gcp/internal/controller/compute/backendservice"
	disk "github.com/upbound/provider-gcp/internal/controller/compute/disk"
	diskiammember "github.com/upbound/provider-gcp/internal/controller/compute/diskiammember"
	diskresourcepolicyattachment "github.com/upbound/provider-gcp/internal/controller/compute/diskresourcepolicyattachment"
	externalvpngateway "github.com/upbound/provider-gcp/internal/controller/compute/externalvpngateway"
	firewall "github.com/upbound/provider-gcp/internal/controller/compute/firewall"
	firewallpolicy "github.com/upbound/provider-gcp/internal/controller/compute/firewallpolicy"
	firewallpolicyassociation "github.com/upbound/provider-gcp/internal/controller/compute/firewallpolicyassociation"
	firewallpolicyrule "github.com/upbound/provider-gcp/internal/controller/compute/firewallpolicyrule"
	forwardingrule "github.com/upbound/provider-gcp/internal/controller/compute/forwardingrule"
	globaladdress "github.com/upbound/provider-gcp/internal/controller/compute/globaladdress"
	globalforwardingrule "github.com/upbound/provider-gcp/internal/controller/compute/globalforwardingrule"
	globalnetworkendpoint "github.com/upbound/provider-gcp/internal/controller/compute/globalnetworkendpoint"
	globalnetworkendpointgroup "github.com/upbound/provider-gcp/internal/controller/compute/globalnetworkendpointgroup"
	havpngateway "github.com/upbound/provider-gcp/internal/controller/compute/havpngateway"
	healthcheck "github.com/upbound/provider-gcp/internal/controller/compute/healthcheck"
	httphealthcheck "github.com/upbound/provider-gcp/internal/controller/compute/httphealthcheck"
	httpshealthcheck "github.com/upbound/provider-gcp/internal/controller/compute/httpshealthcheck"
	image "github.com/upbound/provider-gcp/internal/controller/compute/image"
	imageiammember "github.com/upbound/provider-gcp/internal/controller/compute/imageiammember"
	instance "github.com/upbound/provider-gcp/internal/controller/compute/instance"
	instancefromtemplate "github.com/upbound/provider-gcp/internal/controller/compute/instancefromtemplate"
	instancegroup "github.com/upbound/provider-gcp/internal/controller/compute/instancegroup"
	instancegroupmanager "github.com/upbound/provider-gcp/internal/controller/compute/instancegroupmanager"
	instancegroupnamedport "github.com/upbound/provider-gcp/internal/controller/compute/instancegroupnamedport"
	instanceiammember "github.com/upbound/provider-gcp/internal/controller/compute/instanceiammember"
	instancetemplate "github.com/upbound/provider-gcp/internal/controller/compute/instancetemplate"
	interconnectattachment "github.com/upbound/provider-gcp/internal/controller/compute/interconnectattachment"
	managedsslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/managedsslcertificate"
	network "github.com/upbound/provider-gcp/internal/controller/compute/network"
	networkendpoint "github.com/upbound/provider-gcp/internal/controller/compute/networkendpoint"
	networkendpointgroup "github.com/upbound/provider-gcp/internal/controller/compute/networkendpointgroup"
	networkpeering "github.com/upbound/provider-gcp/internal/controller/compute/networkpeering"
	networkpeeringroutesconfig "github.com/upbound/provider-gcp/internal/controller/compute/networkpeeringroutesconfig"
	nodegroup "github.com/upbound/provider-gcp/internal/controller/compute/nodegroup"
	nodetemplate "github.com/upbound/provider-gcp/internal/controller/compute/nodetemplate"
	packetmirroring "github.com/upbound/provider-gcp/internal/controller/compute/packetmirroring"
	perinstanceconfig "github.com/upbound/provider-gcp/internal/controller/compute/perinstanceconfig"
	projectdefaultnetworktier "github.com/upbound/provider-gcp/internal/controller/compute/projectdefaultnetworktier"
	projectmetadata "github.com/upbound/provider-gcp/internal/controller/compute/projectmetadata"
	projectmetadataitem "github.com/upbound/provider-gcp/internal/controller/compute/projectmetadataitem"
	regionautoscaler "github.com/upbound/provider-gcp/internal/controller/compute/regionautoscaler"
	regionbackendservice "github.com/upbound/provider-gcp/internal/controller/compute/regionbackendservice"
	regiondisk "github.com/upbound/provider-gcp/internal/controller/compute/regiondisk"
	regiondiskiammember "github.com/upbound/provider-gcp/internal/controller/compute/regiondiskiammember"
	regiondiskresourcepolicyattachment "github.com/upbound/provider-gcp/internal/controller/compute/regiondiskresourcepolicyattachment"
	regionhealthcheck "github.com/upbound/provider-gcp/internal/controller/compute/regionhealthcheck"
	regioninstancegroupmanager "github.com/upbound/provider-gcp/internal/controller/compute/regioninstancegroupmanager"
	regionnetworkendpointgroup "github.com/upbound/provider-gcp/internal/controller/compute/regionnetworkendpointgroup"
	regionperinstanceconfig "github.com/upbound/provider-gcp/internal/controller/compute/regionperinstanceconfig"
	regionsslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/regionsslcertificate"
	regiontargethttpproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpsproxy"
	regionurlmap "github.com/upbound/provider-gcp/internal/controller/compute/regionurlmap"
	reservationcompute "github.com/upbound/provider-gcp/internal/controller/compute/reservation"
	resourcepolicy "github.com/upbound/provider-gcp/internal/controller/compute/resourcepolicy"
	route "github.com/upbound/provider-gcp/internal/controller/compute/route"
	router "github.com/upbound/provider-gcp/internal/controller/compute/router"
	routerinterface "github.com/upbound/provider-gcp/internal/controller/compute/routerinterface"
	routernat "github.com/upbound/provider-gcp/internal/controller/compute/routernat"
	securitypolicy "github.com/upbound/provider-gcp/internal/controller/compute/securitypolicy"
	serviceattachment "github.com/upbound/provider-gcp/internal/controller/compute/serviceattachment"
	sslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/sslcertificate"
	subnetwork "github.com/upbound/provider-gcp/internal/controller/compute/subnetwork"
	subnetworkiammember "github.com/upbound/provider-gcp/internal/controller/compute/subnetworkiammember"
	targetgrpcproxy "github.com/upbound/provider-gcp/internal/controller/compute/targetgrpcproxy"
	targethttpproxy "github.com/upbound/provider-gcp/internal/controller/compute/targethttpproxy"
	targethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/targethttpsproxy"
	targetinstance "github.com/upbound/provider-gcp/internal/controller/compute/targetinstance"
	targetpool "github.com/upbound/provider-gcp/internal/controller/compute/targetpool"
	targetsslproxy "github.com/upbound/provider-gcp/internal/controller/compute/targetsslproxy"
	targettcpproxy "github.com/upbound/provider-gcp/internal/controller/compute/targettcpproxy"
	urlmap "github.com/upbound/provider-gcp/internal/controller/compute/urlmap"
	vpngateway "github.com/upbound/provider-gcp/internal/controller/compute/vpngateway"
	vpntunnel "github.com/upbound/provider-gcp/internal/controller/compute/vpntunnel"
	cluster "github.com/upbound/provider-gcp/internal/controller/container/cluster"
	nodepool "github.com/upbound/provider-gcp/internal/controller/container/nodepool"
	registrycontainer "github.com/upbound/provider-gcp/internal/controller/container/registry"
	note "github.com/upbound/provider-gcp/internal/controller/containeranalysis/note"
	clustercontaineraws "github.com/upbound/provider-gcp/internal/controller/containeraws/cluster"
	nodepoolcontaineraws "github.com/upbound/provider-gcp/internal/controller/containeraws/nodepool"
	client "github.com/upbound/provider-gcp/internal/controller/containerazure/client"
	clustercontainerazure "github.com/upbound/provider-gcp/internal/controller/containerazure/cluster"
	nodepoolcontainerazure "github.com/upbound/provider-gcp/internal/controller/containerazure/nodepool"
	entry "github.com/upbound/provider-gcp/internal/controller/datacatalog/entry"
	entrygroup "github.com/upbound/provider-gcp/internal/controller/datacatalog/entrygroup"
	entrygroupiambinding "github.com/upbound/provider-gcp/internal/controller/datacatalog/entrygroupiambinding"
	entrygroupiammember "github.com/upbound/provider-gcp/internal/controller/datacatalog/entrygroupiammember"
	entrygroupiampolicy "github.com/upbound/provider-gcp/internal/controller/datacatalog/entrygroupiampolicy"
	tag "github.com/upbound/provider-gcp/internal/controller/datacatalog/tag"
	tagtemplate "github.com/upbound/provider-gcp/internal/controller/datacatalog/tagtemplate"
	tagtemplateiambinding "github.com/upbound/provider-gcp/internal/controller/datacatalog/tagtemplateiambinding"
	tagtemplateiammember "github.com/upbound/provider-gcp/internal/controller/datacatalog/tagtemplateiammember"
	tagtemplateiampolicy "github.com/upbound/provider-gcp/internal/controller/datacatalog/tagtemplateiampolicy"
	jobdataflow "github.com/upbound/provider-gcp/internal/controller/dataflow/job"
	instancedatafusion "github.com/upbound/provider-gcp/internal/controller/datafusion/instance"
	agent "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/agent"
	entitytype "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/entitytype"
	environmentdialogflowcx "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/environment"
	flow "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/flow"
	intent "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/intent"
	page "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/page"
	version "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/version"
	managedzone "github.com/upbound/provider-gcp/internal/controller/dns/managedzone"
	policy "github.com/upbound/provider-gcp/internal/controller/dns/policy"
	recordset "github.com/upbound/provider-gcp/internal/controller/dns/recordset"
	contact "github.com/upbound/provider-gcp/internal/controller/essentialcontacts/contact"
	triggereventarc "github.com/upbound/provider-gcp/internal/controller/eventarc/trigger"
	instancefilestore "github.com/upbound/provider-gcp/internal/controller/filestore/instance"
	release "github.com/upbound/provider-gcp/internal/controller/firebaserules/release"
	ruleset "github.com/upbound/provider-gcp/internal/controller/firebaserules/ruleset"
	membership "github.com/upbound/provider-gcp/internal/controller/gkehub/membership"
	consentstore "github.com/upbound/provider-gcp/internal/controller/healthcare/consentstore"
	datasethealthcare "github.com/upbound/provider-gcp/internal/controller/healthcare/dataset"
	webbackendserviceiammember "github.com/upbound/provider-gcp/internal/controller/iap/webbackendserviceiammember"
	webiammember "github.com/upbound/provider-gcp/internal/controller/iap/webiammember"
	webtypeappengineiammember "github.com/upbound/provider-gcp/internal/controller/iap/webtypeappengineiammember"
	webtypecomputeiammember "github.com/upbound/provider-gcp/internal/controller/iap/webtypecomputeiammember"
	defaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/defaultsupportedidpconfig"
	inboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/inboundsamlconfig"
	oauthidpconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/oauthidpconfig"
	tenant "github.com/upbound/provider-gcp/internal/controller/identityplatform/tenant"
	tenantdefaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/tenantdefaultsupportedidpconfig"
	tenantinboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/tenantinboundsamlconfig"
	tenantoauthidpconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/tenantoauthidpconfig"
	cryptokey "github.com/upbound/provider-gcp/internal/controller/kms/cryptokey"
	cryptokeyiammember "github.com/upbound/provider-gcp/internal/controller/kms/cryptokeyiammember"
	keyring "github.com/upbound/provider-gcp/internal/controller/kms/keyring"
	keyringiammember "github.com/upbound/provider-gcp/internal/controller/kms/keyringiammember"
	keyringimportjob "github.com/upbound/provider-gcp/internal/controller/kms/keyringimportjob"
	secretciphertext "github.com/upbound/provider-gcp/internal/controller/kms/secretciphertext"
	alertpolicy "github.com/upbound/provider-gcp/internal/controller/monitoring/alertpolicy"
	notificationchannel "github.com/upbound/provider-gcp/internal/controller/monitoring/notificationchannel"
	uptimecheckconfig "github.com/upbound/provider-gcp/internal/controller/monitoring/uptimecheckconfig"
	environmentnotebooks "github.com/upbound/provider-gcp/internal/controller/notebooks/environment"
	instancenotebooks "github.com/upbound/provider-gcp/internal/controller/notebooks/instance"
	instanceiammembernotebooks "github.com/upbound/provider-gcp/internal/controller/notebooks/instanceiammember"
	runtime "github.com/upbound/provider-gcp/internal/controller/notebooks/runtime"
	runtimeiammember "github.com/upbound/provider-gcp/internal/controller/notebooks/runtimeiammember"
	ospolicyassignment "github.com/upbound/provider-gcp/internal/controller/osconfig/ospolicyassignment"
	patchdeployment "github.com/upbound/provider-gcp/internal/controller/osconfig/patchdeployment"
	sshpublickey "github.com/upbound/provider-gcp/internal/controller/oslogin/sshpublickey"
	capool "github.com/upbound/provider-gcp/internal/controller/privateca/capool"
	capooliammember "github.com/upbound/provider-gcp/internal/controller/privateca/capooliammember"
	certificate "github.com/upbound/provider-gcp/internal/controller/privateca/certificate"
	certificateauthority "github.com/upbound/provider-gcp/internal/controller/privateca/certificateauthority"
	certificatetemplate "github.com/upbound/provider-gcp/internal/controller/privateca/certificatetemplate"
	certificatetemplateiammember "github.com/upbound/provider-gcp/internal/controller/privateca/certificatetemplateiammember"
	providerconfig "github.com/upbound/provider-gcp/internal/controller/providerconfig"
	litereservation "github.com/upbound/provider-gcp/internal/controller/pubsub/litereservation"
	litesubscription "github.com/upbound/provider-gcp/internal/controller/pubsub/litesubscription"
	litetopic "github.com/upbound/provider-gcp/internal/controller/pubsub/litetopic"
	schema "github.com/upbound/provider-gcp/internal/controller/pubsub/schema"
	subscription "github.com/upbound/provider-gcp/internal/controller/pubsub/subscription"
	subscriptioniammember "github.com/upbound/provider-gcp/internal/controller/pubsub/subscriptioniammember"
	topic "github.com/upbound/provider-gcp/internal/controller/pubsub/topic"
	topiciammember "github.com/upbound/provider-gcp/internal/controller/pubsub/topiciammember"
	instanceredis "github.com/upbound/provider-gcp/internal/controller/redis/instance"
	secret "github.com/upbound/provider-gcp/internal/controller/secretmanager/secret"
	secretiammember "github.com/upbound/provider-gcp/internal/controller/secretmanager/secretiammember"
	secretversion "github.com/upbound/provider-gcp/internal/controller/secretmanager/secretversion"
	connectionservicenetworking "github.com/upbound/provider-gcp/internal/controller/servicenetworking/connection"
	repository "github.com/upbound/provider-gcp/internal/controller/sourcerepo/repository"
	repositoryiammember "github.com/upbound/provider-gcp/internal/controller/sourcerepo/repositoryiammember"
	database "github.com/upbound/provider-gcp/internal/controller/spanner/database"
	databaseiammember "github.com/upbound/provider-gcp/internal/controller/spanner/databaseiammember"
	instancespanner "github.com/upbound/provider-gcp/internal/controller/spanner/instance"
	instanceiammemberspanner "github.com/upbound/provider-gcp/internal/controller/spanner/instanceiammember"
	databasesql "github.com/upbound/provider-gcp/internal/controller/sql/database"
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

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		connection.Setup,
		dataset.Setup,
		datasetaccess.Setup,
		datasetiambinding.Setup,
		datasetiammember.Setup,
		datasetiampolicy.Setup,
		datatransferconfig.Setup,
		job.Setup,
		reservation.Setup,
		reservationassignment.Setup,
		routine.Setup,
		table.Setup,
		tableiambinding.Setup,
		tableiammember.Setup,
		tableiampolicy.Setup,
		trigger.Setup,
		workerpool.Setup,
		function.Setup,
		functioniammember.Setup,
		device.Setup,
		registry.Setup,
		folder.Setup,
		folderiammember.Setup,
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
		jobcloudscheduler.Setup,
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
		reservationcompute.Setup,
		resourcepolicy.Setup,
		route.Setup,
		router.Setup,
		routerinterface.Setup,
		routernat.Setup,
		securitypolicy.Setup,
		serviceattachment.Setup,
		sslcertificate.Setup,
		subnetwork.Setup,
		subnetworkiammember.Setup,
		targetgrpcproxy.Setup,
		targethttpproxy.Setup,
		targethttpsproxy.Setup,
		targetinstance.Setup,
		targetpool.Setup,
		targetsslproxy.Setup,
		targettcpproxy.Setup,
		urlmap.Setup,
		vpngateway.Setup,
		vpntunnel.Setup,
		cluster.Setup,
		nodepool.Setup,
		registrycontainer.Setup,
		note.Setup,
		clustercontaineraws.Setup,
		nodepoolcontaineraws.Setup,
		client.Setup,
		clustercontainerazure.Setup,
		nodepoolcontainerazure.Setup,
		entry.Setup,
		entrygroup.Setup,
		entrygroupiambinding.Setup,
		entrygroupiammember.Setup,
		entrygroupiampolicy.Setup,
		tag.Setup,
		tagtemplate.Setup,
		tagtemplateiambinding.Setup,
		tagtemplateiammember.Setup,
		tagtemplateiampolicy.Setup,
		jobdataflow.Setup,
		instancedatafusion.Setup,
		agent.Setup,
		entitytype.Setup,
		environmentdialogflowcx.Setup,
		flow.Setup,
		intent.Setup,
		page.Setup,
		version.Setup,
		managedzone.Setup,
		policy.Setup,
		recordset.Setup,
		contact.Setup,
		triggereventarc.Setup,
		instancefilestore.Setup,
		release.Setup,
		ruleset.Setup,
		membership.Setup,
		consentstore.Setup,
		datasethealthcare.Setup,
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
		connectionservicenetworking.Setup,
		repository.Setup,
		repositoryiammember.Setup,
		database.Setup,
		databaseiammember.Setup,
		instancespanner.Setup,
		instanceiammemberspanner.Setup,
		databasesql.Setup,
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
