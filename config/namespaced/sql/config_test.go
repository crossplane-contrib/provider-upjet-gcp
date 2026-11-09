// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package sql

import "testing"

func TestUserNameFromID(t *testing.T) {
	cases := map[string]struct {
		id      string
		want    string
		wantErr bool
	}{
		"BuiltInUserEmptyHost": {
			id:   "postgres//my-instance",
			want: "postgres",
		},
		"MySQLUserWithHost": {
			id:   "admin/%/my-instance",
			want: "admin",
		},
		"ServiceAccountUser": {
			id:   "sa@my-project.iam//my-instance",
			want: "sa@my-project.iam",
		},
		"WorkloadPrincipalWithSlashes": {
			id:   "ns/my-namespace/sa/my-ksa//my-instance",
			want: "ns/my-namespace/sa/my-ksa",
		},
		"Malformed": {
			id:      "only-a-name",
			wantErr: true,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := userNameFromID(tc.id)
			if (err != nil) != tc.wantErr {
				t.Fatalf("userNameFromID(%q) error = %v, wantErr %v", tc.id, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("userNameFromID(%q) = %q, want %q", tc.id, got, tc.want)
			}
		})
	}
}
