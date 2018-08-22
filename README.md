# go-gcp-looging

## Create service account
export APPLICATION=go-gcp-logging
export APP_SA_NAME=gke-$APPLICATION-sa
gcloud iam service-accounts create $APP_SA_NAME --display-name "GKE $APPLICATION Application Service Account"
export APP_SA_EMAIL=`gcloud iam service-accounts list --format='value(email)' --filter='displayName:$APPLICATION Application Service Account'`

## Bind service account policy
export PROJECT=`gcloud config get-value project`

gcloud projects add-iam-policy-binding $PROJECT \
    --member=serviceAccount:${APP_SA_EMAIL} \
    --role=roles/logging.logWriter

## Create service account key and activate it
gcloud iam service-accounts keys create \
    /home/$USER/$APPLICATION-$SA-key.json \
    --iam-account $APP_SA_EMAIL

## Create configmap
kubectl create configmap project-id --from-literal "project-id=${PROJECT}"
kubectl create configmap $APPLICATION-sa --from-literal "sa-email=${APP_SA_EMAIL}"
kubectl create secret generic $APPLICATION --from-file /home/$USER/$APPLICATION-$SA-key.json

envsubst < deployment.template.yaml > deployment.yaml