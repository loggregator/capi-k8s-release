#!/usr/bin/env bash

docker build -f dockerfiles/cloud_controller_ng/Dockerfile -t capi -t $(minikube ip):5000/capi src/
docker push $(minikube ip):5000/capi
kubectl rollout restart deployment/capi
kubectl rollout status deployment/capi -w

# docker save capi -o capi.tar
