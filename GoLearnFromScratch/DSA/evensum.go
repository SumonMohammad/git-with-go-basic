package main

import "fmt"

func main() {
	sumAllEvenNumber := 0

	for i := 2; i <= 100; i += 2 {
		sumAllEvenNumber += i
	}

	fmt.Println(sumAllEvenNumber)
}
