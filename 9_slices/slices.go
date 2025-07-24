package main

import "fmt"

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// unitialized slice is nil

	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2)

	fmt.Println(nums)
}
