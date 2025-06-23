package main

import "fmt"

func average(nums []int) float64 {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
	var avg float64
	avg = float64(sum / len(nums))

	return avg
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	avg := average(nums)
	fmt.Println(avg)
}
