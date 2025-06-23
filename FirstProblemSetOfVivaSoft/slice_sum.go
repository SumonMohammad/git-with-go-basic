package main

import "fmt"

func main() {
	nums := []int{}
	var n int
	fmt.Println("Enter an integer")
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		var num int
		fmt.Println("Enter the slice item :  ")
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	fmt.Println(nums)

	fmt.Println(len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {

		sum += nums[i]
	}
	fmt.Println(sum)
}
