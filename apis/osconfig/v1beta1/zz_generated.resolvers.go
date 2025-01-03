// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *PatchDeployment) ResolveReferences( // ResolveReferences of this PatchDeployment.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.InstanceFilter); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.InstanceFilter[i3].Instances),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.ForProvider.InstanceFilter[i3].InstancesRefs,
				Selector:      mg.Spec.ForProvider.InstanceFilter[i3].InstancesSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.InstanceFilter[i3].Instances")
		}
		mg.Spec.ForProvider.InstanceFilter[i3].Instances = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.InstanceFilter[i3].InstancesRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.InstanceFilter); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.InstanceFilter[i3].Instances),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.InitProvider.InstanceFilter[i3].InstancesRefs,
				Selector:      mg.Spec.InitProvider.InstanceFilter[i3].InstancesSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.InstanceFilter[i3].Instances")
		}
		mg.Spec.InitProvider.InstanceFilter[i3].Instances = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.InstanceFilter[i3].InstancesRefs = mrsp.ResolvedReferences

	}

	return nil
}
