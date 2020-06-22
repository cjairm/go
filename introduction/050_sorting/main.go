package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"Allen", "Queen", "Wayne", "Stark", "Kent"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	sort.Sort(sort.Reverse(sort.IntSlice(xi)))
	fmt.Println(xi)

	fmt.Println("------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
	sort.Sort(sort.Reverse(sort.StringSlice(xs)))
	fmt.Println(xs)
}
