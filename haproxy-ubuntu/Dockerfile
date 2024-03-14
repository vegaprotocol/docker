FROM haproxytech/haproxy-ubuntu:2.9.6
RUN apt-get update \
    && apt-get install -y \
        python3 \
        python3-pip \
    && pip3 install --no-cache \
        requests \
        python-dateutil \
    && rm -rf /var/lib/apt/lists/*
RUN curl https://raw.githubusercontent.com/haproxytech/haproxy-lua-cors/master/lib/cors.lua -o /usr/local/etc/haproxy/cors.lua
