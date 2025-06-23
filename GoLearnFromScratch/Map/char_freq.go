package main

import "fmt"

func main() {
	charFreq := make(map[rune]int)

	text := "programming"

	for _, char := range text {
		charFreq[char]++
	}
	for k, v := range charFreq {
		fmt.Printf(" '%c' = %d \n", k, v)
	}

}
