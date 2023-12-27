#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER managus ENCRYPTED PASSWORD 'managus' LOGIN;
	CREATE DATABASE managus OWNER managus;
EOSQL

psql -v ON_ERROR_STOP=1 --username "managus" --dbname "managus" -f /app/sql/init-db.sql
