package main

import (
	"fmt"
	"sync"
)

func workerPool(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range jobs {
		fmt.Printf("Worker %d processing %d\n", id, num)
		results <- num * 2
	}

}
func main() {
	numbers := []int{7, 3, 4, 2, 4, 5, 6}

	numOfWorkers := 4
	var wg sync.WaitGroup
	jobs := make(chan int, len(numbers))
	results := make(chan int, len(numbers))

	for i := 1; i < numOfWorkers; i++ {
		wg.Add(1)
		go workerPool(i, jobs, results, &wg)
	}

	for _, job := range numbers {
		jobs <- job
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Results :", res)
	}

}
