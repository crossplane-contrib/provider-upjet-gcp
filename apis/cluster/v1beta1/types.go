// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials ProviderCredentials `json:"credentials"`

	// ProjectID is the project name (not numerical ID) of this GCP ProviderConfig.
	ProjectID string `json:"projectID"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;AccessToken;ImpersonateServiceAccount;InjectedIdentity;Environment;Filesystem;Upbound
	Source xpv1.CredentialsSource `json:"source"`

	// Upbound defines the options for authenticating using Upbound as an
	// identity provider.
	Upbound *Upbound `json:"upbound,omitempty"`

	xpv1.CommonCredentialSelectors `json:",inline"`
	// Use GCP service account impersonation for authentication.
	ImpersonateServiceAccountSpec `json:"impersonateServiceAccount,omitempty"`
}

type ImpersonateServiceAccountSpec struct {
	// GCP service account email address
	Name string `json:"name"`
}

// Upbound defines the options for authenticating using Upbound as an identity
// provider.
type Upbound struct {
	// Federation is the configuration for federated identity.
	Federation *Federation `json:"federation,omitempty"`
}

// Federation defines the configuration for federated identity from an external
// provider.
type Federation struct {
	// ProviderID is the fully-qualified identifier for the identity provider on
	// GCP. The format is
	// `projects/<project-id>/locations/global/workloadIdentityPools/<identity-pool>/providers/<identity-provider>`.
	// +kubebuilder:validation:MinLength=1
	ProviderID string `json:"providerID"`
	// ServiceAccount is the email address for the attached service account.
	// +kubebuilder:validation:MinLength=1
	ServiceAccount string `json:"serviceAccount"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a GCP provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,categories={crossplane,providerconfig,gcp}
// +kubebuilder:storageversion
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,providerconfig,gcp}
// +kubebuilder:storageversion
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
