// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package common

import (
	"strings"

	"github.com/pkg/errors"

	jresource "github.com/crossplane/upjet/pkg/resource"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// KeyProject is the key for project in Terraform Provider Configuration
	KeyProject = "project"
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/upbound/provider-gcp/config/common"

	// ExtractResourceIDFuncPath holds the GCP resource ID extractor func name
	ExtractResourceIDFuncPath = "github.com/upbound/provider-gcp/config/common.ExtractResourceID()"
	ExtractProjectIDFuncPath  = "github.com/upbound/provider-gcp/config/common.ExtractProjectID()"
	ExtractFolderIDFuncPath   = "github.com/upbound/provider-gcp/config/common.ExtractFolderID()"
	// VersionV1Beta1 is used for resources that meet the v1beta1 criteria
	// here: https://github.com/upbound/arch/pull/33
	VersionV1Beta1 = "v1beta1"
)

var (
	// PathSelfLinkExtractor is the golang path to SelfLinkExtractor function
	// in this package.
	PathSelfLinkExtractor = SelfPackagePath + ".SelfLinkExtractor()"
)

// SelfLinkExtractor extracts URI of the resources from
// "status.atProvider.selfLink" which is quite common among all GCP resources.
func SelfLinkExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.selfLink")
		if err != nil {
			return ""
		}
		return r
	}
}

// GetNameFromFullyQualifiedID extracts external-name from GCP ID
// Example: projects/{{project}}/zones/{{zone}}/instances/{{name}}
func GetNameFromFullyQualifiedID(tfstate map[string]interface{}) (string, error) {
	id, ok := tfstate["id"].(string)
	if !ok {
		return "", errors.New("cannot get the id field as string in tfstate")
	}
	words := strings.Split(id, "/")
	return words[len(words)-1], nil
}

// GetField returns the value of field as a string in a map[string]interface{},
//
//	fails properly otherwise.
func GetField(from map[string]interface{}, path string) (string, error) {
	return fieldpath.Pave(from).GetString(path)
}

// ExtractResourceID extracts the value of `spec.atProvider.id`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractResourceID() reference.ExtractValueFn {
	return func(mr resource.Managed) string {
		tr, ok := mr.(jresource.Terraformed)
		if !ok {
			return ""
		}
		return tr.GetID()
	}
}

func ExtractProjectID() reference.ExtractValueFn {
	return func(mr resource.Managed) string {
		tr, ok := mr.(jresource.Terraformed)
		if !ok {
			return ""
		}
		return strings.TrimPrefix(tr.GetID(), "projects/")
	}
}

func ExtractFolderID() reference.ExtractValueFn {
	return func(mr resource.Managed) string {
		tr, ok := mr.(jresource.Terraformed)
		if !ok {
			return ""
		}
		return strings.TrimPrefix(tr.GetID(), "folders/")
	}
}
