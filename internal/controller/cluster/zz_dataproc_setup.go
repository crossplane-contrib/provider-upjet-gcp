// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	autoscalingpolicy "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/autoscalingpolicy"
	cluster "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/cluster"
	job "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/job"
	metastoreservice "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/metastoreservice"
	workflowtemplate "github.com/upbound/provider-gcp/internal/controller/cluster/dataproc/workflowtemplate"
)

// Setup_dataproc creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataproc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingpolicy.Setup,
		cluster.Setup,
		job.Setup,
		metastoreservice.Setup,
		workflowtemplate.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dataproc creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dataproc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingpolicy.SetupGated,
		cluster.SetupGated,
		job.SetupGated,
		metastoreservice.SetupGated,
		workflowtemplate.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
