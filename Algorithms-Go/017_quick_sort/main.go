package main

import (
	"fmt"

	"github.com/cjairm/Algorithms-Go/017_quick_sort/algorithm"
)

func main() {
	xi := algorithm.QuickSort([]int{5100, -3, 2, 4, 6, 9, 1, 2, 5, 3, 23}, 0, 0)
	fmt.Println(xi)
}
