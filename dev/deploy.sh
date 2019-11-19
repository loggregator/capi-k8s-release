#!/bin/bash

set -ex

SCRIPT_DIR=$(dirname $0)

#database
helm install capi-database stable/postgresql -f "${SCRIPT_DIR}/postgresql-values.yaml"
#minio
helm install capi-blobstore stable/minio

#capi
helm template "${SCRIPT_DIR}/.." --set-string system_domain=minikube.local -f "${SCRIPT_DIR}/capi-values.yaml" | kubectl apply -f -
