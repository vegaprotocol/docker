# timescaledb

The timescaledb docker image optimized for the [VEGA](http://vegaprotocol.io/) network. 

## Base image

We base our image on the [timescale/timescaledb](https://hub.docker.com/r/timescale/timescaledb). Visit the [docker hub](https://hub.docker.com/r/timescale/timescaledb) to see more details.

## Custom changes

On top of the official [timescale/timescaledb](https://hub.docker.com/r/timescale/timescaledb) we have changed:

Added the `POSTGRES_DBS` variable that allows you to create multiple databases. Example: `-e POSTGRES_DBS="db1,db2,db3"`
Set the `restart_after_crash` option to `no` in the `postgres.conf`