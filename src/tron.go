package main

import (
    "context"
    "crypto/ecdsa"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "math/big"
    "math/rand"
    "time"

    "github.com/TRON-US/go-tron/crypto"
    "github.com/TRON-US/go-tron/rpc"
)

func main() {
    client := rpc.NewTronRPC("https://api.trongrid.io")

    results := make(chan string)

    for {
        privateKey, err := crypto.GenerateKey()
        if err != nil {
            panic(err)
        }

        address := crypto.PubkeyToAddress(privateKey.PublicKey)

        balance, err := client.GetAccountBalance(context.Background(), address.Hex())
        if err != nil {
            panic(err)
        }

        if balance.Int64() > 0 {
            fmt.Printf("The address %s has a balance of %d TRX\n", address.Hex(), balance.Int64())

            privateKeyBytes := crypto.FromECDSA(privateKey)
            err = ioutil.WriteFile("tron_private_key.txt", []byte(hex.EncodeToString(privateKeyBytes)), 0644)
            if err != nil {
                panic(err)
            }

            results <- fmt.Sprintf("The address %s has a balance of %d TRX and the private key has been saved to a local file", address.Hex(), balance.Int64())
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