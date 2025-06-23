package main

import (
	"errors"
	"fmt"
)

func addInt(a, b interface{}) (int, error) {
	x, ok1 := a.(int)
	y, ok2 := b.(int)

	if !ok1 || !ok2 {
		return 0, errors.New("invalid input type")
	}
	return x + y, nil
}

func main() {
	result, err := addInt(2, 3)

	if err != nil {
		fmt.Println("something went wrong")
	} else {
		fmt.Println("the result is :", result)
	}

}
