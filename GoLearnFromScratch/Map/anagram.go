package main

import "fmt"

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range a {
		count[ch]++
	}
	for _, ch := range b {
		count[ch]--
	}
	for _, value := range count {
		if value != 0 {
			return false
		}
	}

	return true
}
func main() {
	fmt.Println(isAnagram("silent", "listen"))
}
