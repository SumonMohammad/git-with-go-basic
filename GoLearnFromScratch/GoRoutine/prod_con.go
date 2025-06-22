package main

import (
	"fmt"
	"sync"
	"time"
)

func Producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		ch <- i
		time.Sleep(200 * time.Microsecond)
	}

	close(ch)

}

func Consumer(ch chan int, out *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Received Number :", num)
		*out = append(*out, num)
	}

}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	var results []int
	wg.Add(2)
	go Producer(ch, &wg)
	go Consumer(ch, &results, &wg)
	wg.Wait()

	fmt.Println("the results of the program :", results)
	fmt.Println("main is finished here")

}
