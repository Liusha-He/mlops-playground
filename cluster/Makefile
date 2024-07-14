SHELL := /bin/sh

flyte:
	flytectl sandbox start

prometheus:
	docker compose up --build -d

run: flyte prometheus
	echo "flyte and seldon are set"

down:
	flytectl sandbox teardown

.PHONY: flyte down prometheus run
