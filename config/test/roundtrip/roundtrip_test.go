package roundtrip

//nolint:typecheck // due to buildtagger constraints

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/apitesting/roundtrip"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-provider-google/google/provider"
	"k8s.io/apimachinery/pkg/runtime"

	clusterapis "github.com/upbound/provider-gcp/v2/apis/cluster"
	containerv1beta1 "github.com/upbound/provider-gcp/v2/apis/cluster/container/v1beta1"
	namespacedapis "github.com/upbound/provider-gcp/v2/apis/namespaced"
	"github.com/upbound/provider-gcp/v2/config"
)

func TestRoundTrip(t *testing.T) {

	schema := provider.Provider()
	ujProviderCluster, err := config.GetProvider(t.Context(), schema, false)
	if err != nil {
		t.Fatalf("GetProvider: %s", err)
	}

	ujProviderNamespaced, err := config.GetNamespacedProvider(t.Context(), schema, false)
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
			// ignore manually-injected SSA index fields at comparison
			cmpopts.IgnoreFields(containerv1beta1.NodePoolNodeConfigInitParameters_2{}, "Index"),
			cmpopts.IgnoreFields(containerv1beta1.NodePoolNodeConfigParameters_2{}, "Index"),
			cmpopts.IgnoreFields(containerv1beta1.NodePoolNodeConfigObservation_2{}, "Index"),
		),
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
