// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.

package v1beta2

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this InboundSAMLConfig.
func (mg *InboundSAMLConfig) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDefaultConfig.
func (mg *ProjectDefaultConfig) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this TenantInboundSAMLConfig.
func (mg *TenantInboundSAMLConfig) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}