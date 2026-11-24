// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package logging

import (
	"net/http"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"

	"github.com/upbound/provider-gcp/v3/config/cluster/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_logging_folder_bucket_config", func(r *config.Resource) {
		r.References["folder"] = config.Reference{
			TerraformName: "google_folder",
			Extractor:     common.ExtractFolderIDFuncPath,
		}
	})
	p.AddResourceConfigurator("google_logging_project_bucket_config", func(r *config.Resource) {
		r.TerraformResource.Read = notFoundAsAbsent(r.TerraformResource.Read) //nolint:staticcheck // wrapping upstream's deprecated Read field
	})
}

// notFoundAsAbsent wraps the Terraform read of a project log bucket so that
// an HTTP 404 reports the bucket as absent instead of failing.
//
// The external name is built from project, location and bucket_id, so the
// provider reads the bucket before it has ever created it. Upstream's read
// returns the 404 as an error (plain Terraform never reads before the first
// create), so observe failed every time and a new bucket was never created.
// Only 404 is treated as absent: a 403 for a project that does not exist yet
// or whose Logging API is disabled still fails the observe.
// See: https://github.com/crossplane-contrib/provider-upjet-gcp/issues/1068
func notFoundAsAbsent(read schema.ReadFunc) schema.ReadFunc { //nolint:staticcheck // upstream resource uses deprecated Read field
	return func(d *schema.ResourceData, meta interface{}) error {
		err := read(d, meta)
		if err != nil && transport_tpg.IsGoogleApiErrorWithCode(err, http.StatusNotFound) {
			d.SetId("")
			return nil
		}
		return err
	}
}
