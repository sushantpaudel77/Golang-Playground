package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("I'm default bro")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("enjoy weekend bro :D")
	default:
		fmt.Println("work hard bro :)")
	}

	// type switch
	whoAmI := func(i any) {
		switch res := i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean")
		case float64:
			fmt.Println("it's a float")
		default:
			fmt.Println("bro i am default too", res)
		}
	}

	whoAmI(22.22)

}
