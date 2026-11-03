#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

# Delete the ReasoningEngineIAMPolicy resource before deleting the ReasoningEngine itself
${KUBECTL} delete reasoningengineiampolicy.vertexai.gcp.upbound.io --all
