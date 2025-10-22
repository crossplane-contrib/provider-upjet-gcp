#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

# terraform-provider-azurerm does not support ad applications 
az ad app create --display-name containerazure-gcp-upbound
export APPLICATION_ID=$(az ad app list --all \
 --query "[?displayName=='containerazure-gcp-upbound'].appId" \
 --output tsv)
# terraform-provider-azurerm does not support ad service principals 
az ad sp create --id "${APPLICATION_ID}"

# Generate the roles manifests
export SERVICE_PRINCIPAL_ID=$(az ad sp list --all  --output tsv \
    --query "[?displayName=='containerazure-gcp-upbound'].{id:id}")
export SUBSCRIPTION_ID=$(az account show --query "id" --output tsv)

envsubst '${APPLICATION_ID},${SERVICE_PRINCIPAL_ID},${SUBSCRIPTION_ID}' < 02-containerazure-roles.yaml.tmpl > 02-containerazure-roles.yaml
