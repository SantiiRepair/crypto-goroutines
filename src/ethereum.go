package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/onrik/ethrpc"
)

func Ethereum(done chan<- bool) {
	rpc := "https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet"

	for {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			panic(err)
		}

		publicKey := privateKey.Public()
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

			privateKeyBytes := crypto.FromECDSA(privateKey)
			err = ioutil.WriteFile("eth_private_key.txt", privateKeyBytes, 0644)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("The address %s has no balance\n", address.Hex())
		}

		rand.Seed(time.Now().UnixNano())
		waitTime := rand.Intn(10) + 1
		fmt.Printf("Waiting %d seconds before generating the next address\n", waitTime)
		time.Sleep(time.Duration(waitTime) * time.Second)

		done <- true
	}
}