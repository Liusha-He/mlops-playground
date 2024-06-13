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

run: flyte cert-manager seldon prometheus grafana
	echo "flyte and seldon are set"

down:
	flytectl sandbox teardown

.PHONY: mlflow seldon flyte down prometheus grafana run down
