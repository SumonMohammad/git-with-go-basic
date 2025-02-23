package main

import "fmt"

func main() {

	var yr int64
	fmt.Println("please enter a year :  ")
	fmt.Scanln(&yr)

	if yr%400 == 0 || (yr%4 == 0 && yr%100 != 0) {
		fmt.Println(yr, "is a leap year")
	} else {
		fmt.Println(yr, "is not a leap year")
	}

}
