package main

import "fmt"

func main() {

	CalculatePayable(192.7, 5)
	CalculatePayable(300, 5)

}
func CalculatePayable(price float64, count int) {
	var discount = 0.95
	var total = price * float64(count)
	if total > 1300 {
		total *= discount
	}
	fmt.Println("payable : ", total)
}
