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

type AccessDatasetObservation struct {
}

type AccessDatasetParameters struct {

	// The dataset this entry applies to
	// Structure is documented below.
	// +kubebuilder:validation:Required
	Dataset []DatasetDatasetParameters `json:"dataset" tf:"dataset,omitempty"`

	// Which resources in the dataset this entry applies to. Currently, only views are supported,
	// but additional target types may be added in the future. Possible values: VIEWS
	// +kubebuilder:validation:Required
	TargetTypes []*string `json:"targetTypes" tf:"target_types,omitempty"`
}

type AccessObservation struct {
}

type AccessParameters struct {

	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Dataset []AccessDatasetParameters `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// An email address of a Google Group to grant access to.
	// +kubebuilder:validation:Optional
	GroupByEmail *string `json:"groupByEmail,omitempty" tf:"group_by_email,omitempty"`

	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles
	// are supported. Predefined roles that have equivalent basic roles
	// are swapped by the API to their basic counterparts. See
	// official docs.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// A special group to grant access to. Possible values include:
	// +kubebuilder:validation:Optional
	SpecialGroup *string `json:"specialGroup,omitempty" tf:"special_group,omitempty"`

	// An email address of a user to grant access to. For example:
	// fred@example.com
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("email",true)
	// +kubebuilder:validation:Optional
	UserByEmail *string `json:"userByEmail,omitempty" tf:"user_by_email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate userByEmail.
	// +kubebuilder:validation:Optional
	UserByEmailRef *v1.Reference `json:"userByEmailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate userByEmail.
	// +kubebuilder:validation:Optional
	UserByEmailSelector *v1.Selector `json:"userByEmailSelector,omitempty" tf:"-"`

	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	View []ViewParameters `json:"view,omitempty" tf:"view,omitempty"`
}

type DatasetDatasetObservation struct {
}

type DatasetDatasetParameters struct {

	// The ID of the dataset containing this table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// The ID of the project containing this table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("project",false)
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Dataset in bigquery to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type DatasetObservation struct {

	// The time when this dataset was created, in milliseconds since the
	// epoch.
	CreationTime *float64 `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// A hash of the resource.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// an identifier for the resource with format projects/{{project}}/datasets/{{dataset_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date when this dataset or any of its tables was last modified, in
	// milliseconds since the epoch.
	LastModifiedTime *float64 `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type DatasetParameters struct {

	// An array of objects that define dataset access for one or more entities.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Access []AccessParameters `json:"access,omitempty" tf:"access,omitempty"`

	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DefaultEncryptionConfiguration []DefaultEncryptionConfigurationParameters `json:"defaultEncryptionConfiguration,omitempty" tf:"default_encryption_configuration,omitempty"`

	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	// +kubebuilder:validation:Optional
	DefaultPartitionExpirationMs *float64 `json:"defaultPartitionExpirationMs,omitempty" tf:"default_partition_expiration_ms,omitempty"`

	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	// +kubebuilder:validation:Optional
	DefaultTableExpirationMs *float64 `json:"defaultTableExpirationMs,omitempty" tf:"default_table_expiration_ms,omitempty"`

	// If set to true, delete all the tables in the
	// dataset when destroying the resource; otherwise,
	// destroying the resource will fail if tables are present.
	// +kubebuilder:validation:Optional
	DeleteContentsOnDestroy *bool `json:"deleteContentsOnDestroy,omitempty" tf:"delete_contents_on_destroy,omitempty"`

	// A user-friendly description of the dataset
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A descriptive name for the dataset
	// +kubebuilder:validation:Optional
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// The labels associated with this dataset. You can use these to
	// organize and group your datasets
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The geographic location where the dataset should reside.
	// See official docs.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DefaultEncryptionConfigurationObservation struct {
}

type DefaultEncryptionConfigurationParameters struct {

	// Describes the Cloud KMS encryption key that will be used to protect destination
	// BigQuery table. The BigQuery Service Account associated with your project requires
	// access to this encryption key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameRef *v1.Reference `json:"kmsKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameSelector *v1.Selector `json:"kmsKeyNameSelector,omitempty" tf:"-"`
}

type ViewObservation struct {
}

type ViewParameters struct {

	// The ID of the dataset containing this table.
	// +kubebuilder:validation:Required
	DatasetID *string `json:"datasetId" tf:"dataset_id,omitempty"`

	// The ID of the project containing this table.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// The ID of the table. The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 1,024 characters.
	// +kubebuilder:validation:Required
	TableID *string `json:"tableId" tf:"table_id,omitempty"`
}

// DatasetSpec defines the desired state of Dataset
type DatasetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetParameters `json:"forProvider"`
}

// DatasetStatus defines the observed state of Dataset.
type DatasetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dataset is the Schema for the Datasets API. Datasets allow you to organize and control access to your tables.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasetSpec   `json:"spec"`
	Status            DatasetStatus `json:"status,omitempty"`
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