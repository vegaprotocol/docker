# FROM postgres:14-bullseye
FROM timescale/timescaledb:2.14.2-pg14

ENV TIMESCALEDB_TELEMETRY "off"
ENV NO_TS_TUNE true
ENV POSTGRES_ALLOW_REMOTE_CONNECTIONS true
ENV POSTGRES_PGHBA_REMOVE_FILTERS true
ENV POSTGRES_PASSWORD "vega"


COPY ./init-db.sh /init-db.sh
COPY 020_add_multiple_databases.sh /docker-entrypoint-initdb.d/

RUN apk add sudo \
    && sudo -u postgres --preserve-env /init-db.sh
RUN timescaledb-tune \
        --yes \
        --memory="4GB" \
        --cpus=2 \
        --wal-disk-size="10GB" \
        --max-bg-workers 50 \
        --max-conns 200;

RUN set -euo; \
    sed -i 's/^restart_after_crash.*//' /var/lib/postgresql/data/postgresql.conf \
        && echo "restart_after_crash=no" >> /var/lib/postgresql/data/postgresql.conf \
    && sed -i 's/^log_statement.*//' /var/lib/postgresql/data/postgresql.conf \
        && echo "log_statement='all'" >> /var/lib/postgresql/data/postgresql.conf \
    ;