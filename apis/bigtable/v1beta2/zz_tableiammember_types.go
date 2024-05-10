// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TableIAMMemberConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TableIAMMemberConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TableIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type TableIAMMemberInitParameters struct {
	Condition *TableIAMMemberConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type TableIAMMemberObservation struct {
	Condition *TableIAMMemberConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Table *string `json:"table,omitempty" tf:"table,omitempty"`
}

type TableIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition *TableIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	Table *string `json:"table" tf:"table,omitempty"`
}

// TableIAMMemberSpec defines the desired state of TableIAMMember
type TableIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableIAMMemberParameters `json:"forProvider"`
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
	InitProvider TableIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// TableIAMMemberStatus defines the observed state of TableIAMMember.
type TableIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TableIAMMember is the Schema for the TableIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TableIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instance) || (has(self.initProvider) && has(self.initProvider.instance))",message="spec.forProvider.instance is a required parameter"
	Spec   TableIAMMemberSpec   `json:"spec"`
	Status TableIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMMemberList contains a list of TableIAMMembers
type TableIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableIAMMember `json:"items"`
}

// Repository type metadata.
var (
	TableIAMMember_Kind             = "TableIAMMember"
	TableIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableIAMMember_Kind}.String()
	TableIAMMember_KindAPIVersion   = TableIAMMember_Kind + "." + CRDGroupVersion.String()
	TableIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(TableIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&TableIAMMember{}, &TableIAMMemberList{})
}