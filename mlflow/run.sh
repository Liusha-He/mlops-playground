#!/bin/sh

echo "$(dirname $PWD)/mlops/registry/artifacts"

./mlflow/wait-for-it.sh localhost:5432 -t 60 -- \
    mlflow server \
    --backend-store-uri postgresql+psycopg2://registry:password@localhost:5432/metadata \
    --default-artifact-root "$(dirname $PWD)/mlops/registry/artifacts" \
    --host localhost \
    --port 5000
