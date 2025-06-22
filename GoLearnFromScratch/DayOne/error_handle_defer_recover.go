package main

import "fmt"

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("❗Recovered from panic:", r)
		}
	}()

	result := a / b // এখানে b=0 হলে panic হবে
	fmt.Println("✅ Result:", result)
}

func main() {
	fmt.Println("🚀 Starting Program")
	safeDivide(10, 2) // OK
	safeDivide(5, 0)  // Causes panic but recovered
	fmt.Println("🏁 Program continues...")
}

//func safeDivide(a, b float64) {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("❗Recovered from panic:", r)
//		}
//	}()
//
//	result := a / b
//	fmt.Println("Result ", result)
//}
//
//func main() {
//	fmt.Println("program is running")
//	safeDivide(2.5, 0)
//
//}
