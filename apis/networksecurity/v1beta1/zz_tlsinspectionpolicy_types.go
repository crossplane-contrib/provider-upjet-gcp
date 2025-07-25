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

type TLSInspectionPolicyInitParameters struct {

	// A CA pool resource used to issue interception certificates.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/privateca/v1beta2.CAPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	CAPool *string `json:"caPool,omitempty" tf:"ca_pool,omitempty"`

	// Reference to a CAPool in privateca to populate caPool.
	// +kubebuilder:validation:Optional
	CAPoolRef *v1.Reference `json:"caPoolRef,omitempty" tf:"-"`

	// Selector for a CAPool in privateca to populate caPool.
	// +kubebuilder:validation:Optional
	CAPoolSelector *v1.Selector `json:"caPoolSelector,omitempty" tf:"-"`

	// List of custom TLS cipher suites selected. This field is valid only if the selected tls_feature_profile is CUSTOM. The compute.SslPoliciesService.ListAvailableFeatures method returns the set of features that can be specified in this list. Note that Secure Web Proxy does not yet honor this field.
	CustomTLSFeatures []*string `json:"customTlsFeatures,omitempty" tf:"custom_tls_features,omitempty"`

	// Free-text description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	ExcludePublicCASet *bool `json:"excludePublicCaSet,omitempty" tf:"exclude_public_ca_set,omitempty"`

	// Minimum TLS version that the firewall should use when negotiating connections with both clients and servers. If this is not set, then the default value is to allow the broadest set of clients and servers (TLS 1.0 or higher). Setting this to more restrictive values may improve security, but may also prevent the firewall from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	// Default value is TLS_VERSION_UNSPECIFIED.
	// Possible values are: TLS_VERSION_UNSPECIFIED, TLS_1_0, TLS_1_1, TLS_1_2, TLS_1_3.
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The selected Profile. If this is not set, then the default value is to allow the broadest set of clients and servers ("PROFILE_COMPATIBLE"). Setting this to more restrictive values may improve security, but may also prevent the TLS inspection proxy from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	// Default value is PROFILE_UNSPECIFIED.
	// Possible values are: PROFILE_UNSPECIFIED, PROFILE_COMPATIBLE, PROFILE_MODERN, PROFILE_RESTRICTED, PROFILE_CUSTOM.
	TLSFeatureProfile *string `json:"tlsFeatureProfile,omitempty" tf:"tls_feature_profile,omitempty"`

	// A TrustConfig resource used when making a connection to the TLS server. This is a relative resource path following the form "projects/{project}/locations/{location}/trustConfigs/{trust_config}". This is necessary to intercept TLS connections to servers with certificates signed by a private CA or self-signed certificates. Trust config and the TLS inspection policy must be in the same region. Note that Secure Web Proxy does not yet honor this field.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta1.TrustConfig
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TrustConfig *string `json:"trustConfig,omitempty" tf:"trust_config,omitempty"`

	// Reference to a TrustConfig in certificatemanager to populate trustConfig.
	// +kubebuilder:validation:Optional
	TrustConfigRef *v1.Reference `json:"trustConfigRef,omitempty" tf:"-"`

	// Selector for a TrustConfig in certificatemanager to populate trustConfig.
	// +kubebuilder:validation:Optional
	TrustConfigSelector *v1.Selector `json:"trustConfigSelector,omitempty" tf:"-"`
}

type TLSInspectionPolicyObservation struct {

	// A CA pool resource used to issue interception certificates.
	CAPool *string `json:"caPool,omitempty" tf:"ca_pool,omitempty"`

	// The timestamp when the resource was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// List of custom TLS cipher suites selected. This field is valid only if the selected tls_feature_profile is CUSTOM. The compute.SslPoliciesService.ListAvailableFeatures method returns the set of features that can be specified in this list. Note that Secure Web Proxy does not yet honor this field.
	CustomTLSFeatures []*string `json:"customTlsFeatures,omitempty" tf:"custom_tls_features,omitempty"`

	// Free-text description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	ExcludePublicCASet *bool `json:"excludePublicCaSet,omitempty" tf:"exclude_public_ca_set,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/tlsInspectionPolicies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the tls inspection policy.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Minimum TLS version that the firewall should use when negotiating connections with both clients and servers. If this is not set, then the default value is to allow the broadest set of clients and servers (TLS 1.0 or higher). Setting this to more restrictive values may improve security, but may also prevent the firewall from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	// Default value is TLS_VERSION_UNSPECIFIED.
	// Possible values are: TLS_VERSION_UNSPECIFIED, TLS_1_0, TLS_1_1, TLS_1_2, TLS_1_3.
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The selected Profile. If this is not set, then the default value is to allow the broadest set of clients and servers ("PROFILE_COMPATIBLE"). Setting this to more restrictive values may improve security, but may also prevent the TLS inspection proxy from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	// Default value is PROFILE_UNSPECIFIED.
	// Possible values are: PROFILE_UNSPECIFIED, PROFILE_COMPATIBLE, PROFILE_MODERN, PROFILE_RESTRICTED, PROFILE_CUSTOM.
	TLSFeatureProfile *string `json:"tlsFeatureProfile,omitempty" tf:"tls_feature_profile,omitempty"`

	// A TrustConfig resource used when making a connection to the TLS server. This is a relative resource path following the form "projects/{project}/locations/{location}/trustConfigs/{trust_config}". This is necessary to intercept TLS connections to servers with certificates signed by a private CA or self-signed certificates. Trust config and the TLS inspection policy must be in the same region. Note that Secure Web Proxy does not yet honor this field.
	TrustConfig *string `json:"trustConfig,omitempty" tf:"trust_config,omitempty"`

	// The timestamp when the resource was updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type TLSInspectionPolicyParameters struct {

	// A CA pool resource used to issue interception certificates.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/privateca/v1beta2.CAPool
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CAPool *string `json:"caPool,omitempty" tf:"ca_pool,omitempty"`

	// Reference to a CAPool in privateca to populate caPool.
	// +kubebuilder:validation:Optional
	CAPoolRef *v1.Reference `json:"caPoolRef,omitempty" tf:"-"`

	// Selector for a CAPool in privateca to populate caPool.
	// +kubebuilder:validation:Optional
	CAPoolSelector *v1.Selector `json:"caPoolSelector,omitempty" tf:"-"`

	// List of custom TLS cipher suites selected. This field is valid only if the selected tls_feature_profile is CUSTOM. The compute.SslPoliciesService.ListAvailableFeatures method returns the set of features that can be specified in this list. Note that Secure Web Proxy does not yet honor this field.
	// +kubebuilder:validation:Optional
	CustomTLSFeatures []*string `json:"customTlsFeatures,omitempty" tf:"custom_tls_features,omitempty"`

	// Free-text description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	// +kubebuilder:validation:Optional
	ExcludePublicCASet *bool `json:"excludePublicCaSet,omitempty" tf:"exclude_public_ca_set,omitempty"`

	// The location of the tls inspection policy.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Minimum TLS version that the firewall should use when negotiating connections with both clients and servers. If this is not set, then the default value is to allow the broadest set of clients and servers (TLS 1.0 or higher). Setting this to more restrictive values may improve security, but may also prevent the firewall from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	// Default value is TLS_VERSION_UNSPECIFIED.
	// Possible values are: TLS_VERSION_UNSPECIFIED, TLS_1_0, TLS_1_1, TLS_1_2, TLS_1_3.
	// +kubebuilder:validation:Optional
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The selected Profile. If this is not set, then the default value is to allow the broadest set of clients and servers ("PROFILE_COMPATIBLE"). Setting this to more restrictive values may improve security, but may also prevent the TLS inspection proxy from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	// Default value is PROFILE_UNSPECIFIED.
	// Possible values are: PROFILE_UNSPECIFIED, PROFILE_COMPATIBLE, PROFILE_MODERN, PROFILE_RESTRICTED, PROFILE_CUSTOM.
	// +kubebuilder:validation:Optional
	TLSFeatureProfile *string `json:"tlsFeatureProfile,omitempty" tf:"tls_feature_profile,omitempty"`

	// A TrustConfig resource used when making a connection to the TLS server. This is a relative resource path following the form "projects/{project}/locations/{location}/trustConfigs/{trust_config}". This is necessary to intercept TLS connections to servers with certificates signed by a private CA or self-signed certificates. Trust config and the TLS inspection policy must be in the same region. Note that Secure Web Proxy does not yet honor this field.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta1.TrustConfig
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TrustConfig *string `json:"trustConfig,omitempty" tf:"trust_config,omitempty"`

	// Reference to a TrustConfig in certificatemanager to populate trustConfig.
	// +kubebuilder:validation:Optional
	TrustConfigRef *v1.Reference `json:"trustConfigRef,omitempty" tf:"-"`

	// Selector for a TrustConfig in certificatemanager to populate trustConfig.
	// +kubebuilder:validation:Optional
	TrustConfigSelector *v1.Selector `json:"trustConfigSelector,omitempty" tf:"-"`
}

// TLSInspectionPolicySpec defines the desired state of TLSInspectionPolicy
type TLSInspectionPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TLSInspectionPolicyParameters `json:"forProvider"`
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
	InitProvider TLSInspectionPolicyInitParameters `json:"initProvider,omitempty"`
}

// TLSInspectionPolicyStatus defines the observed state of TLSInspectionPolicy.
type TLSInspectionPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TLSInspectionPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TLSInspectionPolicy is the Schema for the TLSInspectionPolicys API. The TlsInspectionPolicy resource contains references to CA pools in Certificate Authority Service and associated metadata.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TLSInspectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TLSInspectionPolicySpec   `json:"spec"`
	Status            TLSInspectionPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TLSInspectionPolicyList contains a list of TLSInspectionPolicys
type TLSInspectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TLSInspectionPolicy `json:"items"`
}

// Repository type metadata.
var (
	TLSInspectionPolicy_Kind             = "TLSInspectionPolicy"
	TLSInspectionPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TLSInspectionPolicy_Kind}.String()
	TLSInspectionPolicy_KindAPIVersion   = TLSInspectionPolicy_Kind + "." + CRDGroupVersion.String()
	TLSInspectionPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TLSInspectionPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TLSInspectionPolicy{}, &TLSInspectionPolicyList{})
}
