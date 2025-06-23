package main

import "fmt"

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recovered from panic:", p)
			
		}
	}()

	fmt.Println("Program is starting ....")

	var array [3]int
	pos := 3
	array[pos] = 1
	fmt.Println(array[pos])
	fmt.Println("Program is finished")
}
