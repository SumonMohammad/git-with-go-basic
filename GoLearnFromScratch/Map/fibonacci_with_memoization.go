package main

import "fmt"

func fiboWithMemoization(n int) int {
	if n <= 1 {
		return n
	}
	memo := make([]int, n+1) //slice define how many index are declared even before use
	memo[0], memo[1] = 0, 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
func main() {
	n := 7
	fmt.Println(fiboWithMemoization(n))
}

//ClimbStairs exactly like this just value start from 1 and loop from 3 and the result will be number of ways
