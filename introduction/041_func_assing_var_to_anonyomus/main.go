package main

import "fmt"

func main() {
	sayMyName := func(name string) {
		fmt.Println("Your name is", name)
	}

	sayMyName("Carlos")
}
