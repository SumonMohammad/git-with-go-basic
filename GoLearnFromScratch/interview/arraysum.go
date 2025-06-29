package main

import "fmt"

func arraySum(nums []int, target int) []int {
	start := 0
	currentSum := 0

	for end := 0; end < len(nums); end++ {
		currentSum += nums[end]
		for currentSum > target && start <= end {
			currentSum -= nums[start]
			start++
		}
		if currentSum == target {
			return nums[start : end+1]
		}

	}
	return nil
}
func main() {
	nums := []int{4, 5, 2, 5, 5, 2}
	target := 10
	result := arraySum(nums, target)
	fmt.Println(result)
}
