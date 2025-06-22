// import (
//
//	"fmt"
//
// )
//
//	func sendDataToChannel(ch chan string) {
//		fmt.Println("sending data from channel ")
//
//		ch <- "hey sumon"
//
//		fmt.Println("Data has been sent ")
//
// }
//
//	func getDataFromChannel(ch chan string) {
//		//time.Sleep(2 * time.Second)
//		fmt.Println("String retrieved from channel:", <-ch)
//
// }
//
// func main() {
//
//	ch := make(chan string)
//
//	go sendDataToChannel(ch)
//
//	go getDataFromChannel(ch)
//
//	fmt.Scanln()
//
// }
package main

import (
	"fmt"
	"sync"
	"time"
)

func routine(waitgroup *sync.WaitGroup, number int) {
	defer waitgroup.Done()
	fmt.Printf("Starting routine %d \n", number)
	time.Sleep(time.Second)
	fmt.Printf("Done with routine %d \n", number)
}

func main() {
	fmt.Println("Starting main Goroutine")
	var waitgroup sync.WaitGroup

	for i := 0; i < 5; i++ {
		waitgroup.Add(1)
		go routine(&waitgroup, i)
	}
	waitgroup.Wait()
	fmt.Println("Finishing main Goroutine")
}
