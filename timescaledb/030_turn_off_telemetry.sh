#!/bin/sh

IFS=","
for v in $POSTGRES_DBS
do
   PGPASSWORD=${POSTGRES_PASSWORD} psql -U ${POSTGRES_USER} -c "ALTER DATABASE $v SET timescaledb.telemetry_level to off"
done