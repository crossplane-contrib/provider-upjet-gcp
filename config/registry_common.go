// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	"github.com/crossplane/upjet/pkg/types/name"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
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

func init() {
	// GCP specific acronyms

	// Todo(turkenh): move to Terrajet?
	name.AddAcronym("idp", "IdP")
	name.AddAcronym("oauth", "OAuth")
}
