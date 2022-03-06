ARG ganache_mnemonic="ozone access unlock valid olympic save include omit supply green clown session"

FROM vegaprotocol/smartcontracts-base:v1.0.1
ARG ganache_mnemonic

RUN nohup \
        bash -c "ganache-cli \
            --blockTime 1 \
            --chainId 1440 \
            --networkId 1441 \
            -m \"${ganache_mnemonic}\" \
            --db /app/ganache-db \
            -h 0.0.0.0 \
            -p 8545 &" \
    && GANACHE_MNEMONIC="${ganache_mnemonic}" GANACHE_HOSTNAME=127.0.0.1 ./run