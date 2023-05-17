package main

import (
	"os"	
	"fmt"
	"time"

	tron "github.com/SantiiRepair/crypto-goroutines/tron_utils"
)

func Tron(done chan bool) {
	fromMnemonic, yull := tron.FromMnemonicSeed(tron.Generate(), 0)
	if yull != nil {

		fromBTCEC := fromMnemonic.ToECDSA()

		address := tron.PubkeyToAddress(fromBTCEC.PublicKey)
		fromPublic := fmt.Sprintf("%d", fromBTCEC)
		fmt.Printf("The address %s has no balance\n", address)

		write := os.WriteFile("tron_private_key.txt", []byte(fromPublic), 0644)
		if write != nil {
			panic(write)
		}

		time.Sleep(1 * time.Second)

		done <- true
	}
}
