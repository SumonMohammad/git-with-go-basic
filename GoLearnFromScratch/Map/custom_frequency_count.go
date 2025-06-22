package main

import (
	"fmt"
	"sort"
)
func freqCountDecrease(nums []int) []int{
	freqCount := make(map[int] int )

	for _, count := range nums{
		freqCount[count]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] 
	})
}
func main(){
	n := []int{4,5,6,3,4,3,4,5,4,5,7,2}
	result := freqCountDecrease(n)
	fmt.Println(result)
}