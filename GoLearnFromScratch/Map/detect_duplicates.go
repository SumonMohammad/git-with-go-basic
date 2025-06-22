package main

import "fmt"

func findDuplicates(slice []int) []int {
	numFreq := make(map[int]int)
	duplicates := []int{}
	for _, num := range slice {
		numFreq[num]++
	}
	for num, count := range numFreq {
		if count > 1 {
			duplicates = append(duplicates, num)
		}
	}
	return duplicates
}
func main() {

	numbers := []int{1, 3, 3, 4, 2, 4, 5, 6, 7, 5}

	result := findDuplicates(numbers)

	fmt.Println(result)

}
