package main

import (
	"fmt"
	"strings"
)

func main() {

	sentence := "I Love bangladesh football team over Cricket team"

	words := strings.Fields(sentence)

	wordCount := make(map[string]int)

	for _, count := range words {
		wordCount[count]++
	}

	for word, count := range wordCount {
		fmt.Printf("'%s' = %d \n", word, count)
	}
}
