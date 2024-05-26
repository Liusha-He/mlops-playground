SHELL := /bin/sh

mlflow:
	make -C platform mlflow

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

.PHONY: run platform down compile
