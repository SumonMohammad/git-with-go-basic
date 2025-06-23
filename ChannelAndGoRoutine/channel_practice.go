package main

import "fmt"

func hello(don chan bool) {
	fmt.Println("hello from go routine")
	don <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	//<-done
	fmt.Println("main function")

}
