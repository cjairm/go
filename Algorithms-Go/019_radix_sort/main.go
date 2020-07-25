package main

import (
	"fmt"

	"github.com/cjairm/Algorithms-Go/019_radix_sort/algorithm"
)

func main() {
	numbers := []int32{8, 3, 5, 4, 7, 6, 1, 2, 302, 99, 498}
	algorithm.RadixSort(numbers)
	fmt.Println(numbers)
}
