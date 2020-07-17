package algorithm

import (
	"fmt"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	test1 := MergeSortedArrays([]int{1, 10, 19, 20}, []int{2, 11, 20, 21})

	if !testEq([]int{1, 2, 10, 11, 19, 20, 20, 21}, test1) {
		t.Error("Got: ", test1, "Expected: ", []int{1, 2, 10, 11, 19, 20, 20, 21})
	}

	test2 := MergeSortedArrays([]int{}, []int{1, 2, 3})

	if !testEq([]int{1, 2, 3}, test2) {
		t.Error("Got: ", test2, "Expected: ", []int{1, 2, 3})
	}

	test3 := MergeSortedArrays([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, []int{})

	if !testEq([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, test3) {
		t.Error("Got: ", test3, "Expected: ", []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100})
	}

	test4 := MergeSortedArrays([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, []int{1, 2, 10, 11, 19, 21})

	if !testEq([]int{1, 2, 10, 10, 11, 19, 20, 21, 30, 40, 50, 60, 70, 80, 90, 100}, test4) {
		t.Error("Got: ", test4, "Expected: ", []int{0, 0, 1, 1})
	}

	test5 := MergeSortedArrays([]int{1, 2, 10, 10, 11, 19, 20, 21, 30, 40, 50, 60, 70, 80, 90, 100}, []int{3, 4, 12, 13, 14, 18, 22, 23, 31, 41, 51, 61, 71, 81, 91, 101})

	if !testEq([]int{1, 2, 3, 4, 10, 10, 11, 12, 13, 14, 18, 19, 20, 21, 22, 23, 30, 31, 40, 41, 50, 51, 60, 61, 70, 71, 80, 81, 90, 91, 100, 101}, test5) {
		t.Error("Got: ", test5, "Expected: ", []int{1, 2, 3, 4, 10, 10, 11, 12, 13, 14, 18, 19, 20, 21, 22, 23, 30, 31, 40, 41, 50, 51, 60, 61, 70, 71, 80, 81, 90, 91, 100, 101})
	}
}

func BenchmarkMergeSortedArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSortedArrays([]int{1, 10, 19, 20}, []int{2, 11, 20, 21})
	}
}

func ExampleMergeSortedArrays() {
	fmt.Println(MergeSortedArrays([]int{1, 10, 19, 20}, []int{2, 11, 20, 21}))
	// Output:
	// [1 2 10 11 19 20 20 21]
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
