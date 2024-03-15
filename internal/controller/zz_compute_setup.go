// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
)

// Setup_compute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
