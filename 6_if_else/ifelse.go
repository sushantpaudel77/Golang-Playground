package main

import "fmt"

func main() {

	// name := "John"
	// age := 22

	// if age >= 18 {
	// 	fmt.Printf("%s is an adult.\n", name)
	// } else if age >= 13 {
	// 	fmt.Printf("%s is a teenager.\n", name)
	// } else {
	// 	fmt.Printf("%s is a child.\n", name)
	// }

	var role = "Admin"
	var hasPermissions = false

	if role == "Admin" && hasPermissions == true {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
		if role == "Admin" || hasPermissions == true {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if age := 15; age >= 18 {
		fmt.Println("Person is and adult", age)
	} else if age >= 12 {
		fmt.Println("Person is teeneger", age)
	}
		
}
