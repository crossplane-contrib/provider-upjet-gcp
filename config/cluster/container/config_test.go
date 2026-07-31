// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package container

import (
	"encoding/base64"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/client-go/tools/clientcmd"
)

func TestClusterConnectionDetails(t *testing.T) {
	caPEM := []byte("-----BEGIN CERTIFICATE-----\nprivate-cluster-ca\n-----END CERTIFICATE-----")

	masterAuth := []interface{}{map[string]interface{}{
		"cluster_ca_certificate": base64.StdEncoding.EncodeToString(caPEM),
		"client_certificate":     "",
		"client_key":             "",
	}}
	controlPlaneEndpoints := func(dnsEndpoint string) []interface{} {
		return []interface{}{map[string]interface{}{
			"dns_endpoint_config": []interface{}{map[string]interface{}{
				"endpoint": dnsEndpoint,
			}},
		}}
	}

	type args struct {
		attr map[string]interface{}
	}
	type want struct {
		server string
		caData []byte
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"IPEndpointPinsClusterCA": {
			reason: "An IP endpoint is served by the cluster's own CA, so it must stay pinned.",
			args: args{attr: map[string]interface{}{
				"name":                           "example-cluster",
				"endpoint":                       "34.10.20.30",
				"master_auth":                    masterAuth,
				"control_plane_endpoints_config": controlPlaneEndpoints("uid.europe-west8.gke.goog"),
			}},
			want: want{server: "https://34.10.20.30", caData: caPEM},
		},
		"DNSEndpointOmitsClusterCA": {
			reason: "The DNS endpoint is fronted by Google Front End with a publicly trusted certificate, so pinning the private cluster CA would fail TLS validation.",
			args: args{attr: map[string]interface{}{
				"name":                           "example-cluster",
				"endpoint":                       "uid.europe-west8.gke.goog",
				"master_auth":                    masterAuth,
				"control_plane_endpoints_config": controlPlaneEndpoints("uid.europe-west8.gke.goog"),
			}},
			want: want{server: "https://uid.europe-west8.gke.goog", caData: nil},
		},
		"NoControlPlaneEndpointsConfigPinsClusterCA": {
			reason: "A cluster without controlPlaneEndpointsConfig must keep the pre-existing behaviour.",
			args: args{attr: map[string]interface{}{
				"name":        "example-cluster",
				"endpoint":    "34.10.20.30",
				"master_auth": masterAuth,
			}},
			want: want{server: "https://34.10.20.30", caData: caPEM},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := ClusterConnectionDetails(tc.args.attr)
			if err != nil {
				t.Fatalf("ClusterConnectionDetails(...): unexpected error: %v\nreason: %s", err, tc.reason)
			}
			kc, err := clientcmd.Load(got["kubeconfig"])
			if err != nil {
				t.Fatalf("cannot load generated kubeconfig: %v", err)
			}
			cluster := kc.Clusters[kc.CurrentContext]
			result := want{server: cluster.Server, caData: cluster.CertificateAuthorityData}
			if diff := cmp.Diff(tc.want, result, cmp.AllowUnexported(want{})); diff != "" {
				t.Errorf("ClusterConnectionDetails(...): -want, +got:\n%s\nreason: %s", diff, tc.reason)
			}
		})
	}
}
