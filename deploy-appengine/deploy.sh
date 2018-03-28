#!/bin/bash
# This script build the three docker containers and deploy them on App Engine flexible.

# Get project root directory.
PROJECT_ROOT=$(cd $(dirname "$0")/../ && pwd)
echo "${PROJECT_ROOT}"

# Define google cloud project code.
GCLOUD_PROJECT=grao-199314
echo "GCLOUD_PROJECT=${GCLOUD_PROJECT}"

# Go to root folder
cd "${PROJECT_ROOT}"

# If on CircleCI authenticate to Google Cloud
if [ "${CI}" = true ] ; then
  echo "authenticate to google cloud"
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

# Deploy Docker images to App Engine
function deployImage() {
  # Update app.yaml file to replace templated environment variables.
  perl -pe 's/\$(\{)?([a-zA-Z_]\w*)(?(1)\})/$ENV{$2}/g' "${1}/app.yaml" > "${1}/app.deploy.yaml"
  # Do App Engine deploy
  # The deploy might take more than 5 minutes (https://groups.google.com/forum/#!topic/google-appengine/hZMEkmmObDU)
  yes | gcloud app deploy "${1}/app.deploy.yaml" --image-url="gcr.io/${GCLOUD_PROJECT}/${2}" --project="${GCLOUD_PROJECT}" --promote --stop-previous-version
  rm "${1}/app.deploy.yaml"
}

deployImage rao-front grao-front
deployImage rao-back grao-back
deployImage docd docd

