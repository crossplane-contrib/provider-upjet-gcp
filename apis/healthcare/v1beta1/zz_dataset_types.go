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

type DatasetInitParameters struct {

	// A nested object resource.
	// Structure is documented below.
	EncryptionSpec *EncryptionSpecInitParameters `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// The location for the Dataset.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name for the Dataset.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type DatasetObservation struct {

	// A nested object resource.
	// Structure is documented below.
	EncryptionSpec *EncryptionSpecObservation `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/datasets/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location for the Dataset.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name for the Dataset.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The fully qualified name of this dataset
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type DatasetParameters struct {

	// A nested object resource.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionSpec *EncryptionSpecParameters `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// The location for the Dataset.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name for the Dataset.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type EncryptionSpecInitParameters struct {

	// KMS encryption key that is used to secure this dataset and its sub-resources. The key used for
	// encryption and the dataset must be in the same location. If empty, the default Google encryption
	// key will be used to secure this dataset. The format is
	// projects/{projectId}/locations/{locationId}/keyRings/{keyRingId}/cryptoKeys/{keyId}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta2.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameRef *v1.Reference `json:"kmsKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameSelector *v1.Selector `json:"kmsKeyNameSelector,omitempty" tf:"-"`
}

type EncryptionSpecObservation struct {

	// KMS encryption key that is used to secure this dataset and its sub-resources. The key used for
	// encryption and the dataset must be in the same location. If empty, the default Google encryption
	// key will be used to secure this dataset. The format is
	// projects/{projectId}/locations/{locationId}/keyRings/{keyRingId}/cryptoKeys/{keyId}.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type EncryptionSpecParameters struct {

	// KMS encryption key that is used to secure this dataset and its sub-resources. The key used for
	// encryption and the dataset must be in the same location. If empty, the default Google encryption
	// key will be used to secure this dataset. The format is
	// projects/{projectId}/locations/{locationId}/keyRings/{keyRingId}/cryptoKeys/{keyId}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta2.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameRef *v1.Reference `json:"kmsKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameSelector *v1.Selector `json:"kmsKeyNameSelector,omitempty" tf:"-"`
}

// DatasetSpec defines the desired state of Dataset
type DatasetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetParameters `json:"forProvider"`
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
	InitProvider DatasetInitParameters `json:"initProvider,omitempty"`
}

// DatasetStatus defines the observed state of Dataset.
type DatasetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Dataset is the Schema for the Datasets API. A Healthcare
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DatasetSpec   `json:"spec"`
	Status DatasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetList contains a list of Datasets
type DatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dataset `json:"items"`
}

// Repository type metadata.
var (
	Dataset_Kind             = "Dataset"
	Dataset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dataset_Kind}.String()
	Dataset_KindAPIVersion   = Dataset_Kind + "." + CRDGroupVersion.String()
	Dataset_GroupVersionKind = CRDGroupVersion.WithKind(Dataset_Kind)
)

func init() {
	SchemeBuilder.Register(&Dataset{}, &DatasetList{})
}
