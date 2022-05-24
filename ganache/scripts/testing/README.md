# Testing

This folder contains some examples, experiments and scripts that may be useful for other people or for testing.

### test_tokens.js

Script tries to mint some tokens on ganache.

##### Example usage: 

```bash
docker run ... vegaprotocol/ganache:latest ...;
...
...
cd .../ganache/scripts/testing;
node install;
node test_tokens.js
```

##### Example output is:

```bash
Initial balances for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F: 
  - Vega: 6499972300001650000000000
  - tUSDC: 2060000000000
  - tBTC: 2030000400000

Balances for 0xCcFb83E3b85e70Fd815C2C690B21E48F7170926B: 
  - Vega: 1610000000000
  - tUSDC: 1750000000000
  - tBTC: 400000


Testing mint function
Minting 110000000000 Vega for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F
Minting 120000000000 tUSDC for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F
Minting 130000000000 tBTC for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F
Minting 115000000000 Vega for 0xCcFb83E3b85e70Fd815C2C690B21E48F7170926B
Minting 125000000000 tUSDC for 0xCcFb83E3b85e70Fd815C2C690B21E48F7170926B


Testing faucet function
Giving some amount of tUSDC from faucet for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F
Giving some amount of tBTC from faucet for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F
Giving some amount of tBTC from faucet for 0xCcFb83E3b85e70Fd815C2C690B21E48F7170926B


Testing issue function
Issuing 110000000000 Vega for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F
Issuing 120000000000 tUSDC for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F
Issuing 130000000000 tBTC for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F
Issuing 115000000000 Vega for 0xCcFb83E3b85e70Fd815C2C690B21E48F7170926B
Issuing 125000000000 tUSDC for 0xCcFb83E3b85e70Fd815C2C690B21E48F7170926B


Balances for 0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F: 
  - Vega: 6499972300001870000000000
  - tUSDC: 2310000000000
  - tBTC: 2290000500000

Balances for 0xCcFb83E3b85e70Fd815C2C690B21E48F7170926B: 
  - Vega: 1840000000000
  - tUSDC: 2000000000000
  - tBTC: 500000
```