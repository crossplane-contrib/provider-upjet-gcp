// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesslevel "github.com/upbound/provider-gcp/internal/controller/namespaced/accesscontextmanager/accesslevel"
	accesslevelcondition "github.com/upbound/provider-gcp/internal/controller/namespaced/accesscontextmanager/accesslevelcondition"
	accesspolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/accesscontextmanager/accesspolicy"
	accesspolicyiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/accesscontextmanager/accesspolicyiammember"
	serviceperimeter "github.com/upbound/provider-gcp/internal/controller/namespaced/accesscontextmanager/serviceperimeter"
	serviceperimeterresource "github.com/upbound/provider-gcp/internal/controller/namespaced/accesscontextmanager/serviceperimeterresource"
	domain "github.com/upbound/provider-gcp/internal/controller/namespaced/activedirectory/domain"
	backup "github.com/upbound/provider-gcp/internal/controller/namespaced/alloydb/backup"
	cluster "github.com/upbound/provider-gcp/internal/controller/namespaced/alloydb/cluster"
	instance "github.com/upbound/provider-gcp/internal/controller/namespaced/alloydb/instance"
	addonsconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/addonsconfig"
	endpointattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/endpointattachment"
	envgroup "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/envgroup"
	envgroupattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/envgroupattachment"
	environment "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/environment"
	environmentiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/environmentiammember"
	envkeystore "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/envkeystore"
	envreferences "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/envreferences"
	instanceapigee "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/instance"
	instanceattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/instanceattachment"
	keystoresaliaseskeycertfile "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/keystoresaliaseskeycertfile"
	nataddress "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/nataddress"
	organization "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/organization"
	syncauthorization "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/syncauthorization"
	targetserver "github.com/upbound/provider-gcp/internal/controller/namespaced/apigee/targetserver"
	application "github.com/upbound/provider-gcp/internal/controller/namespaced/appengine/application"
	applicationurldispatchrules "github.com/upbound/provider-gcp/internal/controller/namespaced/appengine/applicationurldispatchrules"
	firewallrule "github.com/upbound/provider-gcp/internal/controller/namespaced/appengine/firewallrule"
	servicenetworksettings "github.com/upbound/provider-gcp/internal/controller/namespaced/appengine/servicenetworksettings"
	standardappversion "github.com/upbound/provider-gcp/internal/controller/namespaced/appengine/standardappversion"
	registryrepository "github.com/upbound/provider-gcp/internal/controller/namespaced/artifact/registryrepository"
	registryrepositoryiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/artifact/registryrepositoryiammember"
	appconnection "github.com/upbound/provider-gcp/internal/controller/namespaced/beyondcorp/appconnection"
	appconnector "github.com/upbound/provider-gcp/internal/controller/namespaced/beyondcorp/appconnector"
	appgateway "github.com/upbound/provider-gcp/internal/controller/namespaced/beyondcorp/appgateway"
	analyticshubdataexchange "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/analyticshubdataexchange"
	analyticshubdataexchangeiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/analyticshubdataexchangeiammember"
	analyticshublisting "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/analyticshublisting"
	analyticshublistingsubscription "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/analyticshublistingsubscription"
	connection "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/connection"
	dataset "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/dataset"
	datasetaccess "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datasetaccess"
	datasetiambinding "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datasetiambinding"
	datasetiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datasetiammember"
	datasetiampolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datasetiampolicy"
	datatransferconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/datatransferconfig"
	job "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/job"
	reservation "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/reservation"
	reservationassignment "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/reservationassignment"
	routine "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/routine"
	table "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/table"
	tableiambinding "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/tableiambinding"
	tableiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/tableiammember"
	tableiampolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/bigquery/tableiampolicy"
	appprofile "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/appprofile"
	garbagecollectionpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/garbagecollectionpolicy"
	instancebigtable "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/instance"
	instanceiambinding "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/instanceiambinding"
	instanceiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/instanceiammember"
	instanceiampolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/instanceiampolicy"
	tablebigtable "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/table"
	tableiambindingbigtable "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/tableiambinding"
	tableiammemberbigtable "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/tableiammember"
	tableiampolicybigtable "github.com/upbound/provider-gcp/internal/controller/namespaced/bigtable/tableiampolicy"
	attestor "github.com/upbound/provider-gcp/internal/controller/namespaced/binaryauthorization/attestor"
	policy "github.com/upbound/provider-gcp/internal/controller/namespaced/binaryauthorization/policy"
	certificate "github.com/upbound/provider-gcp/internal/controller/namespaced/certificatemanager/certificate"
	certificatemap "github.com/upbound/provider-gcp/internal/controller/namespaced/certificatemanager/certificatemap"
	certificatemapentry "github.com/upbound/provider-gcp/internal/controller/namespaced/certificatemanager/certificatemapentry"
	dnsauthorization "github.com/upbound/provider-gcp/internal/controller/namespaced/certificatemanager/dnsauthorization"
	trustconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/certificatemanager/trustconfig"
	idsendpoint "github.com/upbound/provider-gcp/internal/controller/namespaced/cloud/idsendpoint"
	trigger "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudbuild/trigger"
	workerpool "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudbuild/workerpool"
	function "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudfunctions/function"
	functioniammember "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudfunctions/functioniammember"
	functioncloudfunctions2 "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudfunctions2/function"
	folder "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/folder"
	folderiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/folderiammember"
	organizationiamauditconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/organizationiamauditconfig"
	organizationiamcustomrole "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/organizationiamcustomrole"
	organizationiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/organizationiammember"
	project "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/project"
	projectdefaultserviceaccounts "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/projectdefaultserviceaccounts"
	projectiamauditconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/projectiamauditconfig"
	projectiamcustomrole "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/projectiamcustomrole"
	projectiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/projectiammember"
	projectservice "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/projectservice"
	projectusageexportbucket "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/projectusageexportbucket"
	serviceaccount "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudplatform/servicenetworkingpeereddnsdomain"
	domainmapping "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudrun/domainmapping"
	service "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudrun/service"
	serviceiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudrun/serviceiammember"
	v2job "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudrun/v2job"
	v2service "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudrun/v2service"
	jobcloudscheduler "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudscheduler/job"
	queue "github.com/upbound/provider-gcp/internal/controller/namespaced/cloudtasks/queue"
	environmentcomposer "github.com/upbound/provider-gcp/internal/controller/namespaced/composer/environment"
	address "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/address"
	attacheddisk "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/attacheddisk"
	autoscaler "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/autoscaler"
	backendbucket "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/backendbucket"
	backendbucketsignedurlkey "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/backendbucketsignedurlkey"
	backendservice "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/backendservice"
	backendservicesignedurlkey "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/backendservicesignedurlkey"
	disk "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/disk"
	diskiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/diskiammember"
	diskresourcepolicyattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/diskresourcepolicyattachment"
	externalvpngateway "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/externalvpngateway"
	firewall "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/firewall"
	firewallpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/firewallpolicy"
	firewallpolicyassociation "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/firewallpolicyassociation"
	firewallpolicyrule "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/firewallpolicyrule"
	forwardingrule "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/forwardingrule"
	globaladdress "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/globaladdress"
	globalforwardingrule "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/globalforwardingrule"
	globalnetworkendpoint "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/globalnetworkendpoint"
	globalnetworkendpointgroup "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/globalnetworkendpointgroup"
	havpngateway "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/havpngateway"
	healthcheck "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/healthcheck"
	httphealthcheck "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/httphealthcheck"
	httpshealthcheck "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/httpshealthcheck"
	image "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/image"
	imageiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/imageiammember"
	instancecompute "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instance"
	instancefromtemplate "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instancefromtemplate"
	instancegroup "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instancegroup"
	instancegroupmanager "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instancegroupmanager"
	instancegroupnamedport "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instancegroupnamedport"
	instanceiammembercompute "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instanceiammember"
	instancetemplate "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instancetemplate"
	interconnectattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/interconnectattachment"
	managedsslcertificate "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/managedsslcertificate"
	network "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/network"
	networkendpoint "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/networkendpoint"
	networkendpointgroup "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/networkendpointgroup"
	networkfirewallpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/networkfirewallpolicy"
	networkfirewallpolicyassociation "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/networkfirewallpolicyassociation"
	networkfirewallpolicyrule "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/networkfirewallpolicyrule"
	networkpeering "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/networkpeering"
	networkpeeringroutesconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/networkpeeringroutesconfig"
	nodegroup "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/nodegroup"
	nodetemplate "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/nodetemplate"
	packetmirroring "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/packetmirroring"
	perinstanceconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/perinstanceconfig"
	projectdefaultnetworktier "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/projectdefaultnetworktier"
	projectmetadata "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/projectmetadata"
	projectmetadataitem "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/projectmetadataitem"
	regionautoscaler "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionautoscaler"
	regionbackendservice "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionbackendservice"
	regiondisk "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regiondisk"
	regiondiskiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regiondiskiammember"
	regiondiskresourcepolicyattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regiondiskresourcepolicyattachment"
	regionhealthcheck "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionhealthcheck"
	regioninstancegroupmanager "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regioninstancegroupmanager"
	regionnetworkendpoint "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionnetworkendpoint"
	regionnetworkendpointgroup "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionnetworkendpointgroup"
	regionnetworkfirewallpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionnetworkfirewallpolicy"
	regionnetworkfirewallpolicyassociation "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionnetworkfirewallpolicyassociation"
	regionperinstanceconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionperinstanceconfig"
	regionsecuritypolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionsecuritypolicy"
	regionsslcertificate "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionsslcertificate"
	regionsslpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionsslpolicy"
	regiontargethttpproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regiontargethttpsproxy"
	regiontargettcpproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regiontargettcpproxy"
	regionurlmap "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionurlmap"
	reservationcompute "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/reservation"
	resourcepolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/resourcepolicy"
	route "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/route"
	router "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/router"
	routerinterface "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/routerinterface"
	routernat "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/routernat"
	routerpeer "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/routerpeer"
	securitypolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/securitypolicy"
	serviceattachment "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/serviceattachment"
	sharedvpchostproject "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/sharedvpchostproject"
	sharedvpcserviceproject "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/sharedvpcserviceproject"
	snapshot "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/snapshot"
	snapshotiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/snapshotiammember"
	sslcertificate "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/sslcertificate"
	sslpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/sslpolicy"
	subnetwork "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/subnetwork"
	subnetworkiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/subnetworkiammember"
	targetgrpcproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/targetgrpcproxy"
	targethttpproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/targethttpproxy"
	targethttpsproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/targethttpsproxy"
	targetinstance "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/targetinstance"
	targetpool "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/targetpool"
	targetsslproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/targetsslproxy"
	targettcpproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/targettcpproxy"
	urlmap "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/urlmap"
	vpngateway "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/vpngateway"
	vpntunnel "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/vpntunnel"
	clustercontainer "github.com/upbound/provider-gcp/internal/controller/namespaced/container/cluster"
	nodepool "github.com/upbound/provider-gcp/internal/controller/namespaced/container/nodepool"
	registry "github.com/upbound/provider-gcp/internal/controller/namespaced/container/registry"
	note "github.com/upbound/provider-gcp/internal/controller/namespaced/containeranalysis/note"
	clustercontainerattached "github.com/upbound/provider-gcp/internal/controller/namespaced/containerattached/cluster"
	clustercontaineraws "github.com/upbound/provider-gcp/internal/controller/namespaced/containeraws/cluster"
	nodepoolcontaineraws "github.com/upbound/provider-gcp/internal/controller/namespaced/containeraws/nodepool"
	client "github.com/upbound/provider-gcp/internal/controller/namespaced/containerazure/client"
	clustercontainerazure "github.com/upbound/provider-gcp/internal/controller/namespaced/containerazure/cluster"
	nodepoolcontainerazure "github.com/upbound/provider-gcp/internal/controller/namespaced/containerazure/nodepool"
	entry "github.com/upbound/provider-gcp/internal/controller/namespaced/datacatalog/entry"
	entrygroup "github.com/upbound/provider-gcp/internal/controller/namespaced/datacatalog/entrygroup"
	policytag "github.com/upbound/provider-gcp/internal/controller/namespaced/datacatalog/policytag"
	tag "github.com/upbound/provider-gcp/internal/controller/namespaced/datacatalog/tag"
	tagtemplate "github.com/upbound/provider-gcp/internal/controller/namespaced/datacatalog/tagtemplate"
	taxonomy "github.com/upbound/provider-gcp/internal/controller/namespaced/datacatalog/taxonomy"
	jobdataflow "github.com/upbound/provider-gcp/internal/controller/namespaced/dataflow/job"
	instancedatafusion "github.com/upbound/provider-gcp/internal/controller/namespaced/datafusion/instance"
	deidentifytemplate "github.com/upbound/provider-gcp/internal/controller/namespaced/datalossprevention/deidentifytemplate"
	inspecttemplate "github.com/upbound/provider-gcp/internal/controller/namespaced/datalossprevention/inspecttemplate"
	jobtrigger "github.com/upbound/provider-gcp/internal/controller/namespaced/datalossprevention/jobtrigger"
	storedinfotype "github.com/upbound/provider-gcp/internal/controller/namespaced/datalossprevention/storedinfotype"
	aspecttype "github.com/upbound/provider-gcp/internal/controller/namespaced/dataplex/aspecttype"
	asset "github.com/upbound/provider-gcp/internal/controller/namespaced/dataplex/asset"
	lake "github.com/upbound/provider-gcp/internal/controller/namespaced/dataplex/lake"
	lakeiampolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/dataplex/lakeiampolicy"
	zone "github.com/upbound/provider-gcp/internal/controller/namespaced/dataplex/zone"
	autoscalingpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/dataproc/autoscalingpolicy"
	clusterdataproc "github.com/upbound/provider-gcp/internal/controller/namespaced/dataproc/cluster"
	jobdataproc "github.com/upbound/provider-gcp/internal/controller/namespaced/dataproc/job"
	metastoreservice "github.com/upbound/provider-gcp/internal/controller/namespaced/dataproc/metastoreservice"
	workflowtemplate "github.com/upbound/provider-gcp/internal/controller/namespaced/dataproc/workflowtemplate"
	connectionprofile "github.com/upbound/provider-gcp/internal/controller/namespaced/datastream/connectionprofile"
	privateconnection "github.com/upbound/provider-gcp/internal/controller/namespaced/datastream/privateconnection"
	connectaccountconnector "github.com/upbound/provider-gcp/internal/controller/namespaced/developerconnect/connectaccountconnector"
	connectconnection "github.com/upbound/provider-gcp/internal/controller/namespaced/developerconnect/connectconnection"
	connectgitrepositorylink "github.com/upbound/provider-gcp/internal/controller/namespaced/developerconnect/connectgitrepositorylink"
	agent "github.com/upbound/provider-gcp/internal/controller/namespaced/dialogflowcx/agent"
	entitytype "github.com/upbound/provider-gcp/internal/controller/namespaced/dialogflowcx/entitytype"
	environmentdialogflowcx "github.com/upbound/provider-gcp/internal/controller/namespaced/dialogflowcx/environment"
	flow "github.com/upbound/provider-gcp/internal/controller/namespaced/dialogflowcx/flow"
	intent "github.com/upbound/provider-gcp/internal/controller/namespaced/dialogflowcx/intent"
	page "github.com/upbound/provider-gcp/internal/controller/namespaced/dialogflowcx/page"
	version "github.com/upbound/provider-gcp/internal/controller/namespaced/dialogflowcx/version"
	webhook "github.com/upbound/provider-gcp/internal/controller/namespaced/dialogflowcx/webhook"
	managedzone "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/managedzone"
	managedzoneiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/managedzoneiammember"
	policydns "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/policy"
	recordset "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/recordset"
	responsepolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/responsepolicy"
	responsepolicyrule "github.com/upbound/provider-gcp/internal/controller/namespaced/dns/responsepolicyrule"
	processor "github.com/upbound/provider-gcp/internal/controller/namespaced/documentai/processor"
	contact "github.com/upbound/provider-gcp/internal/controller/namespaced/essentialcontacts/contact"
	channel "github.com/upbound/provider-gcp/internal/controller/namespaced/eventarc/channel"
	googlechannelconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/eventarc/googlechannelconfig"
	triggereventarc "github.com/upbound/provider-gcp/internal/controller/namespaced/eventarc/trigger"
	backupfilestore "github.com/upbound/provider-gcp/internal/controller/namespaced/filestore/backup"
	instancefilestore "github.com/upbound/provider-gcp/internal/controller/namespaced/filestore/instance"
	snapshotfilestore "github.com/upbound/provider-gcp/internal/controller/namespaced/filestore/snapshot"
	release "github.com/upbound/provider-gcp/internal/controller/namespaced/firebaserules/release"
	ruleset "github.com/upbound/provider-gcp/internal/controller/namespaced/firebaserules/ruleset"
	coderepositoryindex "github.com/upbound/provider-gcp/internal/controller/namespaced/gemini/coderepositoryindex"
	codetoolssetting "github.com/upbound/provider-gcp/internal/controller/namespaced/gemini/codetoolssetting"
	datasharingwithgooglesetting "github.com/upbound/provider-gcp/internal/controller/namespaced/gemini/datasharingwithgooglesetting"
	geminigcpenablementsetting "github.com/upbound/provider-gcp/internal/controller/namespaced/gemini/geminigcpenablementsetting"
	loggingsetting "github.com/upbound/provider-gcp/internal/controller/namespaced/gemini/loggingsetting"
	releasechannelsetting "github.com/upbound/provider-gcp/internal/controller/namespaced/gemini/releasechannelsetting"
	repositorygroup "github.com/upbound/provider-gcp/internal/controller/namespaced/gemini/repositorygroup"
	backupbackupplan "github.com/upbound/provider-gcp/internal/controller/namespaced/gke/backupbackupplan"
	membership "github.com/upbound/provider-gcp/internal/controller/namespaced/gkehub/membership"
	membershipiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/gkehub/membershipiammember"
	consentstore "github.com/upbound/provider-gcp/internal/controller/namespaced/healthcare/consentstore"
	datasethealthcare "github.com/upbound/provider-gcp/internal/controller/namespaced/healthcare/dataset"
	datasetiammemberhealthcare "github.com/upbound/provider-gcp/internal/controller/namespaced/healthcare/datasetiammember"
	workloadidentitypool "github.com/upbound/provider-gcp/internal/controller/namespaced/iam/workloadidentitypool"
	workloadidentitypoolprovider "github.com/upbound/provider-gcp/internal/controller/namespaced/iam/workloadidentitypoolprovider"
	appengineserviceiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/iap/appengineserviceiammember"
	appengineversioniammember "github.com/upbound/provider-gcp/internal/controller/namespaced/iap/appengineversioniammember"
	tunneliammember "github.com/upbound/provider-gcp/internal/controller/namespaced/iap/tunneliammember"
	webbackendserviceiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/iap/webbackendserviceiammember"
	webiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/iap/webiammember"
	webtypeappengineiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/iap/webtypeappengineiammember"
	webtypecomputeiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/iap/webtypecomputeiammember"
	config "github.com/upbound/provider-gcp/internal/controller/namespaced/identityplatform/config"
	defaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/identityplatform/defaultsupportedidpconfig"
	inboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/identityplatform/inboundsamlconfig"
	oauthidpconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/identityplatform/oauthidpconfig"
	tenant "github.com/upbound/provider-gcp/internal/controller/namespaced/identityplatform/tenant"
	tenantdefaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/identityplatform/tenantdefaultsupportedidpconfig"
	tenantinboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/identityplatform/tenantinboundsamlconfig"
	tenantoauthidpconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/identityplatform/tenantoauthidpconfig"
	cryptokey "github.com/upbound/provider-gcp/internal/controller/namespaced/kms/cryptokey"
	cryptokeyiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/kms/cryptokeyiammember"
	cryptokeyversion "github.com/upbound/provider-gcp/internal/controller/namespaced/kms/cryptokeyversion"
	keyring "github.com/upbound/provider-gcp/internal/controller/namespaced/kms/keyring"
	keyringiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/kms/keyringiammember"
	keyringimportjob "github.com/upbound/provider-gcp/internal/controller/namespaced/kms/keyringimportjob"
	secretciphertext "github.com/upbound/provider-gcp/internal/controller/namespaced/kms/secretciphertext"
	folderbucketconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/logging/folderbucketconfig"
	folderexclusion "github.com/upbound/provider-gcp/internal/controller/namespaced/logging/folderexclusion"
	foldersink "github.com/upbound/provider-gcp/internal/controller/namespaced/logging/foldersink"
	logview "github.com/upbound/provider-gcp/internal/controller/namespaced/logging/logview"
	metric "github.com/upbound/provider-gcp/internal/controller/namespaced/logging/metric"
	projectbucketconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/logging/projectbucketconfig"
	projectexclusion "github.com/upbound/provider-gcp/internal/controller/namespaced/logging/projectexclusion"
	projectsink "github.com/upbound/provider-gcp/internal/controller/namespaced/logging/projectsink"
	instancememcache "github.com/upbound/provider-gcp/internal/controller/namespaced/memcache/instance"
	model "github.com/upbound/provider-gcp/internal/controller/namespaced/mlengine/model"
	alertpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/monitoring/alertpolicy"
	customservice "github.com/upbound/provider-gcp/internal/controller/namespaced/monitoring/customservice"
	dashboard "github.com/upbound/provider-gcp/internal/controller/namespaced/monitoring/dashboard"
	group "github.com/upbound/provider-gcp/internal/controller/namespaced/monitoring/group"
	metricdescriptor "github.com/upbound/provider-gcp/internal/controller/namespaced/monitoring/metricdescriptor"
	notificationchannel "github.com/upbound/provider-gcp/internal/controller/namespaced/monitoring/notificationchannel"
	servicemonitoring "github.com/upbound/provider-gcp/internal/controller/namespaced/monitoring/service"
	slo "github.com/upbound/provider-gcp/internal/controller/namespaced/monitoring/slo"
	uptimecheckconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/monitoring/uptimecheckconfig"
	groupnetworkconnectivity "github.com/upbound/provider-gcp/internal/controller/namespaced/networkconnectivity/group"
	hub "github.com/upbound/provider-gcp/internal/controller/namespaced/networkconnectivity/hub"
	serviceconnectionpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/networkconnectivity/serviceconnectionpolicy"
	spoke "github.com/upbound/provider-gcp/internal/controller/namespaced/networkconnectivity/spoke"
	connectivitytest "github.com/upbound/provider-gcp/internal/controller/namespaced/networkmanagement/connectivitytest"
	addressgroup "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/addressgroup"
	gatewaysecuritypolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/gatewaysecuritypolicy"
	gatewaysecuritypolicyrule "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/gatewaysecuritypolicyrule"
	tlsinspectionpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/tlsinspectionpolicy"
	urllists "github.com/upbound/provider-gcp/internal/controller/namespaced/networksecurity/urllists"
	gateway "github.com/upbound/provider-gcp/internal/controller/namespaced/networkservices/gateway"
	environmentnotebooks "github.com/upbound/provider-gcp/internal/controller/namespaced/notebooks/environment"
	instancenotebooks "github.com/upbound/provider-gcp/internal/controller/namespaced/notebooks/instance"
	instanceiammembernotebooks "github.com/upbound/provider-gcp/internal/controller/namespaced/notebooks/instanceiammember"
	runtime "github.com/upbound/provider-gcp/internal/controller/namespaced/notebooks/runtime"
	runtimeiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/notebooks/runtimeiammember"
	policyorgpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/orgpolicy/policy"
	ospolicyassignment "github.com/upbound/provider-gcp/internal/controller/namespaced/osconfig/ospolicyassignment"
	patchdeployment "github.com/upbound/provider-gcp/internal/controller/namespaced/osconfig/patchdeployment"
	sshpublickey "github.com/upbound/provider-gcp/internal/controller/namespaced/oslogin/sshpublickey"
	capool "github.com/upbound/provider-gcp/internal/controller/namespaced/privateca/capool"
	capooliammember "github.com/upbound/provider-gcp/internal/controller/namespaced/privateca/capooliammember"
	certificateprivateca "github.com/upbound/provider-gcp/internal/controller/namespaced/privateca/certificate"
	certificateauthority "github.com/upbound/provider-gcp/internal/controller/namespaced/privateca/certificateauthority"
	certificatetemplate "github.com/upbound/provider-gcp/internal/controller/namespaced/privateca/certificatetemplate"
	certificatetemplateiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/privateca/certificatetemplateiammember"
	providerconfig "github.com/upbound/provider-gcp/internal/controller/namespaced/providerconfig"
	litereservation "github.com/upbound/provider-gcp/internal/controller/namespaced/pubsub/litereservation"
	litesubscription "github.com/upbound/provider-gcp/internal/controller/namespaced/pubsub/litesubscription"
	litetopic "github.com/upbound/provider-gcp/internal/controller/namespaced/pubsub/litetopic"
	schema "github.com/upbound/provider-gcp/internal/controller/namespaced/pubsub/schema"
	subscription "github.com/upbound/provider-gcp/internal/controller/namespaced/pubsub/subscription"
	subscriptioniammember "github.com/upbound/provider-gcp/internal/controller/namespaced/pubsub/subscriptioniammember"
	topic "github.com/upbound/provider-gcp/internal/controller/namespaced/pubsub/topic"
	topiciammember "github.com/upbound/provider-gcp/internal/controller/namespaced/pubsub/topiciammember"
	clusterredis "github.com/upbound/provider-gcp/internal/controller/namespaced/redis/cluster"
	instanceredis "github.com/upbound/provider-gcp/internal/controller/namespaced/redis/instance"
	secret "github.com/upbound/provider-gcp/internal/controller/namespaced/secretmanager/secret"
	secretiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/secretmanager/secretiammember"
	secretversion "github.com/upbound/provider-gcp/internal/controller/namespaced/secretmanager/secretversion"
	connectionservicenetworking "github.com/upbound/provider-gcp/internal/controller/namespaced/servicenetworking/connection"
	repository "github.com/upbound/provider-gcp/internal/controller/namespaced/sourcerepo/repository"
	repositoryiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/sourcerepo/repositoryiammember"
	database "github.com/upbound/provider-gcp/internal/controller/namespaced/spanner/database"
	databaseiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/spanner/databaseiammember"
	instancespanner "github.com/upbound/provider-gcp/internal/controller/namespaced/spanner/instance"
	instanceiammemberspanner "github.com/upbound/provider-gcp/internal/controller/namespaced/spanner/instanceiammember"
	databasesql "github.com/upbound/provider-gcp/internal/controller/namespaced/sql/database"
	databaseinstance "github.com/upbound/provider-gcp/internal/controller/namespaced/sql/databaseinstance"
	sourcerepresentationinstance "github.com/upbound/provider-gcp/internal/controller/namespaced/sql/sourcerepresentationinstance"
	sslcert "github.com/upbound/provider-gcp/internal/controller/namespaced/sql/sslcert"
	user "github.com/upbound/provider-gcp/internal/controller/namespaced/sql/user"
	bucket "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucket"
	bucketaccesscontrol "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketaccesscontrol"
	bucketacl "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketacl"
	bucketiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketiammember"
	bucketiampolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketiampolicy"
	bucketobject "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/bucketobject"
	defaultobjectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/defaultobjectaccesscontrol"
	defaultobjectacl "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/defaultobjectacl"
	hmackey "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/hmackey"
	managedfolder "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/managedfolder"
	managedfolderiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/managedfolderiammember"
	notification "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/notification"
	objectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/objectaccesscontrol"
	objectacl "github.com/upbound/provider-gcp/internal/controller/namespaced/storage/objectacl"
	agentpool "github.com/upbound/provider-gcp/internal/controller/namespaced/storagetransfer/agentpool"
	locationtagbinding "github.com/upbound/provider-gcp/internal/controller/namespaced/tags/locationtagbinding"
	tagbinding "github.com/upbound/provider-gcp/internal/controller/namespaced/tags/tagbinding"
	tagkey "github.com/upbound/provider-gcp/internal/controller/namespaced/tags/tagkey"
	tagvalue "github.com/upbound/provider-gcp/internal/controller/namespaced/tags/tagvalue"
	node "github.com/upbound/provider-gcp/internal/controller/namespaced/tpu/node"
	datasetvertexai "github.com/upbound/provider-gcp/internal/controller/namespaced/vertexai/dataset"
	featurestore "github.com/upbound/provider-gcp/internal/controller/namespaced/vertexai/featurestore"
	featurestoreentitytype "github.com/upbound/provider-gcp/internal/controller/namespaced/vertexai/featurestoreentitytype"
	tensorboard "github.com/upbound/provider-gcp/internal/controller/namespaced/vertexai/tensorboard"
	connector "github.com/upbound/provider-gcp/internal/controller/namespaced/vpcaccess/connector"
	workflow "github.com/upbound/provider-gcp/internal/controller/namespaced/workflows/workflow"
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
		envkeystore.Setup,
		envreferences.Setup,
		instanceapigee.Setup,
		instanceattachment.Setup,
		keystoresaliaseskeycertfile.Setup,
		nataddress.Setup,
		organization.Setup,
		syncauthorization.Setup,
		targetserver.Setup,
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
		analyticshublistingsubscription.Setup,
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
		regionsecuritypolicy.Setup,
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
		groupnetworkconnectivity.Setup,
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
		managedfolder.Setup,
		managedfolderiammember.Setup,
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
		envkeystore.SetupGated,
		envreferences.SetupGated,
		instanceapigee.SetupGated,
		instanceattachment.SetupGated,
		keystoresaliaseskeycertfile.SetupGated,
		nataddress.SetupGated,
		organization.SetupGated,
		syncauthorization.SetupGated,
		targetserver.SetupGated,
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
		analyticshublistingsubscription.SetupGated,
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
		regionsecuritypolicy.SetupGated,
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
		groupnetworkconnectivity.SetupGated,
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
		managedfolder.SetupGated,
		managedfolderiammember.SetupGated,
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
