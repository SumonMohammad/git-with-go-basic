package main

import "fmt"

func reverse(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func main() {
	result := reverse([]int{1, 2, 3, 4, 5})
	fmt.Println(result)
}
