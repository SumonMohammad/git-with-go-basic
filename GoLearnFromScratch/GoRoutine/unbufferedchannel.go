package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	fmt.Println("data is sending from send data")

	ch <- "Hey Sumon"

	fmt.Println("data is sent from send data")
}
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Data is received from sendData :", <-ch)
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

}
