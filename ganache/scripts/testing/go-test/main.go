package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	vegaTokenAddress     = common.HexToAddress("0x9708FF7510D4A7B9541e1699d15b53Ecb1AFDc54")
	bridgeAddress        = common.HexToAddress("0x9708FF7510D4A7B9541e1699d15b53Ecb1AFDc54")
	contractOwnerAddress = common.HexToAddress("0xEe7D375bcB50C26d52E1A4a472D8822A2A22d94F")

	vegaPubKey = "vega_1"
	amountStr  = "100000000000000000000"
)

func main() {
	amount := new(big.Int)
	amount, ok := amount.SetString(amountStr, 10)
	if !ok {
		log.Fatalf("can not parse %s into big string", amountStr)
	}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to create Ethereum client: %s", err)
	}

	bridge, err := NewBridgeLogic(bridgeAddress, client)
	if err != nil {
		log.Fatalf("Failed to create NewBridgeLogic: %s", err)
	}

	vegaToken, err := NewBaseFaucetToken(vegaTokenAddress, client)
	if err != nil {
		log.Fatalf("Failed to create NewBaseFaucetToken: %s", err)
	}

	// a37f4c2a678aefb5037bf415a826df1540b330b7e471aa54184877ba901b9ef0

	// privateKey, err := crypto.HexToECDSA("a37f4c2a678aefb5037bf415a826df1540b330b7e471aa54184877ba901b9ef0")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bind.NewKeyedTransactor(privateKey)

	balance, err := vegaToken.BalanceOf(&bind.CallOpts{}, contractOwnerAddress)
	if err != nil {
		log.Fatalf("Failed to call first BalanceOf contract: %s", err)
	}
	fmt.Printf("Balance of %q is %q \n", contractOwnerAddress, balance)

	fmt.Printf("Minting token amount %q for %q \n", amount, contractOwnerAddress)
	_, err = vegaToken.Mint(&bind.TransactOpts{From: contractOwnerAddress}, contractOwnerAddress, amount)
	if err != nil {
		log.Fatalf("Failed to call Mint contract: %s", err)
	}

	balance, err = vegaToken.BalanceOf(&bind.CallOpts{}, contractOwnerAddress)
	if err != nil {
		log.Fatalf("Failed to call second BalanceOf contract: %s", err)
	}
	fmt.Printf("Balance of %q after mint is %q \n", contractOwnerAddress, balance)

	fmt.Printf("Approving token %q amount %q for %q \n", vegaTokenAddress, amount, bridgeAddress)
	if _, err := vegaToken.Approve(&bind.TransactOpts{From: contractOwnerAddress}, bridgeAddress, amount); err != nil {
		log.Fatalf("Failed to call Approve contract: %s", err)
	}

	vegaPubKeyByte, err := stringToByte32(vegaPubKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Depositing asset %q amout %q Vega pub key %q", vegaTokenAddress, amount, vegaPubKey)

	_, err = bridge.DepositAsset(&bind.TransactOpts{}, vegaTokenAddress, amount, vegaPubKeyByte)
	if err != nil {
		log.Fatalf("Failed to call DepositAsset contract: %s", err)
	}

	fmt.Println("Asset deposited")
}

func stringToByte32(str string) ([32]byte, error) {
	b := [32]byte{}
	if len(str) > 32 {
		return b, fmt.Errorf("string too long")
	}

	for i, v := range str {
		b[i] = byte(v)
	}

	return b, nil
}
