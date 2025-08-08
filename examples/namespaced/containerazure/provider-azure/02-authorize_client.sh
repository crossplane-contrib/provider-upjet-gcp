#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

AZURE_CLIENT=$1

export APPLICATION_ID=$(az ad app list --all \
 --query "[?displayName=='containerazure-gcp-upbound'].appId" \
 --output tsv)

export CERT=`gcloud container azure clients get-public-cert --location=us-west1 ${AZURE_CLIENT}`

az ad app credential reset --id "${APPLICATION_ID}" --cert "${CERT}" --append

### Create an SSH Key
ssh-keygen -m PEM -t rsa -b 4096 -f ./${AZURE_CLIENT}
