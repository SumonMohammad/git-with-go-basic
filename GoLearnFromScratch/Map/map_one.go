package main

import "fmt"

func main() {

	freq := make(map[rune]int)

	text := "আমিতুমারে  ভালবাসি "

	for _, char := range text {
		freq[char]++
	}

	for k, v := range freq {
		fmt.Printf(" '%c' = %d \n", k, v)
	}
}
