// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesslevel "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevel"
	accesslevelcondition "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevelcondition"
	accesspolicy "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesspolicy"
	accesspolicyiammember "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesspolicyiammember"
	serviceperimeter "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeter"
	serviceperimeterresource "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeterresource"
	domain "github.com/upbound/provider-gcp/internal/controller/activedirectory/domain"
	backup "github.com/upbound/provider-gcp/internal/controller/alloydb/backup"
	cluster "github.com/upbound/provider-gcp/internal/controller/alloydb/cluster"
	instance "github.com/upbound/provider-gcp/internal/controller/alloydb/instance"
	envgroup "github.com/upbound/provider-gcp/internal/controller/apigee/envgroup"
	environment "github.com/upbound/provider-gcp/internal/controller/apigee/environment"
	environmentiammember "github.com/upbound/provider-gcp/internal/controller/apigee/environmentiammember"
	instanceapigee "github.com/upbound/provider-gcp/internal/controller/apigee/instance"
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
	appprofile "github.com/upbound/provider-gcp/internal/controller/bigtable/appprofile"
	garbagecollectionpolicy "github.com/upbound/provider-gcp/internal/controller/bigtable/garbagecollectionpolicy"
	instancebigtable "github.com/upbound/provider-gcp/internal/controller/bigtable/instance"
	instanceiambinding "github.com/upbound/provider-gcp/internal/controller/bigtable/instanceiambinding"
	instanceiammember "github.com/upbound/provider-gcp/internal/controller/bigtable/instanceiammember"
	instanceiampolicy "github.com/upbound/provider-gcp/internal/controller/bigtable/instanceiampolicy"
	tablebigtable "github.com/upbound/provider-gcp/internal/controller/bigtable/table"
	tableiambindingbigtable "github.com/upbound/provider-gcp/internal/controller/bigtable/tableiambinding"
	tableiammemberbigtable "github.com/upbound/provider-gcp/internal/controller/bigtable/tableiammember"
	tableiampolicybigtable "github.com/upbound/provider-gcp/internal/controller/bigtable/tableiampolicy"
	attestor "github.com/upbound/provider-gcp/internal/controller/binaryauthorization/attestor"
	policy "github.com/upbound/provider-gcp/internal/controller/binaryauthorization/policy"
	certificate "github.com/upbound/provider-gcp/internal/controller/certificatemanager/certificate"
	certificatemap "github.com/upbound/provider-gcp/internal/controller/certificatemanager/certificatemap"
	certificatemapentry "github.com/upbound/provider-gcp/internal/controller/certificatemanager/certificatemapentry"
	dnsauthorization "github.com/upbound/provider-gcp/internal/controller/certificatemanager/dnsauthorization"
	idsendpoint "github.com/upbound/provider-gcp/internal/controller/cloud/idsendpoint"
	trigger "github.com/upbound/provider-gcp/internal/controller/cloudbuild/trigger"
	workerpool "github.com/upbound/provider-gcp/internal/controller/cloudbuild/workerpool"
	function "github.com/upbound/provider-gcp/internal/controller/cloudfunctions/function"
	functioniammember "github.com/upbound/provider-gcp/internal/controller/cloudfunctions/functioniammember"
	functioncloudfunctions2 "github.com/upbound/provider-gcp/internal/controller/cloudfunctions2/function"
	folder "github.com/upbound/provider-gcp/internal/controller/cloudplatform/folder"
	folderiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/folderiammember"
	organizationiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cloudplatform/organizationiamauditconfig"
	organizationiamcustomrole "github.com/upbound/provider-gcp/internal/controller/cloudplatform/organizationiamcustomrole"
	organizationiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/organizationiammember"
	project "github.com/upbound/provider-gcp/internal/controller/cloudplatform/project"
	projectdefaultserviceaccounts "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectdefaultserviceaccounts"
	projectiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectiamauditconfig"
	projectiamcustomrole "github.com/upbound/provider-gcp/internal/controller/cloudplatform/projectiamcustomrole"
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
	v2job "github.com/upbound/provider-gcp/internal/controller/cloudrun/v2job"
	v2service "github.com/upbound/provider-gcp/internal/controller/cloudrun/v2service"
	jobcloudscheduler "github.com/upbound/provider-gcp/internal/controller/cloudscheduler/job"
	queue "github.com/upbound/provider-gcp/internal/controller/cloudtasks/queue"
	environmentcomposer "github.com/upbound/provider-gcp/internal/controller/composer/environment"
	address "github.com/upbound/provider-gcp/internal/controller/compute/address"
	attacheddisk "github.com/upbound/provider-gcp/internal/controller/compute/attacheddisk"
	autoscaler "github.com/upbound/provider-gcp/internal/controller/compute/autoscaler"
	backendbucket "github.com/upbound/provider-gcp/internal/controller/compute/backendbucket"
	backendbucketsignedurlkey "github.com/upbound/provider-gcp/internal/controller/compute/backendbucketsignedurlkey"
	backendservice "github.com/upbound/provider-gcp/internal/controller/compute/backendservice"
	backendservicesignedurlkey "github.com/upbound/provider-gcp/internal/controller/compute/backendservicesignedurlkey"
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
	instancecompute "github.com/upbound/provider-gcp/internal/controller/compute/instance"
	instancefromtemplate "github.com/upbound/provider-gcp/internal/controller/compute/instancefromtemplate"
	instancegroup "github.com/upbound/provider-gcp/internal/controller/compute/instancegroup"
	instancegroupmanager "github.com/upbound/provider-gcp/internal/controller/compute/instancegroupmanager"
	instancegroupnamedport "github.com/upbound/provider-gcp/internal/controller/compute/instancegroupnamedport"
	instanceiammembercompute "github.com/upbound/provider-gcp/internal/controller/compute/instanceiammember"
	instancetemplate "github.com/upbound/provider-gcp/internal/controller/compute/instancetemplate"
	interconnectattachment "github.com/upbound/provider-gcp/internal/controller/compute/interconnectattachment"
	managedsslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/managedsslcertificate"
	network "github.com/upbound/provider-gcp/internal/controller/compute/network"
	networkendpoint "github.com/upbound/provider-gcp/internal/controller/compute/networkendpoint"
	networkendpointgroup "github.com/upbound/provider-gcp/internal/controller/compute/networkendpointgroup"
	networkfirewallpolicy "github.com/upbound/provider-gcp/internal/controller/compute/networkfirewallpolicy"
	networkfirewallpolicyassociation "github.com/upbound/provider-gcp/internal/controller/compute/networkfirewallpolicyassociation"
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
	regionnetworkendpoint "github.com/upbound/provider-gcp/internal/controller/compute/regionnetworkendpoint"
	regionnetworkendpointgroup "github.com/upbound/provider-gcp/internal/controller/compute/regionnetworkendpointgroup"
	regionnetworkfirewallpolicy "github.com/upbound/provider-gcp/internal/controller/compute/regionnetworkfirewallpolicy"
	regionnetworkfirewallpolicyassociation "github.com/upbound/provider-gcp/internal/controller/compute/regionnetworkfirewallpolicyassociation"
	regionperinstanceconfig "github.com/upbound/provider-gcp/internal/controller/compute/regionperinstanceconfig"
	regionsslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/regionsslcertificate"
	regiontargethttpproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpsproxy"
	regiontargettcpproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargettcpproxy"
	regionurlmap "github.com/upbound/provider-gcp/internal/controller/compute/regionurlmap"
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
	sslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/sslcertificate"
	sslpolicy "github.com/upbound/provider-gcp/internal/controller/compute/sslpolicy"
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
	clustercontainer "github.com/upbound/provider-gcp/internal/controller/container/cluster"
	nodepool "github.com/upbound/provider-gcp/internal/controller/container/nodepool"
	registry "github.com/upbound/provider-gcp/internal/controller/container/registry"
	note "github.com/upbound/provider-gcp/internal/controller/containeranalysis/note"
	clustercontainerattached "github.com/upbound/provider-gcp/internal/controller/containerattached/cluster"
	clustercontaineraws "github.com/upbound/provider-gcp/internal/controller/containeraws/cluster"
	nodepoolcontaineraws "github.com/upbound/provider-gcp/internal/controller/containeraws/nodepool"
	client "github.com/upbound/provider-gcp/internal/controller/containerazure/client"
	clustercontainerazure "github.com/upbound/provider-gcp/internal/controller/containerazure/cluster"
	nodepoolcontainerazure "github.com/upbound/provider-gcp/internal/controller/containerazure/nodepool"
	entry "github.com/upbound/provider-gcp/internal/controller/datacatalog/entry"
	entrygroup "github.com/upbound/provider-gcp/internal/controller/datacatalog/entrygroup"
	tag "github.com/upbound/provider-gcp/internal/controller/datacatalog/tag"
	tagtemplate "github.com/upbound/provider-gcp/internal/controller/datacatalog/tagtemplate"
	jobdataflow "github.com/upbound/provider-gcp/internal/controller/dataflow/job"
	instancedatafusion "github.com/upbound/provider-gcp/internal/controller/datafusion/instance"
	deidentifytemplate "github.com/upbound/provider-gcp/internal/controller/datalossprevention/deidentifytemplate"
	inspecttemplate "github.com/upbound/provider-gcp/internal/controller/datalossprevention/inspecttemplate"
	jobtrigger "github.com/upbound/provider-gcp/internal/controller/datalossprevention/jobtrigger"
	storedinfotype "github.com/upbound/provider-gcp/internal/controller/datalossprevention/storedinfotype"
	asset "github.com/upbound/provider-gcp/internal/controller/dataplex/asset"
	lake "github.com/upbound/provider-gcp/internal/controller/dataplex/lake"
	zone "github.com/upbound/provider-gcp/internal/controller/dataplex/zone"
	autoscalingpolicy "github.com/upbound/provider-gcp/internal/controller/dataproc/autoscalingpolicy"
	clusterdataproc "github.com/upbound/provider-gcp/internal/controller/dataproc/cluster"
	jobdataproc "github.com/upbound/provider-gcp/internal/controller/dataproc/job"
	metastoreservice "github.com/upbound/provider-gcp/internal/controller/dataproc/metastoreservice"
	workflowtemplate "github.com/upbound/provider-gcp/internal/controller/dataproc/workflowtemplate"
	index "github.com/upbound/provider-gcp/internal/controller/datastore/index"
	connectionprofile "github.com/upbound/provider-gcp/internal/controller/datastream/connectionprofile"
	privateconnection "github.com/upbound/provider-gcp/internal/controller/datastream/privateconnection"
	agent "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/agent"
	entitytype "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/entitytype"
	environmentdialogflowcx "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/environment"
	flow "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/flow"
	intent "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/intent"
	page "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/page"
	version "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/version"
	webhook "github.com/upbound/provider-gcp/internal/controller/dialogflowcx/webhook"
	managedzone "github.com/upbound/provider-gcp/internal/controller/dns/managedzone"
	managedzoneiammember "github.com/upbound/provider-gcp/internal/controller/dns/managedzoneiammember"
	policydns "github.com/upbound/provider-gcp/internal/controller/dns/policy"
	recordset "github.com/upbound/provider-gcp/internal/controller/dns/recordset"
	processor "github.com/upbound/provider-gcp/internal/controller/documentai/processor"
	contact "github.com/upbound/provider-gcp/internal/controller/essentialcontacts/contact"
	channel "github.com/upbound/provider-gcp/internal/controller/eventarc/channel"
	googlechannelconfig "github.com/upbound/provider-gcp/internal/controller/eventarc/googlechannelconfig"
	triggereventarc "github.com/upbound/provider-gcp/internal/controller/eventarc/trigger"
	backupfilestore "github.com/upbound/provider-gcp/internal/controller/filestore/backup"
	instancefilestore "github.com/upbound/provider-gcp/internal/controller/filestore/instance"
	snapshotfilestore "github.com/upbound/provider-gcp/internal/controller/filestore/snapshot"
	release "github.com/upbound/provider-gcp/internal/controller/firebaserules/release"
	ruleset "github.com/upbound/provider-gcp/internal/controller/firebaserules/ruleset"
	backupbackupplan "github.com/upbound/provider-gcp/internal/controller/gke/backupbackupplan"
	membership "github.com/upbound/provider-gcp/internal/controller/gkehub/membership"
	membershipiammember "github.com/upbound/provider-gcp/internal/controller/gkehub/membershipiammember"
	consentstore "github.com/upbound/provider-gcp/internal/controller/healthcare/consentstore"
	datasethealthcare "github.com/upbound/provider-gcp/internal/controller/healthcare/dataset"
	datasetiammemberhealthcare "github.com/upbound/provider-gcp/internal/controller/healthcare/datasetiammember"
	workloadidentitypool "github.com/upbound/provider-gcp/internal/controller/iam/workloadidentitypool"
	workloadidentitypoolprovider "github.com/upbound/provider-gcp/internal/controller/iam/workloadidentitypoolprovider"
	appengineserviceiammember "github.com/upbound/provider-gcp/internal/controller/iap/appengineserviceiammember"
	appengineversioniammember "github.com/upbound/provider-gcp/internal/controller/iap/appengineversioniammember"
	tunneliammember "github.com/upbound/provider-gcp/internal/controller/iap/tunneliammember"
	webbackendserviceiammember "github.com/upbound/provider-gcp/internal/controller/iap/webbackendserviceiammember"
	webiammember "github.com/upbound/provider-gcp/internal/controller/iap/webiammember"
	webtypeappengineiammember "github.com/upbound/provider-gcp/internal/controller/iap/webtypeappengineiammember"
	webtypecomputeiammember "github.com/upbound/provider-gcp/internal/controller/iap/webtypecomputeiammember"
	defaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/defaultsupportedidpconfig"
	inboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/inboundsamlconfig"
	oauthidpconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/oauthidpconfig"
	projectdefaultconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/projectdefaultconfig"
	tenant "github.com/upbound/provider-gcp/internal/controller/identityplatform/tenant"
	tenantdefaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/tenantdefaultsupportedidpconfig"
	tenantinboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/tenantinboundsamlconfig"
	tenantoauthidpconfig "github.com/upbound/provider-gcp/internal/controller/identityplatform/tenantoauthidpconfig"
	cryptokey "github.com/upbound/provider-gcp/internal/controller/kms/cryptokey"
	cryptokeyiammember "github.com/upbound/provider-gcp/internal/controller/kms/cryptokeyiammember"
	cryptokeyversion "github.com/upbound/provider-gcp/internal/controller/kms/cryptokeyversion"
	keyring "github.com/upbound/provider-gcp/internal/controller/kms/keyring"
	keyringiammember "github.com/upbound/provider-gcp/internal/controller/kms/keyringiammember"
	keyringimportjob "github.com/upbound/provider-gcp/internal/controller/kms/keyringimportjob"
	secretciphertext "github.com/upbound/provider-gcp/internal/controller/kms/secretciphertext"
	folderbucketconfig "github.com/upbound/provider-gcp/internal/controller/logging/folderbucketconfig"
	folderexclusion "github.com/upbound/provider-gcp/internal/controller/logging/folderexclusion"
	foldersink "github.com/upbound/provider-gcp/internal/controller/logging/foldersink"
	logview "github.com/upbound/provider-gcp/internal/controller/logging/logview"
	metric "github.com/upbound/provider-gcp/internal/controller/logging/metric"
	projectbucketconfig "github.com/upbound/provider-gcp/internal/controller/logging/projectbucketconfig"
	projectexclusion "github.com/upbound/provider-gcp/internal/controller/logging/projectexclusion"
	projectsink "github.com/upbound/provider-gcp/internal/controller/logging/projectsink"
	instancememcache "github.com/upbound/provider-gcp/internal/controller/memcache/instance"
	model "github.com/upbound/provider-gcp/internal/controller/mlengine/model"
	alertpolicy "github.com/upbound/provider-gcp/internal/controller/monitoring/alertpolicy"
	customservice "github.com/upbound/provider-gcp/internal/controller/monitoring/customservice"
	dashboard "github.com/upbound/provider-gcp/internal/controller/monitoring/dashboard"
	group "github.com/upbound/provider-gcp/internal/controller/monitoring/group"
	metricdescriptor "github.com/upbound/provider-gcp/internal/controller/monitoring/metricdescriptor"
	notificationchannel "github.com/upbound/provider-gcp/internal/controller/monitoring/notificationchannel"
	servicemonitoring "github.com/upbound/provider-gcp/internal/controller/monitoring/service"
	slo "github.com/upbound/provider-gcp/internal/controller/monitoring/slo"
	uptimecheckconfig "github.com/upbound/provider-gcp/internal/controller/monitoring/uptimecheckconfig"
	hub "github.com/upbound/provider-gcp/internal/controller/networkconnectivity/hub"
	spoke "github.com/upbound/provider-gcp/internal/controller/networkconnectivity/spoke"
	connectivitytest "github.com/upbound/provider-gcp/internal/controller/networkmanagement/connectivitytest"
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
	certificateprivateca "github.com/upbound/provider-gcp/internal/controller/privateca/certificate"
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
	hmackey "github.com/upbound/provider-gcp/internal/controller/storage/hmackey"
	notification "github.com/upbound/provider-gcp/internal/controller/storage/notification"
	objectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/storage/objectaccesscontrol"
	objectacl "github.com/upbound/provider-gcp/internal/controller/storage/objectacl"
	agentpool "github.com/upbound/provider-gcp/internal/controller/storagetransfer/agentpool"
	node "github.com/upbound/provider-gcp/internal/controller/tpu/node"
	datasetvertexai "github.com/upbound/provider-gcp/internal/controller/vertexai/dataset"
	featurestore "github.com/upbound/provider-gcp/internal/controller/vertexai/featurestore"
	featurestoreentitytype "github.com/upbound/provider-gcp/internal/controller/vertexai/featurestoreentitytype"
	tensorboard "github.com/upbound/provider-gcp/internal/controller/vertexai/tensorboard"
	connector "github.com/upbound/provider-gcp/internal/controller/vpcaccess/connector"
	workflow "github.com/upbound/provider-gcp/internal/controller/workflows/workflow"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslevel.Setup,
		accesslevelcondition.Setup,
		accesspolicy.Setup,
		accesspolicyiammember.Setup,
		serviceperimeter.Setup,
		serviceperimeterresource.Setup,
		domain.Setup,
		backup.Setup,
		cluster.Setup,
		instance.Setup,
		envgroup.Setup,
		environment.Setup,
		environmentiammember.Setup,
		instanceapigee.Setup,
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
		appprofile.Setup,
		garbagecollectionpolicy.Setup,
		instancebigtable.Setup,
		instanceiambinding.Setup,
		instanceiammember.Setup,
		instanceiampolicy.Setup,
		tablebigtable.Setup,
		tableiambindingbigtable.Setup,
		tableiammemberbigtable.Setup,
		tableiampolicybigtable.Setup,
		attestor.Setup,
		policy.Setup,
		certificate.Setup,
		certificatemap.Setup,
		certificatemapentry.Setup,
		dnsauthorization.Setup,
		idsendpoint.Setup,
		trigger.Setup,
		workerpool.Setup,
		function.Setup,
		functioniammember.Setup,
		functioncloudfunctions2.Setup,
		folder.Setup,
		folderiammember.Setup,
		organizationiamauditconfig.Setup,
		organizationiamcustomrole.Setup,
		organizationiammember.Setup,
		project.Setup,
		projectdefaultserviceaccounts.Setup,
		projectiamauditconfig.Setup,
		projectiamcustomrole.Setup,
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
		v2job.Setup,
		v2service.Setup,
		jobcloudscheduler.Setup,
		queue.Setup,
		environmentcomposer.Setup,
		address.Setup,
		attacheddisk.Setup,
		autoscaler.Setup,
		backendbucket.Setup,
		backendbucketsignedurlkey.Setup,
		backendservice.Setup,
		backendservicesignedurlkey.Setup,
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
		instancecompute.Setup,
		instancefromtemplate.Setup,
		instancegroup.Setup,
		instancegroupmanager.Setup,
		instancegroupnamedport.Setup,
		instanceiammembercompute.Setup,
		instancetemplate.Setup,
		interconnectattachment.Setup,
		managedsslcertificate.Setup,
		network.Setup,
		networkendpoint.Setup,
		networkendpointgroup.Setup,
		networkfirewallpolicy.Setup,
		networkfirewallpolicyassociation.Setup,
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
		regionnetworkendpoint.Setup,
		regionnetworkendpointgroup.Setup,
		regionnetworkfirewallpolicy.Setup,
		regionnetworkfirewallpolicyassociation.Setup,
		regionperinstanceconfig.Setup,
		regionsslcertificate.Setup,
		regiontargethttpproxy.Setup,
		regiontargethttpsproxy.Setup,
		regiontargettcpproxy.Setup,
		regionurlmap.Setup,
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
		sslcertificate.Setup,
		sslpolicy.Setup,
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
		clustercontainer.Setup,
		nodepool.Setup,
		registry.Setup,
		note.Setup,
		clustercontainerattached.Setup,
		clustercontaineraws.Setup,
		nodepoolcontaineraws.Setup,
		client.Setup,
		clustercontainerazure.Setup,
		nodepoolcontainerazure.Setup,
		entry.Setup,
		entrygroup.Setup,
		tag.Setup,
		tagtemplate.Setup,
		jobdataflow.Setup,
		instancedatafusion.Setup,
		deidentifytemplate.Setup,
		inspecttemplate.Setup,
		jobtrigger.Setup,
		storedinfotype.Setup,
		asset.Setup,
		lake.Setup,
		zone.Setup,
		autoscalingpolicy.Setup,
		clusterdataproc.Setup,
		jobdataproc.Setup,
		metastoreservice.Setup,
		workflowtemplate.Setup,
		index.Setup,
		connectionprofile.Setup,
		privateconnection.Setup,
		agent.Setup,
		entitytype.Setup,
		environmentdialogflowcx.Setup,
		flow.Setup,
		intent.Setup,
		page.Setup,
		version.Setup,
		webhook.Setup,
		managedzone.Setup,
		managedzoneiammember.Setup,
		policydns.Setup,
		recordset.Setup,
		processor.Setup,
		contact.Setup,
		channel.Setup,
		googlechannelconfig.Setup,
		triggereventarc.Setup,
		backupfilestore.Setup,
		instancefilestore.Setup,
		snapshotfilestore.Setup,
		release.Setup,
		ruleset.Setup,
		backupbackupplan.Setup,
		membership.Setup,
		membershipiammember.Setup,
		consentstore.Setup,
		datasethealthcare.Setup,
		datasetiammemberhealthcare.Setup,
		workloadidentitypool.Setup,
		workloadidentitypoolprovider.Setup,
		appengineserviceiammember.Setup,
		appengineversioniammember.Setup,
		tunneliammember.Setup,
		webbackendserviceiammember.Setup,
		webiammember.Setup,
		webtypeappengineiammember.Setup,
		webtypecomputeiammember.Setup,
		defaultsupportedidpconfig.Setup,
		inboundsamlconfig.Setup,
		oauthidpconfig.Setup,
		projectdefaultconfig.Setup,
		tenant.Setup,
		tenantdefaultsupportedidpconfig.Setup,
		tenantinboundsamlconfig.Setup,
		tenantoauthidpconfig.Setup,
		cryptokey.Setup,
		cryptokeyiammember.Setup,
		cryptokeyversion.Setup,
		keyring.Setup,
		keyringiammember.Setup,
		keyringimportjob.Setup,
		secretciphertext.Setup,
		folderbucketconfig.Setup,
		folderexclusion.Setup,
		foldersink.Setup,
		logview.Setup,
		metric.Setup,
		projectbucketconfig.Setup,
		projectexclusion.Setup,
		projectsink.Setup,
		instancememcache.Setup,
		model.Setup,
		alertpolicy.Setup,
		customservice.Setup,
		dashboard.Setup,
		group.Setup,
		metricdescriptor.Setup,
		notificationchannel.Setup,
		servicemonitoring.Setup,
		slo.Setup,
		uptimecheckconfig.Setup,
		hub.Setup,
		spoke.Setup,
		connectivitytest.Setup,
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
		certificateprivateca.Setup,
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
		hmackey.Setup,
		notification.Setup,
		objectaccesscontrol.Setup,
		objectacl.Setup,
		agentpool.Setup,
		node.Setup,
		datasetvertexai.Setup,
		featurestore.Setup,
		featurestoreentitytype.Setup,
		tensorboard.Setup,
		connector.Setup,
		workflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
