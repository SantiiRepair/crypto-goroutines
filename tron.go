package crypto_goroutines

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tronprotocol/go-tron"
	"github.com/tronprotocol/go-tron/client"
)

func Tron(done <-chan bool) {

	for {
		select {
		case <-done:
			fmt.Println("Tron() starting...")
			privateKey, err := tron.GeneratePrivateKey()
			if err != nil {
				panic(err)
			}

			publicKey := privateKey.Public()
			publicKeyBytes := publicKey.Bytes()[1:]

			address := hexutil.Encode(publicKeyBytes)
			fmt.Printf("The address %s has no balance\n", address)

			privateKeyBytes := hexutil.Encode(privateKey.Bytes())
			err = ioutil.WriteFile("tron_private_key.txt", []byte(privateKeyBytes), 0644)
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
