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

type AnalyticsHubListingInitParameters struct {

	// Shared dataset i.e. BigQuery dataset source.
	// Structure is documented below.
	BigqueryDataset []BigqueryDatasetInitParameters `json:"bigqueryDataset,omitempty" tf:"bigquery_dataset,omitempty"`

	// Categories of the listing. Up to two categories are allowed.
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// Details of the data provider who owns the source data.
	// Structure is documented below.
	DataProvider []DataProviderInitParameters `json:"dataProvider,omitempty" tf:"data_provider,omitempty"`

	// Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Documentation describing the listing.
	Documentation *string `json:"documentation,omitempty" tf:"documentation,omitempty"`

	// Base64 encoded image representing the listing.
	Icon *string `json:"icon,omitempty" tf:"icon,omitempty"`

	// If true, subscriber email logging is enabled and all queries on the linked dataset will log the email address of the querying user. Once enabled, this setting cannot be turned off.
	LogLinkedDatasetQueryUserEmail *bool `json:"logLinkedDatasetQueryUserEmail,omitempty" tf:"log_linked_dataset_query_user_email,omitempty"`

	// Email or URL of the listing publisher.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Details of the publisher who owns the listing and who can share the source data.
	// Structure is documented below.
	Publisher []PublisherInitParameters `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// Pub/Sub topic source.
	// Structure is documented below.
	PubsubTopic []PubsubTopicInitParameters `json:"pubsubTopic,omitempty" tf:"pubsub_topic,omitempty"`

	// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
	RequestAccess *string `json:"requestAccess,omitempty" tf:"request_access,omitempty"`

	// If set, restricted export configuration will be propagated and enforced on the linked dataset.
	// Structure is documented below.
	RestrictedExportConfig []RestrictedExportConfigInitParameters `json:"restrictedExportConfig,omitempty" tf:"restricted_export_config,omitempty"`
}

type AnalyticsHubListingObservation struct {

	// Shared dataset i.e. BigQuery dataset source.
	// Structure is documented below.
	BigqueryDataset []BigqueryDatasetObservation `json:"bigqueryDataset,omitempty" tf:"bigquery_dataset,omitempty"`

	// Categories of the listing. Up to two categories are allowed.
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeID *string `json:"dataExchangeId,omitempty" tf:"data_exchange_id,omitempty"`

	// Details of the data provider who owns the source data.
	// Structure is documented below.
	DataProvider []DataProviderObservation `json:"dataProvider,omitempty" tf:"data_provider,omitempty"`

	// Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Documentation describing the listing.
	Documentation *string `json:"documentation,omitempty" tf:"documentation,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Base64 encoded image representing the listing.
	Icon *string `json:"icon,omitempty" tf:"icon,omitempty"`

	// The name of the location this data exchange listing.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// If true, subscriber email logging is enabled and all queries on the linked dataset will log the email address of the querying user. Once enabled, this setting cannot be turned off.
	LogLinkedDatasetQueryUserEmail *bool `json:"logLinkedDatasetQueryUserEmail,omitempty" tf:"log_linked_dataset_query_user_email,omitempty"`

	// The resource name of the listing. e.g. "projects/myproject/locations/US/dataExchanges/123/listings/456"
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email or URL of the listing publisher.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Details of the publisher who owns the listing and who can share the source data.
	// Structure is documented below.
	Publisher []PublisherObservation `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// Pub/Sub topic source.
	// Structure is documented below.
	PubsubTopic []PubsubTopicObservation `json:"pubsubTopic,omitempty" tf:"pubsub_topic,omitempty"`

	// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
	RequestAccess *string `json:"requestAccess,omitempty" tf:"request_access,omitempty"`

	// If set, restricted export configuration will be propagated and enforced on the linked dataset.
	// Structure is documented below.
	RestrictedExportConfig []RestrictedExportConfigObservation `json:"restrictedExportConfig,omitempty" tf:"restricted_export_config,omitempty"`
}

type AnalyticsHubListingParameters struct {

	// Shared dataset i.e. BigQuery dataset source.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BigqueryDataset []BigqueryDatasetParameters `json:"bigqueryDataset,omitempty" tf:"bigquery_dataset,omitempty"`

	// Categories of the listing. Up to two categories are allowed.
	// +kubebuilder:validation:Optional
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.AnalyticsHubDataExchange
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("data_exchange_id",false)
	// +kubebuilder:validation:Optional
	DataExchangeID *string `json:"dataExchangeId,omitempty" tf:"data_exchange_id,omitempty"`

	// Reference to a AnalyticsHubDataExchange in bigquery to populate dataExchangeId.
	// +kubebuilder:validation:Optional
	DataExchangeIDRef *v1.Reference `json:"dataExchangeIdRef,omitempty" tf:"-"`

	// Selector for a AnalyticsHubDataExchange in bigquery to populate dataExchangeId.
	// +kubebuilder:validation:Optional
	DataExchangeIDSelector *v1.Selector `json:"dataExchangeIdSelector,omitempty" tf:"-"`

	// Details of the data provider who owns the source data.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DataProvider []DataProviderParameters `json:"dataProvider,omitempty" tf:"data_provider,omitempty"`

	// Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Documentation describing the listing.
	// +kubebuilder:validation:Optional
	Documentation *string `json:"documentation,omitempty" tf:"documentation,omitempty"`

	// Base64 encoded image representing the listing.
	// +kubebuilder:validation:Optional
	Icon *string `json:"icon,omitempty" tf:"icon,omitempty"`

	// The name of the location this data exchange listing.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// If true, subscriber email logging is enabled and all queries on the linked dataset will log the email address of the querying user. Once enabled, this setting cannot be turned off.
	// +kubebuilder:validation:Optional
	LogLinkedDatasetQueryUserEmail *bool `json:"logLinkedDatasetQueryUserEmail,omitempty" tf:"log_linked_dataset_query_user_email,omitempty"`

	// Email or URL of the listing publisher.
	// +kubebuilder:validation:Optional
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Details of the publisher who owns the listing and who can share the source data.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Publisher []PublisherParameters `json:"publisher,omitempty" tf:"publisher,omitempty"`

	// Pub/Sub topic source.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PubsubTopic []PubsubTopicParameters `json:"pubsubTopic,omitempty" tf:"pubsub_topic,omitempty"`

	// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
	// +kubebuilder:validation:Optional
	RequestAccess *string `json:"requestAccess,omitempty" tf:"request_access,omitempty"`

	// If set, restricted export configuration will be propagated and enforced on the linked dataset.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RestrictedExportConfig []RestrictedExportConfigParameters `json:"restrictedExportConfig,omitempty" tf:"restricted_export_config,omitempty"`
}

type BigqueryDatasetInitParameters struct {

	// Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// Reference to a Dataset in bigquery to populate dataset.
	// +kubebuilder:validation:Optional
	DatasetRef *v1.Reference `json:"datasetRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate dataset.
	// +kubebuilder:validation:Optional
	DatasetSelector *v1.Selector `json:"datasetSelector,omitempty" tf:"-"`

	// Resource in this dataset that is selectively shared. This field is required for data clean room exchanges.
	// Structure is documented below.
	SelectedResources []SelectedResourcesInitParameters `json:"selectedResources,omitempty" tf:"selected_resources,omitempty"`
}

type BigqueryDatasetObservation struct {

	// Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// Resource in this dataset that is selectively shared. This field is required for data clean room exchanges.
	// Structure is documented below.
	SelectedResources []SelectedResourcesObservation `json:"selectedResources,omitempty" tf:"selected_resources,omitempty"`
}

type BigqueryDatasetParameters struct {

	// Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// Reference to a Dataset in bigquery to populate dataset.
	// +kubebuilder:validation:Optional
	DatasetRef *v1.Reference `json:"datasetRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate dataset.
	// +kubebuilder:validation:Optional
	DatasetSelector *v1.Selector `json:"datasetSelector,omitempty" tf:"-"`

	// Resource in this dataset that is selectively shared. This field is required for data clean room exchanges.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SelectedResources []SelectedResourcesParameters `json:"selectedResources,omitempty" tf:"selected_resources,omitempty"`
}

type DataProviderInitParameters struct {

	// Name of the data provider.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email or URL of the data provider.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

type DataProviderObservation struct {

	// Name of the data provider.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email or URL of the data provider.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

type DataProviderParameters struct {

	// Name of the data provider.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Email or URL of the data provider.
	// +kubebuilder:validation:Optional
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

type PublisherInitParameters struct {

	// Name of the listing publisher.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email or URL of the listing publisher.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

type PublisherObservation struct {

	// Name of the listing publisher.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email or URL of the listing publisher.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

type PublisherParameters struct {

	// Name of the listing publisher.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Email or URL of the listing publisher.
	// +kubebuilder:validation:Optional
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`
}

type PubsubTopicInitParameters struct {

	// Region hint on where the data might be published. Data affinity regions are modifiable.
	// See https://cloud.google.com/about/locations for full listing of possible Cloud regions.
	// +listType=set
	DataAffinityRegions []*string `json:"dataAffinityRegions,omitempty" tf:"data_affinity_regions,omitempty"`

	// Resource name of the Pub/Sub topic source for this listing. e.g. projects/myproject/topics/topicId
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/pubsub/v1beta2.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// Reference to a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicRef *v1.Reference `json:"topicRef,omitempty" tf:"-"`

	// Selector for a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicSelector *v1.Selector `json:"topicSelector,omitempty" tf:"-"`
}

type PubsubTopicObservation struct {

	// Region hint on where the data might be published. Data affinity regions are modifiable.
	// See https://cloud.google.com/about/locations for full listing of possible Cloud regions.
	// +listType=set
	DataAffinityRegions []*string `json:"dataAffinityRegions,omitempty" tf:"data_affinity_regions,omitempty"`

	// Resource name of the Pub/Sub topic source for this listing. e.g. projects/myproject/topics/topicId
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type PubsubTopicParameters struct {

	// Region hint on where the data might be published. Data affinity regions are modifiable.
	// See https://cloud.google.com/about/locations for full listing of possible Cloud regions.
	// +kubebuilder:validation:Optional
	// +listType=set
	DataAffinityRegions []*string `json:"dataAffinityRegions,omitempty" tf:"data_affinity_regions,omitempty"`

	// Resource name of the Pub/Sub topic source for this listing. e.g. projects/myproject/topics/topicId
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/pubsub/v1beta2.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// Reference to a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicRef *v1.Reference `json:"topicRef,omitempty" tf:"-"`

	// Selector for a Topic in pubsub to populate topic.
	// +kubebuilder:validation:Optional
	TopicSelector *v1.Selector `json:"topicSelector,omitempty" tf:"-"`
}

type RestrictedExportConfigInitParameters struct {

	// If true, enable restricted export.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// If true, restrict export of query result derived from restricted linked dataset table.
	RestrictQueryResult *bool `json:"restrictQueryResult,omitempty" tf:"restrict_query_result,omitempty"`
}

type RestrictedExportConfigObservation struct {

	// If true, enable restricted export.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Output)
	// If true, restrict direct table access(read api/tabledata.list) on linked table.
	RestrictDirectTableAccess *bool `json:"restrictDirectTableAccess,omitempty" tf:"restrict_direct_table_access,omitempty"`

	// If true, restrict export of query result derived from restricted linked dataset table.
	RestrictQueryResult *bool `json:"restrictQueryResult,omitempty" tf:"restrict_query_result,omitempty"`
}

type RestrictedExportConfigParameters struct {

	// If true, enable restricted export.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// If true, restrict export of query result derived from restricted linked dataset table.
	// +kubebuilder:validation:Optional
	RestrictQueryResult *bool `json:"restrictQueryResult,omitempty" tf:"restrict_query_result,omitempty"`
}

type SelectedResourcesInitParameters struct {

	// Format: For table: projects/{projectId}/datasets/{datasetId}/tables/{tableId} Example:"projects/test_project/datasets/test_dataset/tables/test_table"
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Table
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Reference to a Table in bigquery to populate table.
	// +kubebuilder:validation:Optional
	TableRef *v1.Reference `json:"tableRef,omitempty" tf:"-"`

	// Selector for a Table in bigquery to populate table.
	// +kubebuilder:validation:Optional
	TableSelector *v1.Selector `json:"tableSelector,omitempty" tf:"-"`
}

type SelectedResourcesObservation struct {

	// Format: For table: projects/{projectId}/datasets/{datasetId}/tables/{tableId} Example:"projects/test_project/datasets/test_dataset/tables/test_table"
	Table *string `json:"table,omitempty" tf:"table,omitempty"`
}

type SelectedResourcesParameters struct {

	// Format: For table: projects/{projectId}/datasets/{datasetId}/tables/{tableId} Example:"projects/test_project/datasets/test_dataset/tables/test_table"
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta2.Table
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Reference to a Table in bigquery to populate table.
	// +kubebuilder:validation:Optional
	TableRef *v1.Reference `json:"tableRef,omitempty" tf:"-"`

	// Selector for a Table in bigquery to populate table.
	// +kubebuilder:validation:Optional
	TableSelector *v1.Selector `json:"tableSelector,omitempty" tf:"-"`
}

// AnalyticsHubListingSpec defines the desired state of AnalyticsHubListing
type AnalyticsHubListingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyticsHubListingParameters `json:"forProvider"`
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
	InitProvider AnalyticsHubListingInitParameters `json:"initProvider,omitempty"`
}

// AnalyticsHubListingStatus defines the observed state of AnalyticsHubListing.
type AnalyticsHubListingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyticsHubListingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AnalyticsHubListing is the Schema for the AnalyticsHubListings API. A Bigquery Analytics Hub data exchange listing
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AnalyticsHubListing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bigqueryDataset) || (has(self.initProvider) && has(self.initProvider.bigqueryDataset))",message="spec.forProvider.bigqueryDataset is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	Spec   AnalyticsHubListingSpec   `json:"spec"`
	Status AnalyticsHubListingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsHubListingList contains a list of AnalyticsHubListings
type AnalyticsHubListingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyticsHubListing `json:"items"`
}

// Repository type metadata.
var (
	AnalyticsHubListing_Kind             = "AnalyticsHubListing"
	AnalyticsHubListing_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AnalyticsHubListing_Kind}.String()
	AnalyticsHubListing_KindAPIVersion   = AnalyticsHubListing_Kind + "." + CRDGroupVersion.String()
	AnalyticsHubListing_GroupVersionKind = CRDGroupVersion.WithKind(AnalyticsHubListing_Kind)
)

func init() {
	SchemeBuilder.Register(&AnalyticsHubListing{}, &AnalyticsHubListingList{})
}
