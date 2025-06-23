package main

import "fmt"

func recursion(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * recursion(n-1)
}

func main() {
	var n = 5
	result := recursion(n)

	fmt.Println(result)
}
