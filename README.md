# Provider GCP

`provider-gcp` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the 
GCP API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/upbound/provider-gcp/releases):
```
kubectl crossplane install provider upbound/provider-gcp:v0.1.0
```

You can see the API reference [here](https://doc.crds.dev/github.com/upbound/provider-gcp).

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

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://slack.crossplane.io)
* Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
* Email: [info@crossplane.io](mailto:info@crossplane.io)

## Governance and Owners

provider-gcp is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-gcp adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-gcp is under the Apache 2.0 license.

