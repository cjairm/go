package main

import "fmt"

func main() {
	func(n string) {
		fmt.Println("Hello my name is", n)
	}("Carlos Mendez")
}
