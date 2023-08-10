FROM python:3.11.4-slim

ENV VEGA_MARKET_SIM_BRANCH="research-bots"
ENV POETRY_VERSION=1.2.2
ENV WALLET_VERSION=v0.71.2

ENV VEGA_WALLET_HOME=/opt/home
ENV VEGA_WALLET_PATH=/bin/vegawallet
ENV VEGA_USER_WALLET_NAME=vegamarketsim
ENV VEGA_WALLET_TOKENS_PASSPHRASE_FILE=/assets/wallet-passphrase.txt
ENV VEGA_WALLET_TOKENS_FILE=/vega_market_sim/wallet-info.json

RUN apt-get update \
    && apt-get install -y \
        git \
        unzip \
        wget \
        jq \
        make \
        rsync

RUN  pip install "poetry==$POETRY_VERSION" \
    && mkdir /vega_market_sim

WORKDIR /vega_market_sim

# get single commit for vega-market-sim
RUN git init \
    && git remote add origin "https://github.com/vegaprotocol/vega-market-sim.git" \
    && git fetch -q --depth=1 origin "$VEGA_MARKET_SIM_BRANCH" \
    && git reset --hard FETCH_HEAD \
    && poetry install \
    \
    && make networks

COPY ./assets /assets


# Prepare wallet and load pre generated key with known parties
RUN wget -q https://github.com/vegaprotocol/vega/releases/download/$WALLET_VERSION/vegawallet-linux-amd64.zip \
    && unzip vegawallet-linux-amd64.zip \
    && rm -f vegawallet-linux-amd64.zip \
    && mv vegawallet /bin/vegawallet \
    && chmod a+x /bin/vegawallet \
    && vegawallet software version \
    \
    && vegawallet software version \
    && vegawallet init --home $VEGA_WALLET_HOME \
    && cp /assets/wallets/vegamarketsim $VEGA_WALLET_HOME/data/wallets/ \
    && vegawallet key list \
        --wallet vegamarketsim \
        --home $VEGA_WALLET_HOME \
        --output json \
        --passphrase-file /assets/wallet-passphrase.txt \
    | jq '.keys' | tee wallets.json \
    \
    \
    && vegawallet api-token init \
        --home $VEGA_WALLET_HOME \
        --output=json \
        --passphrase-file /assets/wallet-passphrase.txt > api-token-init.json \
    && vegawallet api-token generate \
        --home $VEGA_WALLET_HOME \
        --wallet-name vegamarketsim \
        --wallet-passphrase-file /assets/wallet-passphrase.txt \
        --tokens-passphrase-file /assets/wallet-passphrase.txt \
        --output=json > api-token-generate.json \
    \
    && export WALLET_TOKEN=$(jq -r '.token' api-token-generate.json) \
    && echo "{\"$VEGA_USER_WALLET_NAME\": \"${WALLET_TOKEN}\"}" > $VEGA_WALLET_TOKENS_FILE \
    && cat wallet-info.json api-token-generate.json

RUN vegawallet network import \
        --home $VEGA_WALLET_HOME \
        --from-file ./extern/networks-internal/stagnet1/vegawallet-stagnet1.toml \
    && vegawallet network import \
        --home $VEGA_WALLET_HOME \
        --from-file ./extern/networks-internal/mainnet-mirror/vegawallet-mainnet-mirror.toml \
    && vegawallet network import \
        --home $VEGA_WALLET_HOME \
        --from-file ./extern/networks-internal/fairground/vegawallet-fairground.toml \
    && vegawallet network import \
        --home $VEGA_WALLET_HOME \
        --from-file ./extern/networks-internal/devnet1/vegawallet-devnet1.toml 

ENTRYPOINT ["poetry", "run"]