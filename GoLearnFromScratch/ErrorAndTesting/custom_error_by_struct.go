package main

import "fmt"

type User struct {
	Firstname string
	LastName  string
}
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error : %d %s ", e.Code, e.Message)
}

func myFunc(user User) error {
	if user.Firstname == "" {
		return CustomError{Code: 400, Message: "FirstName is required"}
	}
	if user.LastName == "" {
		return CustomError{Code: 400, Message: "LastName is required"}
	}

	return nil
}

func main() {
	user := &User{}
	user.Firstname = "Sumon"
	user.LastName = "Mohammad"
	err := myFunc(*user)
	fmt.Println(err.Error())
}
