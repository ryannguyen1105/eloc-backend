#!/bin/sh
set -e

echo "Running database migration"
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

echo "Starting Go Application"
exec "$@"

