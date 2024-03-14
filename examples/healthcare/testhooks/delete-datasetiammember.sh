#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

# Delete the DatasetIAMMember resource before deleting the Dataset itself
${KUBECTL} delete datasetiammember.healthcare.gcp.upbound.io --all