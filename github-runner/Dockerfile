FROM summerwind/actions-runner:v2.299.1-ubuntu-20.04
USER root
RUN apt-get update \
    && apt-get install -y \
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
    && wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && apt-get -f -y install ./google-chrome-stable_current_amd64.deb \
    && rm google-chrome-stable_current_amd64.deb \
    && wget -q https://go.dev/dl/go1.19.linux-amd64.tar.gz \
    && tar -xvf go1.19.linux-amd64.tar.gz \
    && rm go1.19.linux-amd64.tar.gz \
    && mv go /usr/local \
    && rm -rf /var/lib/apt/lists/* 
USER runner
RUN echo 'export GOBIN=/usr/local/go/bin' >> $HOME/.profile \
    && echo 'export GOPATH=$HOME/go' >> $HOME/.profile \
    && echo 'export PATH=$PATH:$GOBIN:$GOPATH/bin' >> $HOME/.profile \
    && source /home/runner/.profile
