// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestResourceConfiguratorAppliesExternalNames(t *testing.T) {
	cases := map[string]struct {
		args string
		want bool
	}{
		"TerraformPluginFrameworkResource": {
			args: "google_storage_notification",
			want: true,
		},
		"TerraformPluginFrameworkResourceWithCustomConfig": {
			args: "google_apigee_keystores_aliases_key_cert_file",
			want: true,
		},
		"TerraformPluginSDKResource": {
			args: "google_active_directory_domain",
			want: true,
		},
		"UnknownResource": {
			args: "google_nonexistent_resource",
			want: false,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			r := &config.Resource{Name: tc.args}
			resourceConfigurator()(r)
			got := r.Version == VersionV1Beta1
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("resourceConfigurator() applied: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestFrameworkExternalNamesConfigureNotFoundDiagnostics(t *testing.T) {
	for name := range terraformPluginFrameworkExternalNameConfigs {
		t.Run(name, func(t *testing.T) {
			if terraformPluginFrameworkExternalNameConfigs[name].IsNotFoundDiagnosticFn == nil {
				t.Errorf("external-name configuration for %s must configure IsNotFoundDiagnosticFn: the upstream plugin-framework Read implementations surface not-found conditions as error diagnostics", name)
			}
		})
	}
}

func TestApigeeKeystoresAliasesKeyCertFileIsNotFoundDiagnostic(t *testing.T) {
	fn := apigeeKeystoresAliasesKeyCertFile().IsNotFoundDiagnosticFn
	cases := map[string]struct {
		args []*tfprotov6.Diagnostic
		want bool
	}{
		"NotFound": {
			args: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error when sending HTTP request: ",
				Detail:   "googleapi: Error 404: generic::not_found: alias not organizations/example-org/environments/example-env/keystores/example-keystore/aliases/example-alias found",
			}},
			want: true,
		},
		"NotFoundRawBody": {
			args: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error when sending HTTP request: ",
				Detail:   "googleapi: got HTTP response code 404 with body: not found",
			}},
			want: true,
		},
		"PermissionDenied": {
			args: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error when sending HTTP request: ",
				Detail:   "googleapi: Error 403: permission denied",
			}},
			want: false,
		},
		"OtherSummary": {
			args: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Invalid resource ID",
				Detail:   "googleapi: Error 404: not found",
			}},
			want: false,
		},
		"WarningSeverity": {
			args: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityWarning,
				Summary:  "Error when sending HTTP request: ",
				Detail:   "googleapi: Error 404: not found",
			}},
			want: false,
		},
		"NoDiagnostics": {
			args: nil,
			want: false,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(tc.want, fn(tc.args)); diff != "" {
				t.Errorf("IsNotFoundDiagnosticFn(...): -want, +got:\n%s", diff)
			}
		})
	}
}

func TestStorageNotificationIsNotFoundDiagnostic(t *testing.T) {
	fn := storageNotification().IsNotFoundDiagnosticFn
	cases := map[string]struct {
		args []*tfprotov6.Diagnostic
		want bool
	}{
		"EmptyID": {
			args: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Invalid resource ID",
				Detail:   "invalid storage notification ID format, expected '{bucket}/notificationConfigs/{notification_id}', got ''",
			}},
			want: true,
		},
		"MalformedNonEmptyID": {
			args: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Invalid resource ID",
				Detail:   "invalid storage notification ID format, expected '{bucket}/notificationConfigs/{notification_id}', got 'some-wrong-id'",
			}},
			want: false,
		},
		"OtherSummary": {
			args: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Error reading Storage Notification",
				Detail:   "googleapi: Error 500: internal error, got ''",
			}},
			want: false,
		},
		"NoDiagnostics": {
			args: nil,
			want: false,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(tc.want, fn(tc.args)); diff != "" {
				t.Errorf("IsNotFoundDiagnosticFn(...): -want, +got:\n%s", diff)
			}
		})
	}
}

func TestServiceAccountGetIDFn(t *testing.T) {
	cases := map[string]struct {
		parameters map[string]any
		setup      map[string]any
		want       string
		wantErr    bool
	}{
		"ProjectFromResource": {
			parameters: map[string]any{"project": "my-project"},
			setup:      map[string]any{"configuration": terraform.ProviderConfiguration{"project": "pc-project"}},
			want:       "projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com",
		},
		"ProjectFromProviderConfig": {
			parameters: map[string]any{},
			setup:      map[string]any{"configuration": terraform.ProviderConfiguration{"project": "pc-project"}},
			want:       "projects/pc-project/serviceAccounts/my-sa@pc-project.iam.gserviceaccount.com",
		},
		"DomainScopedProject": {
			parameters: map[string]any{"project": "example.com:my-project"},
			setup:      map[string]any{},
			want:       "projects/example.com:my-project/serviceAccounts/my-sa@my-project.example.com.iam.gserviceaccount.com",
		},
		"DomainScopedProjectFromProviderConfig": {
			parameters: map[string]any{},
			setup:      map[string]any{"configuration": terraform.ProviderConfiguration{"project": "example.com:my-project"}},
			want:       "projects/example.com:my-project/serviceAccounts/my-sa@my-project.example.com.iam.gserviceaccount.com",
		},
		"ProjectFromProviderConfigPlainMap": {
			parameters: map[string]any{},
			setup:      map[string]any{"configuration": map[string]any{"project": "pc-project"}},
			want:       "projects/pc-project/serviceAccounts/my-sa@pc-project.iam.gserviceaccount.com",
		},
		"NoProject": {
			parameters: map[string]any{},
			setup:      map[string]any{},
			wantErr:    true,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := serviceAccount().GetIDFn(context.Background(), "my-sa", tc.parameters, tc.setup)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("GetIDFn(): want error, got %q", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("GetIDFn(): unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetIDFn(): -want, +got:\n%s", diff)
			}
		})
	}
}

func TestServiceAccountGetExternalNameFn(t *testing.T) {
	cases := map[string]struct {
		id   string
		want string
	}{
		"DefaultProject": {
			id:   "projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com",
			want: "my-sa",
		},
		"DomainScopedProject": {
			id:   "projects/s3ns:test-mathieu/serviceAccounts/my-sa@test-mathieu.s3ns.iam.gserviceaccount.com",
			want: "my-sa",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := serviceAccount().GetExternalNameFn(map[string]any{"id": tc.id})
			if err != nil {
				t.Fatalf("GetExternalNameFn(): unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetExternalNameFn(): -want, +got:\n%s", diff)
			}
		})
	}
}
