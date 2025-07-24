package main

import "fmt"

// maps -> hash, obj
func main() {
	// creating map

	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	// fmt.Println(m["name"], m["area"])
	fmt.Println(m["phone"])

	countryCodes := map[string]string{
		"AU": "Australia",
		"US": "United States",
		"FR": "France",
		"NP": "Nepal",
	}

	// Add a new element
	countryCodes["Jp"] = "Japan"

	// Access an element
	fmt.Println("NP:", countryCodes["NP"])

	// Update an element
	countryCodes["FR"] = "French Republic"

	// Delete an element
	delete(countryCodes, "US")

	// Print the whole map
	fmt.Println(countryCodes)
}
