// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/upbound/provider-gcp/config/namespaced"
	"github.com/upbound/provider-gcp/hack"
)

// GetNamespacedProvider returns the namespaced provider configuration
func GetNamespacedProvider(_ context.Context, sdkProvider *schema.Provider, generationProvider bool) (*ujconfig.Provider, error) {
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		ujconfig.WithDefaultResourceOptions(
			groupOverrides(),
			externalNameConfig(),
			defaultVersion(),
			resourceConfigurator(),
			descriptionOverrides(),
		),
		ujconfig.WithRootGroup("gcp.m.upbound.io"),
		ujconfig.WithShortName("gcp"),
		// Comment out the following line to generate all resources.
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithSkipList(skipList),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	registerTerraformConversions(pc)
	// add custom config functions
	for _, configure := range namespaced.ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// resourceList returns the list of resources that have external
// name configured in the specified table.

func registerTerraformConversions(pc *ujconfig.Provider) {
	for n, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}
		// the controller will be reconciling on the CRD API version
		// with the converted API (with embedded objects in place of
		// singleton lists), so we need the appropriate Terraform
		// converter in this case.
		r.TerraformConversions = []config.TerraformConversion{
			config.NewTFSingletonConversion(),
		}
		pc.Resources[n] = r
	}
}
