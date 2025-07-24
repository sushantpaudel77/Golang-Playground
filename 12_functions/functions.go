package main

import "fmt"

// --- Higher-order function ---
func mapSlice(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}

func double(x int) int {
	return x * 2
}

// --- Callback function type and usage ---
type callback func(string)

func greet(name string, cb callback) {
	message := "Hello, " + name
	cb(message) // invoke the callback
}

func printMessage(msg string) {
	fmt.Println("Callback executed with message:", msg)
}

// --- Multiple return values ---
func getLanguages() (string, string, string) {
	return "golang", "java", "python"
}

// --- Function returning a function (closure) ---
func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

// --- main function ---
func main() {
	// Using the closure
	fn := processIt()
	fmt.Println("processIt returned:", fn(5)) // prints 4

	// Multiple return values
	lang1, lang2, _ := getLanguages()
	fmt.Println("Languages:", lang1, lang2)

	// Higher-order function usage
	nums := []int{1, 2, 3, 4, 5}
	doubled := mapSlice(nums, double)
	fmt.Println("Original:", nums)
	fmt.Println("Doubled:", doubled)

	// Callback example
	greet("Sushant", printMessage)
}
