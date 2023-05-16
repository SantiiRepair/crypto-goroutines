package crypto_goroutines

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	tron "github.com/SantiiRepair/crypto-goroutines/tron_utils"
)

func Tron(done <-chan bool) {

	for {
		select {
		case <-done:
			fmt.Println("Tron() starting...")
			mnmonic, err := tron.Generate()
			privateKey, err := keys.FromMnemonicSeedAndPassphrase()
			if err != nil {
				panic(err)
			}

			address := address.PubkeyToAddress(privateKey)
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
