package main

import "fmt"

func main() {
	var discount = 0.95
	var price, total float64
	var count int
	price = 200
	count = 6
	total = price * float64(count)
	if total > 1300 {
		total *= discount
	}
	fmt.Println("payable : ", total)
}
