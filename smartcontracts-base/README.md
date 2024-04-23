# smartcontracts-base

Base docker image for vega network smart contracts. Builds all programs required to start local Ethereum network chain.

## Getting started

1. Start ganache

```
docker run \
    --detach \
    -p 8545:8545 \
    vegaprotocol/ganache:v1.4.0 \
        --miner.blockTime 0.5 \
        --chain.chainId 1440 \
        --chain.networkId 1441 \
        -h 0.0.0.0 \
        -p 8545 \
        --wallet.mnemonic "ozone access unlock valid olympic save include omit supply green clown session" \
        --database.dbPath /app/ganache-db \
        --wallet.accountKeysPath keys.json
```

2. Install dependencies

```
npm install
```

3. Export required env variable

```
export GANACHE_MNEMONIC="ozone access unlock valid olympic save include omit supply green clown session"
export VEGA_CONTRACTS_MIGRATION_WORKDIR=$(pwd)
```

4. Run the migrations

```
./run
```

5. As a result you should see the `addresses.json` file in your local directory

```
ls -als

...
  4 drwxr-xr-x   6 daniel daniel   4096 Apr 23 13:25 .
  4 drwxr-xr-x  13 daniel daniel   4096 Apr 23 12:05 ..
  4 -rw-rw-r--   1 daniel daniel   1741 Apr 23 13:30 addresses.json
...
```