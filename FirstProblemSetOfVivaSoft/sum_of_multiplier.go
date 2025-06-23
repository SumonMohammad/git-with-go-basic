package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter an positive number :   ")
	fmt.Scan(&number)

	sum := 0

	for i := 0; i <= number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
