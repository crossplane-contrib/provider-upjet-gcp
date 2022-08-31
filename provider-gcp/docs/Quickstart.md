# Quickstart

This guide walks through the process to install Upbound Universal Crossplane and install the GCP official provider.

To use this official provider, install Upbound Universal Crossplane into your Kubernetes cluster, install the `Provider`, apply a `ProviderConfig`, and create a *managed resource* in GCP via Kubernetes.

## Create an Upbound.io user account
Create an account on [Upbound.io](https://accounts.upbound.io/register). 

<!-- Find detailed instructions in the [account documentation](/getting-started/create-account). -->

## Install the Up command-line
Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
mv up /usr/local/bin/
```

Verify the version of `up` with `up --version`

```shell
$ up --version
v0.13.0
```

_Note_: official providers only support `up` command-line versions v0.13.0 or later.

## Install Universal Crossplane
Install Upbound Universal Crossplane with the Up command-line.

```shell
$ up uxp install
UXP 1.9.0-up.3 installed
```

Verify the UXP pods are running with `kubectl get pods -n upbound-system`

```shell
$ kubectl get pods -n upbound-system
NAME                                        READY   STATUS    RESTARTS      AGE
crossplane-7fdfbd897c-pmrml                 1/1     Running   0             68m
crossplane-rbac-manager-7d6867bc4d-v7wpb    1/1     Running   0             68m
upbound-bootstrapper-5f47977d54-t8kvk       1/1     Running   0             68m
xgql-7c4b74c458-5bf2q                       1/1     Running   3 (67m ago)   68m
```

## Log in with the Up command-line
Use `up login` to authenticate to the Upbound Marketplace.

It's important to use `-a <your organization>` when logging in. Only accounts belonging to organizations can use official providers.

```shell
$ up login -a my-org
username: my-user
password: 
my-user logged in
```
## Create an Upbound robot account
Upbound robots are identities used for authentication that are independent from a single user and aren’t tied to specific usernames or passwords.

Creating a robot account allows Kubernetes to install an official provider.

Use `up robot create <robot account name>` to create a new robot account.

_Note_: only users logged into an organization can create robot accounts.

```shell
$ up robot create my-robot
my-org/my-robot created
```

## Create an Upbound robot account token
The token associates with a specific robot account and acts as a username and password for authentication.

Generate a token using `up robot token create <robot account> <token name> --output=<file>`.

```shell
$ up robot token create my-robot my-token --output=token.json
my-org/my-robot/my-token created
```

The `output` file is a JSON file containing the robot token's `accessId` and `token`. The `accessId` is the username and `token` is the password for the token.

_Note_: you can't recover a lost robot token. You must delete and recreate the token.


## Create a Kubernetes pull secret
Downloading and installing official providers requires Kubernetes to authenticate to the Upbound Marketplace using a Kubernetes `secret` object.

Using the `up controlplane pull-secret create <secret name> -f <robot token file>` command create an [Upbound robot account](http://docs.upbound.io/cli/command-reference/robot/) account. 

Provide a name for your Kubernetes secret and the robot token JSON file.

_Note_: robot accounts are independent from your account. Your account information is never stored in Kubernetes.

_Note_: you must provide the robot token file or you can't authenticate to install an official provider.  

```shell
$ up controlplane pull-secret create my-upbound-secret -f token.json
my-org/my-upbound-secret created
```

`Up` creates the secret in the `upbound-system` namespace. 

```shell
$ kubectl get secret -n upbound-system
NAME                                         TYPE                             DATA   AGE
my-upbound-secret                            kubernetes.io/dockerconfigjson   1      8m46s
sh.helm.release.v1.universal-crossplane.v1   helm.sh/release.v1               1      21m
upbound-agent-tls                            Opaque                           3      21m
uxp-ca                                       Opaque                           3      21m
xgql-tls                                     Opaque                           3      21m
```

## Install the official GCP provider
<!-- Use the marketplace button -->

Install the official provider into the Kubernetes cluster with a Kubernetes configuration file. 

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-gcp
spec:
  package: xpkg.upbound.io/upbound/provider-gcp:latest
  packagePullSecrets:
    - name: my-upbound-secret
```

_Note_: the `name` of the `packagePullSecrets` must be the same as the name of the Kubernetes secret just created.  

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
$ kubectl get providers
NAME           INSTALLED   HEALTHY   PACKAGE                                       AGE
provider-gcp   True        True      xpkg.upbound.io/upbound/provider-gcp:latest   15s
```

It may take up to 5 minutes to report `HEALTHY`.

If the `packagePullSecrets` is incorrect the provider returns a `401 Unauthorized` error. View the status and error with `kubectl describe provider`.

```yaml
$ kubectl describe provider
Name:         provider-gcp
API Version:  pkg.crossplane.io/v1
Kind:         Provider
# Output truncated
Events:
  Type     Reason         Age              From                                 Message
  ----     ------         ----             ----                                 -------
  Warning  UnpackPackage  1s (x2 over 3s)  packages/provider.pkg.crossplane.io  cannot unpack package: failed to fetch package digest from remote: GET https://xpkg.upbound.io/service/token?scope=repository%!A(MISSING)upbound%!F(MISSING)provider-gcp%!A(MISSING)pull&service=xpkg.upbound.io: unexpected status code 401 Unauthorized
```

## Create a Kubernetes secret
The provider requires credentials to create and manage GCP resources.

### Generate a GCP JSON key file
Create a JSON key file containing the GCP account credentials. GCP provides documentation on [how to create a key file](https://cloud.google.com/iam/docs/creating-managing-service-account-keys).

Here is an example key file:  
```json
{
  "type": "service_account",
  "project_id": "caramel-goat-354919",
  "private_key_id": "e97e40a4a27661f12345678f4bd92139324dbf46",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCwA+6MWRhmcPB3\nF/irb5MDPYAT6BWr7Vu/16U8FbCHk7xtsAWYjKXKHu5mGzum4F781sM0aMCeitlv\n+jr2y7Ny23S9uP5W2kfnD/lfj0EjCdfoaN3m7W0j4DrriJviV6ESeSdb0Ehg+iEW\ngNrkb/ljigYgsSLMuemby5lvJVINUazXJtGUEZew+iAOnI4/j/IrDXPCYVNo5z+b\neiMsDYWfccenWGOQf1hkbVWyKqzsInxu8NQef3tNhoUXNOn+/kgarOA5VTYvFUPr\n2l1P9TxzrcYuL8XK++HjVj5mcNaWXNN+jnFpxjMIJOiDJOZoAo0X7tuCJFXtAZbH\n9P61GjhbAgMBAAECggEARXo31kRw4jbgZFIdASa4hAXpoXHx4/x8Q9yOR4pUNR/2\nt+FMRCv4YTEWb01+nV9hfzISuYRDzBEIxS+jyLkda0/+48i69HOTAD0I9VRppLgE\ne97e40a4a27661f12345678f4bd92139324dbf46+2H7ulQDtbEgfcWpNMQcL2JiFq+WS\neh3H0gHSFFIWGnAM/xofrlhGsN64palZmbt2YiKXcHPT+WgLbD45mT5j9oMYxBJf\nPkUUX5QibSSBQyvNqCgRKHSnsY9yAkoNTbPnEV0clQ4FmSccogyS9uPEocQDefuY\nY7gpwSzjXpaw7tP5scK3NtWmmssi+dwDadfLrKF7oQKBgQDjIZ+jwAggCp7AYB/S\n6dznl5/G28Mw6CIM6kPgFnJ8P/C/Yi2y/OPKFKhMs2ecQI8lJfcvvpU/z+kZizcG\nr/7iRMR/SX8n1eqS8XfWKeBzIdwQmiKyRg2AKelGKljuVtI8sXKv9t6cm8RkWKuZ\n9uVroTCPWGpIrh2EMxLeOrlm0QKBgQDGYxoBvl5GfrOzjhYOa5GBgGYYPdE7kNny\nhpHE9CrPZFIcb5nGMlBCOfV+bqA9ALCXKFCr0eHhTjk9HjHfloxuxDmz34vC0xXG\ncegqfV9GNKZPDctysAlCWW/dMYw4+tzAgoG9Qm13Iyfi2Ikll7vfeMX7fH1cnJs0\nnYpN9LYPawKBgQCwMi09QoMLGDH+2pLVc0ZDAoSYJ3NMRUfk7Paqp784VAHW9bqt\n1zB+W3gTyDjgJdTl5IXVK+tsDUWu4yhUr8LylJY6iDF0HaZTR67HHMVZizLETk4M\nLfvbKKgmHkPO4NtG6gEmMESRCOVZUtAMKFPhIrIhAV2x9CBBpb1FWBjrgQKBgQCj\nkP3WRjDQipJ7DkEdLo9PaJ/EiOND60/m6BCzhGTvjVUt4M22XbFSiRrhXTB8W189\noZ2xrGBCNQ54V7bjE+tBQEQbC8rdnNAtR6kVrzyoU6xzLXp6Wq2nqLnUc4+bQypT\nBscVVfmO6stt+v5Iomvh+l+x05hAjVZh8Sog0AxzdQKBgQCMgMTXt0ZBs0ScrG9v\np5CGa18KC+S3oUOjK/qyACmCqhtd+hKHIxHx3/FQPBWb4rDJRsZHH7C6URR1pHzJ\nmhCWgKGsvYrXkNxtiyPXwnU7PNP9JNuCWa45dr/vE/uxcbccK4JnWJ8+Kk/9LEX0\nmjtDm7wtLVlTswYhP6AP69RoMQ==\n-----END PRIVATE KEY-----\n",
  "client_email": "my-sa-313@caramel-goat-354919.iam.gserviceaccount.com",
  "client_id": "103735491955093092925",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/my-sa-313%40caramel-goat-354919.iam.gserviceaccount.com"
}
```

Save this JSON file as `gcp-credentials.json`.

### Create a Kubernetes secret with GCP credentials
Use `kubectl create secret -n upbound-system` to generate the Kubernetes secret object inside the Universal Crossplane cluster.

`kubectl create secret generic gcp-secret -n upbound-system --from-file=creds=./gcp-credentials.json`

View the secret with `kubectl describe secret`
```shell
$ kubectl describe secret gcp-secret -n upbound-system
Name:         gcp-secret
Namespace:    upbound-system
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
creds:  2334 bytes
```
## Create a ProviderConfig
Create a `ProviderConfig` Kubernetes configuration file to attach the GCP credentials to the installed official provider.

**Note:** the `ProviderConfig` must contain the correct GCP project ID. The project ID must match the `project_id` from the JSON key file.

```yaml
apiVersion: gcp.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  projectID: <PROJECT_ID>
  credentials:
    source: Secret
    secretRef:
      namespace: upbound-system
      name: gcp-secret
      key: creds
```

Apply this configuration with `kubectl apply -f`.

**Note:** the `ProviderConfig` value `spec.secretRef.name` must match the `name` of the secret in `kubectl get secrets -n upbound-system` and `spec.secretRef.key` must match the value in the `Data` section of the secret.

Verify the `ProviderConfig` with `kubectl describe providerconfigs`. 

```yaml
$ kubectl describe providerconfigs
Name:         default
Namespace:
API Version:  gcp.upbound.io/v1beta1
Kind:         ProviderConfig
# Output truncated
Spec:
  Credentials:
    Secret Ref:
      Key:        creds
      Name:       gcp-secret
      Namespace:  upbound-system
    Source:       Secret
  Project ID:     caramel-goat-354919
```

## Create a managed resource
Create a managed resource to verify the provider is functioning. 

This example creates a GCP storage bucket, which requires a globally unique name. 

Generate a unique bucket name from the command line.

`echo "upbound-bucket-"$(head -n 4096 /dev/urandom | openssl sha1 | tail -c 10)`

For example
```
$ echo "upbound-bucket-"$(head -n 4096 /dev/urandom | openssl sha1 | tail -c 10)
upbound-bucket-21e85e732
```

Use this bucket name for `metadata.annotations.crossplane.io/external-name` value.

Create a `Bucket` configuration file. Replace `<BUCKET NAME>` with the `upbound-bucket-` generated name.

```yaml
apiVersion: storage.gcp.upbound.io/v1beta1
kind: Bucket
metadata:
  name: example
  labels:
  annotations:
    crossplane.io/external-name: <BUCKET NAME>
spec:
  forProvider:
    location: US
    storageClass: MULTI_REGIONAL
  providerConfigRef:
    name: default
  deletionPolicy: Delete
```

**Note:** the `spec.providerConfigRef.name` must match the `ProviderConfig` `metadata.name` value.

Apply this configuration with `kubectl apply -f`.

Use `kubectl get managed` to verify bucket creation.

```shell
$ kubectl get bucket
NAME      READY   SYNCED   EXTERNAL-NAME              AGE
example   True    True     upbound-bucket-4a917c947   90s
```

Upbound created the bucket when the values `READY` and `SYNCED` are `True`.

If the `READY` or `SYNCED` are blank or `False` use `kubectl describe` to understand why.

Here is an example of a failure because the `Bucket` `spec.providerConfigRef.name` value doesn't match the `ProviderConfig` `metadata.name`.

```shell
$ kubectl describe bucket
Name:         example
Namespace:
Labels:       <none>
Annotations:  crossplane.io/external-name: upbound-bucket-4a917c947
API Version:  storage.gcp.upbound.io/v1beta1
Kind:         Bucket
# Output truncated
Spec:
  Deletion Policy:  Delete
  For Provider:
    Location:       US
    Storage Class:  MULTI_REGIONAL
  Provider Config Ref:
    Name:  default
Status:
  At Provider:
  Conditions:
    Last Transition Time:  2022-07-26T19:29:35Z
    Message:               connect failed: cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.gcp.upbound.io "default" not found
    Reason:                ReconcileError
    Status:                False
    Type:                  Synced
Events:
  Type     Reason                   Age               From                                                 Message
  ----     ------                   ----              ----                                                 -------
  Warning  CannotConnectToProvider  7s (x5 over 20s)  managed/storage.gcp.upbound.io/v1beta1, kind=bucket  cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.gcp.upbound.io "default" not found
```

The output indicates the `Bucket` is using `ProviderConfig` named `default`. The applied `ProviderConfig` is `my-config`. 

```shell
$ kubectl get providerconfigs
NAME        AGE
my-config   56s
```

## Delete the managed resource
Remove the managed resource by using `kubectl delete -f` with the same `Bucket` object file. Verify the removal of the bucket with `kubectl get bucket`.

```shell
$ kubectl get bucket
No resources found
```
