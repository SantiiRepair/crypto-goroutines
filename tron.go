package crypto_goroutines

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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

func Tron(done <-chan bool, candidate *Creation) {

	for {
		select {
		case <-done:
			fmt.Println("Tron starting...")
			candidate.Mnemonic = tron.Generate()
			fromMnemonic, err := tron.FromMnemonicSeedAndPassphrase(candidate.Mnemonic, candidate.MnemonicPassphrase, 0)
			if err != nil {
				panic(err)
			}
			privateKey := fromMnemonic.ToECDSA()

			address := tron.PubkeyToAddress(privateKey)
			fmt.Printf("The address %s has no balance\n", address)

			err = ioutil.WriteFile("tron_private_key.txt", []byte(privateKey), 0644)
			if err != nil {
				panic(err)
			}

			rand.Seed(time.Now().UnixNano())
			waitTime := rand.Intn(10) + 1
			fmt.Printf("Waiting %d seconds before generating the next address\n", waitTime)
			time.Sleep(time.Duration(waitTime) * time.Second)

		default:
			time.Sleep(1 * time.Second)
		}
	}
}
