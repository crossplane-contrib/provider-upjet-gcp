/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SnapshotEncryptionKeyObservation struct {

	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	// encryption key that protects this resource.
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type SnapshotEncryptionKeyParameters struct {

	// The name of the encryption key that is stored in Google Cloud KMS.
	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	// +kubebuilder:validation:Optional
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	RawKeySecretRef *v1.SecretKeySelector `json:"rawKeySecretRef,omitempty" tf:"-"`
}

type SnapshotObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// Size of the snapshot, specified in GB.
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/snapshots/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fingerprint used for optimistic locking of this resource. Used
	// internally during updates.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// A list of public visible licenses that apply to this snapshot. This
	// can be because the original image had licenses attached (such as a
	// Windows image).  snapshotEncryptionKey nested object Encrypts the
	// snapshot using a customer-supplied encryption key.
	Licenses []*string `json:"licenses,omitempty" tf:"licenses,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Encrypts the snapshot using a customer-supplied encryption key.
	// After you encrypt a snapshot using a customer-supplied key, you must
	// provide the same key if you use the snapshot later. For example, you
	// must provide the encryption key when you create a disk from the
	// encrypted snapshot in a future request.
	// Customer-supplied encryption keys do not protect access to metadata of
	// the snapshot.
	// If you do not provide an encryption key when creating the snapshot,
	// then the snapshot will be encrypted using an automatically generated
	// key and you do not need to provide a key to use the snapshot later.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SnapshotEncryptionKey []SnapshotEncryptionKeyObservation `json:"snapshotEncryptionKey,omitempty" tf:"snapshot_encryption_key,omitempty"`

	// The unique identifier for the resource.
	SnapshotID *float64 `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// A size of the storage used by the snapshot. As snapshots share
	// storage, this number is expected to change with snapshot
	// creation/deletion.
	StorageBytes *float64 `json:"storageBytes,omitempty" tf:"storage_bytes,omitempty"`
}

type SnapshotParameters struct {

	// Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and
	// comply with RFC1035. This is an uncommon option only for advanced
	// service owners who needs to create separate snapshot chains, for
	// example, for chargeback tracking.  When you describe your snapshot
	// resource, this field is visible only if it has a non-empty value.
	// +kubebuilder:validation:Optional
	ChainName *string `json:"chainName,omitempty" tf:"chain_name,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels to apply to this Snapshot.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Encrypts the snapshot using a customer-supplied encryption key.
	// After you encrypt a snapshot using a customer-supplied key, you must
	// provide the same key if you use the snapshot later. For example, you
	// must provide the encryption key when you create a disk from the
	// encrypted snapshot in a future request.
	// Customer-supplied encryption keys do not protect access to metadata of
	// the snapshot.
	// If you do not provide an encryption key when creating the snapshot,
	// then the snapshot will be encrypted using an automatically generated
	// key and you do not need to provide a key to use the snapshot later.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SnapshotEncryptionKey []SnapshotEncryptionKeyParameters `json:"snapshotEncryptionKey,omitempty" tf:"snapshot_encryption_key,omitempty"`

	// A reference to the disk used to create this snapshot.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Disk
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SourceDiskEncryptionKey []SourceDiskEncryptionKeyParameters `json:"sourceDiskEncryptionKey,omitempty" tf:"source_disk_encryption_key,omitempty"`

	// Reference to a Disk in compute to populate sourceDisk.
	// +kubebuilder:validation:Optional
	SourceDiskRef *v1.Reference `json:"sourceDiskRef,omitempty" tf:"-"`

	// Selector for a Disk in compute to populate sourceDisk.
	// +kubebuilder:validation:Optional
	SourceDiskSelector *v1.Selector `json:"sourceDiskSelector,omitempty" tf:"-"`

	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	// +kubebuilder:validation:Optional
	StorageLocations []*string `json:"storageLocations,omitempty" tf:"storage_locations,omitempty"`

	// A reference to the zone where the disk is hosted.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SourceDiskEncryptionKeyObservation struct {
}

type SourceDiskEncryptionKeyParameters struct {

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	// +kubebuilder:validation:Optional
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	RawKeySecretRef *v1.SecretKeySelector `json:"rawKeySecretRef,omitempty" tf:"-"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Snapshot is the Schema for the Snapshots API. Represents a Persistent Disk Snapshot resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}