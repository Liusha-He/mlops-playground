SHELL := /bin/sh

export MODE := LOCAL

install-dev:
	poetry install

install:
	poetry install --no-dev

compile:
	poetry update

run-local:
	gunicorn app.config.wsgi:app --bind localhost:8080

pre-lint:
	poetry run black --line-length 120 app

lint:
	poetry run black --line-length 120 --check app

test:
	pytest -svv tests

coverage:
	pytest --cov=app tests/

run-docker:
	sudo docker compose up --build

migrate:
	python -m app.manage makemigrations

down:
	sudo docker compose down --volumes --remove-orphans

.PHONY: run-local test compile coverage install install-dev migrate down
