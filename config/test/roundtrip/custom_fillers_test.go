// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

//nolint:typecheck // due to buildtagger constraints

package roundtrip

import (
	"sigs.k8s.io/randfill"

	"github.com/crossplane/upjet/v2/pkg/apitesting/roundtrip"
	"k8s.io/utils/ptr"

	storageclusterv1beta3 "github.com/upbound/provider-gcp/v2/apis/cluster/storage/v1beta3"
	storagensv1beta2 "github.com/upbound/provider-gcp/v2/apis/namespaced/storage/v1beta2"
)

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
