package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/name"

	"github.com/upbound/official-providers/provider-gcp/config/accessapproval"
	"github.com/upbound/official-providers/provider-gcp/config/bigtable"
	"github.com/upbound/official-providers/provider-gcp/config/cloudfunctions"
	"github.com/upbound/official-providers/provider-gcp/config/cloudiot"
	"github.com/upbound/official-providers/provider-gcp/config/cloudplatform"
	"github.com/upbound/official-providers/provider-gcp/config/compute"
	"github.com/upbound/official-providers/provider-gcp/config/container"
	"github.com/upbound/official-providers/provider-gcp/config/dataflow"
	"github.com/upbound/official-providers/provider-gcp/config/dataproc"
	"github.com/upbound/official-providers/provider-gcp/config/project"
	"github.com/upbound/official-providers/provider-gcp/config/sql"
	"github.com/upbound/official-providers/provider-gcp/config/storage"
)

const (
	resourcePrefix = "gcp"
	modulePath     = "github.com/upbound/official-providers/provider-gcp"
)

//go:embed schema.json
var providerSchema string

var skipList = []string{
	// Note(turkenh): Following two resources conflicts their singular versions
	// "google_access_context_manager_access_level" and
	// "google_access_context_manager_service_perimeter". Skipping for now.
	"google_access_context_manager_access_levels$",
	"google_access_context_manager_service_perimeters$",
}

var includeList = []string{
	// Storage
	"google_storage_bucket$",

	// Compute
	"google_compute_network$",
	"google_compute_subnetwork$",
	"google_compute_address$",
	"google_compute_firewall$",
	"google_compute_instance$",
	"google_compute_managed_ssl_certificate$",
	"google_compute_router$",
	"google_compute_router_nat$",
	"google_compute_disk$",
	"google_compute_external_vpn_gateway$",
	"google_compute_global_address$",
	"google_compute_global_network_endpoint_group$",
	"google_compute_ha_vpn_gateway$",
	"google_compute_health_check$",
	"google_compute_http_health_check$",
	"google_compute_https_health_check$",
	"google_compute_image$",
	"google_compute_instance_group$",
	"google_compute_instance_template$",
	"google_compute_instance_from_template",

	// Container
	"google_container_cluster",
	"google_container_node_pool",

	// Monitoring
	"google_monitoring_alert_policy",
	"google_monitoring_notification_channel",
	"google_monitoring_uptime_check_config",

	// Memory Store
	"google_redis_instance",

	// CloudPlatform
	"google_folder$",
	"google_project$",
	"google_service_account$",
	"google_service_account_key$",

	// Sql
	"google_sql_.+",
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, "",
		tjconfig.WithDefaultResourceOptions(
			groupOverrides(),
			externalNameConfig(),
			defaultVersion(),
			externalNameConfigurations(),
		),
		tjconfig.WithRootGroup("gcp.upbound.io"),
		tjconfig.WithShortName("gcp"),
		// Comment out the following line to generate all resources.
		tjconfig.WithIncludeList(includeList),
		tjconfig.WithSkipList(skipList))

	for _, configure := range []func(provider *tjconfig.Provider){
		accessapproval.Configure,
		bigtable.Configure,
		cloudfunctions.Configure,
		cloudiot.Configure,
		cloudplatform.Configure,
		container.Configure,
		compute.Configure,
		dataflow.Configure,
		dataproc.Configure,
		project.Configure,
		storage.Configure,
		sql.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

func init() {
	// GCP specific acronyms

	// Todo(turkenh): move to Terrajet?
	name.AddAcronym("idp", "IdP")
	name.AddAcronym("oauth", "OAuth")
}
