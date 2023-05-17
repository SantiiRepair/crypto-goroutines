package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tron "github.com/SantiiRepair/crypto-goroutines/tron_utils"
)

type Creation struct {
	Name               string
	Passphrase         string
	Mnemonic           string
	MnemonicPassphrase string
	HdAccountNumber    *uint32
	HdIndexNumber      *uint32
}

func Tron(done chan bool) {
	mnemonic := tron.Generate()
	fromMnemonic, err := tron.FromMnemonicSeedAndPassphrase(mnemonic, 0)
	if err != nil {
		panic(err)
	}
	fromBTCEC := fromMnemonic.ToECDSA()

	address := tron.PubkeyToAddress(fromBTCEC.PublicKey)
	fromPublic := fmt.Sprintf("%d", fromBTCEC)
	fmt.Printf("The address %s has no balance\n", address)

	write := os.WriteFile("tron_private_key.txt", []byte(fromPublic), 7879)
	if write != nil {
		panic(write)
	}

	rand.Seed(time.Now().UnixNano())
	waitTime := rand.Intn(10) + 1
	fmt.Printf("Waiting %d seconds before generating the next address\n", waitTime)
	time.Sleep(time.Duration(waitTime) * time.Second)

	time.Sleep(1 * time.Second)

	done <- true
}
