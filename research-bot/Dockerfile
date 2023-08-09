FROM python:3.11.4-slim

ENV VEGA_MARKET_SIM_BRANCH="develop"
ENV POETRY_VERSION=1.2.2
ENV WALLET_VERSION=v0.71.2
env WALLET_HOME=/opt/home

RUN apt-get update \
    && apt-get install -y git unzip wget jq

RUN  pip install "poetry==$POETRY_VERSION" \
    && mkdir /vega_market_sim

WORKDIR /vega_market_sim

RUN git clone https://github.com/vegaprotocol/vega-market-sim.git . \
    && ls -als /vega_market_sim \
    && poetry install

COPY ./assets /assets

RUN wget -q https://github.com/vegaprotocol/vega/releases/download/$WALLET_VERSION/vegawallet-linux-amd64.zip \
    && unzip vegawallet-linux-amd64.zip \
    && rm -f vegawallet-linux-amd64.zip \
    && mv vegawallet /bin/vegawallet \
    && chmod a+x /bin/vegawallet \
    && vegawallet software version \
    \
    && vegawallet software version \
    && vegawallet init --home $WALLET_HOME \
    && cp /assets/wallets/vegamarketsim $WALLET_HOME/data/wallets/ \
    && vegawallet key list \
        --wallet vegamarketsim \
        --home $WALLET_HOME \
        --output json \
        --passphrase-file /assets/wallet-passphrase.txt \
    | jq '.keys' | tee wallets.json