package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// unitialized slice is nil

	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// capacity -> maximum numbers of elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)

	// nums := []int{}

	var nums = make([]int, 0, 5)
	nums = append(nums, 2)

	var nums1 = make([]int, len(nums))

	// copy function
	copy(nums1, nums)
	fmt.Println(nums, nums1)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// slice operator

	var num = []int{1, 2, 3}
	// fmt.Println(num[0:2])
	fmt.Println(num[1:])
	fmt.Println(num[:1])

	// slice
	var num1 = []int{1, 2, 3}
	var num2 = []int{3, 2, 1}

	var num3 = []int{1, 2, 3}
	var num4 = []int{1, 2, 3}

	fmt.Println(slices.Equal(num1, num2))
	fmt.Println(slices.Equal(num3, num4))

	var numb = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(numb)
	fmt.Println(len(numb))
}
