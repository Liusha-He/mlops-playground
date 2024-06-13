SHELL := /bin/sh

# setup mlflow
mlflow:
	docker compose -f mlflow/docker-compose.yaml up -d

shut-mlflow:
	docker compose -f mlflow/docker-compose.yaml down --volumes --remove-orphans

# setup mlops cluster
cluster:
	make -C cluster run
	echo "cluster are all set..."

down:
	make -C cluster down

run:
	jupyter notebook

compile:
	poetry update

.PHONY: run cluster down compile mlflow shut-mlflow
