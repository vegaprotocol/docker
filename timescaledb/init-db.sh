#!/bin/bash

source /usr/local/bin/docker-entrypoint.sh

printenv
docker_setup_env
docker_create_db_directories
docker_verify_minimum_env
ls /docker-entrypoint-initdb.d/ > /dev/null
docker_init_database_dir
pg_setup_hba_conf "postgres"

export PGPASSWORD="${PGPASSWORD:-$POSTGRES_PASSWORD}"
docker_temp_server_start "postgres"

docker_setup_db
docker_process_init_files /docker-entrypoint-initdb.d/*

docker_temp_server_stop
unset PGPASSWORD