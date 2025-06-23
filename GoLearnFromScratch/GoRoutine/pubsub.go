package main

import (
	"fmt"
	"time"
)

func ProducerC(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
		time.Sleep(200 * time.Microsecond)
	}
	close(ch)
}

func ConsumerC(ch chan int) {
	for num := range ch {
		fmt.Println("Recieved Number : ", num)
	}
}

func main() {
	ch := make(chan int)
	go ProducerC(ch)
	go ConsumerC(ch)
	time.Sleep(3000 * time.Microsecond)
	fmt.Println("Main go routine is finished")
}
