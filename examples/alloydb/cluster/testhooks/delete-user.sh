#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

# Delete the User resource before deleting the Instance itself
${KUBECTL} delete user.alloydb.gcp.upbound.io --all
