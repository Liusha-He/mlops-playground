#!/bin/sh
set -e

postgres_ready() {
    python << END
import sys

from psycopg2 import connect
from psycopg2.errors import OperationalError

try:
    connect(
        dbname="${DB_NAME}",
        user="${DB_USER}",
        password="${DB_PASSWORD}",
        host="${DB_HOST}",
        port="${DB_PORT}",
    )
except OperationalError:
    sys.exit(-1)
END
}

until postgres_ready
do
  >&2 echo "Waiting for PostgreSQL to become available..."
  sleep 2
done
>&2 echo "PostgreSQL is available"

python -m app.manage migrate
# python -m app.manage collectstatic --noinput

gunicorn app.main:app --bind 0.0.0.0:8080
