SHELL := /bin/sh

flyte:
	flytectl sandbox start

seldon: 
	chmod +x seldon/install.sh
	./seldon/install.sh

cert-manager:
	chmod +c cert-manager/install.sh
	./cert-manager/install.sh

prometheus:
	chmod +x prometheus/install.sh
	./prometheus/install.sh

grafana:
	chmod +x grafana/install.sh
	./grafana/install.sh

cluster: flyte cert-manager seldon prometheus grafana
	echo "flyte and seldon are set"

mlflow:
	sudo docker compose --profile mlflow -f mlflow/docker-compose.yaml up -d

shut-cluster:
	flytectl sandbox teardown

down: shut-cluster
	docker compose -f mlflow/docker-compose.yaml down --volumes --remove-orphans

.PHONY: mlflow fs seldon flyte shut-cluster down prometheus grafana
