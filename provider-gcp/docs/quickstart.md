The Upbound GCP Provider is the officially supported provider for Google Compute Platform (GCP).

View the [GCP Provider Documentation](configuration) for details and configuration options. 

## Quickstart
This guide walks through the process to create an Upbound managed control plane and install the GCP official provider.

To use this official provider, install it into your Upbound control plane, apply a `ProviderConfiguration`, and create a *managed resource* in GCP via Kubernetes.

- [Quickstart](#quickstart)
- [Create an Upbound.io user account](#create-an-upboundio-user-account)
- [Create an Upbound user token](#create-an-upbound-user-token)
- [Create a robot account and robot token](#create-a-robot-account-and-robot-token)
- [Install the Up command-line](#install-the-up-command-line)
- [Log in to Upbound](#log-in-to-upbound)
- [Create a managed control plane](#create-a-managed-control-plane)
- [Connect to the managed control plane](#connect-to-the-managed-control-plane)
- [Create a Kubernetes imagePullSecret for Upbound](#create-a-kubernetes-imagepullsecret-for-upbound)
- [Install the official GCP provider in to the managed control plane](#install-the-official-gcp-provider-in-to-the-managed-control-plane)
- [Create a Kubernetes secret](#create-a-kubernetes-secret)
  - [Generate a GCP JSON key file](#generate-a-gcp-json-key-file)
  - [Create a Kubernetes secret with GCP credentials](#create-a-kubernetes-secret-with-gcp-credentials)
- [Create a ProviderConfig](#create-a-providerconfig)
- [Create a managed resource](#create-a-managed-resource)
- [Delete the managed resource](#delete-the-managed-resource)

## Create an Upbound.io user account
Create an account on [Upbound.io](https://cloud.upbound.io/register). 

<!-- Find detailed instructions in the [account documentation](/getting-started/create-account). -->

## Create an Upbound user token
Authentication to an Upbound managed control plane requires a unique user authentication token.

Generate a user token through the [Upbound Universal Console](https://cloud.upbound.io/).

![Create an API Token](images/create-api-token.gif)

To generate a user token in the Upbound Universal Console:
*<!-- vale Microsoft.FirstPerson = NO -->*
1. Log in to the [Upbound Universal Console](https://cloud.upbound.io) and select **My Account** from the account menu.
2. Select **API Tokens**.
3. Select the **Create New Token** button.
4. Provide a token name.
5. Select **Create Token**.
*<!-- vale Microsoft.FirstPerson = Yes -->*

The Console generates a new token and displays it on screen. Save this token. The Console can't print the token again.

## Create a robot account and robot token
Installing an Official Provider requires an Upbound account and associated _Robot Token_.

To create a robot account and robot token in the Upbound Universal Console:
1. Log in to the [Upbound Universal Console](https://cloud.upbound.io) and select **Create New Organization** from the account menu.
2. Provide a unique **Organization ID** and **Display Name**.
3. Select the organization from the account menu.
4. Select **Admin Console**.
5. Select **Robots** from the left-hand navigation. 
6. Select **Create Robot Account**.
7. Provide a **Name** and optional description.
8. Select **Create Robot**.
9. Select **Create Token**.
10. Provide a **Name** for the token.

The console generates an `Access ID` and `Token` on screen. Save this token. The Console can't print the token again.

<!-- Find detailed instructions in the [Robot account and Robot Token](/upbound-cloud/robot-accounts) documentation.  -->

## Install the Up command-line
Install the [Up command-line](https://cloud.upbound.io/docs/cli/install) to connect to Upbound managed control planes.

```shell
curl -sL "https://cli.upbound.io" | sh
sudo mv up /usr/local/bin/
```

## Log in to Upbound
Log in to Upbound with the `up` command-line.

`up login`

## Create a managed control plane
Create a new managed control plane using the command `up controlplane create <control plane name>`.

For example  
`up controlplane create my-gcp-controlplane`

Verify control plane creation with the command

`up controlplane list`

The `STATUS` starts as `provisioning` and moves to `ready`.

```shell
$ $ up controlplane list
NAME                  ID                                     STATUS
my-gcp-controlplane   4fbada32-08f8-4deb-8db6-19e585db5f28   ready
```
## Connect to the managed control plane
Connecting to a managed control plane requires a `kubeconfig` file to connect to the remote cluster.  

Using the **user token** generated earlier and the control plane ID from `up controlplane list`, generate a kubeconfig context configuration.  

`up controlplane kubeconfig get --token <token> <control plane ID>`

Verify that a new context is available in `kubectl` and is the `CURRENT` context.

```shell
$ kubectl config get-contexts
CURRENT   NAME                                           CLUSTER                                        AUTHINFO                                       NAMESPACE
          kubernetes-admin@kubernetes                    kubernetes                                     kubernetes-admin
*         upbound-4fbada32-08f8-4deb-8db6-19e585db5f28   upbound-4fbada32-08f8-4deb-8db6-19e585db5f28   upbound-4fbada32-08f8-4deb-8db6-19e585db5f28
```

**Note:** change the `CURRENT` context with `kubectl config use-context <context name>`.

Confirm your token's access with any `kubectl` command.

```shell
$ kubectl get pods -A
NAMESPACE        NAME                                      READY   STATUS    RESTARTS   AGE
upbound-system   crossplane-6cc78c56df-vbxxc               1/1     Running   0          117s
upbound-system   crossplane-rbac-manager-c5b5c44c4-x4l5s   1/1     Running   0          117s
upbound-system   upbound-agent-6cf5cfc5bc-x7d8v            1/1     Running   0          115s
upbound-system   upbound-bootstrapper-5b96cf657d-whv7v     1/1     Running   0          117s
upbound-system   xgql-7fbc547668-78j6q                     1/1     Running   1          117s
```

**Note:** if the token is incorrect the `kubectl` command returns an error.
```
$ kubectl get pods -A
Error from server (BadRequest): the server rejected our request for an unknown reason
```

## Create a Kubernetes imagePullSecret for Upbound
Official providers require a Kubernetes `imagePullSecret` to download and install. The credentials for the `imagePullSecret` are from an Upbound robot token. 

Using the **robot token** generated earlier create an `imagePullSecret` with the command `kubectl create secret docker-registry package-pull-secret`.

```shell
kubectl create secret docker-registry package-pull-secret --namespace=crossplane-system --docker-server=xpkg.upbound.io --docker-username=<robot token access ID> --docker-password=<robot token value>
```

Replace `<robot token access ID>` with the `Access ID` of the robot token and `<robot token value>` with the value of the robot token.

Verify the secret with `kubectl get secrets`
```shell
$ kubectl get secrets -n crossplane-system package-pull-secret
NAME                  TYPE                             DATA   AGE
package-pull-secret   kubernetes.io/dockerconfigjson   1      23s
```

## Install the official GCP provider in to the managed control plane
<!-- Use the marketplace button -->

Install the official provider into the managed control plane with a Kubernetes configuration file. 

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-gcp
spec:
  package: xpkg.upbound.io/upbound/provider-gcp:v0.4.1
  packagePullSecrets:
    - name: package-pull-secret
```

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
$ kubectl get providers
NAME           INSTALLED   HEALTHY   PACKAGE                                       AGE
provider-gcp   True        True      xpkg.upbound.io/upbound/provider-gcp:v0.4.1   15s
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
Use `kubectl create secret -n upbound-system` to generate the Kubernetes secret object inside the managed control plane.

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

**Note:** the `ProviderConfg` must contain the correct GCP project ID. The project ID must match the `project_id` from the JSON key file.

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

**Note:** the `Providerconfig` value `spec.secretRef.name` must match the `name` of the secret in `kubectl get secrets -n upbound-system` and `spec.SecretRef.key` must match the value in the `Data` section of the secret.

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
    crossplane.io/external-name: upbound-bucket-4a917c947
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