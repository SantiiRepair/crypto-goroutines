package crypto_goroutines

import (
	"./ethereum"
	"./tron"
	"fmt"
	"plugin"
	"time"
)

func main() {
	p, err := plugin.Open("crypto-goroutines.so")
	if err != nil {
		panic(err)
	}

	ethereumSym, err := p.Lookup("Ethereum")
	if err != nil {
		panic(err)
	}

	tronSym, err := p.Lookup("Tron")
	if err != nil {
		panic(err)
	}

	ethereumFn, ok := ethereumSym.(func(chan<- bool))
	if !ok {
		panic("Ethereum has wrong type")
	}

	tronFn, ok := tronSym.(func(chan<- bool))
	if !ok {
		panic("Tron has wrong type")
	}

	done := make(chan bool)

	for {
		fmt.Println("Starting Ethereum...")
		go ethereumFn(done)

		<-done

		fmt.Println("Starting Tron...")
		go tronFn(done)

		<-done

		time.Sleep(10 * time.Second)
	}
}
