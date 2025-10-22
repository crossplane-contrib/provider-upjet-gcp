// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/kingpin/v2"
	changelogsv1alpha1 "github.com/crossplane/crossplane-runtime/v2/apis/changelogs/proto/v1alpha1"
	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/gate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/customresourcesgate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/statemetrics"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/controller/conversion"
	"github.com/hashicorp/terraform-provider-google/google/provider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	authv1 "k8s.io/api/authorization/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	clusterapis "github.com/upbound/provider-gcp/apis/cluster"
	namespacedapis "github.com/upbound/provider-gcp/apis/namespaced"
	"github.com/upbound/provider-gcp/config"
	resolverapis "github.com/upbound/provider-gcp/internal/apis"
	"github.com/upbound/provider-gcp/internal/bootcheck"
	"github.com/upbound/provider-gcp/internal/clients"
	clustercontroller "github.com/upbound/provider-gcp/internal/controller/cluster"
	namespacedcontroller "github.com/upbound/provider-gcp/internal/controller/namespaced"
	"github.com/upbound/provider-gcp/internal/features"
	"github.com/upbound/provider-gcp/internal/version"
)

const (
	webhookTLSCertDirEnvVar = "WEBHOOK_TLS_CERT_DIR"
	tlsServerCertDirEnvVar  = "TLS_SERVER_CERTS_DIR"
	certsDirEnvVar          = "CERTS_DIR"
	tlsServerCertDir        = "/tls/server"
)

func deprecationAction(flagName string) kingpin.Action {
	return func(c *kingpin.ParseContext) error {
		_, err := fmt.Fprintf(os.Stderr, "warning: Command-line flag %q is deprecated and no longer used. It will be removed in a future release. Please remove it from all of your configurations (ControllerConfigs, etc.).\n", flagName)
		kingpin.FatalIfError(err, "Failed to print the deprecation notice.")
		return nil
	}
}

func init() {
	err := bootcheck.CheckEnv()
	if err != nil {
		log.Fatalf("bootcheck failed. provider will not be started: %v", err)
	}
}

func main() {
	var (
		app                     = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for GCP").DefaultEnvars()
		debug                   = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncInterval            = app.Flag("sync", "Sync interval controls how often all resources will be double checked for drift.").Short('s').Default("1h").Duration()
		pollInterval            = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		pollStateMetricInterval = app.Flag("poll-state-metric", "State metric recording interval").Default("5s").Duration()
		leaderElection          = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate        = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may checked for drift from the desired state.").Default("100").Int()
		webhookPort             = app.Flag("webhook-port", "The port the webhook listens on").Default("9443").Envar("WEBHOOK_PORT").Int()
		metricsBindAddress      = app.Flag("metrics-bind-address", "The address the metrics server listens on").Default(":8080").Envar("METRICS_BIND_ADDRESS").String()
		healthProbeBindAddress  = app.Flag("health-probe-bind-addr", "The address the health/readiness probe server listens on").Default(":8081").Envar("HEALTH_PROBE_BIND_ADDRESS").String()
		changelogsSocketPath    = app.Flag("changelogs-socket-path", "Path for changelogs socket (if enabled)").Default("/var/run/changelogs/changelogs.sock").Envar("CHANGELOGS_SOCKET_PATH").String()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()
		enableChangeLogs         = app.Flag("enable-changelogs", "Enable support for capturing change logs during reconciliation.").Default("false").Envar("ENABLE_CHANGE_LOGS").Bool()

		certsDirSet = false
		// we record whether the command-line option "--certs-dir" was supplied
		// in the registered PreAction for the flag.
		certsDir = app.Flag("certs-dir", "The directory that contains the server key and certificate.").Default(tlsServerCertDir).Envar(certsDirEnvVar).PreAction(func(_ *kingpin.ParseContext) error {
			certsDirSet = true
			return nil
		}).String()

		// now deprecated command-line arguments with the Terraform SDK-based upjet architecture
		_ = app.Flag("namespace", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Namespace used to set as default scope in default secret store config.").Default("crossplane-system").Envar("POD_NAMESPACE").Hidden().Action(deprecationAction("namespace")).String()
		_ = app.Flag("ess-tls-cert-dir", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Path of ESS TLS certificates.").Envar("ESS_TLS_CERTS_DIR").Hidden().Action(deprecationAction("ess-tls-cert-dir")).String()
		_ = app.Flag("enable-external-secret-stores", "[DEPRECATED: This option is no longer used and it will be removed in a future release.] Enable support for ExternalSecretStores.").Default("false").Envar("ENABLE_EXTERNAL_SECRET_STORES").Hidden().Action(deprecationAction("enable-external-secret-stores")).Bool()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))
	log.Default().SetOutput(io.Discard)
	ctrl.SetLogger(zap.New(zap.WriteTo(io.Discard)))

	zl := zap.New(zap.UseDevMode(*debug))
	logr := logging.NewLogrLogger(zl.WithName("provider-gcp"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	// currently, we configure the jitter to be the 5% of the poll interval
	pollJitter := time.Duration(float64(*pollInterval) * 0.05)
	logr.Debug("Starting", "sync-interval", syncInterval.String(),
		"poll-interval", pollInterval.String(), "poll-jitter", pollJitter, "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	// Get the TLS certs directory from the environment variables set by
	// Crossplane if they're available.
	// In older XP versions we used WEBHOOK_TLS_CERT_DIR, in newer versions
	// we use TLS_SERVER_CERTS_DIR. If an explicit certs dir is not supplied
	// via the command-line options, then these environment variables are used
	// instead.
	if !certsDirSet {
		// backwards-compatibility concerns
		xpCertsDir := os.Getenv(certsDirEnvVar)
		if xpCertsDir == "" {
			xpCertsDir = os.Getenv(tlsServerCertDirEnvVar)
		}
		if xpCertsDir == "" {
			xpCertsDir = os.Getenv(webhookTLSCertDirEnvVar)
		}
		// we probably don't need this condition but just to be on the
		// safe side, if we are missing any kingpin machinery details...
		if xpCertsDir != "" {
			*certsDir = xpCertsDir
		}
	}

	mgr, err := ctrl.NewManager(ratelimiter.LimitRESTConfig(cfg, *maxReconcileRate), ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-gcp-datastream",
		Cache: cache.Options{
			SyncPeriod: syncInterval,
		},
		Metrics: metricsserver.Options{
			BindAddress: *metricsBindAddress,
		},
		WebhookServer: webhook.NewServer(
			webhook.Options{
				CertDir: *certsDir,
				Port:    *webhookPort,
			}),
		HealthProbeBindAddress:     *healthProbeBindAddress,
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")
	if len(*certsDir) > 0 {
		kingpin.FatalIfError(mgr.AddReadyzCheck("webhook", mgr.GetWebhookServer().StartedChecker()), "Cannot add webhook server readyz checker to controller manager")
	}

	kingpin.FatalIfError(clusterapis.AddToScheme(mgr.GetScheme()), "Cannot add cluster-scoped GCP APIs to scheme")
	kingpin.FatalIfError(resolverapis.BuildScheme(clusterapis.AddToSchemes), "Cannot register the cluster-scoped Azure APIs with the API resolver's runtime scheme")
	kingpin.FatalIfError(namespacedapis.AddToScheme(mgr.GetScheme()), "Cannot add namespace-scoped GCP APIs to scheme")
	kingpin.FatalIfError(resolverapis.BuildScheme(namespacedapis.AddToSchemes), "Cannot register the namespace-scoped Azure APIs with the API resolver's runtime scheme")
	kingpin.FatalIfError(apiextensionsv1.AddToScheme(mgr.GetScheme()), "Cannot add api-extensions APIs to scheme")

	metricRecorder := managed.NewMRMetricRecorder()
	stateMetrics := statemetrics.NewMRStateMetrics()

	metrics.Registry.MustRegister(metricRecorder)
	metrics.Registry.MustRegister(stateMetrics)

	ctx := context.Background()
	sdkProvider := provider.Provider()
	clusterProvider, err := config.GetProvider(ctx, sdkProvider, false)
	kingpin.FatalIfError(err, "Cannot initialize the cluster provider configuration")
	namespacedProvider, err := config.GetNamespacedProvider(ctx, sdkProvider, false)
	kingpin.FatalIfError(err, "Cannot initialize the namespaced provider configuration")
	clusterOpts := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  logr,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                &feature.Flags{},
			MetricOptions: &xpcontroller.MetricOptions{
				PollStateMetricInterval: *pollStateMetricInterval,
				MRMetrics:               metricRecorder,
				MRStateMetrics:          stateMetrics,
			},
		},
		Provider:              clusterProvider,
		SetupFn:               clients.TerraformSetupBuilder(clusterProvider.TerraformProvider),
		PollJitter:            pollJitter,
		OperationTrackerStore: tjcontroller.NewOperationStore(logr),
		StartWebhooks:         *certsDir != "",
	}

	namespacedOpts := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  logr,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                &feature.Flags{},
			MetricOptions: &xpcontroller.MetricOptions{
				PollStateMetricInterval: *pollStateMetricInterval,
				MRMetrics:               metricRecorder,
				MRStateMetrics:          stateMetrics,
			},
		},
		Provider:              namespacedProvider,
		SetupFn:               clients.TerraformSetupBuilder(namespacedProvider.TerraformProvider),
		PollJitter:            pollJitter,
		OperationTrackerStore: tjcontroller.NewOperationStore(logr),
		StartWebhooks:         *certsDir != "",
	}

	if *enableManagementPolicies {
		clusterOpts.Features.Enable(features.EnableBetaManagementPolicies)
		namespacedOpts.Features.Enable(features.EnableBetaManagementPolicies)
		logr.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	if *enableChangeLogs {
		clusterOpts.Features.Enable(feature.EnableAlphaChangeLogs)
		namespacedOpts.Features.Enable(feature.EnableAlphaChangeLogs)
		logr.Info("Alpha feature enabled", "flag", feature.EnableAlphaChangeLogs)

		conn, err := grpc.NewClient("unix://"+*changelogsSocketPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
		kingpin.FatalIfError(err, "failed to create change logs client connection at %s", *changelogsSocketPath)

		clo := xpcontroller.ChangeLogOptions{
			ChangeLogger: managed.NewGRPCChangeLogger(
				changelogsv1alpha1.NewChangeLogServiceClient(conn),
				managed.WithProviderVersion(fmt.Sprintf("provider-upjet-aws:%s", version.Version))),
		}
		clusterOpts.ChangeLogOptions = &clo
		namespacedOpts.ChangeLogOptions = &clo
	}

	canSafeStart, err := canWatchCRD(ctx, mgr)
	kingpin.FatalIfError(err, "SafeStart precheck failed")
	if canSafeStart {
		crdGate := new(gate.Gate[schema.GroupVersionKind])
		clusterOpts.Gate = crdGate
		namespacedOpts.Gate = crdGate
		kingpin.FatalIfError(customresourcesgate.Setup(mgr, namespacedOpts.Options), "Cannot setup CRD gate")
		kingpin.FatalIfError(clustercontroller.SetupGated_datastream(mgr, clusterOpts), "Cannot setup cluster-scoped GCP controllers")
		kingpin.FatalIfError(namespacedcontroller.SetupGated_datastream(mgr, namespacedOpts), "Cannot setup namespaced GCP controllers")
	} else {
		logr.Info("Provider has missing RBAC permissions for watching CRDs, controller SafeStart capability will be disabled")
		kingpin.FatalIfError(clustercontroller.Setup_datastream(mgr, clusterOpts), "Cannot setup cluster-scoped GCP controllers")
		kingpin.FatalIfError(namespacedcontroller.Setup_datastream(mgr, namespacedOpts), "Cannot setup namespaced GCP controllers")
	}
	kingpin.FatalIfError(conversion.RegisterConversions(clusterOpts.Provider, namespacedOpts.Provider, mgr.GetScheme()), "Cannot initialize the webhook conversion registry")
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}

func canWatchCRD(ctx context.Context, mgr manager.Manager) (bool, error) {
	if err := authv1.AddToScheme(mgr.GetScheme()); err != nil {
		return false, err
	}
	verbs := []string{"get", "list", "watch"}
	for _, verb := range verbs {
		sar := &authv1.SelfSubjectAccessReview{
			Spec: authv1.SelfSubjectAccessReviewSpec{
				ResourceAttributes: &authv1.ResourceAttributes{
					Group:    "apiextensions.k8s.io",
					Resource: "customresourcedefinitions",
					Verb:     verb,
				},
			},
		}
		if err := mgr.GetClient().Create(ctx, sar); err != nil {
			return false, errors.Wrapf(err, "unable to perform RBAC check for verb %s on CustomResourceDefinitions", verbs)
		}
		if !sar.Status.Allowed {
			return false, nil
		}
	}
	return true, nil
}
