// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cryptokey "github.com/upbound/provider-gcp/internal/controller/kms/cryptokey"
	cryptokeyiammember "github.com/upbound/provider-gcp/internal/controller/kms/cryptokeyiammember"
	cryptokeyversion "github.com/upbound/provider-gcp/internal/controller/kms/cryptokeyversion"
	keyring "github.com/upbound/provider-gcp/internal/controller/kms/keyring"
	keyringiammember "github.com/upbound/provider-gcp/internal/controller/kms/keyringiammember"
	keyringimportjob "github.com/upbound/provider-gcp/internal/controller/kms/keyringimportjob"
	secretciphertext "github.com/upbound/provider-gcp/internal/controller/kms/secretciphertext"
)

// Setup_kms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cryptokey.Setup,
		cryptokeyiammember.Setup,
		cryptokeyversion.Setup,
		keyring.Setup,
		keyringiammember.Setup,
		keyringimportjob.Setup,
		secretciphertext.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
