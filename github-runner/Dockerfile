FROM summerwind/actions-runner:v2.296.2-ubuntu-20.04
USER root
RUN apt-get update \
    && apt-get install -y \
        yarn \
        xvfb \
    && rm -rf /var/lib/apt/lists/*
USER runner
