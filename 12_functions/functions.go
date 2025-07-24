package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

func getLanguages() (string, string, string) {
	return "golang", "java", "python"
}

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

// callback function type
type callback func(string)

// function that accepts a callback
func greet(name string, cb callback) {
	message := "Hello, " + name
	cb(message) // invoke the callback
}

// Actual callback function
func printMessage(msg string) {
	fmt.Println("Callback executed with message:", msg)
}


func main() {

	// Passing the function as a callback
	greet("Golang", printMessage)

	// fn := func(a int) int {
	// 	return 2
	// }

	fn := processIt()
	fn(5)

	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)
	// result := add(7, 7)
	// fmt.Println("Result:", result)
}
