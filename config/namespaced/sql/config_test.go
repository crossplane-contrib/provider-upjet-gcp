// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package sql

import "testing"

func TestUserNameFromID(t *testing.T) {
	cases := map[string]struct {
		id, host, instance string
		want               string
		wantErr            bool
	}{
		"BuiltInUserEmptyHost": {
			id: "postgres//my-instance", instance: "my-instance",
			want: "postgres",
		},
		"MySQLUserWithHost": {
			id: "admin/%/my-instance", host: "%", instance: "my-instance",
			want: "admin",
		},
		// The host can contain the separator too; the external name of such
		// an existing user must stay "admin".
		"MySQLUserWithCIDRHost": {
			id: "admin/10.0.0.0/24/my-instance", host: "10.0.0.0/24", instance: "my-instance",
			want: "admin",
		},
		"ServiceAccountUser": {
			id: "sa@my-project.iam//my-instance", instance: "my-instance",
			want: "sa@my-project.iam",
		},
		"WorkloadPrincipalWithSlashes": {
			id: "ns/my-namespace/sa/my-ksa//my-instance", instance: "my-instance",
			want: "ns/my-namespace/sa/my-ksa",
		},
		"WorkloadPrincipalWithCIDRHost": {
			id: "ns/my-namespace/sa/my-ksa/10.0.0.0/24/my-instance", host: "10.0.0.0/24", instance: "my-instance",
			want: "ns/my-namespace/sa/my-ksa",
		},
		// Without host and instance in the state, keep the previous result.
		"NoInstanceInState": {
			id:   "admin/10.0.0.0/24/my-instance",
			want: "admin",
		},
		"Malformed": {
			id:      "/my-instance",
			wantErr: true,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := userNameFromID(tc.id, tc.host, tc.instance)
			if (err != nil) != tc.wantErr {
				t.Fatalf("userNameFromID(%q) error = %v, wantErr %v", tc.id, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("userNameFromID(%q) = %q, want %q", tc.id, got, tc.want)
			}
		})
	}
}
