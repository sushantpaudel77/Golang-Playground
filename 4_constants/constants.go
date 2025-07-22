package main

import "fmt"

// const age = 30

var name string = "golang"

func main() {
		var name = "golang is awesome"

		name = "java is awesome"

		fmt.Println(name)

		const (
			port = 5000
			host = "localhost"
		)

		fmt.Println(port, host)

}