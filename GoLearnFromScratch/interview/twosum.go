package main

import "fmt"

func twoSum(arr []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range arr {
		complement := target - num

		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}
func main() {
	nums := []int{5, 7, 9, 1, 3}
	target := 8
	res := twoSum(nums, target)
	fmt.Println(res)
}
