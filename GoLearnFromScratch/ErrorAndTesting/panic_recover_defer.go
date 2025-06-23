package main

import (
	"fmt"
)

func main() {
	// Test cases
	fmt.Println(safeDivide(10, 2)) // ✅ 5
	fmt.Println(safeDivide(10, 0)) // ❌ panic → recover
	fmt.Println(safeDivide(25, 5)) // ✅ 5
}

func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 0 // fallback value
		}
	}()

	// This line can panic if b == 0
	result = a / b
	return
}
