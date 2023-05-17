package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/onrik/ethrpc"
)

func Ethereum(done chan bool) {
	rpc := "https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet"

	privateKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("Could not cast public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	rpcClient := ethrpc.New(rpc)
	balance, err := rpcClient.EthGetBalance(address.String(), "latest")
	if err != nil {
		panic(err)
	}

	if balance.Int64() > 0 {
		fmt.Printf("The address %s has a balance of %s ETH\n", address.Hex(), balance.String())
		privateKeyBytes := crypto.FromECDSA(privateKeyECDSA)
		privateKey := hex.EncodeToString(privateKeyBytes)
		os.WriteFile("eth_private_key.txt", []byte(privateKey), 0644)
	} else {
		fmt.Printf("The address %s has no balance\n", address.Hex())
	}

	time.Sleep(1 * time.Second)

	done <- true
}
