#!/usr/bin/env bash
set -aeuo pipefail

# Delete the RegistryRepositoryIAMMember resource before deleting the RegistryRepository itself
${KUBECTL} delete registryrepositoryiammember.artifact.gcp.upbound.io --all