const HDWalletProvider = require("@truffle/hdwallet-provider");
const Web3 = require("web3");

let faucetAbi = require("../../../smartcontracts-base/contracts-solc-0.8.8/build/doc/data/files/abi/Base_Faucet_Token_ABI.json");
let bridgeAbi = require("../../../smartcontracts-base/contracts-solc-0.8.8/build/doc/data/files/abi/ERC20_Bridge_Logic_Restricted_ABI.json");

const contractOwner = {
    'private': 'a37f4c2a678aefb5037bf415a826df1540b330b7e471aa54184877ba901b9ef0',
    'public': '0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F'
};
const vegaTokenEthAddress = '0x67175Da1D5e966e40D11c4B2519392B2058373de';

const vegaWalletPubKey = 'party_name_2'

let w3 = new Web3(new HDWalletProvider({
        "mnemonic": {
            "phrase": 
                "ozone access unlock valid olympic save include omit supply green clown session"
        },
        "providerOrUrl": "http://127.0.0.1:8545",
}));

// create instances of erc20 smart contracts
let vegaInstance = new w3.eth.Contract(faucetAbi, vegaTokenEthAddress);
let bridgeInstance = new w3.eth.Contract(bridgeAbi, '0x9708FF7510D4A7B9541e1699d15b53Ecb1AFDc54')

async function getBalance(contract, address) {
    let result = await contract.methods.balanceOf(address).call();

    return result
}

async function mintToken(contract, addressTo, amount, name) {
    console.log('Minting ' + amount + ' ' + name + ' for ' + addressTo);
    await contract.methods.mint(addressTo, amount).send({
        from: contractOwner.public,
      });
}

async function depositAsset(contract, addressSource, amount, vegaPubKey) {
    let s = await contract.methods.deposit_asset(addressSource, amount, Web3.utils.fromAscii(vegaPubKey)).send({
        from: contractOwner.public,
      });

    console.log(s)
}

const amount = 115000000000

async function main() {
    let initRandomAccVegaBalance = await getBalance(vegaInstance, contractOwner.public);

    console.log("Initial balances for " + contractOwner.public + ": ");
    console.log("  - Vega: " + initRandomAccVegaBalance);
    console.log("");
    console.log("");

    console.log("Mint ")
    await mintToken(vegaInstance, contractOwner.public, amount, 'Vega');

    console.log("Balances for " + contractOwner.public + ": ");
    console.log("  - Random account: " + await getBalance(vegaInstance, contractOwner.public));
    console.log("");
    console.log("");

    await vegaInstance.methods.approve(bridgeInstance.options.address, amount).send({
        from: contractOwner.public,
    })

    console.log("Depositing token from", contractOwner.public, "to " + vegaWalletPubKey)
    await depositAsset(bridgeInstance, vegaInstance.options.address, amount, vegaWalletPubKey)
    console.log("Done")

    console.log("  - Random account after deposit: " + await getBalance(vegaInstance, contractOwner.public));
    console.log("");
    console.log("");

    process.exit(0);
}


main();