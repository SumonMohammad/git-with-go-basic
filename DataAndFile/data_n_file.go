package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//var boolean = true
	//boolB, _ := json.Marshal(boolean)
	//for i := range boolB {
	//	fmt.Println(string(boolB[i]))
	//}
	//
	//intB, _ := json.Marshal(1)
	//fmt.Println(string(intB))
	//fltB, _ := json.Marshal(3.14)
	//fmt.Println(string(fltB))
	//strB, _ := json.Marshal("gopher")
	//fmt.Println(string(strB))

	mapToJSON()
}

func arrayToJSON() {

	slcD := []string{"one", "two", "three"}
	fmt.Println("Before", slcD)

	slcN, _ := json.Marshal(slcD)
	fmt.Println("After", string(slcN))

}

func mapToJSON() {
	mapA := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("Before", mapA)
	mapB, _ := json.Marshal(mapA)
	fmt.Println("After", string(mapB))
}
