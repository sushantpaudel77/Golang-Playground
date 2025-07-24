package main

import "fmt"

// iterating over data structure
func main() {
	nums := []int{5, 6, 7}

	sum := 0

	for _, num := range nums {
		sum += num
		fmt.Println(sum)
	}

	fmt.Println()

	for i, num := range nums {
		fmt.Println(num, i)
	}

	// without range
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	myMap := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range myMap {
		fmt.Printf("%s: %s\n", k, v)
	}

	for _, v := range myMap {
		fmt.Println(v)
	}
}
