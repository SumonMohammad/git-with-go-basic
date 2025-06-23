package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	first := 1
	second := 2

	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}

	return second

}
func main() {
	numOfStairs := 10

	fmt.Println(climbStairs(numOfStairs))
}

//Time: O(n) Space: O(1)
