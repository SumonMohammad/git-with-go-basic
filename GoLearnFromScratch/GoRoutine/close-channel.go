package main

import (
	"fmt"
	"time"
)

func worker(cancel <-chan struct{}) {
	fmt.Println("starting worker")

	defer fmt.Println("worker is closed")

	for {
		select {
		case <-cancel:
			fmt.Println("worker closing")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("working")
		}

	}

}
func main() {
	cancel := make(chan struct{})
	go worker(cancel)
	time.Sleep(5 * time.Second)
	fmt.Println("Cancelling worker")
	close(cancel)
	time.Sleep(2 * time.Second)
	fmt.Println("main is stopping now")

}
