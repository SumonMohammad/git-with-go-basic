package main

import "fmt"

type User struct {
	Name string
}

func (u User) renameValue(newName string) {
	u.Name = newName
}

func (u *User) renamePointer(newName string) {
	u.Name = newName
}
func main() {

	u := User{Name: "Sumon"}

	u.renameValue("hey")
	fmt.Println(u.Name)

	u.renamePointer("raj")
	fmt.Println(u.Name)

}
