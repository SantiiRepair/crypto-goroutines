package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "io/ioutil"
    "math/big"
    "math/rand"
    "time"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/onrik/ethrpc"
)

func main() {
    rpc := "https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet"
    client, err := ethclient.Dial(rpc)
    if err != nil {
        panic(err)
    }

    results := make(chan string)

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

        rpcClient := ethrpc.New("https://mainnet.infura.io/v3/your-project-id")
        balance, err := rpcClient.EthGetBalance(address, "latest")
        if err != nil {
            panic(err)
        }

        if balance.Int64() > 0 {
            fmt.Printf("The address %s has a balance of %s ETH\n", address.Hex(), balance.String())

            privateKeyBytes := crypto.FromECDSA(privateKey)
            err = ioutil.WriteFile("private_key.txt", privateKeyBytes, 0644)
            if err != nil {
                panic(err)
            }

            results <- fmt.Sprintf("The address %s has a balance of %s ETH and the private key has been saved to a local file", address.Hex(), balance.String())
        } else {
            fmt.Printf("The address %s has no balance\n", address.Hex())

            results <- fmt.Sprintf("The address %s has no balance", address.Hex())
        }

        rand.Seed(time.Now().UnixNano())
        waitTime := rand.Intn(10) + 1 
        fmt.Printf("Waiting %d seconds before generating the next address\n", waitTime)
        time.Sleep(time.Duration(waitTime) * time.Second)

        select {
        case message := <-results:
            fmt.Printf("%s\n", message)
        case <-time.After(5 * time.Second):
            fmt.Println("Balance check took too long, skipping to the next address")
        }
    }
}