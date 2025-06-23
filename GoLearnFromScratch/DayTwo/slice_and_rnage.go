package main

import "fmt"

func DoSomeArithmaticOp(nums []int) {
	sum := 0

	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)

}
func main() {
	nums := []int{1, 2, 3}
	DoSomeArithmaticOp(nums)
}
