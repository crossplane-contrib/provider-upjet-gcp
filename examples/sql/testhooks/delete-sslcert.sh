#!/usr/bin/env bash
set -aeuo pipefail

# Note(turkenf): The SSLCert resource must be deleted before the DatabaseInstance
# resource to be successfully deleted. This is a workaround for this
# problem to make automated tests to work.
${KUBECTL} delete sslcert.sql.gcp.upbound.io -l testing.upbound.io/example-name=example_cert