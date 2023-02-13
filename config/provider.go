// Copyright 2022 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/registry/reference"
	"github.com/upbound/upjet/pkg/types/name"

	"github.com/upbound/provider-gcp/config/accessapproval"
	"github.com/upbound/provider-gcp/config/beyondcorp"
	"github.com/upbound/provider-gcp/config/bigquery"
	"github.com/upbound/provider-gcp/config/bigtable"
	composer "github.com/upbound/provider-gcp/config/cloudcomposer"
	"github.com/upbound/provider-gcp/config/cloudfunctions"
	"github.com/upbound/provider-gcp/config/cloudiot"
	"github.com/upbound/provider-gcp/config/cloudplatform"
	"github.com/upbound/provider-gcp/config/cloudrun"
	"github.com/upbound/provider-gcp/config/cloudscheduler"
	"github.com/upbound/provider-gcp/config/cloudtasks"
	"github.com/upbound/provider-gcp/config/compute"
	"github.com/upbound/provider-gcp/config/container"
	"github.com/upbound/provider-gcp/config/containeraws"
	"github.com/upbound/provider-gcp/config/containerazure"
	"github.com/upbound/provider-gcp/config/dataflow"
	"github.com/upbound/provider-gcp/config/dataproc"
	"github.com/upbound/provider-gcp/config/dns"
	"github.com/upbound/provider-gcp/config/endpoints"
	"github.com/upbound/provider-gcp/config/firebaserules"
	"github.com/upbound/provider-gcp/config/gameservices"
	"github.com/upbound/provider-gcp/config/gkehub"
	"github.com/upbound/provider-gcp/config/healthcare"
	"github.com/upbound/provider-gcp/config/iap"
	"github.com/upbound/provider-gcp/config/identityplatform"
	"github.com/upbound/provider-gcp/config/kms"
	"github.com/upbound/provider-gcp/config/notebooks"
	"github.com/upbound/provider-gcp/config/oslogin"
	"github.com/upbound/provider-gcp/config/privateca"
	"github.com/upbound/provider-gcp/config/project"
	"github.com/upbound/provider-gcp/config/pubsub"
	"github.com/upbound/provider-gcp/config/redis"
	"github.com/upbound/provider-gcp/config/secretmanager"
	"github.com/upbound/provider-gcp/config/servicenetworking"
	"github.com/upbound/provider-gcp/config/sourcerepo"
	"github.com/upbound/provider-gcp/config/spanner"
	"github.com/upbound/provider-gcp/config/sql"
	"github.com/upbound/provider-gcp/config/storage"
	"github.com/upbound/provider-gcp/config/tpu"
	"github.com/upbound/provider-gcp/config/vertexai"
	"github.com/upbound/provider-gcp/config/vpcaccess"
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

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		tjconfig.WithDefaultResourceOptions(
			groupOverrides(),
			externalNameConfig(),
			defaultVersion(),
			externalNameConfigurations(),
			descriptionOverrides(),
		),
		tjconfig.WithRootGroup("gcp.upbound.io"),
		tjconfig.WithShortName("gcp"),
		// Comment out the following line to generate all resources.
		tjconfig.WithIncludeList(resourcesWithExternalNameConfig()),
		tjconfig.WithReferenceInjectors([]tjconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		tjconfig.WithSkipList(skipList))

	for _, configure := range []func(provider *tjconfig.Provider){
		accessapproval.Configure,
		bigtable.Configure,
		composer.Configure,
		cloudfunctions.Configure,
		cloudiot.Configure,
		cloudplatform.Configure,
		cloudrun.Configure,
		cloudscheduler.Configure,
		cloudtasks.Configure,
		containeraws.Configure,
		containerazure.Configure,
		compute.Configure,
		container.Configure,
		dataflow.Configure,
		dataproc.Configure,
		dns.Configure,
		endpoints.Configure,
		firebaserules.Configure,
		gameservices.Configure,
		iap.Configure,
		identityplatform.Configure,
		kms.Configure,
		notebooks.Configure,
		privateca.Configure,
		oslogin.Configure,
		project.Configure,
		pubsub.Configure,
		redis.Configure,
		secretmanager.Configure,
		servicenetworking.Configure,
		sourcerepo.Configure,
		spanner.Configure,
		storage.Configure,
		sql.Configure,
		redis.Configure,
		bigquery.Configure,
		beyondcorp.Configure,
		vertexai.Configure,
		tpu.Configure,
		vpcaccess.Configure,
		healthcare.Configure,
		gkehub.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// resourcesWithExternalNameConfig returns the list of resources that have external
// name configured in ExternalNameConfigs table.
func resourcesWithExternalNameConfig() []string {
	l := make([]string, len(externalNameConfigs))
	i := 0
	for n := range externalNameConfigs {
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
