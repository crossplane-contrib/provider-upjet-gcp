#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

# Delete the RegistryRepositoryIAMMember resource before deleting the RegistryRepository itself
${KUBECTL} delete registryrepositoryiammember.artifact.gcp.upbound.io --all