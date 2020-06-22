package main

import "fmt"

func main() {
	v := "Hello world"
	fmt.Println(`The address in memory of "`+v+`" is: `, &v)
}
