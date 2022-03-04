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

// Main function
async function run_deposits() {
  try {
    let cfg = require(configfilename);
    let abi = require(cfg.token_abi_filename);
    let w3 = new Web3(new HDWalletProvider(cfg.HDWalletProvider));
    let accounts = await w3.eth.getAccounts();
    let from_addr = accounts[0];
    let balance_pre = Web3.utils.fromWei(await w3.eth.getBalance(from_addr));
    console.log("balance: " + balance_pre);
    for (let k = 0; k < cfg.bridges.length && run; k++) {
      let bridge_cfg = cfg.bridges[k];
      // Object.entries(cfg.bridges).forEach(async (bridge, index) => {
      for (let j = 0; j < bridge_cfg.tokens.length && run; j++) {
        let token_cfg = bridge_cfg.tokens[j];
        let token_instance = new w3.eth.Contract(abi, token_cfg.address);
        let r = token_cfg.recipients;
        console.log(
          `bridge:${bridge_cfg.type} token:${token_cfg.name} (${token_cfg.address}) ` +
            `amount:${token_cfg.amount} total_recipients:${r.length}`
        );
        for (let i = 0; i < r.length && run; i += cfg.ethereum.chunksize) {
          let recip_slice = r.slice(i, i + cfg.ethereum.chunksize);
          let admin_bulk_func;
          switch (bridge_cfg.type) {
            case "erc20":
              admin_bulk_func = token_instance.methods.admin_deposit_bulk;
              break;
            case "stake":
              admin_bulk_func = token_instance.methods.admin_stake_bulk;
              break;
            default:
              console.log(`Invalid bridge type: ${bridge_cfg.type}`);
              process.exit(1);
          }
          let result = await admin_bulk_func(
            token_cfg.amount,
            bridge_cfg.address,
            recip_slice
          ).send({
            from: from_addr,
            gasPrice: cfg.ethereum.gasPrice,
            gas: cfg.ethereum.gas,
          });

          console.log(
            `bridge:${bridge_cfg.type} token:${token_cfg.name} (${token_cfg.address}) ` +
              `block:${result.blockNumber} tx:${result.transactionHash} ` +
              `gas:${result.cumulativeGasUsed} recipients:${recip_slice.length}`
          );
          if (cfg.ethereum.sleep > 0 && i < r.length - cfg.ethereum.chunksize) {
            console.log("sleep " + cfg.ethereum.sleep + "ms");
            await sleep(cfg.ethereum.sleep);
          }
        } // i
        if (
          cfg.ethereum.sleep > 0 &&
          r.length > 0 &&
          j < bridge_cfg.tokens.length - 1
        ) {
          console.log("sleep " + cfg.ethereum.sleep + " ms");
          await sleep(cfg.ethereum.sleep);
        }
      } // j
    } // k // );
    let balance_post = Web3.utils.fromWei(await w3.eth.getBalance(from_addr));
    console.log(`balance: ${balance_post}`);
    console.log(
      "cost: " + Math.round((balance_pre - balance_post) * 10000) / 10000
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

run_deposits();
