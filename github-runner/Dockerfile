FROM summerwind/actions-runner:v2.301.1-ubuntu-22.04
USER root

RUN curl -s https://deb.nodesource.com/setup_16.x | bash \
    && curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - \
    && echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list \
    && add-apt-repository ppa:longsleep/golang-backports \
    && apt-get update \
    && apt-get install -y \
    nodejs\
    yarn \
    libgtk2.0-0 \
    libgtk-3-0 \
    libgbm-dev \
    libnotify-dev \
    libgconf-2-4 \
    libnss3 \
    libxss1 \
    libasound2 \
    libxtst6 \
    xauth \
    xvfb \
    libnss3-dev \
    wget \
    golang-go \
    && wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && apt-get -f -y install ./google-chrome-stable_current_amd64.deb \
    && rm google-chrome-stable_current_amd64.deb \
    && rm -rf /var/lib/apt/lists/* 

# ENV variables are here to allow cache previous step and speet the build up
ENV VEGA_VERSION=v0.70.1
ENV VEGACAPSULE_VERSION=v0.70.1
ENV NOMAD_VERSION=1.3.1

# Install vegacapsule
RUN wget -q https://github.com/vegaprotocol/vegacapsule/releases/download/$VEGACAPSULE_VERSION/vegacapsule-linux-amd64.zip \
    && unzip vegacapsule-linux-amd64.zip \
    && mv vegacapsule /usr/local/bin/ \
    && rm vegacapsule-linux-amd64.zip

# Install Nomad
RUN wget -q https://releases.hashicorp.com/nomad/${NOMAD_VERSION}/nomad_${NOMAD_VERSION}_linux_amd64.zip \
    && unzip nomad_${NOMAD_VERSION}_linux_amd64.zip \
    && mv nomad /usr/local/bin/ \
    && rm nomad_${NOMAD_VERSION}_linux_amd64.zip

# Install Vega CLI
RUN wget https://github.com/vegaprotocol/vega/releases/download/$VEGA_VERSION/vega-linux-amd64.zip \
    && unzip vega-linux-amd64.zip \
    && mv vega /usr/local/bin/ \
    && rm vega-linux-amd64.zip 

RUN chmod +x /usr/local/bin/vega /usr/local/bin/vegacapsule /usr/local/bin/nomad

USER runner
ENV GOBIN=$HOME/go/bin \
    PATH=$PATH:$HOME/go/bin
RUN mkdir -p $GOBIN \
    && mkdir -p ~/.vegacapsule \
    && ln -s /usr/local/bin/nomad ~/.vegacapsule/nomad_${NOMAD_VERSION}

RUN vegacapsule version \
    && vega version \
    && nomad version
