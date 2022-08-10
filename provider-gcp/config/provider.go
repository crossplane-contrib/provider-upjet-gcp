package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/registry/reference"
	"github.com/upbound/upjet/pkg/types/name"

	"github.com/upbound/official-providers/provider-gcp/config/accessapproval"
	"github.com/upbound/official-providers/provider-gcp/config/bigtable"
	"github.com/upbound/official-providers/provider-gcp/config/cloudfunctions"
	"github.com/upbound/official-providers/provider-gcp/config/cloudiot"
	"github.com/upbound/official-providers/provider-gcp/config/cloudplatform"
	"github.com/upbound/official-providers/provider-gcp/config/cloudscheduler"
	"github.com/upbound/official-providers/provider-gcp/config/cloudtasks"
	"github.com/upbound/official-providers/provider-gcp/config/compute"
	"github.com/upbound/official-providers/provider-gcp/config/container"
	"github.com/upbound/official-providers/provider-gcp/config/dataflow"
	"github.com/upbound/official-providers/provider-gcp/config/dataproc"
	"github.com/upbound/official-providers/provider-gcp/config/dns"
	"github.com/upbound/official-providers/provider-gcp/config/project"
	"github.com/upbound/official-providers/provider-gcp/config/pubsub"
	"github.com/upbound/official-providers/provider-gcp/config/redis"
	"github.com/upbound/official-providers/provider-gcp/config/secretmanager"
	"github.com/upbound/official-providers/provider-gcp/config/sql"
	"github.com/upbound/official-providers/provider-gcp/config/storage"
)

const (
	resourcePrefix = "gcp"
	modulePath     = "github.com/upbound/official-providers/provider-gcp"
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
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		tjconfig.WithDefaultResourceOptions(
			groupOverrides(),
			externalNameConfig(),
			defaultVersion(),
			externalNameConfigurations(),
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
		cloudfunctions.Configure,
		cloudiot.Configure,
		cloudplatform.Configure,
		cloudscheduler.Configure,
		cloudtasks.Configure,
		container.Configure,
		compute.Configure,
		dataflow.Configure,
		dataproc.Configure,
		dns.Configure,
		project.Configure,
		pubsub.Configure,
		redis.Configure,
		secretmanager.Configure,
		storage.Configure,
		sql.Configure,
		redis.Configure,
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
