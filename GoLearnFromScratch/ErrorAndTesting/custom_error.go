package main

import "fmt"

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func myErrorFunc() error {
	return MyError("My Custom Error")
}

func main() {
	err := myErrorFunc()
	if err != nil {
		fmt.Println("Error", err)
	}
}
