# ganache

## BIP39

Network is created with the following metadata:

- Mnemonic: `ozone access unlock valid olympic save include omit supply green clown session`
- BIP39 Seed: `e1dc04fb86551bacc618d6cae5cf707a2533937a72b952ee59867ac5ab33fbc0c5595aaf19072d8f771ac28fb159e5d9d2b34e78fddee8cf4159a1063f38023c`

## Usage

```
docker run vegaprotocol/ganache:latest ganache-cli \
    --blockTime 1 \
    --chainId 1440 \
    --networkId 1441 \
    -h 0.0.0.0 \
    -p 8545 \
	-m "ozone access unlock valid olympic save include omit supply green clown session" \
	--db /app/ganache-db
```