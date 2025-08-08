// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

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
	instance "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instance"
	instancefromtemplate "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instancefromtemplate"
	instancegroup "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instancegroup"
	instancegroupmanager "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instancegroupmanager"
	instancegroupnamedport "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instancegroupnamedport"
	instanceiammember "github.com/upbound/provider-gcp/internal/controller/cluster/compute/instanceiammember"
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
	reservation "github.com/upbound/provider-gcp/internal/controller/cluster/compute/reservation"
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

// SetupGated_compute creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
		instance.SetupGated,
		instancefromtemplate.SetupGated,
		instancegroup.SetupGated,
		instancegroupmanager.SetupGated,
		instancegroupnamedport.SetupGated,
		instanceiammember.SetupGated,
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
		reservation.SetupGated,
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
