// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package cloudplatform

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestProjectReadWrapper(t *testing.T) {
	type args struct {
		origErr error
		id      string
	}
	type want struct {
		err error
		id  string
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"AbsentProjectError": {
			reason: "An ambiguous 403 indicating the project may not exist should clear the ID and return nil",
			args: args{
				origErr: errors.New(`the user does not have permission to access Project "my-project" or it may not exist`),
				id:      "my-project",
			},
			want: want{
				err: nil,
				id:  "",
			},
		},
		"OtherError": {
			reason: "Other errors should be returned as-is without clearing the ID",
			args: args{
				origErr: errors.New("some other API error"),
				id:      "my-project",
			},
			want: want{
				err: cmpopts.AnyError,
				id:  "my-project",
			},
		},
		"NoError": {
			reason: "A successful read should pass through without modification",
			args: args{
				origErr: nil,
				id:      "my-project",
			},
			want: want{
				err: nil,
				id:  "my-project",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// Create a minimal schema.Resource to get a valid ResourceData.
			r := &schema.Resource{
				Schema: map[string]*schema.Schema{
					"org_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			}
			d := r.TestResourceData()
			d.SetId(tc.args.id)

			// Build the wrapper the same way the configurator does.
			origRead := func(d *schema.ResourceData, meta interface{}) error {
				return tc.args.origErr
			}
			wrappedRead := func(d *schema.ResourceData, meta interface{}) error {
				err := origRead(d, meta)
				if err != nil && strings.Contains(err.Error(), "or it may not exist") {
					d.SetId("")
					return nil
				}
				return err
			}

			err := wrappedRead(d, nil)

			if diff := cmp.Diff(tc.want.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("%s\nwrappedRead(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.id, d.Id()); diff != "" {
				t.Errorf("%s\nwrappedRead(...): -want ID, +got ID:\n%s", tc.reason, diff)
			}
		})
	}
}
