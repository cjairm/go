package algorithm

import (
	"fmt"
	"testing"
)

func TestGetPivotQuickSort(t *testing.T) {
	test1, _ := GetPivotQuickSort([]int{5, 2, 1, 8, 4, 7, 6, 3}, 0, 0)

	if test1 != 4 {
		t.Error("Got: ", test1, "Expected: ", 4)
	}

	test2, _ := GetPivotQuickSort([]int{}, 0, 0)

	if test2 != 0 {
		t.Error("Got: ", test2, "Expected: ", 0)
	}

	test3, _ := GetPivotQuickSort([]int{9, 1, 2, 3, 10, 11}, 0, 0)

	if test3 != 3 {
		t.Error("Got: ", test3, "Expected: ", 3)
	}

	test4, _ := GetPivotQuickSort([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 0, 0)

	if test4 != 0 {
		t.Error("Got: ", test4, "Expected: ", 0)
	}

	test5, _ := GetPivotQuickSort([]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 0, 0)

	if test5 != 9 {
		t.Error("Got: ", test5, "Expected: ", 9)
	}
}

func BenchmarkGetPivotQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPivotQuickSort([]int{5, 2, 1, 8, 4, 7, 6, 3}, 0, 0)
	}
}

func ExampleGetPivotQuickSort() {
	fmt.Println(GetPivotQuickSort([]int{5, 2, 1, 8, 4, 7, 6, 3}, 0, 0))
	// Output:
	// 4 [3 2 1 4 5 7 6 8]
}
