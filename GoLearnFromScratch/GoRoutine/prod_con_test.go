package main

import (
	"sync"
	"testing"
)

//func TestProducerConsumer(t *testing.T) {
//	ch := make(chan int)
//	var results []int
//	var wg sync.WaitGroup
//
//	wg.Add(2) // One for Producer, one for Consumer
//
//	go Producer(ch, &wg)
//	go Consumer(ch, &results, &wg)
//
//	wg.Wait() // Wait until both finish
//
//	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
//
//	if len(results) != len(expected) {
//		t.Fatalf("Expected %d results, got %d", len(expected), len(results))
//	}
//
//	for i := range expected {
//		if results[i] != expected[i] {
//			t.Errorf("Mismatch at index %d: expected %d, got %d", i, expected[i], results[i])
//		}
//	}
//}

func TestProdCon(t *testing.T) {
	ch := make(chan int)
	var wg *sync.WaitGroup
	var results []int

	defer wg.Done()

	go Producer(ch, wg)
	go Consumer(ch, &results, wg)

	wg.Wait()
}
