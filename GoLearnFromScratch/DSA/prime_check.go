package main

import "fmt"

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%i+2 == 0 {
			return false
		}
	}

	return true
}

func main() {
	nums := []int{2, 3, 5, 119, 33, 19, 473, 127, 201}

	for _, num := range nums {
		if isPrime(num) {
			fmt.Printf("the number %d is prime \n", num)
		} else {
			fmt.Printf("the number %d is not prime \n", num)
		}
	}
}
