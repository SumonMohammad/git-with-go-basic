package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello is starting by calling by go routine :")
	time.Sleep(10 * time.Second)
	fmt.Println(" ghum theke uthche ")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("here is main starting :")
	go hello(done)

	<-done
	fmt.Println("main close hobe ekhon , hoiche?:")

}
