package main

import "sync"

func makeCube(n int, ch chan map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := make(map[int]int)
	result[n] = n * n * n
	ch <- result
}

func main() {
	nums := []int{8, 4, 3, 5, 3, 5, 2, 6}

	var wg sync.WaitGroup

	ch := make(chan map[int]int, len(nums))

	for _, num := range nums {
		wg.Add(1)
		go makeCube(num, ch, &wg)
	}
	wg.Wait()
	close(ch)

	finalResult := make(map[int]int)
	for m := range finalResult {

	}
}
