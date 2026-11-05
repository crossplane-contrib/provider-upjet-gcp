// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package compute

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfcompute "github.com/hashicorp/terraform-provider-google/google/services/compute"
)

func TestNetworkFirewallPolicyRuleReadWrapper(t *testing.T) {
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
		"NotFoundError": {
			reason: "A 400 error indicating the rule does not exist should clear the ID and return nil",
			args: args{
				origErr: errors.New(`Error when reading or editing ComputeNetworkFirewallPolicyRule "projects/my-project/global/firewallPolicies/my-policy/rules/328": googleapi: Error 400: Invalid value for field 'priority': '328'. The firewall policy does not contain a rule at priority 328., invalid`),
				id:      "projects/my-project/global/firewallPolicies/my-policy/rules/328",
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
				id:      "projects/my-project/global/firewallPolicies/my-policy/rules/328",
			},
			want: want{
				err: cmpopts.AnyError,
				id:  "projects/my-project/global/firewallPolicies/my-policy/rules/328",
			},
		},
		"NoError": {
			reason: "A successful read should pass through without modification",
			args: args{
				origErr: nil,
				id:      "projects/my-project/global/firewallPolicies/my-policy/rules/328",
			},
			want: want{
				err: nil,
				id:  "projects/my-project/global/firewallPolicies/my-policy/rules/328",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// Create a minimal schema.Resource to get a valid ResourceData.
			r := &schema.Resource{
				Schema: map[string]*schema.Schema{
					"priority": {
						Type:     schema.TypeInt,
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
				if err != nil && strings.Contains(err.Error(), "does not contain a rule at priority") {
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

func TestReservationShareDiff(t *testing.T) {
	sm := schema.InternalMap(tfcompute.ResourceComputeReservation().Schema)

	config := func(shareSettings any) *terraform.ResourceConfig {
		raw := map[string]any{
			"name":    "reservation",
			"project": "my-project",
			"zone":    "us-central1-a",
			"specific_reservation": []any{map[string]any{
				"count":               1,
				"instance_properties": []any{map[string]any{"machine_type": "n2-standard-2"}},
			}},
		}
		if shareSettings != nil {
			raw["share_settings"] = shareSettings
		}
		return terraform.NewResourceConfigRaw(raw)
	}
	state := func(shareSettings any) *terraform.InstanceState {
		data := schema.TestResourceDataRaw(t, sm, config(shareSettings).Raw)
		data.SetId("projects/my-project/zones/us-central1-a/reservations/reservation")
		return data.State()
	}
	projects := func(ids ...string) []any {
		m := make([]any, 0, len(ids))
		for _, id := range ids {
			m = append(m, map[string]any{"id": id, "project_id": id})
		}
		return m
	}
	sharedState := state([]any{map[string]any{"share_type": "SPECIFIC_PROJECTS", "project_map": projects("111111111111")}})

	type want struct {
		shareType   string
		projects    []string
		countDiff   bool
		requiresNew bool
	}
	cases := map[string]struct {
		reason string
		state  *terraform.InstanceState
		config *terraform.ResourceConfig
		want   want
	}{
		"CreateSpecificProjects": {
			reason: "Creating with SPECIFIC_PROJECTS must keep the count so Create sends shareSettings.",
			config: config([]any{map[string]any{"share_type": "SPECIFIC_PROJECTS", "project_map": projects("111111111111")}}),
			want:   want{shareType: "SPECIFIC_PROJECTS", projects: []string{"111111111111"}, countDiff: true, requiresNew: true},
		},
		"CreateLocal": {
			reason: "Creating with an explicit LOCAL share type must keep the count.",
			config: config([]any{map[string]any{"share_type": "LOCAL"}}),
			want:   want{shareType: "LOCAL", countDiff: true, requiresNew: true},
		},
		"UnsetInConfigAndState": {
			reason: "The phantom computed count diff for a block set nowhere must be dropped.",
			state:  state(nil),
			config: config(nil),
			want:   want{},
		},
		"AddSharedProject": {
			reason: "Adding a project to an already shared reservation must stay an in-place update.",
			state:  sharedState,
			config: config([]any{map[string]any{"share_type": "SPECIFIC_PROJECTS", "project_map": projects("111111111111", "222222222222")}}),
			want:   want{shareType: "SPECIFIC_PROJECTS", projects: []string{"111111111111", "222222222222"}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			diff, err := sm.Diff(context.Background(), tc.state, tc.config, nil, nil, false)
			if err != nil {
				t.Fatalf("Diff(...): %v", err)
			}
			diff, err = reservationShareDiff(diff, tc.state, tc.config)
			if err != nil {
				t.Fatalf("reservationShareDiff(...): %v", err)
			}

			_, countDiff := diff.Attributes["share_settings.#"]
			got := want{countDiff: countDiff, requiresNew: diff.RequiresNew()}
			data, err := sm.Data(tc.state, diff)
			if err != nil {
				t.Fatalf("Data(...): %v", err)
			}
			if blocks := data.Get("share_settings").([]any); len(blocks) == 1 {
				block := blocks[0].(map[string]any)
				got.shareType = block["share_type"].(string)
				for _, p := range block["project_map"].(*schema.Set).List() {
					got.projects = append(got.projects, p.(map[string]any)["id"].(string))
				}
			}

			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(want{}), cmpopts.SortSlices(func(a, b string) bool { return a < b })); diff != "" {
				t.Errorf("%s\nreservationShareDiff(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
