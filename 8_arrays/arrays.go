package main

import (
	"fmt"
	"sort"
)

// numbered sequences of specific length
func main() {
	// zero values
	// var nums [4]int

	// var vals [4]bool

	// vals[2] = true

	// fmt.Println(vals)

	// nums[0] = 1

	// fmt.Println(nums)

	// fmt.Println(nums[0])

	// fmt.Println(len(nums))

	var name [3]string
	name[0] = "Golang"
	name[1] = "is"
	name[2] = "awesome"
	fmt.Println(name)

	// Assign values

	var numbers [5]int

	numbers[2] = 10
	numbers[0] = 40
	numbers[1] = 20
	numbers[4] = 50
	numbers[3] = 30

	// Print the whole arrays
	fmt.Println("Array:", numbers)

	// length of arr
	fmt.Println(len(numbers))

	// Arrays are not directly  sortable with sort.Ints
	// So we convert it to a slice
	sortedNumbers := numbers[:] // convert array to slice

	sort.Ints(sortedNumbers)

	// Print sorted array
	fmt.Println("Sorted:", sortedNumbers)

	// classic loop through array
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index", i, "value:", numbers[i])
	}

	// new version of loop
	for i, val := range numbers {
		fmt.Println("Index", i, "value:", val)
	}

}
