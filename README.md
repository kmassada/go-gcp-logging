# go-gcp-looging

## Build pre-reqs

```shell
export _REPO_PREFIX=makz-labs
export APPLICATION=go-gcp-logging
export REPO_NAME=$APPLICATION
```

therefore you need a Cloud Storage bucket with this prefix

```shell
gs://${PROJECT_ID}_${_REPO_PREFIX}
```

and you need a gcr prefix that matches this:

```shell
gcr.io/${PROJECT_ID}/${_REPO_PREFIX}/
```

## Build

When using github's integration to build this repo. `REPO_NAME` and `SHORT_SHA` are auto-populated.

`_REPO_PREFIX` and `TAG_NAME` are declared in the `cloudbuild.yaml`.

`PROJECT_ID` is the project gcloud/build repo is authenticated against: `$ gcloud config set project VALUE`

When running in cloudshell we manually set our substitution variables.

```shell
gcloud builds submit \
    --substitutions _REPO_PREFIX=makz-labs,REPO_NAME=go-gcp-logging,TAG_NAME=cli,SHORT_SHA=clisha \
    --config cloudbuild.yaml
```

## Application credentials

This section, I create a service account, download a key for it, this key will be injected into my application.

### Create service account

```shell
export APP_SA_NAME=gke-$APPLICATION-sa
gcloud iam service-accounts create $APP_SA_NAME --display-name "GKE $APPLICATION Application Service Account"
export APP_SA_EMAIL=`gcloud iam service-accounts list --format='value(email)' --filter='displayName:GKE '$A
PPLICATION' Application Service Account'`
```

### Bind service account policy

```shell
export PROJECT_ID=`gcloud config get-value project`

gcloud projects add-iam-policy-binding $PROJECT_ID \
    --member=serviceAccount:${APP_SA_EMAIL} \
    --role=roles/logging.logWriter
```

### Create service account key and activate it

```shell
gcloud iam service-accounts keys create \
    /home/$USER/$APPLICATION-sa-key.json \
    --iam-account $APP_SA_EMAIL
```

## Application Bootstrap

In this section I create configmaps from the variables we've been gathering to start our application

### Create configmap

```shell
kubectl create configmap project-id --from-literal "project-id=${PROJECT_ID}"
kubectl create configmap $APPLICATION-sa --from-literal "sa-email=${APP_SA_EMAIL}"
kubectl create secret generic $APPLICATION --from-file /home/$USER/$APPLICATION-sa-key.json
```

TODO: Replace this by helm

```shell
envsubst < deployment.template.yaml > deployment.yaml
```

NOTE: envsubt is part of `gettext-base` on ubuntu.

Apply config

```shell
kubectl apply -f deployment.yaml
```
