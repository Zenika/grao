#!/bin/bash

# Get project root directory
PROJECT_ROOT=$(cd $(dirname "$0")/../ && pwd)
echo "${PROJECT_ROOT}"

GCLOUD_PROJECT=grao-199314
echo "GCLOUD_PROJECT=${GCLOUD_PROJECT}"

# Go to root folder
cd "${PROJECT_ROOT}"

if [ "${CI}" = true ] ; then
  echo "Authenticate to google cloud"
  # write authentication token to file
  echo ${GCLOUD_SERVICE_KEY} | base64 --decode --ignore-garbage > ${HOME}/gcloud-service-key.json
  # authenticate to gcloud
  gcloud auth activate-service-account --key-file=${HOME}/gcloud-service-key.json
  gcloud config set project ${GCLOUD_PROJECT}
fi

# Build Docker images
function buildImage() {
  docker build -t "gcr.io/${GCLOUD_PROJECT}/${1}" ${2}/
}

buildImage "grao-front" rao-front
buildImage "grao-back" rao-back
buildImage "docd" docd

# Push Docker images to gcr.io
function pushImage() {
  echo "push image ${1}"
  gcloud docker -- push "gcr.io/${GCLOUD_PROJECT}/${1}:latest" > /dev/null
}

pushImage "grao-front"
pushImage "grao-back"
pushImage "docd"

# Deploy images
function deployImage() {
  perl -pe 's/\$(\{)?([a-zA-Z_]\w*)(?(1)\})/$ENV{$2}/g' "${1}/app.yaml" > "${1}/app.deploy.yaml"
  gcloud app deploy "${1}/app.deploy.yaml" --image-url="gcr.io/${GCLOUD_PROJECT}/${2}" --project="${GCLOUD_PROJECT}" --promote --stop-previous-version
  rm "${1}/app.deploy.yaml"
}

yes | deployImage rao-front grao-front
yes | deployImage rao-back grao-back
yes | deployImage docd docd

