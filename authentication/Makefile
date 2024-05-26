SHELL := /bin/bash

export IMAGE?=user-service
export TAG?=latest

run:
	swag init --dir \
	./src/cmd/api,\
	./src/controller &
	sudo docker compose up --build

down:
	sudo rm -rf db-data
	docker compose down

build:
	CGO_ENABLED=0 go build -o userService ./src

docs: 
	swag init --dir \
	./src/cmd/api,\
	./src/controller

test:
	go test -v ./tests/...

docker: docs
	docker build . -t ${IMAGE}:${TAG}

.PHONY: run down build docs docker
