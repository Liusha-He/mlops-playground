#!/usr/bin/env bash

helm repo add grafana https://grafana.github.io/helm-charts

helm install grafana grafana/grafana \
    --namespace grafana \
    --create-namespace

kubectl expose service grafana \
    --namespace grafana \
    --type=NodePort \
    --target-port=3000 \
    --name=grafana-np

# print grafana admin secret
kubectl get secret --namespace grafana grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
