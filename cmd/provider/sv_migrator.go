// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	authv1 "k8s.io/api/authorization/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const rootGroup = "gcp.upbound.io"

// RunStorageVersionMigration runs storage version migration scoped to
// shortGroup (e.g. "compute"), targeting only CRDs in
// "<shortGroup>.gcp.upbound.io". Pass an empty shortGroup to migrate all CRDs
// under the root group (used by the monolith binary).
func RunStorageVersionMigration(ctx context.Context, logr logging.Logger, mgr manager.Manager, shortGroup string) error {
	// authv1 is required for the permission check inside CRDMigrator but is
	// not yet registered in the scheme at the point this is called.
	if err := authv1.AddToScheme(mgr.GetScheme()); err != nil {
		return errors.Wrap(err, "failed to add authv1 to scheme")
	}

	kubeClient, err := client.New(mgr.GetConfig(), client.Options{Scheme: mgr.GetScheme()})
	if err != nil {
		return errors.Wrap(err, "failed to create kube client for storage version migration")
	}

	var opts []ujconfig.CRDMigratorOption
	if shortGroup != "" {
		opts = append(opts, ujconfig.WithShortGroup(shortGroup))
	}

	migrator := ujconfig.NewCRDMigrator(rootGroup, opts...)
	return errors.Wrap(migrator.Run(ctx, logr, kubeClient), "failed to run storage version migrator")
}
