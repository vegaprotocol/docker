# timescaledb

The timescaledb docker image optimized for the [VEGA](http://vegaprotocol.io/) network.

## Custom changes

On top of the official [timescale/timescaledb](https://hub.docker.com/r/timescale/timescaledb) we have changed:

- Pre run optimizations with the following parameters:
  - Memory: 4GB
  - CPUs: 2
  - WAL disk size: 10GB
  - Max background workers: 50
  - Max connections: 200
- Created 30 databases with the following names: vega0, vega1, vega2, ..., vega29
- TimescaleDB telemetry is off
- Updated PostgreSQL settings: 
  - restart_after_crash: no


Because everything is initialized when the docker container is built, you cannot use the environment variables from the base PostgreSQL image.
