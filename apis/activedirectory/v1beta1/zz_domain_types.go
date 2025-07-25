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

type DomainInitParameters struct {

	// The name of delegated administrator account used to perform Active Directory operations.
	// If not specified, setupadmin will be used.
	Admin *string `json:"admin,omitempty" tf:"admin,omitempty"`

	// The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	// +listType=set
	AuthorizedNetworks []*string `json:"authorizedNetworks,omitempty" tf:"authorized_networks,omitempty"`

	// Defaults to true.
	// When the field is set to false, deleting the domain is allowed.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions
	// of https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Resource labels that can contain user-provided metadata
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	ReservedIPRange *string `json:"reservedIpRange,omitempty" tf:"reserved_ip_range,omitempty"`
}

type DomainObservation struct {

	// The name of delegated administrator account used to perform Active Directory operations.
	// If not specified, setupadmin will be used.
	Admin *string `json:"admin,omitempty" tf:"admin,omitempty"`

	// The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	// +listType=set
	AuthorizedNetworks []*string `json:"authorizedNetworks,omitempty" tf:"authorized_networks,omitempty"`

	// Defaults to true.
	// When the field is set to false, deleting the domain is allowed.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions
	// of https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// The fully-qualified domain name of the exposed domain used by clients to connect to the service.
	// Similar to what would be chosen for an Active Directory set up on an internal network.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource labels that can contain user-provided metadata
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`

	// The unique name of the domain using the format: projects/{project}/locations/global/domains/{domainName}.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	ReservedIPRange *string `json:"reservedIpRange,omitempty" tf:"reserved_ip_range,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`
}

type DomainParameters struct {

	// The name of delegated administrator account used to perform Active Directory operations.
	// If not specified, setupadmin will be used.
	// +kubebuilder:validation:Optional
	Admin *string `json:"admin,omitempty" tf:"admin,omitempty"`

	// The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	// +kubebuilder:validation:Optional
	// +listType=set
	AuthorizedNetworks []*string `json:"authorizedNetworks,omitempty" tf:"authorized_networks,omitempty"`

	// Defaults to true.
	// When the field is set to false, deleting the domain is allowed.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions
	// of https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Resource labels that can contain user-provided metadata
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	// +kubebuilder:validation:Optional
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	// +kubebuilder:validation:Optional
	ReservedIPRange *string `json:"reservedIpRange,omitempty" tf:"reserved_ip_range,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
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
	InitProvider DomainInitParameters `json:"initProvider,omitempty"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Domain is the Schema for the Domains API. Creates a Microsoft AD domain
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domainName) || (has(self.initProvider) && has(self.initProvider.domainName))",message="spec.forProvider.domainName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.locations) || (has(self.initProvider) && has(self.initProvider.locations))",message="spec.forProvider.locations is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.reservedIpRange) || (has(self.initProvider) && has(self.initProvider.reservedIpRange))",message="spec.forProvider.reservedIpRange is a required parameter"
	Spec   DomainSpec   `json:"spec"`
	Status DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
