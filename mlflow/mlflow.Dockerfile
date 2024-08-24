FROM python:3.10.11-slim

EXPOSE 5000

ARG SERVER_DIR=./

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt && \
    mkdir -p ${SERVER_DIR}

COPY wait-for-it.sh ${SERVER_DIR}/
RUN chmod +x ${SERVER_DIR}/wait-for-it.sh

WORKDIR ${SERVER_DIR}

RUN mkdir -p mlflow

USER root

RUN chown -R 1000:1000 /mlflow

USER 1000

# wait as long as 60 seconds for posgres to init db, which takes a lot of time particularly in mac osx
CMD ./wait-for-it.sh ${MLFLOW_BACKEND_STORE}:${POSTGRES_PORT} -t ${WAIT_FOR_IT_TIMEOUT} -- \
    mlflow server \
    --backend-store-uri postgresql+psycopg2://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${MLFLOW_BACKEND_STORE}:${POSTGRES_PORT}/${POSTGRES_DATABASE} \
    --default-artifact-root ${MLFLOW_ARTIFACT_STORE} \
    --host ${MLFLOW_TRACKING_SERVER_HOST} \
    --port ${MLFLOW_TRACKING_SERVER_PORT}
