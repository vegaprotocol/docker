#!/bin/bash

export VEGA_POSTGRES_USER="vega"
export VEGA_POSTGRES_PASSWORD="vega"
export VEGA_POSTGRES_DBS="vega0,vega1,vega2,vega3,vega4,vega5,vega6,vega7,vega8,vega9,vega10,vega11,vega12,vega13,vega14,vega15,vega16,vega17,vega18,vega19,vega20,vega21,vega22,vega23,vega24,vega25,vega26,vega27,vega28,vega29"


if [ ! -z "${VEGA_POSTGRES_USER}" ]; then
   psql -c "CREATE USER ${VEGA_POSTGRES_USER} WITH SUPERUSER PASSWORD '${VEGA_POSTGRES_PASSWORD}';"
fi;

psql -c "CREATE DATABASE ${VEGA_POSTGRES_USER}" -c "GRANT ALL PRIVILEGES ON DATABASE ${VEGA_POSTGRES_USER} TO ${VEGA_POSTGRES_USER};"

IFS=","
for v in $VEGA_POSTGRES_DBS; do
   psql -c "CREATE DATABASE $v with OWNER='${VEGA_POSTGRES_USER}'" -c "GRANT ALL PRIVILEGES ON DATABASE $v TO ${VEGA_POSTGRES_USER};"
done