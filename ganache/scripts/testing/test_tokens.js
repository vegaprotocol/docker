const HDWalletProvider = require("@truffle/hdwallet-provider");
const Web3 = require("web3");
var assert = require('assert');


let abi = require("./../../../smartcontracts-base/contracts-solc-0.8.8/build/doc/data/files/abi/Base_Faucet_Token_ABI.json");
const contractOwner = {
    'private': 'a37f4c2a678aefb5037bf415a826df1540b330b7e471aa54184877ba901b9ef0',
    'public': '0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F'
};
const randomAccount = '0xCcFb83E3b85e70Fd815C2C690B21E48F7170926B';

const vegaAddress = '0xFcE3C7CBba976414621887F2D762e7fB0f90b5c1';
const tUSDCEthAddress = '0x9dFD86Aa69c1ad549568c9d51bf2dB27cbDeF2cD';
const tBTCEthAddress = '0x67175Da1D5e966e40D11c4B2519392B2058373de';


// let w3 = new Web3(new HDWalletProvider([contractOwner.private], "http://127.0.0.1:8545"));
let w3 = new Web3(new HDWalletProvider({
        "mnemonic": {
            "phrase": 
                "ozone access unlock valid olympic save include omit supply green clown session"
        },
        "providerOrUrl": "http://localhost:8545",
}));

// create instances of erc20 smart contracts
let vegaInstance = new w3.eth.Contract(abi, vegaAddress);
let tUSDCInstance = new w3.eth.Contract(abi, tUSDCEthAddress);
let tBTCInstance = new w3.eth.Contract(abi, tBTCEthAddress);


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


async function issueToken(contract, addressTo, amount, name) {
    console.log('Issuing ' + amount + ' ' + name + ' for ' + addressTo);
    await contract.methods.issue(addressTo, amount).send({
        from: contractOwner.public,
      });
}

async function faucet(contract, addressTo, name) {
    console.log('Giving some amount of ' + name + ' from faucet for ' + addressTo);
    await contract.methods.faucet().send({
        from: addressTo,
      });
}

async function main() {
    let initOwnerVegaBalance = await getBalance(vegaInstance, contractOwner.public);
    let initOwnerTUSDCBalance = await getBalance(tUSDCInstance, contractOwner.public);
    let initOwnerTBTCBalance = await getBalance(tBTCInstance, contractOwner.public);

    let initRandomAccVegaBalance = await getBalance(vegaInstance, randomAccount);
    let initRandomAccTUSDCBalance = await getBalance(tUSDCInstance, randomAccount);
    let initRandomAccTBTCBalance = await getBalance(tBTCInstance, randomAccount);


    console.log("Initial balances for " + contractOwner.public + ": ");
    console.log("  - Vega: " + initOwnerVegaBalance);
    console.log("  - tUSDC: " + initOwnerTUSDCBalance);
    console.log("  - tBTC: " + initOwnerTBTCBalance);
    console.log("");

    console.log("Balances for " + randomAccount + ": ");
    console.log("  - Vega: " + initRandomAccVegaBalance);
    console.log("  - tUSDC: " + initRandomAccTUSDCBalance);
    console.log("  - tBTC: " + initRandomAccTBTCBalance);
    console.log("");
    console.log("");

    console.log("Testing mint function")
    await mintToken(vegaInstance, contractOwner.public, 110000000000, 'Vega');
    await mintToken(tUSDCInstance, contractOwner.public, 120000000000, 'tUSDC');
    await mintToken(tBTCInstance, contractOwner.public, 130000000000, 'tBTC');

    await mintToken(vegaInstance, randomAccount, 115000000000, 'Vega');
    await mintToken(tUSDCInstance, randomAccount, 125000000000, 'tUSDC');


    let newOwnerVegaBalance = await getBalance(vegaInstance, contractOwner.public);
    let newOwnerTUSDCBalance = await getBalance(tUSDCInstance, contractOwner.public);
    let newOwnerTBTCBalance = await getBalance(tBTCInstance, contractOwner.public);

    let newRandomAccVegaBalance = await getBalance(vegaInstance, randomAccount);
    let newRandomAccTUSDCBalance = await getBalance(tUSDCInstance, randomAccount);
    let newRandomAccTBTCBalance = await getBalance(tBTCInstance, randomAccount);
    

    assert(BigInt(newOwnerVegaBalance) == (BigInt(initOwnerVegaBalance) + BigInt(110000000000)));
    assert(BigInt(newOwnerTUSDCBalance) == (BigInt(initOwnerTUSDCBalance) + BigInt(120000000000)));
    assert(BigInt(newOwnerTBTCBalance) == (BigInt(initOwnerTBTCBalance) + BigInt(130000000000)));

    assert(BigInt(newRandomAccVegaBalance) ==  (BigInt(initRandomAccVegaBalance) + BigInt(115000000000)));
    assert(BigInt(newRandomAccTUSDCBalance) ==  (BigInt(initRandomAccTUSDCBalance) + BigInt(125000000000)));
    assert(BigInt(newRandomAccTBTCBalance) ==  (BigInt(initRandomAccTBTCBalance)));
    console.log("");
    console.log("");


    console.log('Testing faucet function');

    initOwnerVegaBalance = newOwnerVegaBalance;
    initOwnerTUSDCBalance = newOwnerTUSDCBalance;
    initOwnerTBTCBalance = newOwnerTBTCBalance;

    initRandomAccVegaBalance = newRandomAccVegaBalance;
    initRandomAccTUSDCBalance = newRandomAccTUSDCBalance;
    initRandomAccTBTCBalance = newRandomAccTBTCBalance;

    await faucet(tUSDCInstance, contractOwner.public, 'tUSDC');
    await faucet(tBTCInstance, contractOwner.public, 'tBTC');
    await faucet(tBTCInstance, randomAccount, 'tBTC');

    newOwnerVegaBalance = await getBalance(vegaInstance, contractOwner.public);
    newOwnerTUSDCBalance = await getBalance(tUSDCInstance, contractOwner.public);
    newOwnerTBTCBalance = await getBalance(tBTCInstance, contractOwner.public);

    assert(BigInt(initOwnerVegaBalance) == BigInt(newOwnerVegaBalance));
    assert(BigInt(initOwnerTUSDCBalance) < BigInt(newOwnerTUSDCBalance));
    assert(BigInt(initOwnerTBTCBalance) < BigInt(newOwnerTBTCBalance));


    initOwnerVegaBalance = newOwnerVegaBalance;
    initOwnerTUSDCBalance = newOwnerTUSDCBalance;
    initOwnerTBTCBalance = newOwnerTBTCBalance;
    console.log("");
    console.log("");

    
    console.log("Testing issue function")



    await issueToken(vegaInstance, contractOwner.public, 110000000000, 'Vega');
    await issueToken(tUSDCInstance, contractOwner.public, 120000000000, 'tUSDC');
    await issueToken(tBTCInstance, contractOwner.public, 130000000000, 'tBTC');

    await issueToken(vegaInstance, randomAccount, 115000000000, 'Vega');
    await issueToken(tUSDCInstance, randomAccount, 125000000000, 'tUSDC');


    newOwnerVegaBalance = await getBalance(vegaInstance, contractOwner.public);
    newOwnerTUSDCBalance = await getBalance(tUSDCInstance, contractOwner.public);
    newOwnerTBTCBalance = await getBalance(tBTCInstance, contractOwner.public);

    newRandomAccVegaBalance = await getBalance(vegaInstance, randomAccount);
    newRandomAccTUSDCBalance = await getBalance(tUSDCInstance, randomAccount);
    newRandomAccTBTCBalance = await getBalance(tBTCInstance, randomAccount);

    
    console.log("");
    console.log("");
    console.log("Balances for " + contractOwner.public + ": ");
    console.log("  - Vega: " + await getBalance(vegaInstance, contractOwner.public));
    console.log("  - tUSDC: " + await getBalance(tUSDCInstance, contractOwner.public));
    console.log("  - tBTC: " + await getBalance(tBTCInstance, contractOwner.public));
    console.log("");

    console.log("Balances for " + randomAccount + ": ");
    console.log("  - Vega: " + await getBalance(vegaInstance, randomAccount));
    console.log("  - tUSDC: " + await getBalance(tUSDCInstance, randomAccount));
    console.log("  - tBTC: " + await getBalance(tBTCInstance, randomAccount));


    process.exit(0);
}


main();