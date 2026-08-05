package roundtrip

//nolint:typecheck // due to buildtagger constraints

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/apitesting/roundtrip"
	"github.com/hashicorp/terraform-provider-google/google/fwprovider"
	"github.com/hashicorp/terraform-provider-google/google/provider"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/randfill"

	clusterapis "github.com/upbound/provider-gcp/v2/apis/cluster"
	storageclusterv1beta3 "github.com/upbound/provider-gcp/v2/apis/cluster/storage/v1beta3"
	namespacedapis "github.com/upbound/provider-gcp/v2/apis/namespaced"
	storagensv1beta2 "github.com/upbound/provider-gcp/v2/apis/namespaced/storage/v1beta2"
	"github.com/upbound/provider-gcp/v2/config"
)

func TestRoundTrip(t *testing.T) {

	schema := provider.Provider()
	fwProvider := fwprovider.New(schema)
	ujProviderCluster, err := config.GetProvider(t.Context(), schema, fwProvider, false)
	if err != nil {
		t.Fatalf("GetProvider: %s", err)
	}

	ujProviderNamespaced, err := config.GetNamespacedProvider(t.Context(), schema, fwProvider, false)
	if err != nil {
		t.Fatalf("GetNamespacedProvider: %s", err)
	}

	testScheme := runtime.NewScheme()

	if err := clusterapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("cluster-scoped apis AddToScheme: %s", err)
	}

	if err := namespacedapis.AddToScheme(testScheme); err != nil {
		t.Fatalf("namespaced apis AddToScheme: %s", err)
	}

	rt, err := roundtrip.NewRoundTripTest(ujProviderCluster, ujProviderNamespaced, testScheme,
		roundtrip.WithFuzzerConfig(
			roundtrip.FuzzerIterations(10),
			roundtrip.FuzzerNilChance(0)),
		roundtrip.WithFuzzerConfig(
			roundtrip.FuzzerIterations(30),
			roundtrip.FuzzerNilChance(0.3)),
		roundtrip.WithComparisonOptions(
			roundtrip.EquateEmptyAndSingleZeroSlice(),
			roundtrip.EquateNilAndZeroValuePtr(),
		),
		roundtrip.WithExtraFuzzFuncs(gcpCustomFuzzers...),
	)
	if err != nil {
		t.Fatalf("NewRoundTripTest: %s", err)
	}

	t.Run("TestSerializationRoundtrip", func(t *testing.T) {
		rt.TestSerializationRoundtrip(t)
	})

	t.Run("TestConversionRoundtrip", func(t *testing.T) {
		rt.TestConversionRoundtrip(t)
	})
}

var gcpCustomFuzzers = []roundtrip.FuzzFunc{
	fuzzClusterStorageBucketV1Beta3,
	fuzzNamespacedStorageBucketV1Beta2,
}

// fuzzClusterStorageBucketV1Beta3 ensures retentionPolicy.retentionPeriod is a
// valid numeric string so that StringToFloat conversion to the v1beta2 spoke succeeds.
func fuzzClusterStorageBucketV1Beta3(s *storageclusterv1beta3.Bucket, c randfill.Continue) {
	c.Fill(s)
	if s.Spec.ForProvider.RetentionPolicy != nil && s.Spec.ForProvider.RetentionPolicy.RetentionPeriod != nil {
		s.Spec.ForProvider.RetentionPolicy.RetentionPeriod = ptr.To("3600")
	}
	if s.Spec.InitProvider.RetentionPolicy != nil && s.Spec.InitProvider.RetentionPolicy.RetentionPeriod != nil {
		s.Spec.InitProvider.RetentionPolicy.RetentionPeriod = ptr.To("3600")
	}
	if s.Status.AtProvider.RetentionPolicy != nil && s.Status.AtProvider.RetentionPolicy.RetentionPeriod != nil {
		s.Status.AtProvider.RetentionPolicy.RetentionPeriod = ptr.To("3600")
	}
}

// fuzzNamespacedStorageBucketV1Beta2 ensures retentionPolicy.retentionPeriod is a
// valid numeric string so that StringToFloat conversion to the v1beta1 spoke succeeds.
func fuzzNamespacedStorageBucketV1Beta2(s *storagensv1beta2.Bucket, c randfill.Continue) {
	c.Fill(s)
	if s.Spec.ForProvider.RetentionPolicy != nil && s.Spec.ForProvider.RetentionPolicy.RetentionPeriod != nil {
		s.Spec.ForProvider.RetentionPolicy.RetentionPeriod = ptr.To("3600")
	}
	if s.Spec.InitProvider.RetentionPolicy != nil && s.Spec.InitProvider.RetentionPolicy.RetentionPeriod != nil {
		s.Spec.InitProvider.RetentionPolicy.RetentionPeriod = ptr.To("3600")
	}
	if s.Status.AtProvider.RetentionPolicy != nil && s.Status.AtProvider.RetentionPolicy.RetentionPeriod != nil {
		s.Status.AtProvider.RetentionPolicy.RetentionPeriod = ptr.To("3600")
	}
}
