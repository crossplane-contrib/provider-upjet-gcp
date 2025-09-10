// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesslevel "github.com/upbound/provider-gcp/internal/controller/cluster/accesscontextmanager/accesslevel"
	accesslevelcondition "github.com/upbound/provider-gcp/internal/controller/cluster/accesscontextmanager/accesslevelcondition"
	accesspolicy "github.com/upbound/provider-gcp/internal/controller/cluster/accesscontextmanager/accesspolicy"
	accesspolicyiammember "github.com/upbound/provider-gcp/internal/controller/cluster/accesscontextmanager/accesspolicyiammember"
	serviceperimeter "github.com/upbound/provider-gcp/internal/controller/cluster/accesscontextmanager/serviceperimeter"
	serviceperimeterresource "github.com/upbound/provider-gcp/internal/controller/cluster/accesscontextmanager/serviceperimeterresource"
	domain "github.com/upbound/provider-gcp/internal/controller/cluster/activedirectory/domain"
	backup "github.com/upbound/provider-gcp/internal/controller/cluster/alloydb/backup"
	cluster "github.com/upbound/provider-gcp/internal/controller/cluster/alloydb/cluster"
	instance "github.com/upbound/provider-gcp/internal/controller/cluster/alloydb/instance"
	addonsconfig "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/addonsconfig"
	endpointattachment "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/endpointattachment"
	envgroup "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/envgroup"
	envgroupattachment "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/envgroupattachment"
	environment "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/environment"
	environmentiammember "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/environmentiammember"
	instanceapigee "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/instance"
	instanceattachment "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/instanceattachment"
	nataddress "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/nataddress"
	organization "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/organization"
	syncauthorization "github.com/upbound/provider-gcp/internal/controller/cluster/apigee/syncauthorization"
	application "github.com/upbound/provider-gcp/internal/controller/cluster/appengine/application"
	applicationurldispatchrules "github.com/upbound/provider-gcp/internal/controller/cluster/appengine/applicationurldispatchrules"
	firewallrule "github.com/upbound/provider-gcp/internal/controller/cluster/appengine/firewallrule"
	servicenetworksettings "github.com/upbound/provider-gcp/internal/controller/cluster/appengine/servicenetworksettings"
	standardappversion "github.com/upbound/provider-gcp/internal/controller/cluster/appengine/standardappversion"
	registryrepository "github.com/upbound/provider-gcp/internal/controller/cluster/artifact/registryrepository"
	registryrepositoryiammember "github.com/upbound/provider-gcp/internal/controller/cluster/artifact/registryrepositoryiammember"
	appconnection "github.com/upbound/provider-gcp/internal/controller/cluster/beyondcorp/appconnection"
	appconnector "github.com/upbound/provider-gcp/internal/controller/cluster/beyondcorp/appconnector"
	appgateway "github.com/upbound/provider-gcp/internal/controller/cluster/beyondcorp/appgateway"
	analyticshubdataexchange "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/analyticshubdataexchange"
	analyticshubdataexchangeiammember "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/analyticshubdataexchangeiammember"
	analyticshublisting "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/analyticshublisting"
	connection "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/connection"
	dataset "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/dataset"
	datasetaccess "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/datasetaccess"
	datasetiambinding "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/datasetiambinding"
	datasetiammember "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/datasetiammember"
	datasetiampolicy "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/datasetiampolicy"
	datatransferconfig "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/datatransferconfig"
	job "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/job"
	reservation "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/reservation"
	reservationassignment "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/reservationassignment"
	routine "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/routine"
	table "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/table"
	tableiambinding "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/tableiambinding"
	tableiammember "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/tableiammember"
	tableiampolicy "github.com/upbound/provider-gcp/internal/controller/cluster/bigquery/tableiampolicy"
	appprofile "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/appprofile"
	garbagecollectionpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/garbagecollectionpolicy"
	instancebigtable "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/instance"
	instanceiambinding "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/instanceiambinding"
	instanceiammember "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/instanceiammember"
	instanceiampolicy "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/instanceiampolicy"
	tablebigtable "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/table"
	tableiambindingbigtable "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/tableiambinding"
	tableiammemberbigtable "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/tableiammember"
	tableiampolicybigtable "github.com/upbound/provider-gcp/internal/controller/cluster/bigtable/tableiampolicy"
	attestor "github.com/upbound/provider-gcp/internal/controller/cluster/binaryauthorization/attestor"
	policy "github.com/upbound/provider-gcp/internal/controller/cluster/binaryauthorization/policy"
	certificate "github.com/upbound/provider-gcp/internal/controller/cluster/certificatemanager/certificate"
	certificatemap "github.com/upbound/provider-gcp/internal/controller/cluster/certificatemanager/certificatemap"
	certificatemapentry "github.com/upbound/provider-gcp/internal/controller/cluster/certificatemanager/certificatemapentry"
	dnsauthorization "github.com/upbound/provider-gcp/internal/controller/cluster/certificatemanager/dnsauthorization"
	trustconfig "github.com/upbound/provider-gcp/internal/controller/cluster/certificatemanager/trustconfig"
	idsendpoint "github.com/upbound/provider-gcp/internal/controller/cluster/cloud/idsendpoint"
	trigger "github.com/upbound/provider-gcp/internal/controller/cluster/cloudbuild/trigger"
	workerpool "github.com/upbound/provider-gcp/internal/controller/cluster/cloudbuild/workerpool"
	function "github.com/upbound/provider-gcp/internal/controller/cluster/cloudfunctions/function"
	functioniammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudfunctions/functioniammember"
	functioncloudfunctions2 "github.com/upbound/provider-gcp/internal/controller/cluster/cloudfunctions2/function"
	folder "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/folder"
	folderiammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/folderiammember"
	organizationiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/organizationiamauditconfig"
	organizationiamcustomrole "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/organizationiamcustomrole"
	organizationiammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/organizationiammember"
	project "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/project"
	projectdefaultserviceaccounts "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectdefaultserviceaccounts"
	projectiamauditconfig "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectiamauditconfig"
	projectiamcustomrole "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectiamcustomrole"
	projectiammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectiammember"
	projectservice "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectservice"
	projectusageexportbucket "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/projectusageexportbucket"
	serviceaccount "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/provider-gcp/internal/controller/cluster/cloudplatform/servicenetworkingpeereddnsdomain"
	domainmapping "github.com/upbound/provider-gcp/internal/controller/cluster/cloudrun/domainmapping"
	service "github.com/upbound/provider-gcp/internal/controller/cluster/cloudrun/service"
	serviceiammember "github.com/upbound/provider-gcp/internal/controller/cluster/cloudrun/serviceiammember"
	v2job "github.com/upbound/provider-gcp/internal/controller/cluster/cloudrun/v2job"
	v2service "github.com/upbound/provider-gcp/internal/controller/cluster/cloudrun/v2service"
	jobcloudscheduler "github.com/upbound/provider-gcp/internal/controller/cluster/cloudscheduler/job"
	queue "github.com/upbound/provider-gcp/internal/controller/cluster/cloudtasks/queue"
	environmentcomposer "github.com/upbound/provider-gcp/internal/controller/cluster/composer/environment"
	address "github.com/upbound/provider-gcp/internal/controller/cluster/compute/address"
	attacheddisk "github.com/upbound/provider-gcp/internal/controller/cluster/compute/attacheddisk"
	autoscaler "github.com/upbound/provider-gcp/internal/controller/cluster/compute/autoscaler"
	backendbucket "github.com/upbound/provider-gcp/internal/controller/cluster/compute/backendbucket"
	backendbucketsignedurlkey "github.com/upbound/provider-gcp/internal/controller/cluster/compute/backendbucketsignedurlkey"
	backendservice "github.com/upbound/provider-gcp/internal/controller/cluster/compute/backendservice"
	backendservicesignedurlkey "github.com/upbound/provider-gcp/internal/controller/cluster/compute/backendservicesignedurlkey"
	disk "github.com/upbound/provider-gcp/internal/controller/cluster/compute/disk"
	diskiammember "github.com/upbound/provider-gcp/internal/controller/cluster/compute/diskiammember"
	diskresourcepolicyattachment "github.com/upbound/provider-gcp/internal/controller/cluster/compute/diskresourcepolicyattachment"
	externalvpngateway "github.com/upbound/provider-gcp/internal/controller/cluster/compute/externalvpngateway"
	firewall "github.com/upbound/provider-gcp/internal/controller/cluster/compute/firewall"
	firewallpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/firewallpolicy"
	firewallpolicyassociation "github.com/upbound/provider-gcp/internal/controller/cluster/compute/firewallpolicyassociation"
	firewallpolicyrule "github.com/upbound/provider-gcp/internal/controller/cluster/compute/firewallpolicyrule"
	forwardingrule "github.com/upbound/provider-gcp/internal/controller/cluster/compute/forwardingrule"
	globaladdress "github.com/upbound/provider-gcp/internal/controller/cluster/compute/globaladdress"
	globalforwardingrule "github.com/upbound/provider-gcp/internal/controller/cluster/compute/globalforwardingrule"
	globalnetworkendpoint "github.com/upbound/provider-gcp/internal/controller/cluster/compute/globalnetworkendpoint"
	globalnetworkendpointgroup "github.com/upbound/provider-gcp/internal/controller/cluster/compute/globalnetworkendpointgroup"
	havpngateway "github.com/upbound/provider-gcp/internal/controller/cluster/compute/havpngateway"
	healthcheck "github.com/upbound/provider-gcp/internal/controller/cluster/compute/healthcheck"
	httphealthcheck "github.com/upbound/provider-gcp/internal/controller/cluster/compute/httphealthcheck"
	httpshealthcheck "github.com/upbound/provider-gcp/internal/controller/cluster/compute/httpshealthcheck"
	image "github.com/upbound/provider-gcp/internal/controller/cluster/compute/image"
	imageiammember "github.com/upbound/provider-gcp/internal/controller/cluster/compute/imageiammember"
	instancecompute "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instance"
	instancefromtemplate "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instancefromtemplate"
	instancegroup "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instancegroup"
	instancegroupmanager "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instancegroupmanager"
	instancegroupnamedport "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instancegroupnamedport"
	instanceiammembercompute "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instanceiammember"
	instancetemplate "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instancetemplate"
	interconnectattachment "github.com/upbound/provider-gcp/internal/controller/cluster/compute/interconnectattachment"
	managedsslcertificate "github.com/upbound/provider-gcp/internal/controller/cluster/compute/managedsslcertificate"
	network "github.com/upbound/provider-gcp/internal/controller/cluster/compute/network"
	networkendpoint "github.com/upbound/provider-gcp/internal/controller/cluster/compute/networkendpoint"
	networkendpointgroup "github.com/upbound/provider-gcp/internal/controller/cluster/compute/networkendpointgroup"
	networkfirewallpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/networkfirewallpolicy"
	networkfirewallpolicyassociation "github.com/upbound/provider-gcp/internal/controller/cluster/compute/networkfirewallpolicyassociation"
	networkfirewallpolicyrule "github.com/upbound/provider-gcp/internal/controller/cluster/compute/networkfirewallpolicyrule"
	networkpeering "github.com/upbound/provider-gcp/internal/controller/cluster/compute/networkpeering"
	networkpeeringroutesconfig "github.com/upbound/provider-gcp/internal/controller/cluster/compute/networkpeeringroutesconfig"
	nodegroup "github.com/upbound/provider-gcp/internal/controller/cluster/compute/nodegroup"
	nodetemplate "github.com/upbound/provider-gcp/internal/controller/cluster/compute/nodetemplate"
	packetmirroring "github.com/upbound/provider-gcp/internal/controller/cluster/compute/packetmirroring"
	perinstanceconfig "github.com/upbound/provider-gcp/internal/controller/cluster/compute/perinstanceconfig"
	projectdefaultnetworktier "github.com/upbound/provider-gcp/internal/controller/cluster/compute/projectdefaultnetworktier"
	projectmetadata "github.com/upbound/provider-gcp/internal/controller/cluster/compute/projectmetadata"
	projectmetadataitem "github.com/upbound/provider-gcp/internal/controller/cluster/compute/projectmetadataitem"
	regionautoscaler "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionautoscaler"
	regionbackendservice "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionbackendservice"
	regiondisk "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regiondisk"
	regiondiskiammember "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regiondiskiammember"
	regiondiskresourcepolicyattachment "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regiondiskresourcepolicyattachment"
	regionhealthcheck "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionhealthcheck"
	regioninstancegroupmanager "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regioninstancegroupmanager"
	regionnetworkendpoint "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionnetworkendpoint"
	regionnetworkendpointgroup "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionnetworkendpointgroup"
	regionnetworkfirewallpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionnetworkfirewallpolicy"
	regionnetworkfirewallpolicyassociation "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionnetworkfirewallpolicyassociation"
	regionperinstanceconfig "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionperinstanceconfig"
	regionsslcertificate "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionsslcertificate"
	regionsslpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionsslpolicy"
	regiontargethttpproxy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regiontargethttpsproxy"
	regiontargettcpproxy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regiontargettcpproxy"
	regionurlmap "github.com/upbound/provider-gcp/internal/controller/cluster/compute/regionurlmap"
	reservationcompute "github.com/upbound/provider-gcp/internal/controller/cluster/compute/reservation"
	resourcepolicy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/resourcepolicy"
	route "github.com/upbound/provider-gcp/internal/controller/cluster/compute/route"
	router "github.com/upbound/provider-gcp/internal/controller/cluster/compute/router"
	routerinterface "github.com/upbound/provider-gcp/internal/controller/cluster/compute/routerinterface"
	routernat "github.com/upbound/provider-gcp/internal/controller/cluster/compute/routernat"
	routerpeer "github.com/upbound/provider-gcp/internal/controller/cluster/compute/routerpeer"
	securitypolicy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/securitypolicy"
	serviceattachment "github.com/upbound/provider-gcp/internal/controller/cluster/compute/serviceattachment"
	sharedvpchostproject "github.com/upbound/provider-gcp/internal/controller/cluster/compute/sharedvpchostproject"
	sharedvpcserviceproject "github.com/upbound/provider-gcp/internal/controller/cluster/compute/sharedvpcserviceproject"
	snapshot "github.com/upbound/provider-gcp/internal/controller/cluster/compute/snapshot"
	snapshotiammember "github.com/upbound/provider-gcp/internal/controller/cluster/compute/snapshotiammember"
	sslcertificate "github.com/upbound/provider-gcp/internal/controller/cluster/compute/sslcertificate"
	sslpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/sslpolicy"
	subnetwork "github.com/upbound/provider-gcp/internal/controller/cluster/compute/subnetwork"
	subnetworkiammember "github.com/upbound/provider-gcp/internal/controller/cluster/compute/subnetworkiammember"
	targetgrpcproxy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/targetgrpcproxy"
	targethttpproxy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/targethttpproxy"
	targethttpsproxy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/targethttpsproxy"
	targetinstance "github.com/upbound/provider-gcp/internal/controller/cluster/compute/targetinstance"
	targetpool "github.com/upbound/provider-gcp/internal/controller/cluster/compute/targetpool"
	targetsslproxy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/targetsslproxy"
	targettcpproxy "github.com/upbound/provider-gcp/internal/controller/cluster/compute/targettcpproxy"
	urlmap "github.com/upbound/provider-gcp/internal/controller/cluster/compute/urlmap"
	vpngateway "github.com/upbound/provider-gcp/internal/controller/cluster/compute/vpngateway"
	vpntunnel "github.com/upbound/provider-gcp/internal/controller/cluster/compute/vpntunnel"
	clustercontainer "github.com/upbound/provider-gcp/internal/controller/cluster/container/cluster"
	nodepool "github.com/upbound/provider-gcp/internal/controller/cluster/container/nodepool"
	registry "github.com/upbound/provider-gcp/internal/controller/cluster/container/registry"
	note "github.com/upbound/provider-gcp/internal/controller/cluster/containeranalysis/note"
	clustercontainerattached "github.com/upbound/provider-gcp/internal/controller/cluster/containerattached/cluster"
	clustercontaineraws "github.com/upbound/provider-gcp/internal/controller/cluster/containeraws/cluster"
	nodepoolcontaineraws "github.com/upbound/provider-gcp/internal/controller/cluster/containeraws/nodepool"
	client "github.com/upbound/provider-gcp/internal/controller/cluster/containerazure/client"
	clustercontainerazure "github.com/upbound/provider-gcp/internal/controller/cluster/containerazure/cluster"
	nodepoolcontainerazure "github.com/upbound/provider-gcp/internal/controller/cluster/containerazure/nodepool"
	entry "github.com/upbound/provider-gcp/internal/controller/cluster/datacatalog/entry"
	entrygroup "github.com/upbound/provider-gcp/internal/controller/cluster/datacatalog/entrygroup"
	policytag "github.com/upbound/provider-gcp/internal/controller/cluster/datacatalog/policytag"
	tag "github.com/upbound/provider-gcp/internal/controller/cluster/datacatalog/tag"
	tagtemplate "github.com/upbound/provider-gcp/internal/controller/cluster/datacatalog/tagtemplate"
	taxonomy "github.com/upbound/provider-gcp/internal/controller/cluster/datacatalog/taxonomy"
	jobdataflow "github.com/upbound/provider-gcp/internal/controller/cluster/dataflow/job"
	instancedatafusion "github.com/upbound/provider-gcp/internal/controller/cluster/datafusion/instance"
	deidentifytemplate "github.com/upbound/provider-gcp/internal/controller/cluster/datalossprevention/deidentifytemplate"
	inspecttemplate "github.com/upbound/provider-gcp/internal/controller/cluster/datalossprevention/inspecttemplate"
	jobtrigger "github.com/upbound/provider-gcp/internal/controller/cluster/datalossprevention/jobtrigger"
	storedinfotype "github.com/upbound/provider-gcp/internal/controller/cluster/datalossprevention/storedinfotype"
	aspecttype "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/aspecttype"
	asset "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/asset"
	lake "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/lake"
	lakeiampolicy "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/lakeiampolicy"
	zone "github.com/upbound/provider-gcp/internal/controller/cluster/dataplex/zone"
	autoscalingpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/autoscalingpolicy"
	clusterdataproc "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/cluster"
	jobdataproc "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/job"
	metastoreservice "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/metastoreservice"
	workflowtemplate "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/workflowtemplate"
	connectionprofile "github.com/upbound/provider-gcp/internal/controller/cluster/datastream/connectionprofile"
	privateconnection "github.com/upbound/provider-gcp/internal/controller/cluster/datastream/privateconnection"
	connectaccountconnector "github.com/upbound/provider-gcp/internal/controller/cluster/developerconnect/connectaccountconnector"
	connectconnection "github.com/upbound/provider-gcp/internal/controller/cluster/developerconnect/connectconnection"
	connectgitrepositorylink "github.com/upbound/provider-gcp/internal/controller/cluster/developerconnect/connectgitrepositorylink"
	agent "github.com/upbound/provider-gcp/internal/controller/cluster/dialogflowcx/agent"
	entitytype "github.com/upbound/provider-gcp/internal/controller/cluster/dialogflowcx/entitytype"
	environmentdialogflowcx "github.com/upbound/provider-gcp/internal/controller/cluster/dialogflowcx/environment"
	flow "github.com/upbound/provider-gcp/internal/controller/cluster/dialogflowcx/flow"
	intent "github.com/upbound/provider-gcp/internal/controller/cluster/dialogflowcx/intent"
	page "github.com/upbound/provider-gcp/internal/controller/cluster/dialogflowcx/page"
	version "github.com/upbound/provider-gcp/internal/controller/cluster/dialogflowcx/version"
	webhook "github.com/upbound/provider-gcp/internal/controller/cluster/dialogflowcx/webhook"
	managedzone "github.com/upbound/provider-gcp/internal/controller/cluster/dns/managedzone"
	managedzoneiammember "github.com/upbound/provider-gcp/internal/controller/cluster/dns/managedzoneiammember"
	policydns "github.com/upbound/provider-gcp/internal/controller/cluster/dns/policy"
	recordset "github.com/upbound/provider-gcp/internal/controller/cluster/dns/recordset"
	responsepolicy "github.com/upbound/provider-gcp/internal/controller/cluster/dns/responsepolicy"
	responsepolicyrule "github.com/upbound/provider-gcp/internal/controller/cluster/dns/responsepolicyrule"
	processor "github.com/upbound/provider-gcp/internal/controller/cluster/documentai/processor"
	contact "github.com/upbound/provider-gcp/internal/controller/cluster/essentialcontacts/contact"
	channel "github.com/upbound/provider-gcp/internal/controller/cluster/eventarc/channel"
	googlechannelconfig "github.com/upbound/provider-gcp/internal/controller/cluster/eventarc/googlechannelconfig"
	triggereventarc "github.com/upbound/provider-gcp/internal/controller/cluster/eventarc/trigger"
	backupfilestore "github.com/upbound/provider-gcp/internal/controller/cluster/filestore/backup"
	instancefilestore "github.com/upbound/provider-gcp/internal/controller/cluster/filestore/instance"
	snapshotfilestore "github.com/upbound/provider-gcp/internal/controller/cluster/filestore/snapshot"
	release "github.com/upbound/provider-gcp/internal/controller/cluster/firebaserules/release"
	ruleset "github.com/upbound/provider-gcp/internal/controller/cluster/firebaserules/ruleset"
	coderepositoryindex "github.com/upbound/provider-gcp/internal/controller/cluster/gemini/coderepositoryindex"
	codetoolssetting "github.com/upbound/provider-gcp/internal/controller/cluster/gemini/codetoolssetting"
	datasharingwithgooglesetting "github.com/upbound/provider-gcp/internal/controller/cluster/gemini/datasharingwithgooglesetting"
	geminigcpenablementsetting "github.com/upbound/provider-gcp/internal/controller/cluster/gemini/geminigcpenablementsetting"
	loggingsetting "github.com/upbound/provider-gcp/internal/controller/cluster/gemini/loggingsetting"
	releasechannelsetting "github.com/upbound/provider-gcp/internal/controller/cluster/gemini/releasechannelsetting"
	repositorygroup "github.com/upbound/provider-gcp/internal/controller/cluster/gemini/repositorygroup"
	backupbackupplan "github.com/upbound/provider-gcp/internal/controller/cluster/gke/backupbackupplan"
	membership "github.com/upbound/provider-gcp/internal/controller/cluster/gkehub/membership"
	membershipiammember "github.com/upbound/provider-gcp/internal/controller/cluster/gkehub/membershipiammember"
	consentstore "github.com/upbound/provider-gcp/internal/controller/cluster/healthcare/consentstore"
	datasethealthcare "github.com/upbound/provider-gcp/internal/controller/cluster/healthcare/dataset"
	datasetiammemberhealthcare "github.com/upbound/provider-gcp/internal/controller/cluster/healthcare/datasetiammember"
	workloadidentitypool "github.com/upbound/provider-gcp/internal/controller/cluster/iam/workloadidentitypool"
	workloadidentitypoolprovider "github.com/upbound/provider-gcp/internal/controller/cluster/iam/workloadidentitypoolprovider"
	appengineserviceiammember "github.com/upbound/provider-gcp/internal/controller/cluster/iap/appengineserviceiammember"
	appengineversioniammember "github.com/upbound/provider-gcp/internal/controller/cluster/iap/appengineversioniammember"
	tunneliammember "github.com/upbound/provider-gcp/internal/controller/cluster/iap/tunneliammember"
	webbackendserviceiammember "github.com/upbound/provider-gcp/internal/controller/cluster/iap/webbackendserviceiammember"
	webiammember "github.com/upbound/provider-gcp/internal/controller/cluster/iap/webiammember"
	webtypeappengineiammember "github.com/upbound/provider-gcp/internal/controller/cluster/iap/webtypeappengineiammember"
	webtypecomputeiammember "github.com/upbound/provider-gcp/internal/controller/cluster/iap/webtypecomputeiammember"
	config "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/config"
	defaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/defaultsupportedidpconfig"
	inboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/inboundsamlconfig"
	oauthidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/oauthidpconfig"
	tenant "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenant"
	tenantdefaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenantdefaultsupportedidpconfig"
	tenantinboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenantinboundsamlconfig"
	tenantoauthidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenantoauthidpconfig"
	cryptokey "github.com/upbound/provider-gcp/internal/controller/cluster/kms/cryptokey"
	cryptokeyiammember "github.com/upbound/provider-gcp/internal/controller/cluster/kms/cryptokeyiammember"
	cryptokeyversion "github.com/upbound/provider-gcp/internal/controller/cluster/kms/cryptokeyversion"
	keyhandle "github.com/upbound/provider-gcp/internal/controller/cluster/kms/keyhandle"
	keyring "github.com/upbound/provider-gcp/internal/controller/cluster/kms/keyring"
	keyringiammember "github.com/upbound/provider-gcp/internal/controller/cluster/kms/keyringiammember"
	keyringimportjob "github.com/upbound/provider-gcp/internal/controller/cluster/kms/keyringimportjob"
	secretciphertext "github.com/upbound/provider-gcp/internal/controller/cluster/kms/secretciphertext"
	folderbucketconfig "github.com/upbound/provider-gcp/internal/controller/cluster/logging/folderbucketconfig"
	folderexclusion "github.com/upbound/provider-gcp/internal/controller/cluster/logging/folderexclusion"
	foldersink "github.com/upbound/provider-gcp/internal/controller/cluster/logging/foldersink"
	logview "github.com/upbound/provider-gcp/internal/controller/cluster/logging/logview"
	metric "github.com/upbound/provider-gcp/internal/controller/cluster/logging/metric"
	projectbucketconfig "github.com/upbound/provider-gcp/internal/controller/cluster/logging/projectbucketconfig"
	projectexclusion "github.com/upbound/provider-gcp/internal/controller/cluster/logging/projectexclusion"
	projectsink "github.com/upbound/provider-gcp/internal/controller/cluster/logging/projectsink"
	instancememcache "github.com/upbound/provider-gcp/internal/controller/cluster/memcache/instance"
	model "github.com/upbound/provider-gcp/internal/controller/cluster/mlengine/model"
	alertpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/monitoring/alertpolicy"
	customservice "github.com/upbound/provider-gcp/internal/controller/cluster/monitoring/customservice"
	dashboard "github.com/upbound/provider-gcp/internal/controller/cluster/monitoring/dashboard"
	group "github.com/upbound/provider-gcp/internal/controller/cluster/monitoring/group"
	metricdescriptor "github.com/upbound/provider-gcp/internal/controller/cluster/monitoring/metricdescriptor"
	notificationchannel "github.com/upbound/provider-gcp/internal/controller/cluster/monitoring/notificationchannel"
	servicemonitoring "github.com/upbound/provider-gcp/internal/controller/cluster/monitoring/service"
	slo "github.com/upbound/provider-gcp/internal/controller/cluster/monitoring/slo"
	uptimecheckconfig "github.com/upbound/provider-gcp/internal/controller/cluster/monitoring/uptimecheckconfig"
	hub "github.com/upbound/provider-gcp/internal/controller/cluster/networkconnectivity/hub"
	serviceconnectionpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/networkconnectivity/serviceconnectionpolicy"
	spoke "github.com/upbound/provider-gcp/internal/controller/cluster/networkconnectivity/spoke"
	connectivitytest "github.com/upbound/provider-gcp/internal/controller/cluster/networkmanagement/connectivitytest"
	addressgroup "github.com/upbound/provider-gcp/internal/controller/cluster/networksecurity/addressgroup"
	gatewaysecuritypolicy "github.com/upbound/provider-gcp/internal/controller/cluster/networksecurity/gatewaysecuritypolicy"
	gatewaysecuritypolicyrule "github.com/upbound/provider-gcp/internal/controller/cluster/networksecurity/gatewaysecuritypolicyrule"
	tlsinspectionpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/networksecurity/tlsinspectionpolicy"
	urllists "github.com/upbound/provider-gcp/internal/controller/cluster/networksecurity/urllists"
	gateway "github.com/upbound/provider-gcp/internal/controller/cluster/networkservices/gateway"
	environmentnotebooks "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/environment"
	instancenotebooks "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/instance"
	instanceiammembernotebooks "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/instanceiammember"
	runtime "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/runtime"
	runtimeiammember "github.com/upbound/provider-gcp/internal/controller/cluster/notebooks/runtimeiammember"
	policyorgpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/orgpolicy/policy"
	ospolicyassignment "github.com/upbound/provider-gcp/internal/controller/cluster/osconfig/ospolicyassignment"
	patchdeployment "github.com/upbound/provider-gcp/internal/controller/cluster/osconfig/patchdeployment"
	sshpublickey "github.com/upbound/provider-gcp/internal/controller/cluster/oslogin/sshpublickey"
	capool "github.com/upbound/provider-gcp/internal/controller/cluster/privateca/capool"
	capooliammember "github.com/upbound/provider-gcp/internal/controller/cluster/privateca/capooliammember"
	certificateprivateca "github.com/upbound/provider-gcp/internal/controller/cluster/privateca/certificate"
	certificateauthority "github.com/upbound/provider-gcp/internal/controller/cluster/privateca/certificateauthority"
	certificatetemplate "github.com/upbound/provider-gcp/internal/controller/cluster/privateca/certificatetemplate"
	certificatetemplateiammember "github.com/upbound/provider-gcp/internal/controller/cluster/privateca/certificatetemplateiammember"
	providerconfig "github.com/upbound/provider-gcp/internal/controller/cluster/providerconfig"
	litereservation "github.com/upbound/provider-gcp/internal/controller/cluster/pubsub/litereservation"
	litesubscription "github.com/upbound/provider-gcp/internal/controller/cluster/pubsub/litesubscription"
	litetopic "github.com/upbound/provider-gcp/internal/controller/cluster/pubsub/litetopic"
	schema "github.com/upbound/provider-gcp/internal/controller/cluster/pubsub/schema"
	subscription "github.com/upbound/provider-gcp/internal/controller/cluster/pubsub/subscription"
	subscriptioniammember "github.com/upbound/provider-gcp/internal/controller/cluster/pubsub/subscriptioniammember"
	topic "github.com/upbound/provider-gcp/internal/controller/cluster/pubsub/topic"
	topiciammember "github.com/upbound/provider-gcp/internal/controller/cluster/pubsub/topiciammember"
	clusterredis "github.com/upbound/provider-gcp/internal/controller/cluster/redis/cluster"
	instanceredis "github.com/upbound/provider-gcp/internal/controller/cluster/redis/instance"
	secret "github.com/upbound/provider-gcp/internal/controller/cluster/secretmanager/secret"
	secretiammember "github.com/upbound/provider-gcp/internal/controller/cluster/secretmanager/secretiammember"
	secretversion "github.com/upbound/provider-gcp/internal/controller/cluster/secretmanager/secretversion"
	connectionservicenetworking "github.com/upbound/provider-gcp/internal/controller/cluster/servicenetworking/connection"
	repository "github.com/upbound/provider-gcp/internal/controller/cluster/sourcerepo/repository"
	repositoryiammember "github.com/upbound/provider-gcp/internal/controller/cluster/sourcerepo/repositoryiammember"
	database "github.com/upbound/provider-gcp/internal/controller/cluster/spanner/database"
	databaseiammember "github.com/upbound/provider-gcp/internal/controller/cluster/spanner/databaseiammember"
	instancespanner "github.com/upbound/provider-gcp/internal/controller/cluster/spanner/instance"
	instanceiammemberspanner "github.com/upbound/provider-gcp/internal/controller/cluster/spanner/instanceiammember"
	databasesql "github.com/upbound/provider-gcp/internal/controller/cluster/sql/database"
	databaseinstance "github.com/upbound/provider-gcp/internal/controller/cluster/sql/databaseinstance"
	sourcerepresentationinstance "github.com/upbound/provider-gcp/internal/controller/cluster/sql/sourcerepresentationinstance"
	sslcert "github.com/upbound/provider-gcp/internal/controller/cluster/sql/sslcert"
	user "github.com/upbound/provider-gcp/internal/controller/cluster/sql/user"
	bucket "github.com/upbound/provider-gcp/internal/controller/cluster/storage/bucket"
	bucketaccesscontrol "github.com/upbound/provider-gcp/internal/controller/cluster/storage/bucketaccesscontrol"
	bucketacl "github.com/upbound/provider-gcp/internal/controller/cluster/storage/bucketacl"
	bucketiammember "github.com/upbound/provider-gcp/internal/controller/cluster/storage/bucketiammember"
	bucketiampolicy "github.com/upbound/provider-gcp/internal/controller/cluster/storage/bucketiampolicy"
	bucketobject "github.com/upbound/provider-gcp/internal/controller/cluster/storage/bucketobject"
	defaultobjectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/cluster/storage/defaultobjectaccesscontrol"
	defaultobjectacl "github.com/upbound/provider-gcp/internal/controller/cluster/storage/defaultobjectacl"
	hmackey "github.com/upbound/provider-gcp/internal/controller/cluster/storage/hmackey"
	notification "github.com/upbound/provider-gcp/internal/controller/cluster/storage/notification"
	objectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/cluster/storage/objectaccesscontrol"
	objectacl "github.com/upbound/provider-gcp/internal/controller/cluster/storage/objectacl"
	agentpool "github.com/upbound/provider-gcp/internal/controller/cluster/storagetransfer/agentpool"
	locationtagbinding "github.com/upbound/provider-gcp/internal/controller/cluster/tags/locationtagbinding"
	tagbinding "github.com/upbound/provider-gcp/internal/controller/cluster/tags/tagbinding"
	tagkey "github.com/upbound/provider-gcp/internal/controller/cluster/tags/tagkey"
	tagvalue "github.com/upbound/provider-gcp/internal/controller/cluster/tags/tagvalue"
	node "github.com/upbound/provider-gcp/internal/controller/cluster/tpu/node"
	datasetvertexai "github.com/upbound/provider-gcp/internal/controller/cluster/vertexai/dataset"
	featurestore "github.com/upbound/provider-gcp/internal/controller/cluster/vertexai/featurestore"
	featurestoreentitytype "github.com/upbound/provider-gcp/internal/controller/cluster/vertexai/featurestoreentitytype"
	tensorboard "github.com/upbound/provider-gcp/internal/controller/cluster/vertexai/tensorboard"
	connector "github.com/upbound/provider-gcp/internal/controller/cluster/vpcaccess/connector"
	workflow "github.com/upbound/provider-gcp/internal/controller/cluster/workflows/workflow"
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
		addonsconfig.Setup,
		endpointattachment.Setup,
		envgroup.Setup,
		envgroupattachment.Setup,
		environment.Setup,
		environmentiammember.Setup,
		instanceapigee.Setup,
		instanceattachment.Setup,
		nataddress.Setup,
		organization.Setup,
		syncauthorization.Setup,
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
		trustconfig.Setup,
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
		networkfirewallpolicyrule.Setup,
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
		regionsslpolicy.Setup,
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
		policytag.Setup,
		tag.Setup,
		tagtemplate.Setup,
		taxonomy.Setup,
		jobdataflow.Setup,
		instancedatafusion.Setup,
		deidentifytemplate.Setup,
		inspecttemplate.Setup,
		jobtrigger.Setup,
		storedinfotype.Setup,
		aspecttype.Setup,
		asset.Setup,
		lake.Setup,
		lakeiampolicy.Setup,
		zone.Setup,
		autoscalingpolicy.Setup,
		clusterdataproc.Setup,
		jobdataproc.Setup,
		metastoreservice.Setup,
		workflowtemplate.Setup,
		connectionprofile.Setup,
		privateconnection.Setup,
		connectaccountconnector.Setup,
		connectconnection.Setup,
		connectgitrepositorylink.Setup,
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
		responsepolicy.Setup,
		responsepolicyrule.Setup,
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
		coderepositoryindex.Setup,
		codetoolssetting.Setup,
		datasharingwithgooglesetting.Setup,
		geminigcpenablementsetting.Setup,
		loggingsetting.Setup,
		releasechannelsetting.Setup,
		repositorygroup.Setup,
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
		config.Setup,
		defaultsupportedidpconfig.Setup,
		inboundsamlconfig.Setup,
		oauthidpconfig.Setup,
		tenant.Setup,
		tenantdefaultsupportedidpconfig.Setup,
		tenantinboundsamlconfig.Setup,
		tenantoauthidpconfig.Setup,
		cryptokey.Setup,
		cryptokeyiammember.Setup,
		cryptokeyversion.Setup,
		keyhandle.Setup,
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
		serviceconnectionpolicy.Setup,
		spoke.Setup,
		connectivitytest.Setup,
		addressgroup.Setup,
		gatewaysecuritypolicy.Setup,
		gatewaysecuritypolicyrule.Setup,
		tlsinspectionpolicy.Setup,
		urllists.Setup,
		gateway.Setup,
		environmentnotebooks.Setup,
		instancenotebooks.Setup,
		instanceiammembernotebooks.Setup,
		runtime.Setup,
		runtimeiammember.Setup,
		policyorgpolicy.Setup,
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
		clusterredis.Setup,
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
		bucketiampolicy.Setup,
		bucketobject.Setup,
		defaultobjectaccesscontrol.Setup,
		defaultobjectacl.Setup,
		hmackey.Setup,
		notification.Setup,
		objectaccesscontrol.Setup,
		objectacl.Setup,
		agentpool.Setup,
		locationtagbinding.Setup,
		tagbinding.Setup,
		tagkey.Setup,
		tagvalue.Setup,
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

// SetupGated_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslevel.SetupGated,
		accesslevelcondition.SetupGated,
		accesspolicy.SetupGated,
		accesspolicyiammember.SetupGated,
		serviceperimeter.SetupGated,
		serviceperimeterresource.SetupGated,
		domain.SetupGated,
		backup.SetupGated,
		cluster.SetupGated,
		instance.SetupGated,
		addonsconfig.SetupGated,
		endpointattachment.SetupGated,
		envgroup.SetupGated,
		envgroupattachment.SetupGated,
		environment.SetupGated,
		environmentiammember.SetupGated,
		instanceapigee.SetupGated,
		instanceattachment.SetupGated,
		nataddress.SetupGated,
		organization.SetupGated,
		syncauthorization.SetupGated,
		application.SetupGated,
		applicationurldispatchrules.SetupGated,
		firewallrule.SetupGated,
		servicenetworksettings.SetupGated,
		standardappversion.SetupGated,
		registryrepository.SetupGated,
		registryrepositoryiammember.SetupGated,
		appconnection.SetupGated,
		appconnector.SetupGated,
		appgateway.SetupGated,
		analyticshubdataexchange.SetupGated,
		analyticshubdataexchangeiammember.SetupGated,
		analyticshublisting.SetupGated,
		connection.SetupGated,
		dataset.SetupGated,
		datasetaccess.SetupGated,
		datasetiambinding.SetupGated,
		datasetiammember.SetupGated,
		datasetiampolicy.SetupGated,
		datatransferconfig.SetupGated,
		job.SetupGated,
		reservation.SetupGated,
		reservationassignment.SetupGated,
		routine.SetupGated,
		table.SetupGated,
		tableiambinding.SetupGated,
		tableiammember.SetupGated,
		tableiampolicy.SetupGated,
		appprofile.SetupGated,
		garbagecollectionpolicy.SetupGated,
		instancebigtable.SetupGated,
		instanceiambinding.SetupGated,
		instanceiammember.SetupGated,
		instanceiampolicy.SetupGated,
		tablebigtable.SetupGated,
		tableiambindingbigtable.SetupGated,
		tableiammemberbigtable.SetupGated,
		tableiampolicybigtable.SetupGated,
		attestor.SetupGated,
		policy.SetupGated,
		certificate.SetupGated,
		certificatemap.SetupGated,
		certificatemapentry.SetupGated,
		dnsauthorization.SetupGated,
		trustconfig.SetupGated,
		idsendpoint.SetupGated,
		trigger.SetupGated,
		workerpool.SetupGated,
		function.SetupGated,
		functioniammember.SetupGated,
		functioncloudfunctions2.SetupGated,
		folder.SetupGated,
		folderiammember.SetupGated,
		organizationiamauditconfig.SetupGated,
		organizationiamcustomrole.SetupGated,
		organizationiammember.SetupGated,
		project.SetupGated,
		projectdefaultserviceaccounts.SetupGated,
		projectiamauditconfig.SetupGated,
		projectiamcustomrole.SetupGated,
		projectiammember.SetupGated,
		projectservice.SetupGated,
		projectusageexportbucket.SetupGated,
		serviceaccount.SetupGated,
		serviceaccountiammember.SetupGated,
		serviceaccountkey.SetupGated,
		servicenetworkingpeereddnsdomain.SetupGated,
		domainmapping.SetupGated,
		service.SetupGated,
		serviceiammember.SetupGated,
		v2job.SetupGated,
		v2service.SetupGated,
		jobcloudscheduler.SetupGated,
		queue.SetupGated,
		environmentcomposer.SetupGated,
		address.SetupGated,
		attacheddisk.SetupGated,
		autoscaler.SetupGated,
		backendbucket.SetupGated,
		backendbucketsignedurlkey.SetupGated,
		backendservice.SetupGated,
		backendservicesignedurlkey.SetupGated,
		disk.SetupGated,
		diskiammember.SetupGated,
		diskresourcepolicyattachment.SetupGated,
		externalvpngateway.SetupGated,
		firewall.SetupGated,
		firewallpolicy.SetupGated,
		firewallpolicyassociation.SetupGated,
		firewallpolicyrule.SetupGated,
		forwardingrule.SetupGated,
		globaladdress.SetupGated,
		globalforwardingrule.SetupGated,
		globalnetworkendpoint.SetupGated,
		globalnetworkendpointgroup.SetupGated,
		havpngateway.SetupGated,
		healthcheck.SetupGated,
		httphealthcheck.SetupGated,
		httpshealthcheck.SetupGated,
		image.SetupGated,
		imageiammember.SetupGated,
		instancecompute.SetupGated,
		instancefromtemplate.SetupGated,
		instancegroup.SetupGated,
		instancegroupmanager.SetupGated,
		instancegroupnamedport.SetupGated,
		instanceiammembercompute.SetupGated,
		instancetemplate.SetupGated,
		interconnectattachment.SetupGated,
		managedsslcertificate.SetupGated,
		network.SetupGated,
		networkendpoint.SetupGated,
		networkendpointgroup.SetupGated,
		networkfirewallpolicy.SetupGated,
		networkfirewallpolicyassociation.SetupGated,
		networkfirewallpolicyrule.SetupGated,
		networkpeering.SetupGated,
		networkpeeringroutesconfig.SetupGated,
		nodegroup.SetupGated,
		nodetemplate.SetupGated,
		packetmirroring.SetupGated,
		perinstanceconfig.SetupGated,
		projectdefaultnetworktier.SetupGated,
		projectmetadata.SetupGated,
		projectmetadataitem.SetupGated,
		regionautoscaler.SetupGated,
		regionbackendservice.SetupGated,
		regiondisk.SetupGated,
		regiondiskiammember.SetupGated,
		regiondiskresourcepolicyattachment.SetupGated,
		regionhealthcheck.SetupGated,
		regioninstancegroupmanager.SetupGated,
		regionnetworkendpoint.SetupGated,
		regionnetworkendpointgroup.SetupGated,
		regionnetworkfirewallpolicy.SetupGated,
		regionnetworkfirewallpolicyassociation.SetupGated,
		regionperinstanceconfig.SetupGated,
		regionsslcertificate.SetupGated,
		regionsslpolicy.SetupGated,
		regiontargethttpproxy.SetupGated,
		regiontargethttpsproxy.SetupGated,
		regiontargettcpproxy.SetupGated,
		regionurlmap.SetupGated,
		reservationcompute.SetupGated,
		resourcepolicy.SetupGated,
		route.SetupGated,
		router.SetupGated,
		routerinterface.SetupGated,
		routernat.SetupGated,
		routerpeer.SetupGated,
		securitypolicy.SetupGated,
		serviceattachment.SetupGated,
		sharedvpchostproject.SetupGated,
		sharedvpcserviceproject.SetupGated,
		snapshot.SetupGated,
		snapshotiammember.SetupGated,
		sslcertificate.SetupGated,
		sslpolicy.SetupGated,
		subnetwork.SetupGated,
		subnetworkiammember.SetupGated,
		targetgrpcproxy.SetupGated,
		targethttpproxy.SetupGated,
		targethttpsproxy.SetupGated,
		targetinstance.SetupGated,
		targetpool.SetupGated,
		targetsslproxy.SetupGated,
		targettcpproxy.SetupGated,
		urlmap.SetupGated,
		vpngateway.SetupGated,
		vpntunnel.SetupGated,
		clustercontainer.SetupGated,
		nodepool.SetupGated,
		registry.SetupGated,
		note.SetupGated,
		clustercontainerattached.SetupGated,
		clustercontaineraws.SetupGated,
		nodepoolcontaineraws.SetupGated,
		client.SetupGated,
		clustercontainerazure.SetupGated,
		nodepoolcontainerazure.SetupGated,
		entry.SetupGated,
		entrygroup.SetupGated,
		policytag.SetupGated,
		tag.SetupGated,
		tagtemplate.SetupGated,
		taxonomy.SetupGated,
		jobdataflow.SetupGated,
		instancedatafusion.SetupGated,
		deidentifytemplate.SetupGated,
		inspecttemplate.SetupGated,
		jobtrigger.SetupGated,
		storedinfotype.SetupGated,
		aspecttype.SetupGated,
		asset.SetupGated,
		lake.SetupGated,
		lakeiampolicy.SetupGated,
		zone.SetupGated,
		autoscalingpolicy.SetupGated,
		clusterdataproc.SetupGated,
		jobdataproc.SetupGated,
		metastoreservice.SetupGated,
		workflowtemplate.SetupGated,
		connectionprofile.SetupGated,
		privateconnection.SetupGated,
		connectaccountconnector.SetupGated,
		connectconnection.SetupGated,
		connectgitrepositorylink.SetupGated,
		agent.SetupGated,
		entitytype.SetupGated,
		environmentdialogflowcx.SetupGated,
		flow.SetupGated,
		intent.SetupGated,
		page.SetupGated,
		version.SetupGated,
		webhook.SetupGated,
		managedzone.SetupGated,
		managedzoneiammember.SetupGated,
		policydns.SetupGated,
		recordset.SetupGated,
		responsepolicy.SetupGated,
		responsepolicyrule.SetupGated,
		processor.SetupGated,
		contact.SetupGated,
		channel.SetupGated,
		googlechannelconfig.SetupGated,
		triggereventarc.SetupGated,
		backupfilestore.SetupGated,
		instancefilestore.SetupGated,
		snapshotfilestore.SetupGated,
		release.SetupGated,
		ruleset.SetupGated,
		coderepositoryindex.SetupGated,
		codetoolssetting.SetupGated,
		datasharingwithgooglesetting.SetupGated,
		geminigcpenablementsetting.SetupGated,
		loggingsetting.SetupGated,
		releasechannelsetting.SetupGated,
		repositorygroup.SetupGated,
		backupbackupplan.SetupGated,
		membership.SetupGated,
		membershipiammember.SetupGated,
		consentstore.SetupGated,
		datasethealthcare.SetupGated,
		datasetiammemberhealthcare.SetupGated,
		workloadidentitypool.SetupGated,
		workloadidentitypoolprovider.SetupGated,
		appengineserviceiammember.SetupGated,
		appengineversioniammember.SetupGated,
		tunneliammember.SetupGated,
		webbackendserviceiammember.SetupGated,
		webiammember.SetupGated,
		webtypeappengineiammember.SetupGated,
		webtypecomputeiammember.SetupGated,
		config.SetupGated,
		defaultsupportedidpconfig.SetupGated,
		inboundsamlconfig.SetupGated,
		oauthidpconfig.SetupGated,
		tenant.SetupGated,
		tenantdefaultsupportedidpconfig.SetupGated,
		tenantinboundsamlconfig.SetupGated,
		tenantoauthidpconfig.SetupGated,
		cryptokey.SetupGated,
		cryptokeyiammember.SetupGated,
		cryptokeyversion.SetupGated,
		keyhandle.SetupGated,
		keyring.SetupGated,
		keyringiammember.SetupGated,
		keyringimportjob.SetupGated,
		secretciphertext.SetupGated,
		folderbucketconfig.SetupGated,
		folderexclusion.SetupGated,
		foldersink.SetupGated,
		logview.SetupGated,
		metric.SetupGated,
		projectbucketconfig.SetupGated,
		projectexclusion.SetupGated,
		projectsink.SetupGated,
		instancememcache.SetupGated,
		model.SetupGated,
		alertpolicy.SetupGated,
		customservice.SetupGated,
		dashboard.SetupGated,
		group.SetupGated,
		metricdescriptor.SetupGated,
		notificationchannel.SetupGated,
		servicemonitoring.SetupGated,
		slo.SetupGated,
		uptimecheckconfig.SetupGated,
		hub.SetupGated,
		serviceconnectionpolicy.SetupGated,
		spoke.SetupGated,
		connectivitytest.SetupGated,
		addressgroup.SetupGated,
		gatewaysecuritypolicy.SetupGated,
		gatewaysecuritypolicyrule.SetupGated,
		tlsinspectionpolicy.SetupGated,
		urllists.SetupGated,
		gateway.SetupGated,
		environmentnotebooks.SetupGated,
		instancenotebooks.SetupGated,
		instanceiammembernotebooks.SetupGated,
		runtime.SetupGated,
		runtimeiammember.SetupGated,
		policyorgpolicy.SetupGated,
		ospolicyassignment.SetupGated,
		patchdeployment.SetupGated,
		sshpublickey.SetupGated,
		capool.SetupGated,
		capooliammember.SetupGated,
		certificateprivateca.SetupGated,
		certificateauthority.SetupGated,
		certificatetemplate.SetupGated,
		certificatetemplateiammember.SetupGated,
		providerconfig.SetupGated,
		litereservation.SetupGated,
		litesubscription.SetupGated,
		litetopic.SetupGated,
		schema.SetupGated,
		subscription.SetupGated,
		subscriptioniammember.SetupGated,
		topic.SetupGated,
		topiciammember.SetupGated,
		clusterredis.SetupGated,
		instanceredis.SetupGated,
		secret.SetupGated,
		secretiammember.SetupGated,
		secretversion.SetupGated,
		connectionservicenetworking.SetupGated,
		repository.SetupGated,
		repositoryiammember.SetupGated,
		database.SetupGated,
		databaseiammember.SetupGated,
		instancespanner.SetupGated,
		instanceiammemberspanner.SetupGated,
		databasesql.SetupGated,
		databaseinstance.SetupGated,
		sourcerepresentationinstance.SetupGated,
		sslcert.SetupGated,
		user.SetupGated,
		bucket.SetupGated,
		bucketaccesscontrol.SetupGated,
		bucketacl.SetupGated,
		bucketiammember.SetupGated,
		bucketiampolicy.SetupGated,
		bucketobject.SetupGated,
		defaultobjectaccesscontrol.SetupGated,
		defaultobjectacl.SetupGated,
		hmackey.SetupGated,
		notification.SetupGated,
		objectaccesscontrol.SetupGated,
		objectacl.SetupGated,
		agentpool.SetupGated,
		locationtagbinding.SetupGated,
		tagbinding.SetupGated,
		tagkey.SetupGated,
		tagvalue.SetupGated,
		node.SetupGated,
		datasetvertexai.SetupGated,
		featurestore.SetupGated,
		featurestoreentitytype.SetupGated,
		tensorboard.SetupGated,
		connector.SetupGated,
		workflow.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
