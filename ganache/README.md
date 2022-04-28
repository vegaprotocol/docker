# ganache

## BIP39

Network is created with the following metadata:

- Mnemonic: `ozone access unlock valid olympic save include omit supply green clown session`
- BIP39 Seed: `e1dc04fb86551bacc618d6cae5cf707a2533937a72b952ee59867ac5ab33fbc0c5595aaf19072d8f771ac28fb159e5d9d2b34e78fddee8cf4159a1063f38023c`

## Usage

```
docker run --name vega-eth-ganache vegaprotocol/ganache:latest ganache-cli \
    --blockTime 1 \
    --chainId 1440 \
    --networkId 1441 \
    -h 0.0.0.0 \
	-m "ozone access unlock valid olympic save include omit supply green clown session" \
	--db /app/ganache-db
```

### Useful information

You can find all deployed smartcontracts addresses with the following command:

```
docker exec vega-eth-ganache cat /app/addresses.json
```

#### Smartcontracts owner

- ETH Public: `0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F`
- ETH Private: `a37f4c2a678aefb5037bf415a826df1540b330b7e471aa54184877ba901b9ef0`

#### Smartcontracts addresses

- MultisigControl: `0xdEcdA30fd3449718304eA201A8f220eBdE25dd1E`
- ERC20 Asset Pool: `0xAa1eDb6C25e6B5ff2c8EdAf68757Ae557178E6eE`
- ERC20 Bridge Logic 1: `0x9708FF7510D4A7B9541e1699d15b53Ecb1AFDc54`
- ERC20 Bridge Logic 2: `0x29e1eA1cfb78f7c34802C90198Cc24aDcBBE4AD0`
- ERC20 Vesting: `0xF41bD86d462D36b997C0bbb4D97a0a3382f205B7`
- Staking Bridge: `0x9135f5afd6F055e731bca2348429482eE614CFfA`

#### ERC20 Tokens

- tBTC
    - ETH: `0xb63D135B0a6854EEb765d69ca36210cC70BECAE0`
    - Vega: `0x5cfa87844724df6069b94e4c8a6f03af21907d7bc251593d08e4251043ee9f7c`
- tDAI
    - ETH: `0x879B84eCA313D62CE4e5ED717939B42cBa9e53cb`
    - Vega: `0x6d9d35f657589e40ddfb448b7ad4a7463b66efb307527fedd2aa7df1bbd5ea61`
- tEURO
    - ETH: `0x7ccE194dAEf2A4e5C23C78C9330D4c907eCA6980`
    - Vega: `0x8b52d4a3a4b0ffe733cddbc2b67be273816cfeb6ca4c8b339bac03ffba08e4e4`
- tUSDC
    - ETH: `0x1b8a1B6CBE5c93609b46D1829Cc7f3Cb8eeE23a0`
    - Vega: `0x993ed98f4f770d91a796faab1738551193ba45c62341d20597df70fea6704ede`
- VEGA
    - ETH: `0x67175Da1D5e966e40D11c4B2519392B2058373de`
    - Vega: `0xb4f2726571fbe8e33b442dc92ed2d7f0d810e21835b7371a7915a365f07ccd9b`
- VEGAv1
    - ETH: `0x8fa21D653C1bF17741055f00dD55663Bc52a8362`
    - Vega: `0xc1607f28ec1d0a0b36842c8327101b18de2c5f172585870912f5959145a9176c`

#### Available accounts

```
Available Accounts
==================
(0) 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F (100 ETH)
(1) 0xCcFb83E3b85e70Fd815C2C690B21E48F7170926B (100 ETH)
(2) 0xD1813b218cDD4ab2E25Ca0c502Aa773A066A14eE (100 ETH)
(3) 0x55Bb84D3FA3A1D8Bd86643F66eaCdfde812fAce5 (100 ETH)
(4) 0x9BF07EC0474b84747aA496BC4776152C054E2187 (100 ETH)
(5) 0xfA88929915924c7deFEAf92A8266c0C8Abc3003a (100 ETH)
(6) 0x2397CbF493E640ef1019114D46A3F63f4E72e68a (100 ETH)
(7) 0x85e42DC1ACd19e1e9A434b2C4CE40232C86ea39D (100 ETH)
(8) 0xeA13f56532dED7Fa42b8008AeA96f5a4B3Da5E06 (100 ETH)
(9) 0x0b755ae366FE0368F5b0D1879002B06e9Af40f26 (100 ETH)

Private Keys
==================
(0) 0xa37f4c2a678aefb5037bf415a826df1540b330b7e471aa54184877ba901b9ef0
(1) 0x394fd255dd4a391531ab8258de6e9e1650a200dcd6260706bea5e5dfa52fa020
(2) 0xb26e103455cfa29fbbd8f856b2fa67cbe68cd6ce3d8e708c9080a957e6c45ecb
(3) 0x421a3441670c2c27fd0bcb146c729e9c4c8958b5622ac75112697581b33a7d67
(4) 0xdfbd82abc9c8a804c565cd31417683d5f013c27fb18d323925bc5ce9a9cd725d
(5) 0xe4c564e0caee5d16c05e6c886ad969e708ff89c9ebba8ac642703163840996ca
(6) 0x1d55412ce6e92d3d427928742946b451888414c5ab7520781dc580a20f1c28c7
(7) 0xafe378ddf5ff1092b8653cb37b0d9df5585e0ee475ea0f96844db003d8a0e279
(8) 0x9770b0443b973e72cd74b5f87f53fe53f2dba5aee8b0a110789eb6d3d9c70ec6
(9) 0xc8b01eaa62245b9c229cecc6bc306324c066767eeefda82cf9d3764eb1bae2b7
```