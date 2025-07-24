package main

import "fmt"

// Callback function type
type callback func(string)

// Function that accepts a callback
func greet(name string, cb callback) {
	message := "Hello, " + name
	cb(message) // invoke the callback
}

// Actual callback function
func printMessage(msg string) {
	fmt.Println("Callback executed with message:", msg)
}
