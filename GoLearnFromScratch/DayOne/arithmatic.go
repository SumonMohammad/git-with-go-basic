package main

import "fmt"

func main() {
	a := 9
	b := 6
	if b == 0 {
		fmt.Println("not possible")
	} else {
		var r = float64(a) / float64(b)
		fmt.Println("Division", r)
	}

	//swapping value without 3rd variable
	fmt.Println(a, "and ", b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, "and ", b)

}
