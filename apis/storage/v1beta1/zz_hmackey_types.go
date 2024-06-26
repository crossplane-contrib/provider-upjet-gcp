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

type HMACKeyInitParameters struct {

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The email address of the key's associated service account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("email",true)
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailRef *v1.Reference `json:"serviceAccountEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailSelector *v1.Selector `json:"serviceAccountEmailSelector,omitempty" tf:"-"`

	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is ACTIVE.
	// Possible values are: ACTIVE, INACTIVE.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type HMACKeyObservation struct {

	// The access ID of the HMAC Key.
	AccessID *string `json:"accessId,omitempty" tf:"access_id,omitempty"`

	// an identifier for the resource with format projects/{{project}}/hmacKeys/{{access_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The email address of the key's associated service account.
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is ACTIVE.
	// Possible values are: ACTIVE, INACTIVE.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// 'The creation time of the HMAC key in RFC 3339 format. '
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// 'The last modification time of the HMAC key metadata in RFC 3339 format.'
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type HMACKeyParameters struct {

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The email address of the key's associated service account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("email",true)
	// +kubebuilder:validation:Optional
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailRef *v1.Reference `json:"serviceAccountEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate serviceAccountEmail.
	// +kubebuilder:validation:Optional
	ServiceAccountEmailSelector *v1.Selector `json:"serviceAccountEmailSelector,omitempty" tf:"-"`

	// The state of the key. Can be set to one of ACTIVE, INACTIVE.
	// Default value is ACTIVE.
	// Possible values are: ACTIVE, INACTIVE.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

// HMACKeySpec defines the desired state of HMACKey
type HMACKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HMACKeyParameters `json:"forProvider"`
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
	InitProvider HMACKeyInitParameters `json:"initProvider,omitempty"`
}

// HMACKeyStatus defines the observed state of HMACKey.
type HMACKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HMACKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HMACKey is the Schema for the HMACKeys API. The hmacKeys resource represents an HMAC key within Cloud Storage.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type HMACKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HMACKeySpec   `json:"spec"`
	Status            HMACKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HMACKeyList contains a list of HMACKeys
type HMACKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HMACKey `json:"items"`
}

// Repository type metadata.
var (
	HMACKey_Kind             = "HMACKey"
	HMACKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HMACKey_Kind}.String()
	HMACKey_KindAPIVersion   = HMACKey_Kind + "." + CRDGroupVersion.String()
	HMACKey_GroupVersionKind = CRDGroupVersion.WithKind(HMACKey_Kind)
)

func init() {
	SchemeBuilder.Register(&HMACKey{}, &HMACKeyList{})
}
