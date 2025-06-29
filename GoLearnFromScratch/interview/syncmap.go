package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	write := func(start int) {
		defer wg.Done()
		for i := start; i < start+5; i++ {
			m.Store(i, i) // concurrent-safe
		}
	}

	wg.Add(2)
	go write(0)
	go write(5)
	wg.Wait()

	// Read all entries
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, "=", v)
		return true
	})
}
