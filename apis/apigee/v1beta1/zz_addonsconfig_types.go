// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type APISecurityConfigInitParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type APISecurityConfigObservation struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Output)
	// Time at which the Connectors Platform add-on expires in milliseconds since epoch. If unspecified, the add-on will never expire.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`
}

type APISecurityConfigParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AddonsConfigAddonsConfigInitParameters struct {

	// Configuration for the API Security add-on.
	// Structure is documented below.
	APISecurityConfig *APISecurityConfigInitParameters `json:"apiSecurityConfig,omitempty" tf:"api_security_config,omitempty"`

	// Configuration for the Advanced API Ops add-on.
	// Structure is documented below.
	AdvancedAPIOpsConfig *AdvancedAPIOpsConfigInitParameters `json:"advancedApiOpsConfig,omitempty" tf:"advanced_api_ops_config,omitempty"`

	// Configuration for the Monetization add-on.
	// Structure is documented below.
	ConnectorsPlatformConfig *ConnectorsPlatformConfigInitParameters `json:"connectorsPlatformConfig,omitempty" tf:"connectors_platform_config,omitempty"`

	// Configuration for the Integration add-on.
	// Structure is documented below.
	IntegrationConfig *IntegrationConfigInitParameters `json:"integrationConfig,omitempty" tf:"integration_config,omitempty"`

	// Configuration for the Monetization add-on.
	// Structure is documented below.
	MonetizationConfig *MonetizationConfigInitParameters `json:"monetizationConfig,omitempty" tf:"monetization_config,omitempty"`
}

type AddonsConfigAddonsConfigObservation struct {

	// Configuration for the API Security add-on.
	// Structure is documented below.
	APISecurityConfig *APISecurityConfigObservation `json:"apiSecurityConfig,omitempty" tf:"api_security_config,omitempty"`

	// Configuration for the Advanced API Ops add-on.
	// Structure is documented below.
	AdvancedAPIOpsConfig *AdvancedAPIOpsConfigObservation `json:"advancedApiOpsConfig,omitempty" tf:"advanced_api_ops_config,omitempty"`

	// Configuration for the Monetization add-on.
	// Structure is documented below.
	ConnectorsPlatformConfig *ConnectorsPlatformConfigObservation `json:"connectorsPlatformConfig,omitempty" tf:"connectors_platform_config,omitempty"`

	// Configuration for the Integration add-on.
	// Structure is documented below.
	IntegrationConfig *IntegrationConfigObservation `json:"integrationConfig,omitempty" tf:"integration_config,omitempty"`

	// Configuration for the Monetization add-on.
	// Structure is documented below.
	MonetizationConfig *MonetizationConfigObservation `json:"monetizationConfig,omitempty" tf:"monetization_config,omitempty"`
}

type AddonsConfigAddonsConfigParameters struct {

	// Configuration for the API Security add-on.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	APISecurityConfig *APISecurityConfigParameters `json:"apiSecurityConfig,omitempty" tf:"api_security_config,omitempty"`

	// Configuration for the Advanced API Ops add-on.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AdvancedAPIOpsConfig *AdvancedAPIOpsConfigParameters `json:"advancedApiOpsConfig,omitempty" tf:"advanced_api_ops_config,omitempty"`

	// Configuration for the Monetization add-on.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ConnectorsPlatformConfig *ConnectorsPlatformConfigParameters `json:"connectorsPlatformConfig,omitempty" tf:"connectors_platform_config,omitempty"`

	// Configuration for the Integration add-on.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	IntegrationConfig *IntegrationConfigParameters `json:"integrationConfig,omitempty" tf:"integration_config,omitempty"`

	// Configuration for the Monetization add-on.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MonetizationConfig *MonetizationConfigParameters `json:"monetizationConfig,omitempty" tf:"monetization_config,omitempty"`
}

type AddonsConfigInitParameters struct {

	// Addon configurations of the Apigee organization.
	// Structure is documented below.
	AddonsConfig *AddonsConfigAddonsConfigInitParameters `json:"addonsConfig,omitempty" tf:"addons_config,omitempty"`

	// Name of the Apigee organization.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta2.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Reference to a Organization in apigee to populate org.
	// +kubebuilder:validation:Optional
	OrgRef *v1.Reference `json:"orgRef,omitempty" tf:"-"`

	// Selector for a Organization in apigee to populate org.
	// +kubebuilder:validation:Optional
	OrgSelector *v1.Selector `json:"orgSelector,omitempty" tf:"-"`
}

type AddonsConfigObservation struct {

	// Addon configurations of the Apigee organization.
	// Structure is documented below.
	AddonsConfig *AddonsConfigAddonsConfigObservation `json:"addonsConfig,omitempty" tf:"addons_config,omitempty"`

	// an identifier for the resource with format organizations/{{org}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Apigee organization.
	Org *string `json:"org,omitempty" tf:"org,omitempty"`
}

type AddonsConfigParameters struct {

	// Addon configurations of the Apigee organization.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AddonsConfig *AddonsConfigAddonsConfigParameters `json:"addonsConfig,omitempty" tf:"addons_config,omitempty"`

	// Name of the Apigee organization.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta2.Organization
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Reference to a Organization in apigee to populate org.
	// +kubebuilder:validation:Optional
	OrgRef *v1.Reference `json:"orgRef,omitempty" tf:"-"`

	// Selector for a Organization in apigee to populate org.
	// +kubebuilder:validation:Optional
	OrgSelector *v1.Selector `json:"orgSelector,omitempty" tf:"-"`
}

type AdvancedAPIOpsConfigInitParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdvancedAPIOpsConfigObservation struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdvancedAPIOpsConfigParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ConnectorsPlatformConfigInitParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ConnectorsPlatformConfigObservation struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Output)
	// Time at which the Connectors Platform add-on expires in milliseconds since epoch. If unspecified, the add-on will never expire.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`
}

type ConnectorsPlatformConfigParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type IntegrationConfigInitParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type IntegrationConfigObservation struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type IntegrationConfigParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type MonetizationConfigInitParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type MonetizationConfigObservation struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type MonetizationConfigParameters struct {

	// Flag that specifies whether the Monetization add-on is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

// AddonsConfigSpec defines the desired state of AddonsConfig
type AddonsConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddonsConfigParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AddonsConfigInitParameters `json:"initProvider,omitempty"`
}

// AddonsConfigStatus defines the observed state of AddonsConfig.
type AddonsConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddonsConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AddonsConfig is the Schema for the AddonsConfigs API. Configures the add-ons for the Apigee organization.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AddonsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddonsConfigSpec   `json:"spec"`
	Status            AddonsConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddonsConfigList contains a list of AddonsConfigs
type AddonsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddonsConfig `json:"items"`
}

// Repository type metadata.
var (
	AddonsConfig_Kind             = "AddonsConfig"
	AddonsConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AddonsConfig_Kind}.String()
	AddonsConfig_KindAPIVersion   = AddonsConfig_Kind + "." + CRDGroupVersion.String()
	AddonsConfig_GroupVersionKind = CRDGroupVersion.WithKind(AddonsConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&AddonsConfig{}, &AddonsConfigList{})
}
