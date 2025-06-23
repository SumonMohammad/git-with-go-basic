package main

import "fmt"

func makeSquare(ch chan int) {
	for num := range ch {
		fmt.Println("Here the square of %d is %d : ", num, num*num)
	}
}

func main() {
	ch := make(chan int)
	go makeSquare(ch)

	fmt.Println("Enter an integer to make square by channel :  ")

	for {
		var num int

		fmt.Scanln(&num)

		ch <- num

	}
	fmt.Println("Program finished")
}
