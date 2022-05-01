const abi = require("ethereumjs-abi");
const crypto = require("crypto");
const ethUtil = require("ethereumjs-util");
const HDWalletProvider = require("@truffle/hdwallet-provider");
const Web3 = require("web3");

let configfilename;
for (let i = 2; i < process.argv.length; i++) {
  switch (process.argv[i]) {
    case "--config":
      configfilename = process.argv[i + 1];
      i++;
      break;
    default:
      console.log("Unrecognised argument: " + process.argv[i]);
      process.exit(1);
  }
}
if (configfilename === undefined) {
  console.log("Please specify a config filename with --config");
  process.exit(1);
}

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

// Copied from migrations/2_contracts.js
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

async function set_threshold(
  multisigcontrol_instance,
  new_threshold,
  sender_address,
  validator_privkeys,
  gas,
  gasPrice
) {
  let ms = multisign(
    ["uint16"],
    [new_threshold],
    "set_threshold",
    sender_address,
    validator_privkeys
  );
  await multisigcontrol_instance.methods
    .set_threshold(new_threshold, ms.nonce, ms.sigs)
    .send({
      gas: gas,
      gasPrice: gasPrice,
      from: sender_address,
    });
}

async function add_signer(
  multisigcontrol_instance,
  new_signer,
  sender_address,
  validator_privkeys,
  gas,
  gasPrice
) {
  let ms = multisign(
    ["address"],
    [new_signer],
    "add_signer",
    sender_address,
    validator_privkeys
  );
  await multisigcontrol_instance.methods
    .add_signer(new_signer, ms.nonce, ms.sigs)
    .send({
      gas: gas,
      gasPrice: gasPrice,
      from: sender_address,
    });
}

async function remove_signer(
  multisigcontrol_instance,
  old_signer,
  sender_address,
  validator_privkeys,
  gas,
  gasPrice
) {
  let ms = multisign(
    ["address"],
    [old_signer],
    "remove_signer",
    sender_address,
    validator_privkeys
  );
  await multisigcontrol_instance.methods
    .remove_signer(old_signer, ms.nonce, ms.sigs)
    .send({
      gas: gas,
      gasPrice: gasPrice,
      from: sender_address,
    });
}

// Main function
async function add_signers() {
  // ASSUMPTION: The contract creator is a signer.
  try {
    let cfg = require(configfilename);
    let abi = require(cfg.multisigcontrol.abi_filename);
    let w3 = new Web3(new HDWalletProvider(cfg.HDWalletProvider));
    let multisigcontrol_instance = new w3.eth.Contract(
      abi,
      cfg.multisigcontrol.address
    );

    // Set the threshold to 1/1000 so we only need one signature for later operations.
    await set_threshold(
      multisigcontrol_instance,
      1, // 0 < threshold <= 1000
      cfg.source.pubkey,
      [Buffer.from(cfg.source.privkey, "hex")],
      cfg.ethereum.gas,
      cfg.ethereum.gasPrice
    );

    for (let i = 0; i < cfg.new_signers.length && run; i++) {
      // Add a signer
      await add_signer(
        multisigcontrol_instance,
        cfg.new_signers[i],
        cfg.source.pubkey,
        [Buffer.from(cfg.source.privkey, "hex")],
        cfg.ethereum.gas,
        cfg.ethereum.gasPrice
      );
      console.log("Added MultisigControl signer: " + cfg.new_signers[i]);
    }

    // remove original signer so that only validators are on the
    await remove_signer(
      multisigcontrol_instance,
      cfg.source.pubkey,
      cfg.source.pubkey,
      [Buffer.from(cfg.source.privkey, "hex")],
      cfg.ethereum.gas,
      cfg.ethereum.gasPrice
    );
    console.log("Removed MultisigControl signer: " + cfg.source.pubkey);

    // Finally, set the threshold to the proper value
    await set_threshold(
      multisigcontrol_instance,
      500, // 0 < threshold <= 1000
      cfg.source.pubkey,
      [Buffer.from(cfg.new_priv_signers[0], "hex")],
      cfg.ethereum.gas,
      cfg.ethereum.gasPrice
    );

    process.exit(0);
  } catch (e) {
    console.error("Caught an exception: " + e + JSON.stringify(e));
    process.exit(1);
  }
}

let run = true;

process.on("SIGTERM", () => {
  console.info("Caught SIGTERM");
  run = false;
});

process.on("SIGINT", function () {
  console.log("Caught SIGINT");
  run = false;
});

add_signers();
