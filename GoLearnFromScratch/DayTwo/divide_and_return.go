package main

import "fmt"

func divide(a, b int) (float64, error) {
	if b == 0 {
		fmt.Println("not possible")
	} else {
		var res = float64(a) / float64(b)
	}
	return res, nil
}

func main() {
	divide(10, 4)
}
