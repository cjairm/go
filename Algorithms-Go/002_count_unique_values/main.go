package main

import (
	"fmt"

	"github.com/cjairm/Algorithms-Go/002_count_unique_values/algorithm"
)

func main() {
	x := []int{1, 1, 1, 2, 2}
	fmt.Println(algorithm.GetCountOfUniqueValues(x))
}
