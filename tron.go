package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	tron "github.com/SantiiRepair/crypto-goroutines/tron_utils"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
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

		url := fmt.Sprintf("https://api.trongrid.io/v1/accounts/%s", address)

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
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
			privateKeyBytes := crypto.FromECDSA(fromBTCEC)
			hexPrivateKey := hex.EncodeToString(privateKeyBytes)
			b := []byte(hexPrivateKey)
			privateKey := base58.Encode(b)
			os.WriteFile("tron_private_key.txt", []byte(privateKey), 0644)
		} else {
			fmt.Printf("The address %s has no balance\n", address)
		}

		time.Sleep(1 * time.Second)

		done <- true
	}
}
