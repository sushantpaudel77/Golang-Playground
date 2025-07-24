package main

import "fmt"

// maps -> hash, obj
func main() {
	// creating map

	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["name"], m["area"])
	m := make(map[string]int)
	m["age"] = 30
	m["price"] = 50
	// fmt.Println(m["phone"])

	fmt.Println(len(m))
	clear(m)

	myMap := map[string]int{
		"price": 70,
		"phone": 50,
	}

	_, ok := myMap["phone"]

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println()

	countryCodes := map[string]string{
		"AU": "Australia",
		"US": "United States",
		"FR": "France",
		"NP": "Nepal",

	}

	// Add a new element
	countryCodes["JP"] = "Japan"
	// Access an element
	fmt.Println("NP:", countryCodes["NP"])

	// Update an element
	countryCodes["FR"] = "French Republic"

	// Delete an element
	delete(countryCodes, "US")

	// Print the whole map
	fmt.Println(countryCodes)


}
