#!/usr/bin/env bash

helm repo add jetstack https://charts.jetstack.io 
helm repo update

echo "========== Setup Cert Manager =========="

helm install certmgr jetstack/cert-manager \
    --set installCRDs=true \
    --version v1.8.0 \
    --namespace cert-manager \
    --create-namespace
