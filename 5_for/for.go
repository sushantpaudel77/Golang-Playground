package main

import "fmt"

func main() {
	// while loop
	//  i := 1
	//  for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	//  }

	// infinite loop
	// for {
	// 	fmt.Println("Go")
	// }

	// classic for loop
	for i := 1; i <= 5; i++ {
		// break

		// if i == 2 {
		// 	continue
		// }
		
		// fmt.Printf("%d. %s\n", i, "Golang is awesome!")
	}

	// loop with range 

	for i := range 5 {
		fmt.Println(i)
	}
}