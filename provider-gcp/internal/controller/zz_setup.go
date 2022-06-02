/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	folder "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/folder"
	project "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/project"
	serviceaccount "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountkey "github.com/upbound/official-providers/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	address "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/address"
	firewall "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/firewall"
	instance "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/instance"
	managedsslcertificate "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/managedsslcertificate"
	network "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/network"
	router "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/router"
	routernat "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/routernat"
	subnetwork "github.com/upbound/official-providers/provider-gcp/internal/controller/compute/subnetwork"
	cluster "github.com/upbound/official-providers/provider-gcp/internal/controller/container/cluster"
	nodepool "github.com/upbound/official-providers/provider-gcp/internal/controller/container/nodepool"
	alertpolicy "github.com/upbound/official-providers/provider-gcp/internal/controller/monitoring/alertpolicy"
	notificationchannel "github.com/upbound/official-providers/provider-gcp/internal/controller/monitoring/notificationchannel"
	uptimecheckconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/monitoring/uptimecheckconfig"
	providerconfig "github.com/upbound/official-providers/provider-gcp/internal/controller/providerconfig"
	instanceredis "github.com/upbound/official-providers/provider-gcp/internal/controller/redis/instance"
	database "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/database"
	databaseinstance "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/databaseinstance"
	sourcerepresentationinstance "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/sourcerepresentationinstance"
	sslcert "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/sslcert"
	user "github.com/upbound/official-providers/provider-gcp/internal/controller/sql/user"
	bucket "github.com/upbound/official-providers/provider-gcp/internal/controller/storage/bucket"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		folder.Setup,
		project.Setup,
		serviceaccount.Setup,
		serviceaccountkey.Setup,
		address.Setup,
		firewall.Setup,
		instance.Setup,
		managedsslcertificate.Setup,
		network.Setup,
		router.Setup,
		routernat.Setup,
		subnetwork.Setup,
		cluster.Setup,
		nodepool.Setup,
		alertpolicy.Setup,
		notificationchannel.Setup,
		uptimecheckconfig.Setup,
		providerconfig.Setup,
		instanceredis.Setup,
		database.Setup,
		databaseinstance.Setup,
		sourcerepresentationinstance.Setup,
		sslcert.Setup,
		user.Setup,
		bucket.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
