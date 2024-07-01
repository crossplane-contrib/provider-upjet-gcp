// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package common

import (
	"context"
	"strings"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/password"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/upjet/pkg/config"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

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
	// ErrGetPasswordSecret is an error string for failing to get password secret
	ErrGetPasswordSecret = "cannot get password secret"
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

// PasswordGenerator  returns an InitializerFn that will generate a password
// for a resource if the toggle field is set to true and the secret referenced
// by the secretRefFieldPath is not found or does not have content corresponding
// to the password key.
func PasswordGenerator(secretRefFieldPath, toggleFieldPath string) config.NewInitializerFn { //nolint:gocyclo
	return func(client client.Client) managed.Initializer {
		return managed.InitializerFn(func(ctx context.Context, mg resource.Managed) error {
			paved, err := fieldpath.PaveObject(mg)
			if err != nil {
				return errors.Wrap(err, "cannot convert runtime object to unstructured object")
			}
			sel := &v1.SecretKeySelector{}
			if err := paved.GetValueInto(secretRefFieldPath, sel); err != nil {
				return errors.Wrapf(resource.Ignore(fieldpath.IsNotFound, err), "cannot unmarshal %s into a secret key selector", secretRefFieldPath)
			}
			s := &corev1.Secret{}
			if err := client.Get(ctx, types.NamespacedName{Namespace: sel.Namespace, Name: sel.Name}, s); resource.IgnoreNotFound(err) != nil {
				return errors.Wrap(err, ErrGetPasswordSecret)
			}
			if err == nil && len(s.Data[sel.Key]) != 0 {
				// Password is already set.
				return nil
			}
			// At this point, either the secret doesn't exist, or it doesn't
			// have the password filled.
			if gen, err := paved.GetBool(toggleFieldPath); err != nil || !gen {
				// If there is error, then we return that.
				// If the toggle field is not set to true, then we return nil.
				// Because we don't want to generate a password if the user
				// doesn't want to.
				return errors.Wrapf(resource.Ignore(fieldpath.IsNotFound, err), "cannot get the value of %s", toggleFieldPath)
			}
			pw, err := password.Generate()
			if err != nil {
				return errors.Wrap(err, "cannot generate password")
			}
			s.SetName(sel.Name)
			s.SetNamespace(sel.Namespace)
			if !meta.WasCreated(s) {
				// We don't want to own the Secret if it is created by someone
				// else, otherwise the deletion of the managed resource will
				// delete the Secret that we didn't create in the first place.
				meta.AddOwnerReference(s, meta.AsOwner(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind())))
			}
			if s.Data == nil {
				s.Data = make(map[string][]byte, 1)
			}
			s.Data[sel.Key] = []byte(pw)
			return errors.Wrap(resource.NewAPIPatchingApplicator(client).Apply(ctx, s), "cannot apply password secret")
		})
	}
}
