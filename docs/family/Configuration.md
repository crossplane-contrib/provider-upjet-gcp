---
title: Configuration
weight: 2
---

# GCP official provider-family documentation
Upbound supports and maintains the Upbound GCP official provider-family.

## Install the provider-family-gcp
### Prerequisites
#### Upbound Up command-line
The Upbound Up command-line simplifies configuration and management of Upbound
Universal Crossplane (UXP) and interacts with the Upbound Marketplace to manage
users and accounts.

Install `up` with the command:
```shell
curl -sL "https://cli.upbound.io" | sh
```
More information about the Up command-line is available in the [Upbound Up
documentation](https://docs.upbound.io/cli/).

#### Upbound Universal Crossplane
UXP is the Upbound official enterprise-grade distribution of Crossplane for
self-hosted control planes.

Install UXP into your Kubernetes cluster using the Up command-line.

```shell
up uxp install
```

Find more information in the [Upbound UXP documentation](https://docs.upbound.io/uxp/).

### Install the provider-family-gcp

Install the Upbound official GCP provider-family with the following configuration file.
For instance, let's install the `provider-gcp-storage`.

_Note_: The first provider installed of a family also installs an extra provider-family Provider.
The provider-family provider manages the ProviderConfig for all other providers in the same family.

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-gcp-storage
spec:
  package: xpkg.upbound.io/upbound/provider-gcp-storage:<version>
```

Define the `provider-gcp-storage` version with `spec.package`.

Install the `provider-gcp-storage` with `kubectl apply -f`.

Verify the configuration with `kubectl get providers`.

```shell
$ kubectl get providers
NAME                          INSTALLED   HEALTHY   PACKAGE                                                AGE
provider-gcp-storage          True        True      xpkg.upbound.io/upbound/provider-gcp-storage:v0.36.0   78s
upbound-provider-family-gcp   True        True      xpkg.upbound.io/upbound/provider-family-gcp:v0.36.0    70s
```

View the Crossplane [Provider CRD definition](https://doc.crds.dev/github.com/crossplane/crossplane/pkg.crossplane.io/Provider/v1) to view all available `Provider` options.

If you are going to use your own registry please check [Install Providers in an offline environment](https://docs.upbound.io/providers/provider-families/#installing-a-provider-family:~:text=services%20to%20install.-,Install%20Providers%20in%20an%20offline%20environment,-View%20the%20installed).

## Configure the provider-family-gcp
The GCP provider-family requires credentials for authentication to Google Cloud Platform.
This can be done in one of the following ways:

- Authenticating using a base-64 encoded service account key in a Kubernetes
  `Secret`.
- Authenticating using [Workload Identity](https://cloud.google.com/kubernetes-engine/docs/concepts/workload-identity).

### Generate a Kubernetes secret
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
Use the JSON file to generate a Kubernetes secret.

`kubectl create secret generic gcp-secret --from-file=creds=./<JSON file name>`

### Create a ProviderConfig object
Apply the secret in a `ProviderConfig` Kubernetes configuration file.

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

**Note:** the `spec.credentials.secretRef.name` must match the `name` in the `kubectl create secret generic <name>` command.

## Authenticating with Workload Identity

Using Workload Identity requires some additional setup.
Many of the steps can also be found in the [documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity).

### Steps

These steps assume you already have a running GKE cluster which has already
enabled Workload Identity and has a sufficiently large node pool.

#### 0. Prepare your variables

In the following sections, you'll need to name your resources.
Define the variables below with any names valid in Kubernetes or GCP so that you
can smoothly set it up:

```console
$ PROJECT_ID=<YOUR_GCP_PROJECT_ID>                               # e.g.) <project_id>
$ PROVIDER_GCP=<YOUR_PROVIDER_GCP_NAME>                          # e.g.) provider-gcp-storage
$ VERSION=<YOUR_PROVIDER_GCP_VERSION>                            # e.g.) 0.36.0
$ GCP_SERVICE_ACCOUNT=<YOUR_CROSSPLANE_GCP_SERVICE_ACCOUNT_NAME> # e.g.) <gcp_sa_name>
$ ROLE=<YOUR_ROLE_FOR_CROSSPLANE_GCP_SERVICE_ACCOUNT>            # e.g.) roles/cloudsql.admin
$ CONTROLLER_CONFIG=<YOUR_CONTROLLER_CONFIG_NAME>                # e.g.) gcp-config (Optional)
```

#### 1. Configure service accounts to use Workload Identity

Create a GCP service account, which will be used for provisioning actual
infrastructure in GCP, and grant IAM roles you need for accessing the Google
Cloud APIs:

```console
$ gcloud iam service-accounts create ${GCP_SERVICE_ACCOUNT}
$ gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member "serviceAccount:${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com" \
    --role ${ROLE} \
    --project ${PROJECT_ID}
```

Get the name of your current `ProviderRevision` of this provider:

```console
$ REVISION=$(kubectl get providers.pkg.crossplane.io ${PROVIDER_GCP} -o jsonpath="{.status.currentRevision}")
```

Next, you'll configure IAM to use Workload Identity.
In this step, you can choose one of the following options to configure service accounts:

- [Option 1] Use a Kubernetes `ServiceAccount` managed by a provider's controller.
- [Option 2] Use a Kubernetes `ServiceAccount` which you created and is specified to `.spec.serviceAccountName`
  in a [`ControllerConfig`](https://doc.crds.dev/github.com/crossplane/crossplane/pkg.crossplane.io/ControllerConfig/v1alpha1@v1.6.2).

##### 2.1. [Option 1] Use a controller-managed `ServiceAccount`

Specify a Kubernetes `ServiceAccount` with the revision you got in the last
step:

```console
$ KUBERNETES_SERVICE_ACCOUNT=${REVISION}
```

##### 2.1. [Option 2] Use a user-managed `ServiceAccount`

Name your Kubernetes `ServiceAccount`:

```console
$ KUBERNETES_SERVICE_ACCOUNT=<YOUR_KUBERNETES_SERVICE_ACCOUNT>
```

Create a `ServiceAccount`, `ControllerConfig`, and `ClusterRoleBinding`:

```console
$ cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name: ${KUBERNETES_SERVICE_ACCOUNT}
  namespace: upbound-system
---
apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: ${CONTROLLER_CONFIG}
spec:
  serviceAccountName: ${KUBERNETES_SERVICE_ACCOUNT}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: crossplane:provider:${PROVIDER_GCP}:system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: crossplane:provider:${REVISION}:system
subjects:
- kind: ServiceAccount
  name: ${KUBERNETES_SERVICE_ACCOUNT}
  namespace: upbound-system
EOF
```

#### 2.2. Allow the Kubernetes `ServiceAccount` to impersonate the GCP service account

Grant `roles/iam.workloadIdentityUser` to the GCP service account:

```console
$ gcloud iam service-accounts add-iam-policy-binding \
    ${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com \
    --role roles/iam.workloadIdentityUser \
    --member "serviceAccount:${PROJECT_ID}.svc.id.goog[upbound-system/${KUBERNETES_SERVICE_ACCOUNT}]" \
    --project ${PROJECT_ID}
```

Annotate the `ServiceAccount` with the email address of the GCP service account:

```console
$ kubectl annotate serviceaccount ${KUBERNETES_SERVICE_ACCOUNT} \
    iam.gke.io/gcp-service-account=${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com \
    -n upbound-system
```

### 3. Configure a `ProviderConfig`

Create a `ProviderConfig` with `InjectedIdentity` in `.spec.credentials.source`:

```console
$ cat <<EOF | kubectl apply -f -
apiVersion: gcp.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  projectID: ${PROJECT_ID}
  credentials:
    source: InjectedIdentity
EOF
```

### 4. Next steps

Now that you have configured `provider-gcp-storage` with Workload Identity supported.

## Authenticate with Impersonated Service Account
#### 0. Setup Variables

```console
$ CLIENT_GCP_SERVICE_ACCOUNT=<YOUR_CLIENT_GCP_SERVICE_ACCOUNT>                               # e.g.) impersonate-example
```

#### 1. Setup Provider with Workload Identity

Please follow steps 0, 1, and 2 in [Authenticating with Workload Identity](#authenticating-with-workload-identity)

#### 2. Create Client Service Account

Create a GCP service account, which will be used for provisioning actual
infrastructure in GCP, and grant IAM roles you need for accessing the Google
Cloud APIs. Remove role previously created for ${GCP_SERVICE_ACCOUNT}:

```console
$ gcloud iam service-accounts create ${CLIENT_GCP_SERVICE_ACCOUNT}
$ gcloud projects remove-iam-policy-binding ${PROJECT_ID} \
    --member "serviceAccount:${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com" \
    --role ${ROLE} \
    --project ${PROJECT_ID}
$ gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member "serviceAccount:${CLIENT_GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com" \
    --role ${ROLE} \
    --project ${PROJECT_ID}
$ gcloud iam service-accounts add-iam-policy-binding \
    ${CLIENT_GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com \
    --role roles/iam.serviceAccountTokenCreator \
    --member "serviceAccount:${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com" \
    --project ${PROJECT_ID}
```

### 3. Configure a `ProviderConfig`

Create a `ProviderConfig` with `ImpersonateServiceAccount` in `.spec.credentials.source` and specify a GCP service account in `.spec.credentials.serviceAccount`:

```console
$ cat <<EOF | kubectl apply -f -
apiVersion: gcp.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  projectID: ${PROJECT_ID}
  credentials:
    source: ImpersonateServiceAccount
    impersonateServiceAccount:
      name: ${CLIENT_GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com
EOF
```

### 4. Next steps

Now you have configured `provider-gcp-storage` with Impersonated Service Account support.

## Authenticating with Access Tokens

Using temporary Access Tokens will require a process to regenerate an access token before it expires. Luckily we can use a Kubernetes CronJob to fulfill that.

**DISCLAIMER**

*The following method will only work if running the provider in a GKE cluster on GCP. This is because the creation of access tokens requires a service account with Workload Identity enabled.*

### Steps

#### 0. Prepare your variables

In the following sections, you'll need to name your resources.
Define the variables below with any names valid in Kubernetes or GCP so that you
can smoothly set it up:

```console
$ PROJECT_ID=<YOUR_GCP_PROJECT_ID>                               # e.g.) acme-prod
$ REGION=<YOUR_GCP_REGION>                                       # e.g.) us-central1
$ CLUSTER_NAME=<YOUR_CLUSTER_NAME>                               # e.g.) demo
$ GCP_SERVICE_ACCOUNT=<YOUR_CROSSPLANE_GCP_SERVICE_ACCOUNT_NAME> # e.g.) crossplane
$ ROLE=<YOUR_ROLE_FOR_CROSSPLANE_GCP_SERVICE_ACCOUNT>            # e.g.) roles/editor
$ KUBERNETES_SERVICE_ACCOUNT=<YOUR_KUBERNETES_SERVICE_ACCOUNT>   # e.g.) token-generator
$ NAMESPACE=<YOUR_KUBERNETES_NAMESPACE>                          # e.g.) default
$ SECRET_NAME=<YOUR_CREDENTIALS_SECRET_NAME>                     # e.g.) gcp-credentials
$ SECRET_KEY=<NAME_OF_KEY_IN_SECRET>                             # e.g.) token
$ PROVIDER_GCP=<YOUR_PROVIDER_GCP_NAME>                          # e.g.) provider-gcp-storage
$ VERSION=<YOUR_PROVIDER_GCP_VERSION>                            # e.g.) v0.36.0
```

#### 1. Create a GKE cluster with Workload Identity Enabled
Create a default vpc if one does not already exist
```console
$ gcloud compute networks create default \
    --subnet-mode=auto \
    --bgp-routing-mode=global \
    --project=${PROJECT_ID}
```
Create a cloud router
```console
$ gcloud compute routers create ${CLUSTER_NAME} \
    --project=${PROJECT_ID} \
    --network=default \
    --region=${REGION}
```
Create a cloud nat
```console
$ gcloud compute routers nats create ${CLUSTER_NAME} \
   --router=${CLUSTER_NAME} \
   --region=${REGION} \
   --auto-allocate-nat-external-ips \
   --nat-all-subnet-ip-ranges \
   --project=${PROJECT_ID}
```
Create the cluster
```console
$ gcloud container clusters create ${CLUSTER_NAME} \
  --region=${REGION} \
  --workload-pool=${PROJECT_ID}.svc.id.goog \
  --create-subnetwork name=gke \
  --enable-ip-alias \
  --enable-private-nodes \
  --no-enable-master-authorized-networks \
  --enable-master-global-access \
  --master-ipv4-cidr=172.16.0.32/28 \
  --num-nodes=1 \
  --project=${PROJECT_ID}
```
Get the cluster credentials
```console
$ gcloud container clusters get-credentials ${CLUSTER_NAME} --region=${REGION} --project=${PROJECT_ID}
```

#### 2. Configure service accounts to use Workload Identity

Create a GCP service account, which will be used for provisioning actual
infrastructure in GCP, and grant IAM roles you need for accessing the Google
Cloud APIs:

```console
$ gcloud iam service-accounts create ${GCP_SERVICE_ACCOUNT} \
  --project=${PROJECT_ID}
```
```console
$ gcloud projects add-iam-policy-binding ${PROJECT_ID} \
  --member="serviceAccount:${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com" \
  --role=${ROLE} \
  --project=${PROJECT_ID}
```

#### 3. Create resources to generate an access-token
Create the Kubernetes service account, RBAC, and CronJob to generate the temporary access-token

**NOTE:** Ensure your kube context is pointing to the cluster created above

```console
$ cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name:  ${KUBERNETES_SERVICE_ACCOUNT}
  namespace: ${NAMESPACE}
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ${KUBERNETES_SERVICE_ACCOUNT}-sync
  namespace: ${NAMESPACE}
rules:
- apiGroups: [""]
  resources:
  - secrets
  verbs:
  - get
  - create
  - patch
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ${KUBERNETES_SERVICE_ACCOUNT}-sync-rb
  namespace: ${NAMESPACE}
subjects:
- kind: ServiceAccount
  name: ${KUBERNETES_SERVICE_ACCOUNT}
roleRef:
  kind: Role
  name: ${KUBERNETES_SERVICE_ACCOUNT}-sync
  apiGroup: ""
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: ${KUBERNETES_SERVICE_ACCOUNT}-credentials-sync
  namespace: ${NAMESPACE}
spec:
  suspend: false
  schedule: "*/45 * * * *"
  failedJobsHistoryLimit: 1
  successfulJobsHistoryLimit: 1
  concurrencyPolicy: Forbid
  startingDeadlineSeconds: 1800
  jobTemplate:
    spec:
      template:
        spec:
          serviceAccountName:  ${KUBERNETES_SERVICE_ACCOUNT}
          restartPolicy: Never
          containers:
            - image: google/cloud-sdk:debian_component_based
              name: create-access-token
              imagePullPolicy: IfNotPresent
              livenessProbe:
                exec:
                  command:
                  - gcloud
                  - version
              readinessProbe:
                exec:
                  command:
                  - gcloud
                  - version
              env:
                - name: SECRET_NAME
                  value: ${SECRET_NAME}
                - name: SECRET_KEY
                  value: ${SECRET_KEY}
              command:
                - /bin/bash
                - -ce
                - |-
                  kubectl create secret generic $SECRET_NAME \
                    --dry-run=client \
                    --from-literal=$SECRET_KEY=$(gcloud auth print-access-token) \
                    -o yaml | kubectl apply -f -
              resources:
                requests:
                  cpu: 250m
                  memory: 256Mi
                limits:
                  cpu: 500m
                  memory: 512Mi
EOF
```
Grant `roles/iam.workloadIdentityUser` to the GCP service account:

```console
$ gcloud iam service-accounts add-iam-policy-binding \
    ${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com \
    --role=roles/iam.workloadIdentityUser \
    --member="serviceAccount:${PROJECT_ID}.svc.id.goog[${NAMESPACE}/${KUBERNETES_SERVICE_ACCOUNT}]" \
    --project=${PROJECT_ID}
```

Annotate the `ServiceAccount` with the email address of the GCP service account:

```console
$ kubectl annotate serviceaccount ${KUBERNETES_SERVICE_ACCOUNT} \
    iam.gke.io/gcp-service-account=${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com \
    -n ${NAMESPACE}
```

#### 4. Create initial Access Token
```console
$ kubectl -n ${NAMESPACE} create job --from=cronjob/${KUBERNETES_SERVICE_ACCOUNT}-credentials-sync cred-sync-001
```

#### 5. Install Crossplane

Install Crossplane from `stable` channel:

```bash
$ helm repo add crossplane-stable https://charts.crossplane.io/stable
$ helm install crossplane --create-namespace --namespace crossplane-system crossplane-stable/crossplane
```

`provider-gcp-storage` can be installed with either the [Crossplane CLI](https://crossplane.io/docs/v1.6/getting-started/install-configure.html#install-crossplane-cli)
or a `Provider` resource as below:

```console
$ cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-gcp-storage
spec:
  package: xpkg.upbound.io/upbound/provider-gcp-storage:${VERSION}
EOF
```

#### 6. Create ProviderConfig
```console
$ cat <<EOF | kubectl apply -f -
apiVersion: gcp.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  projectID: ${PROJECT_ID}
  credentials:
    source: AccessToken
    secretRef:
      name: ${SECRET_NAME}
      namespace: ${NAMESPACE}
      key: ${SECRET_KEY}
EOF
```

### 7. Next steps

Now that you have configured `provider-gcp-storage` with Access Tokens supported,
you can [provision infrastructure](https://crossplane.io/docs/v1.6/getting-started/provision-infrastructure).

#### 8. Clean up
Delete GKE cluster
```console
$ gcloud container clusters delete ${CLUSTER_NAME} \
  --region=${REGION} \
  --project=${PROJECT_ID}
```
Delete cloud nat
```console
$ gcloud compute routers nats delete ${CLUSTER_NAME} \
  --router=${CLUSTER_NAME} \
  --region=${REGION} \
  --project=${PROJECT_ID}
```
Delete cloud router
```console
$ gcloud compute routers delete ${CLUSTER_NAME} \
  --region=${REGION} \
  --project=${PROJECT_ID}
```
Delete VPC
```console
$ gcloud compute networks delete default \
  --project=${PROJECT_ID}
```
Delete project IAM bindings
```console
$ gcloud projects remove-iam-policy-binding ${PROJECT_ID} \
  --member="serviceAccount:${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com" \
  --role=${ROLE}
```
Delete service account
```console
$ gcloud iam service-accounts delete ${GCP_SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com
```
