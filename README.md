# Provider GCP

`provider-gcp` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the 
GCP API.

## Getting Started

Follow the quick start guide [here](https://marketplace.upbound.io/providers/upbound/provider-gcp/latest/docs/quickstart).
You can find a detailed API reference with all CRDs and examples [here](https://marketplace.upbound.io/providers/upbound/provider-gcp)

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build image:

```console
make image
```

Push image:

```console
make push
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/upbound/provider-gcp/issues).

## Contact

Please open a Github issue for all requests. If you need to reach out to Upbound, you can do so via the following channels:

* Slack: [#upbound](https://crossplane.slack.com/archives/C01TRKD4623) channel in [Crossplane Slack](https://slack.crossplane.io)
* Twitter: [@upbound_io](https://twitter.com/upbound_io)
* Email: [support@upbound.io](mailto:support@upbound.io)

## Licensing

Provider GCP is under [the Apache 2.0 license](LICENSE) with [notice](NOTICE).

