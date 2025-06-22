package main

import "fmt"

func reverseSlice(arr []int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left] //এটা Go এর tuple assignment → একই সাথে দুইটা index এর মান অদল-বদল হয়

		left++
		right--
	}
	return arr
}

func main() {
	arr := []int{5, 6, 2, 3, 4}
	fmt.Printf("Main  earray :", arr)
	rev := reverseSlice(arr)
	fmt.Printf("Reverse earray :", rev)
}
