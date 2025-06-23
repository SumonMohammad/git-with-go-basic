package main

import "fmt"

func main() {
	n := 6
	for i := 0; i < n; i++ {
		for j := i; j < 2*i; j++ {
			fmt.Print("**")
		}
		fmt.Println()
	}
}
