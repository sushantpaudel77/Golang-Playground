package main

import "fmt"

// iterating over data structure
func main() {
	nums := []int{5, 6, 7}

	for _, num := range nums {
		fmt.Println(num)
	}

	// without range
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
}
