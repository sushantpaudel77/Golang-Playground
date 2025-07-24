package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {

	result := add(7, 7)
	fmt.Println("Result:", result)
}
