FROM timescale/timescaledb:2.7.1-pg14

RUN set -euo; \
    sed -i 's/^restart_after_crash.*//' /usr/local/share/postgresql/postgresql.conf.sample\
        && echo "restart_after_crash=no" >> /usr/local/share/postgresql/postgresql.conf.sample;

COPY 020_add_multiple_databases.sh /docker-entrypoint-initdb.d/