package main

func main() {
	done := make(chan bool)

	for {
		go Ethereum(done)
		<-done

		go Tron(done)
		<-done
	}
}
