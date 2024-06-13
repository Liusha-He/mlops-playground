SHELL := /bin/sh

# setup mlflow
mlflow:
	docker compose -f mlflow/docker-compose.yaml up -d

shut-mlflow:
	docker compose -f mlflow/docker-compose.yaml down --volumes --remove-orphans

# setup mlops cluster
cluster:
	make -C platform cluster

platform: cluster mlflow
	echo "cluster are all set..."

down:
	make -C platform down

run:
	jupyter notebook

compile:
	poetry update

.PHONY: run platform down compile mlflow shut-mlflow
