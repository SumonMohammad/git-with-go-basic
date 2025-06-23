package main

import "fmt"

func binarySearch(arr []int, target, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if target == arr[mid] {
		return mid
	} else if target < arr[mid] {
		return binarySearch(arr, target, low, mid-1)
	} else {
		return binarySearch(arr, target, mid+1, high)
	}

}

func main() {

	arr := []int{3, 4, 5, 6, 7, 8, 9, 23, 55, 77, 88, 99}
	target := 10
	result := binarySearch(arr, target, 0, len(arr)-1)
	if result != -1 {
		fmt.Printf("✅ Found %d at index %d\n", target, result)
	} else {
		fmt.Printf("❌ %d not found in array\n", target)
	}
}
