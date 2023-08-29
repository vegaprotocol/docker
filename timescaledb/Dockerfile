FROM timescale/timescaledb:2.8.1-pg14-bitnami

ENV TIMESCALEDB_TELEMETRY "off"
ENV NO_TS_TUNE true
ENV POSTGRESQL_ALLOW_REMOTE_CONNECTIONS true
ENV POSTGRESQL_PGHBA_REMOVE_FILTERS true
ENV VEGA_POSTGRES_USER "vega"
ENV VEGA_POSTGRES_PASSWORD "vega"
ENV POSTGRES_DBS "vega0,vega1,vega2,vega3,vega4,vega5,vega6,vega7,vega8,vega9,vega10,vega11,vega12,vega13,vega14,vega15"


# RUN apt-get install sed; \
#     set -euo; \
#     sed -i 's/^restart_after_crash.*//' /usr/local/share/postgresql/postgresql.conf.sample\
#         && echo "restart_after_crash=no" >> /usr/local/share/postgresql/postgresql.conf.sample;

COPY 020_add_multiple_databases.sh /docker-entrypoint-initdb.d/