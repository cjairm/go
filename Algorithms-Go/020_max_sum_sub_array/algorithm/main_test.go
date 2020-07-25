package algorithm

import (
	"fmt"
	"testing"
)

func TestMaxSumSubArray(t *testing.T) {
	test1 := MaxSumSubArray([]int{100, 200, 300, 400}, 2)

	if test1 != 700 {
		t.Error("Got: ", test1, "Expected: ", 700)
	}

	test2 := MaxSumSubArray([]int{1, 4, 2, 10, 23, 3, 1, 0, 20}, 4)

	if test2 != 39 {
		t.Error("Got: ", test2, "Expected: ", 39)
	}

	test3 := MaxSumSubArray([]int{-3, 4, 0, -2, 6, -1}, 2)

	if test3 != 5 {
		t.Error("Got: ", test3, "Expected: ", 5)
	}

	test4 := MaxSumSubArray([]int{3, -2, 7, -4, 1, -1, 4, -2, 1}, 2)

	if test4 != 5 {
		t.Error("Got: ", test4, "Expected: ", 5)
	}

	test5 := MaxSumSubArray([]int{2, 3}, 3)

	if test5 != 0 {
		t.Error("Got: ", test5, "Expected: ", 0)
	}
}

func BenchmarkMaxSumSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSumSubArray([]int{1, 4, 2, 10, 23, 3, 1, 0, 20}, 4)
	}
}

func ExampleMaxSumSubArray() {
	fmt.Println(MaxSumSubArray([]int{1, 4, 2, 10, 23, 3, 1, 0, 20}, 4))
	// Output:
	// 39
}
