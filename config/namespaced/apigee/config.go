// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package apigee

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/upbound/provider-gcp/v3/config/namespaced/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) { //nolint:gocyclo
	p.AddResourceConfigurator("google_apigee_api", configureAPI)
	p.AddResourceConfigurator("google_apigee_endpoint_attachment", func(r *config.Resource) {
		r.References["service_attachment"] = config.Reference{
			TerraformName: "google_compute_service_attachment",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_apigee_env_keystore", func(r *config.Resource) {
		r.References["env_id"] = config.Reference{
			TerraformName: "google_apigee_environment",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_apigee_env_references", func(r *config.Resource) {
		r.References["env_id"] = config.Reference{
			TerraformName: "google_apigee_environment",
			Extractor:     common.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_apigee_keystores_aliases_key_cert_file", func(r *config.Resource) {
		// org_id requires Organization ID without organization/ prefix, so extract name attribute
		r.References["org_id"] = config.Reference{
			TerraformName: "google_apigee_organization",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("name",true)`,
		}
	})
	p.AddResourceConfigurator("google_apigee_addons_config", func(r *config.Resource) {
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}
			// when all the addons configuration options are disabled,
			// upstream TF provider returns an empty list.
			// when the configuration explicitly specifies all as false,
			// it causes a diff, although they are
			// semantically equivalent. in such situation, ignore the diff.
			isAddonsConfigEmptyState := state.Attributes["addons_config.#"] == "0"
			addonsConfigurationsFields := []string{"advanced_api_ops_config", "api_security_config",
				"connectors_platform_config", "integration_config", "monetization_config"}

			hasActualDiff := false
			for _, field := range addonsConfigurationsFields {
				addonDiff, ok := diff.Attributes[fmt.Sprintf("addons_config.0.%s.0.enabled", field)]
				if ok && (addonDiff.Old != "" || addonDiff.New != "false") {
					hasActualDiff = true
					break
				}
			}
			if isAddonsConfigEmptyState && !hasActualDiff {
				for k := range diff.Attributes {
					if strings.HasPrefix(k, "addons_config.") {
						delete(diff.Attributes, k)
					}
				}
			}
			return diff, nil
		}
	})
}

func configureAPI(r *config.Resource) {
	// org_id requires Organization ID without organization/ prefix, so extract name attribute
	r.References["org_id"] = config.Reference{
		TerraformName: "google_apigee_organization",
		Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("name",true)`,
	}
	// Terraform reads the bundle from a local file, which a managed resource
	// cannot supply. The spec carries the bundle itself and
	// injectConfigBundlePath hands Terraform the path of a copy.
	r.TerraformConfigurationInjector = injectConfigBundlePath
	r.MetaResource.ArgumentDocs["config_bundle"] = configBundleDescription
	r.TerraformResource.Schema["config_bundle"].Description = ""
	// The observed config_bundle is the path of that copy, which must never
	// be late-initialized into the spec as if it were a bundle.
	r.LateInitializer = config.LateInitializer{
		IgnoredFields: []string{"config_bundle"},
	}
	// detect_md5hash lets a Terraform user signal a changed file behind an
	// unchanged path. Here the path changes with the bundle, so the hash is
	// left to the provider.
	config.MoveToStatus(r.TerraformResource, "detect_md5hash")
}

const (
	configBundleDescription = "The API proxy configuration bundle: the content of the zip archive, base64 encoded. " +
		"The provider writes the bundle to a temporary file and status.atProvider reports the path of that file."

	configBundleDirName = "provider-gcp-apigee-api-config-bundles"

	// Every reconcile refreshes the modification time of its bundle file, so
	// the retention only has to outlast the moment between a reconcile
	// storing the file and Terraform reading it.
	configBundleRetention = time.Hour
)

// injectConfigBundlePath replaces the base64 encoded bundle in config_bundle
// with the path of a file that holds the decoded bundle, which is what the
// google_apigee_api Terraform resource expects. The file is named after the
// digest of the bundle, so a changed bundle is a changed path and Terraform
// plans an update for it.
func injectConfigBundlePath(_ map[string]any, tfMap map[string]any) error {
	encoded, ok := tfMap["config_bundle"].(string)
	if !ok || encoded == "" {
		return nil
	}
	bundle, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return errors.Wrap(err, "cannot decode config_bundle, it must be a base64 encoded zip archive")
	}
	path, err := storeConfigBundle(filepath.Join(os.TempDir(), configBundleDirName), bundle, time.Now())
	if err != nil {
		return errors.Wrap(err, "cannot store the decoded config_bundle")
	}
	tfMap["config_bundle"] = path
	return nil
}

// storeConfigBundle makes sure that dir holds a file with the given bundle and
// returns its path. Storing a new bundle prunes the files that no reconcile
// has asked for within configBundleRetention.
func storeConfigBundle(dir string, bundle []byte, now time.Time) (string, error) {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", err
	}
	path := filepath.Join(dir, fmt.Sprintf("%x.zip", sha256.Sum256(bundle)))
	err := os.Chtimes(path, now, now)
	if err == nil {
		return path, nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return "", err
	}
	pruneConfigBundles(dir, now.Add(-configBundleRetention))

	// A rename is atomic, so a concurrent reconcile of the same bundle never
	// reads a partially written file.
	tmp, err := os.CreateTemp(dir, "bundle-*.tmp")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name()) //nolint:errcheck // the file is gone after a successful rename
	if _, err := tmp.Write(bundle); err != nil {
		_ = tmp.Close()
		return "", err
	}
	if err := tmp.Close(); err != nil {
		return "", err
	}
	if err := os.Chtimes(tmp.Name(), now, now); err != nil {
		return "", err
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		return "", err
	}
	return path, nil
}

// pruneConfigBundles removes the files in dir that were last modified before
// the cutoff. It is best effort: a file that cannot be removed now is retried
// the next time a bundle is stored.
func pruneConfigBundles(dir string, cutoff time.Time) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, e := range entries {
		info, err := e.Info()
		if err != nil || !info.ModTime().Before(cutoff) {
			continue
		}
		_ = os.Remove(filepath.Join(dir, e.Name()))
	}
}
