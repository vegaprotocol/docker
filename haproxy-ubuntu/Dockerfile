FROM haproxytech/haproxy-ubuntu:2.7.0
RUN apt-get update \
    && apt-get install -y \
        python3 \
        python3-pip \
    && pip3 install --no-cache \
        requests \
        python-dateutil \
    && rm -rf /var/lib/apt/lists/*
