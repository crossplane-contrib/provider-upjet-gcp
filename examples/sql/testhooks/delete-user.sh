#!/usr/bin/env bash
set -aeuo pipefail

# Note(turkenf): We are getting (Invalid request since instance is not running) exception if user for
# the instance got deleted before the user resource. This is a workaround for this
# problem to make automated tests to work.
${KUBECTL} delete user.sql.gcp.upbound.io -l testing.upbound.io/example-name=example_user