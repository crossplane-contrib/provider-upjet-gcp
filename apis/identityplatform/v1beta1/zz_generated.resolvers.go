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
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Config.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Config) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Project", "ProjectList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
			Extract:      resource.ExtractParamPath("project_id", false),
			Reference:    mg.Spec.ForProvider.ProjectRef,
			Selector:     mg.Spec.ForProvider.ProjectSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Project", "ProjectList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Project),
			Extract:      resource.ExtractParamPath("project_id", false),
			Reference:    mg.Spec.InitProvider.ProjectRef,
			Selector:     mg.Spec.InitProvider.ProjectSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Project")
	}
	mg.Spec.InitProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TenantDefaultSupportedIdPConfig.
func (mg *TenantDefaultSupportedIdPConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("identityplatform.gcp.upbound.io", "v1beta1", "Tenant", "TenantList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Tenant),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.TenantRef,
			Selector:     mg.Spec.ForProvider.TenantSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Tenant")
	}
	mg.Spec.ForProvider.Tenant = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("identityplatform.gcp.upbound.io", "v1beta1", "Tenant", "TenantList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Tenant),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.InitProvider.TenantRef,
			Selector:     mg.Spec.InitProvider.TenantSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Tenant")
	}
	mg.Spec.InitProvider.Tenant = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TenantRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("identityplatform.gcp.upbound.io", "v1beta1", "Tenant", "TenantList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Tenant),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.TenantRef,
			Selector:     mg.Spec.ForProvider.TenantSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Tenant")
	}
	mg.Spec.ForProvider.Tenant = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("identityplatform.gcp.upbound.io", "v1beta1", "Tenant", "TenantList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Tenant),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.InitProvider.TenantRef,
			Selector:     mg.Spec.InitProvider.TenantSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Tenant")
	}
	mg.Spec.InitProvider.Tenant = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TenantRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TenantOAuthIdPConfig.
func (mg *TenantOAuthIdPConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("identityplatform.gcp.upbound.io", "v1beta1", "Tenant", "TenantList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Tenant),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.TenantRef,
			Selector:     mg.Spec.ForProvider.TenantSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Tenant")
	}
	mg.Spec.ForProvider.Tenant = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("identityplatform.gcp.upbound.io", "v1beta1", "Tenant", "TenantList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Tenant),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.InitProvider.TenantRef,
			Selector:     mg.Spec.InitProvider.TenantSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Tenant")
	}
	mg.Spec.InitProvider.Tenant = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TenantRef = rsp.ResolvedReference

	return nil
}
