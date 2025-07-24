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

func main() {

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
