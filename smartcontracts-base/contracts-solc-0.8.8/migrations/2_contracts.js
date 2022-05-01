const abi = require("ethereumjs-abi");
const crypto = require("crypto");
const ethUtil = require("ethereumjs-util");
const fs = require("fs");
const ethers = require('ethers');

const Base_Faucet_Token = artifacts.require("Base_Faucet_Token");
const MultisigControl = artifacts.require("MultisigControl");
const ERC20_Asset_Pool = artifacts.require("ERC20_Asset_Pool");
const ERC20_Bridge_Logic_Restricted = artifacts.require("ERC20_Bridge_Logic_Restricted");

function multisign(
  param_types,
  params,
  function_name,
  sender,
  validator_privkeys
) {
  let nonce = new ethUtil.BN(crypto.randomBytes(32));
  params.push(nonce);
  param_types.push("uint256");
  params.push(function_name);
  param_types.push("string");
  let encoded = abi.rawEncode(
    ["bytes", "address"],
    [abi.rawEncode(param_types, params), sender]
  );
  let msg_hash = ethUtil.keccak256(encoded);
  let sigs = "0x";
  for (let privkey of validator_privkeys) {
    let sig = ethUtil.ecsign(msg_hash, privkey);
    sigs += sig.r.toString("hex");
    sigs += sig.s.toString("hex");
    sigs += sig.v.toString(16);
  }
  return {
    nonce: nonce,
    sigs: sigs,
  };
}

async function erc20_asset_pool_set_bridge_address(
  erc20_asset_pool_instance,
  bridge_address,
  validator_privkeys // list
) {
  let ms = multisign(
    ["address"],
    [bridge_address],
    "set_bridge_address",
    erc20_asset_pool_instance.address,
    validator_privkeys
  );
  await erc20_asset_pool_instance.set_bridge_address(
    bridge_address,
    ms.nonce,
    ms.sigs
  );
}

async function list_asset_on_bridge(
  bridge_instance,
  asset_address,
  cfg,
  validator_privkeys
) {
  // function list_asset(address asset_source, bytes32 vega_asset_id, uint256 lifetime_limit, uint256 withdraw_threshold, uint256 nonce, bytes memory signatures) public virtual;
  let ms = multisign(
    ["address", "bytes32", "uint256", "uint256"],
    [asset_address, cfg.vega_id, cfg.lifetime_limit, cfg.withdraw_threshold],
    "list_asset",
    bridge_instance.address,
    validator_privkeys
  );

  return await bridge_instance.list_asset(
    asset_address,
    cfg.vega_id, 
    cfg.lifetime_limit, 
    cfg.withdraw_threshold,
    ms.nonce,
    ms.sigs
  );
}

module.exports = async function (deployer) {
  // Contracts
  await deployer.deploy(MultisigControl);
  await deployer.deploy(ERC20_Asset_Pool, MultisigControl.address);
  let erc20_bridge_1 = await deployer.deploy(
    ERC20_Bridge_Logic_Restricted,
    ERC20_Asset_Pool.address
  );
  let erc20_bridge_2 = await deployer.deploy(
    ERC20_Bridge_Logic_Restricted,
    ERC20_Asset_Pool.address
  );
  let erc20_asset_pool_instance = await ERC20_Asset_Pool.deployed();

  const ganmacheMnemonic = process.env.GANACHE_MNEMONIC;

  if (!ganmacheMnemonic || ganmacheMnemonic.length <= 0) {
    console.log("Error: GANACHE_MNEMONIC env variable not set.")
    process.exit(2);
  }

  const wallet = ethers.Wallet.fromMnemonic(ganmacheMnemonic)
  let privkey = wallet.privateKey.substr(2);
  let pubkey = wallet.address;

  let initial_validators = [Buffer.from(privkey, "hex")];

  await erc20_asset_pool_set_bridge_address(
    erc20_asset_pool_instance,
    erc20_bridge_1.address,
    initial_validators
  );

  addresses = {
    addr0: { priv: privkey, pub: pubkey },
    MultisigControl: { Ethereum: MultisigControl.address },
    ERC20_Asset_Pool: { Ethereum: ERC20_Asset_Pool.address },
    erc20_bridge_1: { Ethereum: erc20_bridge_1.address },
    erc20_bridge_2: { Ethereum: erc20_bridge_2.address },
  };

  // Tokens
  let token_config = require("./token_config.json");
  for (let i = 0; i < token_config.length; i++) {
    let token_contract = await deployer.deploy(
      Base_Faucet_Token,
      token_config[i].name,
      token_config[i].symbol,
      token_config[i].decimals,
      token_config[i].total_supply_whole_tokens,
      token_config[i].faucet_amount
    );
    addresses[token_config[i].symbol] = {
      Ethereum: token_contract.address,
      Vega: token_config[i].vega_id,
    };
  }

  // Listing tokens on bridge
  for (let i = 0; i < token_config.length; i++) {
    try {
      let result = await list_asset_on_bridge(
        erc20_bridge_1,
        addresses[token_config[i].symbol].Ethereum,
        token_config[i],
        initial_validators
      );
      console.log(
        `Listed ${token_config[i].symbol} block ${result.receipt.blockNumber} ` +
          `tx ${result.receipt.transactionHash} gas ${result.receipt.cumulativeGasUsed}`
      );
    } catch (e) {
      console.log(
        "Caught an exception trying to list " +
          token_config[i].symbol +
          ": " +
          e
      );
      process.exit(1);
    }
  }

  // New migrations go just above this comment.

  // Print and save the address data
  console.log(addresses)
  fs.writeFileSync(
    "/app/addresses.json",
    JSON.stringify(addresses, null, 2)
  );
};
