// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	addonsconfig "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/addonsconfig"
	api "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/api"
	apiproduct "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/apiproduct"
	appgroup "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/appgroup"
	controlplaneaccess "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/controlplaneaccess"
	developer "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/developer"
	dnszone "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/dnszone"
	endpointattachment "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/endpointattachment"
	envgroup "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/envgroup"
	envgroupattachment "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/envgroupattachment"
	environment "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/environment"
	environmentaddonsconfig "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/environmentaddonsconfig"
	environmentiambinding "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/environmentiambinding"
	environmentiammember "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/environmentiammember"
	environmentiampolicy "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/environmentiampolicy"
	environmentkeyvaluemaps "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/environmentkeyvaluemaps"
	environmentkeyvaluemapsentries "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/environmentkeyvaluemapsentries"
	envkeystore "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/envkeystore"
	envreferences "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/envreferences"
	flowhook "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/flowhook"
	instance "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/instance"
	instanceattachment "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/instanceattachment"
	keystoresaliaseskeycertfile "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/keystoresaliaseskeycertfile"
	keystoresaliasespkcs12 "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/keystoresaliasespkcs12"
	keystoresaliasesselfsignedcert "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/keystoresaliasesselfsignedcert"
	nataddress "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/nataddress"
	organization "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/organization"
	securityaction "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/securityaction"
	securitymonitoringcondition "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/securitymonitoringcondition"
	securityprofilev2 "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/securityprofilev2"
	sharedflow "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/sharedflow"
	sharedflowdeployment "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/sharedflowdeployment"
	syncauthorization "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/syncauthorization"
	targetserver "github.com/upbound/provider-gcp/v2/internal/controller/namespaced/apigee/targetserver"
)

// Setup_apigee creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apigee(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addonsconfig.Setup,
		api.Setup,
		apiproduct.Setup,
		appgroup.Setup,
		controlplaneaccess.Setup,
		developer.Setup,
		dnszone.Setup,
		endpointattachment.Setup,
		envgroup.Setup,
		envgroupattachment.Setup,
		environment.Setup,
		environmentaddonsconfig.Setup,
		environmentiambinding.Setup,
		environmentiammember.Setup,
		environmentiampolicy.Setup,
		environmentkeyvaluemaps.Setup,
		environmentkeyvaluemapsentries.Setup,
		envkeystore.Setup,
		envreferences.Setup,
		flowhook.Setup,
		instance.Setup,
		instanceattachment.Setup,
		keystoresaliaseskeycertfile.Setup,
		keystoresaliasespkcs12.Setup,
		keystoresaliasesselfsignedcert.Setup,
		nataddress.Setup,
		organization.Setup,
		securityaction.Setup,
		securitymonitoringcondition.Setup,
		securityprofilev2.Setup,
		sharedflow.Setup,
		sharedflowdeployment.Setup,
		syncauthorization.Setup,
		targetserver.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_apigee creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apigee(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addonsconfig.SetupGated,
		api.SetupGated,
		apiproduct.SetupGated,
		appgroup.SetupGated,
		controlplaneaccess.SetupGated,
		developer.SetupGated,
		dnszone.SetupGated,
		endpointattachment.SetupGated,
		envgroup.SetupGated,
		envgroupattachment.SetupGated,
		environment.SetupGated,
		environmentaddonsconfig.SetupGated,
		environmentiambinding.SetupGated,
		environmentiammember.SetupGated,
		environmentiampolicy.SetupGated,
		environmentkeyvaluemaps.SetupGated,
		environmentkeyvaluemapsentries.SetupGated,
		envkeystore.SetupGated,
		envreferences.SetupGated,
		flowhook.SetupGated,
		instance.SetupGated,
		instanceattachment.SetupGated,
		keystoresaliaseskeycertfile.SetupGated,
		keystoresaliasespkcs12.SetupGated,
		keystoresaliasesselfsignedcert.SetupGated,
		nataddress.SetupGated,
		organization.SetupGated,
		securityaction.SetupGated,
		securitymonitoringcondition.SetupGated,
		securityprofilev2.SetupGated,
		sharedflow.SetupGated,
		sharedflowdeployment.SetupGated,
		syncauthorization.SetupGated,
		targetserver.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
