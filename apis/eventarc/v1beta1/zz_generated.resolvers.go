// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Trigger.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Trigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Destination); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Destination[i3].CloudRunService); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudrun.gcp.upbound.io", "v1beta1", "Service", "ServiceList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Destination[i3].CloudRunService[i4].Service),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.Destination[i3].CloudRunService[i4].ServiceRef,
					Selector:     mg.Spec.ForProvider.Destination[i3].CloudRunService[i4].ServiceSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Destination[i3].CloudRunService[i4].Service")
			}
			mg.Spec.ForProvider.Destination[i3].CloudRunService[i4].Service = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Destination[i3].CloudRunService[i4].ServiceRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Destination); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Destination[i3].CloudRunService); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudrun.gcp.upbound.io", "v1beta1", "Service", "ServiceList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Destination[i3].CloudRunService[i4].Service),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.Destination[i3].CloudRunService[i4].ServiceRef,
					Selector:     mg.Spec.InitProvider.Destination[i3].CloudRunService[i4].ServiceSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Destination[i3].CloudRunService[i4].Service")
			}
			mg.Spec.InitProvider.Destination[i3].CloudRunService[i4].Service = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Destination[i3].CloudRunService[i4].ServiceRef = rsp.ResolvedReference

		}
	}

	return nil
}