package main

import "fmt"

func LongestSubString(s string) int {
	m := make(map[byte]int)
	start := 0
	maxLen := 0

	for end := 0; end < len(s)-1; end++ {
		ch := s[end]
		if idx, ok := m[ch]; ok && idx >= start {
			start = idx + 1
		}
		m[ch] = end

		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
	}
	return maxLen
}
func main() {
	result := LongestSubString("abcabcdefui")
	fmt.Println("longest substring :", result)

}
