#!/usr/bin/env bash

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm install prometheus prometheus-community/prometheus \
    --namespace prometheus \
    --create-namespace

kubectl expose service prometheus-server \
    -n prometheus \
    --type=NodePort \
    --target-port=9090 \
    --name=prometheus-server-np
