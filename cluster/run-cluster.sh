#!/usr/bin/env bash

flytectl demo start

# set cert-manager
helm install certmgr jetstack/cert-manager --set installCRDs=true --version v1.8.0 --namespace cert-manager --create-namespace

# set  kserve
kubectl create namespace kserve
helm install kserve-crd oci://ghcr.io/kserve/charts/kserve-crd \
  --version v0.13.0 \
  --namespace kserve
helm install kserve oci://ghcr.io/kserve/charts/kserve \
  --version v0.13.0 \
  --namespace kserve
