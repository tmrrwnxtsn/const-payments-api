#!/bin/bash -e

APP_ENV=${APP_ENV:-local}

echo "[$(date)] Running entrypoint script in the '${APP_ENV}' environment..."

CONFIG_FILE=./configs/${APP_ENV}.yml

if [[ -z ${DSN} ]]; then
  # shellcheck disable=SC2155
  export DSN=$(sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' "${CONFIG_FILE}")
fi

echo "[$(date)] Running DB migrations..."
migrate -database "${APP_DSN}" -path ./migrations up

echo "[$(date)] Starting API server..."
./server -config "${CONFIG_FILE}"
