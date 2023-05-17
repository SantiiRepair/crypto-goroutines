package main

import (
	"encoding/json"
	"fmt"
	tron "github.com/SantiiRepair/crypto-goroutines/tron_utils"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Account struct {
	Address string `json:"address"`
	Balance uint64 `json:"balance"`
}

func Tron(done chan bool) {
	fromMnemonic, yull := tron.FromMnemonicToPrivateKey(tron.Generate(), 0)
	if yull != nil {

		fromBTCEC := fromMnemonic.ToECDSA()

		address := tron.PubkeyToAddress(fromBTCEC.PublicKey)
		fromPublic := fmt.Sprintf("%d", fromBTCEC)

		url := fmt.Sprintf("https://api.trongrid.io/v1/accounts/%s", address)

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var account Account
		err = json.Unmarshal(body, &account)
		if err != nil {
			panic(err)
		}

		if account.Balance > 0 {
			fmt.Printf("TRX balance from account %s: %d\n", account.Address, account.Balance)
			os.WriteFile("tron_private_key.txt", []byte(fromPublic), 0644)
		} else {
			fmt.Printf("The address %s has no balance\n", address)
		}

		time.Sleep(1 * time.Second)

		done <- true
	}
}
