package main

import (
	"fmt"

	"github.com/cjairm/Algorithms-Go/015_get_pivot_quick_sort/algorithm"
)

func main() {
	pos, xi := algorithm.GetPivotQuickSort([]int{5, 2, 1, 8, 4, 7, 6, 3}, 0, 0)
	fmt.Println(pos, xi)
}
