SHELL := /bin/sh

# setup mlflow
mlflow:
	docker compose -f mlflow/docker-compose.yaml up --build -d && ./mlflow/run.sh

shut-mlflow:
	docker compose -f mlflow/docker-compose.yaml down

# setup mlops cluster
cluster:
	make -C cluster run
	echo "cluster are all set..."

shut-cluster:
	make -C cluster down

trace:
	docker compose -f prometheus/docker-compose.yaml up --build -d

shut-trace:
	docker compose -f prometheus/docker-compose.yaml down

# clear the environment
clear-mlflow:
	docker compose -f mlflow/docker-compose.yaml down --volumes --remove-orphans

clear-trace:
	docker compose -f prometheus/docker-compose.yaml down --volumes --remove-orphans

clear: clear-mlflow clear-trace down

up: cluster trace mlflow

down: shut-cluster shut-trace shut-mlflow

lab:
	jupyter lab

.PHONY: cluster shut-cluster mlflow shut-mlflow clear-mlflow trace shut-trace clear-trace clear
