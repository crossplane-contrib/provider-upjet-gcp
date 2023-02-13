#!/usr/bin/env bash
set -aeuo pipefail

# Delete the DatasetIAMMember resource before deleting the Dataset itself
${KUBECTL} delete datasetiammember.healthcare.gcp.upbound.io --all