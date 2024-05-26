FROM python:3.10-slim

WORKDIR /usr/app

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN apt-get update
RUN pip install --upgrade pip
RUN pip install poetry==1.3.2

COPY ./pyproject.toml ./poetry.lock ./
RUN poetry config virtualenvs.create false && poetry install --no-dev

COPY . .

RUN chmod +x app-entrypoint.sh
ENTRYPOINT ["./app-entrypoint.sh"]
