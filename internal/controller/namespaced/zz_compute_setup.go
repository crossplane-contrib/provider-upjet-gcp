// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
	instance "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instance"
	instancefromtemplate "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instancefromtemplate"
	instancegroup "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instancegroup"
	instancegroupmanager "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instancegroupmanager"
	instancegroupnamedport "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instancegroupnamedport"
	instanceiammember "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/instanceiammember"
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
	regionsslcertificate "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionsslcertificate"
	regionsslpolicy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionsslpolicy"
	regiontargethttpproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regiontargethttpsproxy"
	regiontargettcpproxy "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regiontargettcpproxy"
	regionurlmap "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/regionurlmap"
	reservation "github.com/upbound/provider-gcp/internal/controller/namespaced/compute/reservation"
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
