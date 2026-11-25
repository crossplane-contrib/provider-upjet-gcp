// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apigee

import (
	"context"
	"encoding/base64"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfapigee "github.com/hashicorp/terraform-provider-google/google/services/apigee"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

const (
	bundleFileName      = "1e6ed65d77d6364eeaed5a745ba5c4985ae2b700dd85d7cf7f027bdf294a33fc.zip"
	otherBundleFileName = "53033a07c7692390658742368631254742f6fbb634db8a78f5a947460adbbd19.zip"
)

func Test_injectConfigBundlePath(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("TMPDIR", tmp)
	bundlePath := filepath.Join(tmp, configBundleDirName, bundleFileName)

	type args struct {
		tfMap map[string]any
	}
	type want struct {
		tfMap  map[string]any
		bundle []byte
		err    error
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"NoConfigBundle": {
			reason: "An observe-only resource has no bundle, so there is nothing to store",
			args:   args{tfMap: map[string]any{"name": "api"}},
			want:   want{tfMap: map[string]any{"name": "api"}},
		},
		"Base64Bundle": {
			reason: "The bundle is decoded into a file named after its digest and Terraform gets the path",
			args:   args{tfMap: map[string]any{"config_bundle": "YnVuZGxl"}},
			want: want{
				tfMap:  map[string]any{"config_bundle": bundlePath},
				bundle: []byte("bundle"),
			},
		},
		"Base64BundleWithLineBreaks": {
			reason: "A bundle wrapped by the base64 tool is the same bundle",
			args:   args{tfMap: map[string]any{"config_bundle": "YnVu\nZGxl\n"}},
			want: want{
				tfMap:  map[string]any{"config_bundle": bundlePath},
				bundle: []byte("bundle"),
			},
		},
		"InvalidBase64": {
			reason: "A value that is not base64 is rejected and left untouched",
			args:   args{tfMap: map[string]any{"config_bundle": "%%%"}},
			want: want{
				tfMap: map[string]any{"config_bundle": "%%%"},
				err:   base64.CorruptInputError(0),
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := injectConfigBundlePath(nil, tc.args.tfMap)
			if diff := cmp.Diff(tc.want.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\ninjectConfigBundlePath(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.tfMap, tc.args.tfMap); diff != "" {
				t.Errorf("\n%s\ninjectConfigBundlePath(...): -want tfMap, +got tfMap:\n%s", tc.reason, diff)
			}
			if tc.want.bundle == nil {
				return
			}
			bundle, err := os.ReadFile(tc.args.tfMap["config_bundle"].(string))
			if err != nil {
				t.Fatalf("\n%s\ncannot read the stored bundle: %v", tc.reason, err)
			}
			if diff := cmp.Diff(tc.want.bundle, bundle); diff != "" {
				t.Errorf("\n%s\ninjectConfigBundlePath(...): -want bundle, +got bundle:\n%s", tc.reason, diff)
			}
		})
	}
}

func Test_storeConfigBundle(t *testing.T) {
	now := time.Date(2026, time.November, 23, 12, 0, 0, 0, time.UTC)

	type args struct {
		files map[string]time.Duration
	}
	type want struct {
		files map[string]time.Duration
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"NewBundle": {
			reason: "A bundle that is not stored yet is written",
			args:   args{files: map[string]time.Duration{}},
			want:   want{files: map[string]time.Duration{bundleFileName: 0}},
		},
		"StoredBundle": {
			reason: "A stored bundle is reused and marked as recently used",
			args:   args{files: map[string]time.Duration{bundleFileName: 2 * configBundleRetention}},
			want:   want{files: map[string]time.Duration{bundleFileName: 0}},
		},
		"NewBundlePrunesUnusedBundles": {
			reason: "Storing a new bundle removes the bundles that were not used within the retention",
			args: args{files: map[string]time.Duration{
				otherBundleFileName: 2 * configBundleRetention,
				"recently-used.zip": configBundleRetention / 2,
			}},
			want: want{files: map[string]time.Duration{
				bundleFileName:      0,
				"recently-used.zip": configBundleRetention / 2,
			}},
		},
		"StoredBundleKeepsUnusedBundles": {
			reason: "Pruning only runs when a new bundle is stored",
			args: args{files: map[string]time.Duration{
				bundleFileName:      configBundleRetention / 2,
				otherBundleFileName: 2 * configBundleRetention,
			}},
			want: want{files: map[string]time.Duration{
				bundleFileName:      0,
				otherBundleFileName: 2 * configBundleRetention,
			}},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			for file, age := range tc.args.files {
				path := filepath.Join(dir, file)
				if err := os.WriteFile(path, []byte("bundle"), 0o600); err != nil {
					t.Fatal(err)
				}
				if err := os.Chtimes(path, now.Add(-age), now.Add(-age)); err != nil {
					t.Fatal(err)
				}
			}

			path, err := storeConfigBundle(dir, []byte("bundle"), now)
			if err != nil {
				t.Fatalf("\n%s\nstoreConfigBundle(...): unexpected error: %v", tc.reason, err)
			}
			if diff := cmp.Diff(filepath.Join(dir, bundleFileName), path); diff != "" {
				t.Errorf("\n%s\nstoreConfigBundle(...): -want path, +got path:\n%s", tc.reason, diff)
			}

			entries, err := os.ReadDir(dir)
			if err != nil {
				t.Fatal(err)
			}
			files := map[string]time.Duration{}
			for _, e := range entries {
				info, err := e.Info()
				if err != nil {
					t.Fatal(err)
				}
				files[e.Name()] = now.Sub(info.ModTime())
			}
			if diff := cmp.Diff(tc.want.files, files); diff != "" {
				t.Errorf("\n%s\nstoreConfigBundle(...): -want files, +got files:\n%s", tc.reason, diff)
			}
		})
	}
}

// TestAPIConfigBundleDiff plans the google_apigee_api Terraform resource the
// way upjet does, with the bundle path injected. Every update of this resource
// uploads the bundle as a new API proxy revision, so an unchanged bundle must
// plan nothing.
func TestAPIConfigBundleDiff(t *testing.T) {
	t.Setenv("TMPDIR", t.TempDir())

	r := &config.Resource{
		TerraformResource: tfapigee.ResourceApigeeApi(),
		References:        config.References{},
		MetaResource:      &registry.Resource{ArgumentDocs: map[string]string{}},
	}
	configureAPI(r)

	bundle := base64.StdEncoding.EncodeToString([]byte("bundle"))
	otherBundle := base64.StdEncoding.EncodeToString([]byte("other bundle"))

	type args struct {
		uploadedBundle  string
		modifiedOutside bool
		specBundle      string
	}
	type want struct {
		diff []string
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"UnchangedBundle": {
			reason: "The uploaded bundle is the bundle in the spec, so there is nothing to update",
			args:   args{uploadedBundle: bundle, specBundle: bundle},
			want:   want{diff: []string{}},
		},
		"ChangedBundle": {
			reason: "A new bundle in the spec is uploaded as a new revision",
			args:   args{uploadedBundle: bundle, specBundle: otherBundle},
			want:   want{diff: []string{"config_bundle", "latest_revision_id", "meta_data.#", "revision.#"}},
		},
		"BundleModifiedOutside": {
			reason: "A revision that was not uploaded by the provider is replaced with the bundle in the spec",
			args:   args{uploadedBundle: bundle, modifiedOutside: true, specBundle: bundle},
			want:   want{diff: []string{"latest_revision_id", "meta_data.#", "revision.#"}},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			uploaded := injectedParams(t, tc.args.uploadedBundle)
			hash := tpgresource.GetFileMd5Hash(uploaded["config_bundle"].(string))
			if tc.args.modifiedOutside {
				// What the Terraform resource stores on read when the last
				// modification time of the API proxy has moved.
				hash = "UNKNOWN"
			}
			state := map[string]any{
				"id":                 "organizations/org/apis/api",
				"name":               "api",
				"org_id":             "org",
				"config_bundle":      uploaded["config_bundle"],
				"md5hash":            hash,
				"detect_md5hash":     hash,
				"deletion_policy":    "DELETE",
				"latest_revision_id": "1",
				"revision":           []any{"1"},
				"meta_data": []any{map[string]any{
					"created_at":       "1",
					"last_modified_at": "1",
					"sub_type":         "Proxy",
				}},
			}

			got := diffKeys(t, r.TerraformResource, state, injectedParams(t, tc.args.specBundle))
			if diff := cmp.Diff(tc.want.diff, got); diff != "" {
				t.Errorf("\n%s\n-want diff, +got diff:\n%s", tc.reason, diff)
			}
		})
	}
}

func injectedParams(t *testing.T, bundle string) map[string]any {
	t.Helper()
	params := map[string]any{
		"name":          "api",
		"org_id":        "org",
		"config_bundle": bundle,
	}
	if err := injectConfigBundlePath(nil, params); err != nil {
		t.Fatal(err)
	}
	return params
}

// diffKeys returns the sorted attribute keys of the diff that upjet's
// Terraform plugin SDK client computes for the given state and parameters.
func diffKeys(t *testing.T, r *schema.Resource, state, params map[string]any) []string {
	t.Helper()
	block := r.CoreConfigSchema()
	rawConfig, err := schema.JSONMapToStateValue(params, block)
	if err != nil {
		t.Fatal(err)
	}
	rawState, err := schema.JSONMapToStateValue(state, block)
	if err != nil {
		t.Fatal(err)
	}
	s, err := r.ShimInstanceStateFromValue(rawState)
	if err != nil {
		t.Fatal(err)
	}
	s.RawConfig = rawConfig
	s.RawState = rawState
	s.RawPlan = rawState

	diff, err := schema.InternalMap(r.Schema).Diff(context.Background(), s, terraform.NewResourceConfigRaw(params), r.CustomizeDiff, &transport_tpg.Config{}, false)
	if err != nil {
		t.Fatal(err)
	}
	keys := []string{}
	if diff == nil {
		return keys
	}
	for k := range diff.Attributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
