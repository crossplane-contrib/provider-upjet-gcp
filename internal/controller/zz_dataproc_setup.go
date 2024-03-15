// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	autoscalingpolicy "github.com/upbound/provider-gcp/internal/controller/dataproc/autoscalingpolicy"
	cluster "github.com/upbound/provider-gcp/internal/controller/dataproc/cluster"
	job "github.com/upbound/provider-gcp/internal/controller/dataproc/job"
	metastoreservice "github.com/upbound/provider-gcp/internal/controller/dataproc/metastoreservice"
	workflowtemplate "github.com/upbound/provider-gcp/internal/controller/dataproc/workflowtemplate"
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
