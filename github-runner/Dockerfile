FROM summerwind/actions-runner:v2.299.1-ubuntu-20.04
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
    # && wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    # && apt-get -f -y install ./google-chrome-stable_current_amd64.deb \
    # && rm google-chrome-stable_current_amd64.deb \
    && rm -rf /var/lib/apt/lists/* 
USER runner
