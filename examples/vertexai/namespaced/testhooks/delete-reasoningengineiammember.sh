#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

# Delete the ReasoningEngineIAMMember resource before deleting the ReasoningEngine itself
${KUBECTL} delete reasoningengineiammember.vertexai.gcp.m.upbound.io --all --all-namespaces
