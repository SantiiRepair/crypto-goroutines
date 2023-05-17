package main

import "fmt"

func main() {
	done := make(chan bool)

	for {
		fmt.Println("Ethereum starting...")
		go Ethereum(done)
		<-done

		fmt.Println("Tron starting...")
		go Tron(done)
		<-done
	}
}
