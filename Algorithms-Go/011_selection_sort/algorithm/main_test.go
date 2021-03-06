package algorithm

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	test1 := SelectionSort([]int{1, 10, 19, 20, 2, 3, 10, 18, 10})

	if !testEq([]int{1, 2, 3, 10, 10, 10, 18, 19, 20}, test1) {
		t.Error("Got: ", test1, "Expected: ", []int{1, 2, 3, 10, 10, 10, 18, 19, 20})
	}

	test2 := SelectionSort([]int{})

	if !testEq([]int{}, test2) {
		t.Error("Got: ", test2, "Expected: ", []int{})
	}

	test3 := SelectionSort([]int{9, 10, 1, 2})

	if !testEq([]int{1, 2, 9, 10}, test3) {
		t.Error("Got: ", test3, "Expected: ", []int{1, 2, 9, 10})
	}

	test4 := SelectionSort([]int{1, 0, 0, 1})

	if !testEq([]int{0, 0, 1, 1}, test4) {
		t.Error("Got: ", test4, "Expected: ", []int{0, 0, 1, 1})
	}

	test5 := SelectionSort([]int{100, 20, 90, 80, 10, 0, 30, 70, 50, 60, 11, 23, 34, 45, 56, 67, 78, 99})

	if !testEq([]int{0, 10, 11, 20, 23, 30, 34, 45, 50, 56, 60, 67, 70, 78, 80, 90, 99, 100}, test5) {
		t.Error("Got: ", test5, "Expected: ", []int{0, 10, 11, 20, 23, 30, 34, 45, 50, 56, 60, 67, 70, 78, 80, 90, 99, 100})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{1, 10, 19, 20, 2, 3, 10, 18, 10})
	}
}

func ExampleSelectionSort() {
	fmt.Println(SelectionSort([]int{1, 10, 19, 20, 2, 3, 10, 18, 10}))
	// Output:
	// [1 2 3 10 10 10 18 19 20]
}

func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	// The length of the slices has to be the same
	if len(a) != len(b) {
		return false
	}

	// Check element by element
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
