package algorithm

import (
	"fmt"
	"testing"
)

func TestSearchNumber(t *testing.T) {
	test1 := SearchNumber([]int{-10, -5, -1, 0, 1, 2, 3})

	if test1 == 0 {
		t.Error("Got: ", test1, "Expected: ", []int{-1, 1})
	}

	test2 := SearchNumber([]int{-20, -18, -16, -14, -12, -10, -8, -6, -4, -2, 0, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19})

	if test2 != 0 {
		t.Error("Got: ", test2, "Expected: ", nil)
	}

	test3 := SearchNumber([]int{})

	if test3 != 0 {
		t.Error("Got: ", test3, "Expected: ", nil)
	}

	test4 := SearchNumber([]int{1, 2})

	if test4 != 0 {
		t.Error("Got: ", test4, "Expected: ", nil)
	}

	test5 := SearchNumber([]int{-20, -18, -16, -14, -12, -10, -8, -6, -4, -2, 0, 1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 19})

	if test5 == 0 {
		t.Error("Got: ", test5, "Expected: ", []int{-2, 2})
	}
}

func BenchmarkSearchNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchNumber([]int{-4, -3, -1, 0, 1, 2, 5})
	}
}

func ExampleSearchNumber() {
	fmt.Println(SearchNumber([]int{-4, -3, -1, 0, 1, 2, 5}))
	// Output:
	// [-1 1]
}
