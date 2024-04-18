<!--
SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>

SPDX-License-Identifier: CC-BY-4.0
-->

# Official Provider GCP

<div align="center">

![CI](https://github.com/crossplane-contrib/provider-upjet-gcp/workflows/CI/badge.svg)
[![GitHub release](https://img.shields.io/github/release/crossplane-contrib/provider-upjet-gcp/all.svg)](https://github.com/crossplane-contrib/provider-upjet-gcp/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/crossplane-contrib/provider-upjet-gcp)](https://goreportcard.com/report/github.com/crossplane-contrib/provider-upjet-gcp)
[![Contributors](https://img.shields.io/github/contributors/crossplane-contrib/provider-upjet-gcp)](https://github.com/crossplane-contrib/provider-upjet-gcp/graphs/contributors)
[![Slack](https://img.shields.io/badge/Slack-4A154B?logo=slack)](https://crossplane.slack.com/archives/C05E7EVM459)
[![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/crossplane_io)](https://twitter.com/crossplane_io)

</div>

Provider GCP is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the 
GCP API.

## Getting Started

Follow the quick start guide [here](https://marketplace.upbound.io/providers/upbound/provider-gcp/latest/docs/quickstart).
You can find a detailed API reference for all the managed resources with examples in the [Upbound Marketplace](https://marketplace.upbound.io/providers/upbound/provider-gcp/latest/managed-resources).

For getting more information about resource consumption and monitoring
the upjet runtime, please see [Sizing Guide](https://github.com/crossplane/upjet/blob/v0.10.0/docs/sizing-guide.md)
and [Monitoring Guide](https://github.com/crossplane/upjet/blob/main/docs/monitoring.md)

## Contributing

For the general contribution guide, see [Upjet Contribution Guide](https://github.com/crossplane/upjet/blob/main/CONTRIBUTING.md)

If you'd like to learn how to use Upjet, see [Usage Guide](https://github.com/crossplane/upjet/tree/main/docs).

### Add a New Resource

Follow the guide [here](https://github.com/crossplane/upjet/blob/v0.10.0/docs/add-new-resource-short.md).

### Local Development

1. Check out the provider repo, crossplane-contrib/provider-upjet-gcp, and go to the project
directory on your local machine.

2. Do the necessary changes locally and please make sure you have comitted all of them.

3. Run the `make load-pkg` command for family providers as follows.

If you want to build any of the family resource providers (e.g., `provider-gcp-cloudplatform`, `provider-gcp-container`),  you can do:

```shell
make load-pkg PROVIDERS="cloudplatform container" REPO="index.docker.io/<your-dockerhub-org>"
```
*Note: Installable images for the family config provider will be created by default.*

For more information, follow the guide [here](https://marketplace.upbound.io/providers/upbound/provider-family-gcp/latest/docs/development).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-upjet-gcp/issues/new/choose).

## Contact

[#upjet-provider-gcp](https://crossplane.slack.com/archives/C05E7EVM459) channel in
[Crossplane Slack](https://slack.crossplane.io)

## Licensing

Provider GCP is under [the Apache 2.0 license](LICENSE) with [notice](NOTICE).

