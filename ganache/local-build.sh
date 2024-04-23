#!/bin/bash


cd ../smartcontracts-base
docker build -t vegaprotocol/smartcontracts-base:v1.5.0 .

cd ../ganache
docker build --progress plain -t vegaprotocol/ganache:v1.5.0 .


docker run vegaprotocol/ganache:v1.5.0 \
        --miner.blockTime 0.5 \
        --chain.chainId 1440 \
        --chain.networkId 1441 \
        -h 0.0.0.0 \
        -p 8545 \
        --wallet.mnemonic "ozone access unlock valid olympic save include omit supply green clown session" \
        --database.dbPath /app/ganache-db \
        --wallet.accountKeysPath keys.json