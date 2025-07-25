// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/config/conversion"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/crossplane/upjet/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	"github.com/crossplane/upjet/pkg/types/name"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/upbound/provider-gcp/hack"
)

const (
	resourcePrefix = "gcp"
	modulePath     = "github.com/upbound/provider-gcp"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte

	// oldSingletonListAPIs is a newline-delimited list of Terraform resource
	// names with converted singleton list APIs with at least CRD API version
	// containing the old singleton list API. This is to prevent the API
	// conversion for the newly added resources whose CRD APIs will already
	// use embedded objects instead of the singleton lists and thus, will
	// not possess a CRD API version with the singleton list. Thus, for
	// the newly added resources (resources added after the singleton lists
	// have been converted), we do not need the CRD API conversion
	// functions that convert between singleton lists and embedded objects,
	// but we need only the Terraform conversion functions.
	// This list is immutable and represents the set of resources with the
	// already generated CRD API versions with now converted singleton lists.
	// Because new resources should never have singleton lists in their
	// generated APIs, there should be no need to add them to this list.
	// However, bugs might result in exceptions in the future.
	// Please see:
	// https://github.com/crossplane-contrib/provider-upjet-gcp/pull/508
	// for more context on singleton list to embedded object conversions.
	//go:embed old-singleton-list-apis.txt
	oldSingletonListAPIs string
)

var skipList = []string{
	// Note(turkenh): Following two resources conflicts their singular versions
	// "google_access_context_manager_access_level" and
	// "google_access_context_manager_service_perimeter". Skipping for now.
	"google_access_context_manager_access_levels$",
	"google_access_context_manager_service_perimeters$",
	// Note(piotr): Following resources are potentially dangerous to implement
	// details in: https://github.com/upbound/official-providers/issues/587
	"google_kms_crypto_key_iam_policy",
	"google_kms_crypto_key_iam_binding",
	"google_kms_key_ring_iam_policy",
	"google_kms_key_ring_iam_binding",
	"google_cloudfunctions_function_iam_policy",
	"google_cloudfunctions_function_iam_binding",
	"google_compute_region_disk_iam_policy",
	"google_compute_region_disk_iam_binding",
	// Note(donovamuller): Following resources are potentially dangerous to implement
	// details in: https://github.com/upbound/official-providers/issues/521
	"google_project_iam_policy",
	"google_project_iam_binding",
	"google_organization_iam_binding",
	"google_service_account_iam_policy",
	"google_service_account_iam_binding",
	"google_cloud_run_service_iam_policy",
	"google_cloud_run_service_iam_binding",
	"google_pubsub_topic_iam_policy",
	"google_pubsub_topic_iam_binding",
	"google_pubsub_subscription_iam_policy",
	"google_pubsub_subscription_iam_binding",
	"google_compute_disk_iam_policy",
	"google_compute_disk_iam_binding",
	"google_compute_instance_iam_policy",
	"google_compute_instance_iam_binding",
	"google_compute_image_iam_policy",
	"google_compute_image_iam_binding",
	"google_notebooks_instance_iam_policy",
	"google_notebooks_instance_iam_binding",
	"google_notebooks_runtime_iam_policy",
	"google_notebooks_runtime_iam_binding",
	"google_secret_manager_secret_iam_policy",
	"google_secret_manager_secret_iam_binding",
	"google_sourcerepo_repository_iam_policy",
	"google_sourcerepo_repository_iam_binding",
	"google_spanner_instance_iam_policy",
	"google_spanner_instance_iam_binding",
	"google_spanner_database_iam_policy",
	"google_spanner_database_iam_binding",
	"google_compute_subnetwork_iam_policy",
	"google_compute_subnetwork_iam_binding",
	"google_endpoints_service_iam_policy",
	"google_endpoints_service_iam_binding",
	"google_endpoints_service_consumers_iam_policy",
	"google_endpoints_service_consumers_iam_binding",
}

// workaround for the TF Google v4.77.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, sdkProvider *schema.Provider, generationProvider bool) (*ujconfig.Provider, error) {
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
		ujconfig.WithRootGroup("gcp.upbound.io"),
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

	bumpVersionsWithEmbeddedLists(pc)
	// add custom config functions
	for _, configure := range ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// resourceList returns the list of resources that have external
// name configured in the specified table.
func resourceList(t map[string]ujconfig.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}

func bumpVersionsWithEmbeddedLists(pc *ujconfig.Provider) {
	l := strings.Split(strings.TrimSpace(oldSingletonListAPIs), "\n")
	oldSLAPIs := make(map[string]struct{}, len(l))
	for _, n := range l {
		oldSLAPIs[n] = struct{}{}
	}

	for n, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}

		if _, ok := oldSLAPIs[n]; ok {
			r.Version = "v1beta2"
			r.PreviousVersions = []string{"v1beta1"}
			// we would like to set the storage version to v1beta1 to facilitate
			// downgrades.
			r.SetCRDStorageVersion("v1beta1")
			// because the controller reconciles on the API version with the singleton list API,
			// no need for a Terraform conversion.
			r.ControllerReconcileVersion = "v1beta1"
			r.Conversions = []conversion.Conversion{
				conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, conversion.DefaultPathPrefixes(), r.CRDListConversionPaths()...),
				conversion.NewSingletonListConversion("v1beta1", "v1beta2", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToEmbeddedObject),
				conversion.NewSingletonListConversion("v1beta2", "v1beta1", conversion.DefaultPathPrefixes(), r.CRDListConversionPaths(), conversion.ToSingletonList)}
		} else {
			// the controller will be reconciling on the CRD API version
			// with the converted API (with embedded objects in place of
			// singleton lists), so we need the appropriate Terraform
			// converter in this case.
			r.TerraformConversions = []config.TerraformConversion{
				config.NewTFSingletonConversion(),
			}
		}
		pc.Resources[n] = r
	}
}

func init() {
	// GCP specific acronyms

	// Todo(turkenh): move to Terrajet?
	name.AddAcronym("idp", "IdP")
	name.AddAcronym("oauth", "OAuth")
}

// Configure configures the specified Provider.
type Configure func(provider *config.Provider)

// Configurator is a registry for provider Configs.
type Configurator []Configure

// AddConfig adds a Config to the Configurator registry.
func (c *Configurator) AddConfig(conf Configure) {
	*c = append(*c, conf)
}

// ProviderConfiguration is a global registry to be used by
// the resource providers to register their Config functions.
var ProviderConfiguration = Configurator{}
