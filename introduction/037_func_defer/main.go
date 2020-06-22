package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("I'm the main function")
}

func foo() {
	defer func() {
		fmt.Println("I'm a defer anonymous funtion inside \"foo\"")
	}()

	fmt.Println("Hello... I'm exiting from this funtion \"foo\"")
}
