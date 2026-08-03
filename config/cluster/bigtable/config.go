// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package bigtable

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_bigtable_gc_policy", func(r *config.Resource) {
		r.Kind = "GarbageCollectionPolicy"
	})

	p.AddResourceConfigurator("google_bigtable_table_iam_binding", func(r *config.Resource) {
		r.Version = "v1beta3"
		r.PreviousVersions = []string{"v1beta2"}
		r.SetCRDStorageVersion("v1beta2")
		r.ControllerReconcileVersion = "v1beta2" //nolint:staticcheck // still handling the deprecated behavior
		r.Conversions = append(r.Conversions,
			conversion.NewFieldRenameConversion("v1beta2", "spec.forProvider.instance", "v1beta3", "spec.forProvider.instanceName"),
			conversion.NewFieldRenameConversion("v1beta2", "spec.initProvider.instance", "v1beta3", "spec.initProvider.instanceName"),
			conversion.NewFieldRenameConversion("v1beta2", "status.atProvider.instance", "v1beta3", "status.atProvider.instanceName"),
			conversion.NewFieldRenameConversion("v1beta3", "spec.forProvider.instanceName", "v1beta2", "spec.forProvider.instance"),
			conversion.NewFieldRenameConversion("v1beta3", "spec.initProvider.instanceName", "v1beta2", "spec.initProvider.instance"),
			conversion.NewFieldRenameConversion("v1beta3", "status.atProvider.instanceName", "v1beta2", "status.atProvider.instance"),
		)
		r.TerraformConversions = append(r.TerraformConversions, newInstanceNameConverter())
	})

	p.AddResourceConfigurator("google_bigtable_table_iam_member", func(r *config.Resource) {
		r.Version = "v1beta3"
		r.PreviousVersions = []string{"v1beta2"}
		r.SetCRDStorageVersion("v1beta2")
		r.ControllerReconcileVersion = "v1beta2" //nolint:staticcheck // still handling the deprecated behavior
		r.Conversions = append(r.Conversions,
			conversion.NewFieldRenameConversion("v1beta2", "spec.forProvider.instance", "v1beta3", "spec.forProvider.instanceName"),
			conversion.NewFieldRenameConversion("v1beta2", "spec.initProvider.instance", "v1beta3", "spec.initProvider.instanceName"),
			conversion.NewFieldRenameConversion("v1beta2", "status.atProvider.instance", "v1beta3", "status.atProvider.instanceName"),
			conversion.NewFieldRenameConversion("v1beta3", "spec.forProvider.instanceName", "v1beta2", "spec.forProvider.instance"),
			conversion.NewFieldRenameConversion("v1beta3", "spec.initProvider.instanceName", "v1beta2", "spec.initProvider.instance"),
			conversion.NewFieldRenameConversion("v1beta3", "status.atProvider.instanceName", "v1beta2", "status.atProvider.instance"),
		)
		r.TerraformConversions = append(r.TerraformConversions, newInstanceNameConverter())
	})

	p.AddResourceConfigurator("google_bigtable_table_iam_policy", func(r *config.Resource) {
		r.Version = "v1beta2"
		r.PreviousVersions = []string{"v1beta1"}
		r.SetCRDStorageVersion("v1beta1")
		r.ControllerReconcileVersion = "v1beta1" //nolint:staticcheck // still handling the deprecated behavior
		r.Conversions = append(r.Conversions,
			conversion.NewFieldRenameConversion("v1beta1", "spec.forProvider.instance", "v1beta2", "spec.forProvider.instanceName"),
			conversion.NewFieldRenameConversion("v1beta1", "spec.initProvider.instance", "v1beta2", "spec.initProvider.instanceName"),
			conversion.NewFieldRenameConversion("v1beta1", "status.atProvider.instance", "v1beta2", "status.atProvider.instanceName"),
			conversion.NewFieldRenameConversion("v1beta2", "spec.forProvider.instanceName", "v1beta1", "spec.forProvider.instance"),
			conversion.NewFieldRenameConversion("v1beta2", "spec.initProvider.instanceName", "v1beta1", "spec.initProvider.instance"),
			conversion.NewFieldRenameConversion("v1beta2", "status.atProvider.instanceName", "v1beta1", "status.atProvider.instance"),
		)
		r.TerraformConversions = append(r.TerraformConversions, newInstanceNameConverter())
	})
}

type instanceNameConverter struct{}

func newInstanceNameConverter() config.TerraformConversion {
	return instanceNameConverter{}
}

func (s instanceNameConverter) Convert(params map[string]any, r *config.Resource, mode config.Mode) (map[string]any, error) { //nolint:gocyclo // easier to follow as a unit
	pv := fieldpath.Pave(params)
	switch mode {
	case config.ToTerraform:
		v, err := pv.GetValue("instance")
		if err != nil && !fieldpath.IsNotFound(err) {
			return nil, errors.Wrapf(err, "cannot get instance value")
		} else if err == nil {
			if err := pv.SetValue("instance_name", v); err != nil {
				return nil, errors.Wrapf(err, "cannot set instance_name")
			}
			if err := pv.DeleteField("instance"); err != nil {
				return nil, errors.Wrapf(err, "cannot delete instance")
			}
		}
		return params, nil
	case config.FromTerraform:
		v, err := pv.GetValue("instance_name")
		if err != nil && !fieldpath.IsNotFound(err) {
			return nil, errors.Wrapf(err, "cannot get instance_name value")
		} else if err == nil {
			if err := pv.SetValue("instance", v); err != nil {
				return nil, errors.Wrapf(err, "cannot set instance")
			}
			if err := pv.DeleteField("instance_name"); err != nil {
				return nil, errors.Wrapf(err, "cannot delete instance_name")
			}
		}
		return params, nil
	default:
		return nil, errors.Errorf("unsupported mode %s", mode)
	}
}
