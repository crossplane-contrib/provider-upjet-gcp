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

	// ResolveReferences of this Backup.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Backup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.ClusterNameRef,
			Selector:     mg.Spec.ForProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterName),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.InitProvider.ClusterNameRef,
			Selector:     mg.Spec.InitProvider.ClusterNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterName")
	}
	mg.Spec.InitProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkConfig[i3].Network),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.NetworkConfig[i3].NetworkRef,
				Selector:     mg.Spec.ForProvider.NetworkConfig[i3].NetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfig[i3].Network")
		}
		mg.Spec.ForProvider.NetworkConfig[i3].Network = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkConfig[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RestoreBackupSource); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Backup", "BackupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RestoreBackupSource[i3].BackupName),
				Extract:      resource.ExtractParamPath("name", true),
				Reference:    mg.Spec.ForProvider.RestoreBackupSource[i3].BackupNameRef,
				Selector:     mg.Spec.ForProvider.RestoreBackupSource[i3].BackupNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RestoreBackupSource[i3].BackupName")
		}
		mg.Spec.ForProvider.RestoreBackupSource[i3].BackupName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RestoreBackupSource[i3].BackupNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RestoreContinuousBackupSource); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Cluster", "ClusterList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RestoreContinuousBackupSource[i3].Cluster),
				Extract:      resource.ExtractParamPath("name", true),
				Reference:    mg.Spec.ForProvider.RestoreContinuousBackupSource[i3].ClusterRef,
				Selector:     mg.Spec.ForProvider.RestoreContinuousBackupSource[i3].ClusterSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RestoreContinuousBackupSource[i3].Cluster")
		}
		mg.Spec.ForProvider.RestoreContinuousBackupSource[i3].Cluster = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RestoreContinuousBackupSource[i3].ClusterRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SecondaryConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Cluster", "ClusterList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecondaryConfig[i3].PrimaryClusterName),
				Extract:      resource.ExtractParamPath("name", true),
				Reference:    mg.Spec.ForProvider.SecondaryConfig[i3].PrimaryClusterNameRef,
				Selector:     mg.Spec.ForProvider.SecondaryConfig[i3].PrimaryClusterNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SecondaryConfig[i3].PrimaryClusterName")
		}
		mg.Spec.ForProvider.SecondaryConfig[i3].PrimaryClusterName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SecondaryConfig[i3].PrimaryClusterNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.NetworkConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkConfig[i3].Network),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.NetworkConfig[i3].NetworkRef,
				Selector:     mg.Spec.InitProvider.NetworkConfig[i3].NetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.NetworkConfig[i3].Network")
		}
		mg.Spec.InitProvider.NetworkConfig[i3].Network = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.NetworkConfig[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.RestoreBackupSource); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Backup", "BackupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RestoreBackupSource[i3].BackupName),
				Extract:      resource.ExtractParamPath("name", true),
				Reference:    mg.Spec.InitProvider.RestoreBackupSource[i3].BackupNameRef,
				Selector:     mg.Spec.InitProvider.RestoreBackupSource[i3].BackupNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.RestoreBackupSource[i3].BackupName")
		}
		mg.Spec.InitProvider.RestoreBackupSource[i3].BackupName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.RestoreBackupSource[i3].BackupNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.RestoreContinuousBackupSource); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Cluster", "ClusterList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RestoreContinuousBackupSource[i3].Cluster),
				Extract:      resource.ExtractParamPath("name", true),
				Reference:    mg.Spec.InitProvider.RestoreContinuousBackupSource[i3].ClusterRef,
				Selector:     mg.Spec.InitProvider.RestoreContinuousBackupSource[i3].ClusterSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.RestoreContinuousBackupSource[i3].Cluster")
		}
		mg.Spec.InitProvider.RestoreContinuousBackupSource[i3].Cluster = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.RestoreContinuousBackupSource[i3].ClusterRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.SecondaryConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Cluster", "ClusterList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecondaryConfig[i3].PrimaryClusterName),
				Extract:      resource.ExtractParamPath("name", true),
				Reference:    mg.Spec.InitProvider.SecondaryConfig[i3].PrimaryClusterNameRef,
				Selector:     mg.Spec.InitProvider.SecondaryConfig[i3].PrimaryClusterNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SecondaryConfig[i3].PrimaryClusterName")
		}
		mg.Spec.InitProvider.SecondaryConfig[i3].PrimaryClusterName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SecondaryConfig[i3].PrimaryClusterNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Cluster),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.ClusterRef,
			Selector:     mg.Spec.ForProvider.ClusterSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Cluster")
	}
	mg.Spec.ForProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceType),
			Extract:      resource.ExtractParamPath("cluster_type", false),
			Reference:    mg.Spec.ForProvider.InstanceTypeRef,
			Selector:     mg.Spec.ForProvider.InstanceTypeSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceType")
	}
	mg.Spec.ForProvider.InstanceType = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceTypeRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("alloydb.gcp.upbound.io", "v1beta1", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceType),
			Extract:      resource.ExtractParamPath("cluster_type", false),
			Reference:    mg.Spec.InitProvider.InstanceTypeRef,
			Selector:     mg.Spec.InitProvider.InstanceTypeSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceType")
	}
	mg.Spec.InitProvider.InstanceType = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceTypeRef = rsp.ResolvedReference

	return nil
}
