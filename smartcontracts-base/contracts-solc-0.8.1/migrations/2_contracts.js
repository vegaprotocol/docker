const fs = require("fs");

const ERC20_Vesting = artifacts.require("ERC20_Vesting");
const Vega_Staking_Bridge = artifacts.require("Vega_Staking_Bridge");

module.exports = async function (deployer) {


  const workDir = process.env.VEGA_CONTRACTS_MIGRATION_WORKDIR || "/app";


  // should already have addresses deployed from before
  let rawdata = fs.readFileSync(workDir + '/addresses.json');
  let addresses = JSON.parse(rawdata);
  
  await deployer.deploy(
    ERC20_Vesting,
    addresses["VEGAv1"].Ethereum,
    addresses["VEGA"].Ethereum,
    [],
    []
  );
  
  
  // setup the staking bridge.
  await deployer.deploy(
    Vega_Staking_Bridge,
    addresses["VEGA"].Ethereum
    );

  
  // add contract addresses
  addresses["erc20_vesting"] = { Ethereum: ERC20_Vesting.address };
  addresses["staking_bridge"] = { Ethereum: Vega_Staking_Bridge.address };
  
  // Print and save the address data
  console.log(addresses)
  fs.writeFileSync(
    workDir + "/addresses.json",
    JSON.stringify(addresses, null, 2)
  );
};
