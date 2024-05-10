---
title: Development
weight: 3
---

# Building the Official Provider Family Images Locally

1. Check out the provider repo, crossplane-contrib/provider-upjet-gcp, and go to the project
directory on your local machine.

2. Do the necessary changes locally and please make sure you have comitted all of them.

3. Run the `make load-pkg` command for family providers as follows.

If you want to build any of the family resource providers (e.g., `provider-gcp-cloudplatform`, `provider-gcp-container`),  you can do:

```shell
make load-pkg PROVIDERS="cloudplatform container" REPO="index.docker.io/<your-dockerhub-org>"
```
*Note: Installable images for the family config provider will be created by default.*

After the command is completed, you will see the installable images for `provider-gcp-cloudplatform`, `provider-gcp-container`
and family config provider `provider-family-gcp`:

```shell
> docker image ls
REPOSITORY                                 TAG                         IMAGE ID       CREATED        SIZE
turkenf/provider-family-gcp-amd64          v1.1.0-rc.0.42.g42cf069b9   baac1f7e697d   6 minutes ago  179MB
turkenf/provider-gcp-cloudplatform-amd64   v1.1.0-rc.0.42.g42cf069b9   fe2106877bf9   6 minutes ago  181MB
turkenf/provider-gcp-container-amd64       v1.1.0-rc.0.42.g42cf069b9   94f8d11c2709   6 minutes ago  181MB
turkenf/provider-gcp-cloudplatform-arm64   v1.1.0-rc.0.42.g42cf069b9   edbf18e9d01f   6 minutes ago  176MB
turkenf/provider-gcp-container-arm64       v1.1.0-rc.0.42.g42cf069b9   0d3ab205ef66   6 minutes ago  176MB
turkenf/provider-family-gcp-arm64          v1.1.0-rc.0.42.g42cf069b9   a6ef52d83555   6 minutes ago  175MB
```

## Running Providers with Locally Built Images

One way to install locally built images is pushing them to Dockerhub and installing from there. For example, if you want to 
install the `cloudplatform` package on a arm64 arch:

1. Push provider images to Dockerhub:

```shell
# push the images to your dockerhub org
docker push <your-dockerhub-org>/provider-family-gcp-arm64:v1.1.0-rc.0.42.g42cf069b9
docker push <your-dockerhub-org>/provider-gcp-cloudplatform-arm64:v1.1.0-rc.0.42.g42cf069b9
```

2. Install the provider packages by setting `skipDependencyResolution: true` with the following yaml:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
    name: upbound-provider-gcp-config
spec:
    package: index.docker.io/turkenf/provider-family-gcp-arm64:v1.1.0-rc.0.42.g42cf069b9
    skipDependencyResolution: true
---
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
    name: upbound-provider-gcp-cloudplatform
spec:
    package: index.docker.io/turkenf/provider-gcp-cloudplatform-arm64:v1.1.0-rc.0.42.g42cf069b9
    skipDependencyResolution: true
```

3. Verify that providers are installed and healthy:

```shell
> kubectl get providers.pkg.crossplane.io
NAME                                 INSTALLED   HEALTHY   PACKAGE                                                                              AGE
upbound-provider-gcp-cloudplatform   True        True      index.docker.io/turkenf/provider-gcp-cloudplatform-arm64:v1.1.0-rc.0.42.g42cf069b9   115s
upbound-provider-gcp-config          True        True      index.docker.io/turkenf/provider-family-gcp-arm64:v1.1.0-rc.0.42.g42cf069b9          115s
```

*Note: If you don't want to set `skipDependencyResolution: true`, please follow the steps below:*

- Tag the family config provider image as provider-family-gcp

```shell
# get rid of the arch to resolve package dependencies 
docker tag <your-dockerhub-org>/provider-family-gcp-arm64:v1.1.0-rc.0.42.g42cf069b9 <your-dockerhub-org>/provider-family-gcp:v1.1.0-rc.0.42.g42cf069b9
# push the images to your dockerhub org
docker push <your-dockerhub-org>/provider-family-gcp:v1.1.0-rc.0.42.g42cf069b9
docker push <your-dockerhub-org>/provider-gcp-cloudplatform-arm64:v1.1.0-rc.0.42.g42cf069b9
```

- Install the provider package:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
    name: upbound-provider-gcp-cloudplatform
spec:
    package: index.docker.io/turkenf/provider-gcp-cloudplatform-arm64:v1.1.0-rc.0.42.g42cf069b9
```

- Verify that providers are installed and healthy:

```shell
> kubectl get providers.pkg.crossplane.io
NAME                                 INSTALLED   HEALTHY   PACKAGE                                                                              AGE
turkenf-provider-family-gcp          True        True      index.docker.io/turkenf/provider-family-gcp:v1.1.0-rc.0.42.g42cf069b9                73s
upbound-provider-gcp-cloudplatform   True        True      index.docker.io/turkenf/provider-gcp-cloudplatform-arm64:v1.1.0-rc.0.42.g42cf069b9   2m1s
```