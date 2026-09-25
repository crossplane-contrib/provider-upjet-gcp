// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package compute

import (
	"testing"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/googleapi"
)

const (
	assocMsgSubstring = "An association with that name does not exist"
	assocID           = "projects/my-project/global/firewallPolicies/1234567890123456789/associations/my-network"
)

// wrapAsRead wraps err with the message the Terraform provider's
// HandleNotFoundError prepends to non-404 errors before they reach the
// resource configurators.
func wrapAsRead(resource string, err error) error {
	return errors.Wrap(err, "Error when reading or editing "+resource)
}

func TestBadRequestAsNotFound(t *testing.T) {
	type args struct {
		msgSubstring string
		origErr      error
		id           string
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
		"RuleNotFound": {
			reason: "A 400 error indicating the rule does not exist should clear the ID and return nil",
			args: args{
				msgSubstring: "does not contain a rule at priority",
				origErr: wrapAsRead(`ComputeNetworkFirewallPolicyRule "projects/my-project/global/firewallPolicies/my-policy/rules/328"`, &googleapi.Error{
					Code:    400,
					Message: "Invalid value for field 'priority': '328'. The firewall policy does not contain a rule at priority 328.",
					Errors:  []googleapi.ErrorItem{{Reason: "invalid", Message: "Invalid value for field 'priority': '328'. The firewall policy does not contain a rule at priority 328."}},
				}),
				id: "projects/my-project/global/firewallPolicies/my-policy/rules/328",
			},
			want: want{
				err: nil,
				id:  "",
			},
		},
		"AssociationNotFound": {
			reason: "A 400 error indicating the association does not exist should clear the ID and return nil",
			args: args{
				msgSubstring: assocMsgSubstring,
				origErr: wrapAsRead(`ComputeNetworkFirewallPolicyAssociation "projects/my-project/global/firewallPolicies/1234567890123456789/associations/my-network"`, &googleapi.Error{
					Code:    400,
					Message: "Invalid value for field 'name': 'my-network'. An association with that name does not exist.",
					Errors:  []googleapi.ErrorItem{{Reason: "invalid", Message: "Invalid value for field 'name': 'my-network'. An association with that name does not exist."}},
				}),
				id: assocID,
			},
			want: want{
				err: nil,
				id:  "",
			},
		},
		"BadRequestOtherMessage": {
			reason: "A 400 error with an unrelated message should be returned as-is without clearing the ID",
			args: args{
				msgSubstring: assocMsgSubstring,
				origErr: wrapAsRead(`ComputeNetworkFirewallPolicyAssociation "projects/my-project/global/firewallPolicies/1234567890123456789/associations/my-network"`, &googleapi.Error{
					Code:    400,
					Message: "Invalid value for field 'firewallPolicy': 'not-a-policy'.",
					Errors:  []googleapi.ErrorItem{{Reason: "invalid", Message: "Invalid value for field 'firewallPolicy': 'not-a-policy'."}},
				}),
				id: assocID,
			},
			want: want{
				err: cmpopts.AnyError,
				id:  assocID,
			},
		},
		"MatchingMessageWrongCode": {
			reason: "A non-400 error should be returned as-is even when the message matches",
			args: args{
				msgSubstring: assocMsgSubstring,
				origErr: wrapAsRead(`ComputeNetworkFirewallPolicyAssociation "projects/my-project/global/firewallPolicies/1234567890123456789/associations/my-network"`, &googleapi.Error{
					Code:    409,
					Message: "An association with that name does not exist.",
				}),
				id: assocID,
			},
			want: want{
				err: cmpopts.AnyError,
				id:  assocID,
			},
		},
		"RenderedStringError": {
			reason: "A plain string error is not a googleapi.Error and should be returned as-is even when the text matches",
			args: args{
				msgSubstring: assocMsgSubstring,
				origErr:      errors.New(`Error when reading or editing ComputeNetworkFirewallPolicyAssociation "projects/my-project/global/firewallPolicies/1234567890123456789/associations/my-network": googleapi: Error 400: Invalid value for field 'name': 'my-network'. An association with that name does not exist., invalid`),
				id:           assocID,
			},
			want: want{
				err: cmpopts.AnyError,
				id:  assocID,
			},
		},
		"OtherError": {
			reason: "Other errors should be returned as-is without clearing the ID",
			args: args{
				msgSubstring: assocMsgSubstring,
				origErr:      errors.New("some other API error"),
				id:           assocID,
			},
			want: want{
				err: cmpopts.AnyError,
				id:  assocID,
			},
		},
		"NoError": {
			reason: "A successful read should pass through without modification",
			args: args{
				msgSubstring: assocMsgSubstring,
				origErr:      nil,
				id:           assocID,
			},
			want: want{
				err: nil,
				id:  assocID,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := &schema.Resource{Schema: map[string]*schema.Schema{}}
			d := r.TestResourceData()
			d.SetId(tc.args.id)

			read := badRequestAsNotFound(func(d *schema.ResourceData, meta interface{}) error {
				return tc.args.origErr
			}, tc.args.msgSubstring)

			err := read(d, nil)

			if diff := cmp.Diff(tc.want.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("%s\nread(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.id, d.Id()); diff != "" {
				t.Errorf("%s\nread(...): -want ID, +got ID:\n%s", tc.reason, diff)
			}
		})
	}
}
