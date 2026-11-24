// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package logging

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/googleapi"
)

func TestNotFoundAsAbsent(t *testing.T) {
	const id = "projects/my-project/locations/us-east1/buckets/regional"
	type want struct {
		err error
		id  string
	}
	cases := map[string]struct {
		reason  string
		origErr error
		want    want
	}{
		"NotFound": {
			reason:  "A 404 for a bucket that does not exist yet should clear the ID and return nil so the bucket is created",
			origErr: &googleapi.Error{Code: http.StatusNotFound, Message: "Bucket regional in location us-east1 does not exist"},
			want:    want{err: nil, id: ""},
		},
		"WrappedNotFound": {
			reason:  "A wrapped 404 should also be treated as absent",
			origErr: fmt.Errorf("reading bucket: %w", &googleapi.Error{Code: http.StatusNotFound}),
			want:    want{err: nil, id: ""},
		},
		"Forbidden": {
			reason:  "A 403 should be returned as-is without clearing the ID",
			origErr: &googleapi.Error{Code: http.StatusForbidden},
			want:    want{err: cmpopts.AnyError, id: id},
		},
		"OtherError": {
			reason:  "Errors that are not Google API errors should be returned as-is",
			origErr: errors.New("some other error"),
			want:    want{err: cmpopts.AnyError, id: id},
		},
		"NoError": {
			reason:  "A successful read should pass through without modification",
			origErr: nil,
			want:    want{err: nil, id: id},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			d := (&schema.Resource{Schema: map[string]*schema.Schema{}}).TestResourceData()
			d.SetId(id)

			read := notFoundAsAbsent(func(*schema.ResourceData, interface{}) error {
				return tc.origErr
			})
			err := read(d, nil)

			if diff := cmp.Diff(tc.want.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("%s\nnotFoundAsAbsent(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.id, d.Id()); diff != "" {
				t.Errorf("%s\nnotFoundAsAbsent(...): -want ID, +got ID:\n%s", tc.reason, diff)
			}
		})
	}
}
