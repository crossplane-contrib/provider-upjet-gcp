#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

UPTEST_GCP_PROJECT=${UPTEST_GCP_PROJECT:-official-provider-testing}

echo "Running setup.sh"

if [[ -n "${UPTEST_CLOUD_CREDENTIALS:-}" ]]; then
  echo "Creating cloud credential secret..."
  ${KUBECTL} -n upbound-system create secret generic provider-secret --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

  echo "Creating a default provider config..."
  cat <<EOF | ${KUBECTL} apply -f -
apiVersion: gcp.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  projectID: ${UPTEST_GCP_PROJECT}
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: upbound-system
      key: credentials
EOF

  echo "Creating a default cluster provider config..."
  cat <<EOF | ${KUBECTL} apply -f -
apiVersion: gcp.upbound.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  projectID: ${UPTEST_GCP_PROJECT}
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: upbound-system
      key: credentials
EOF
fi