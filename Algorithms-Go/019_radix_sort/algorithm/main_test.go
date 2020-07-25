package algorithm

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	test1 := []int32{1, 10, 19, 20, 2, 3, 10, 18, 10}
	RadixSort(test1)

	if !testEq([]int32{1, 2, 3, 10, 10, 10, 18, 19, 20}, test1) {
		t.Error("Got: ", test1, "Expected: ", []int32{1, 2, 3, 10, 10, 10, 18, 19, 20})
	}

	test2 := []int32{}
	RadixSort(test2)

	if !testEq([]int32{}, test2) {
		t.Error("Got: ", test2, "Expected: ", []int32{})
	}

	test3 := []int32{9, 10, 1, 2}
	RadixSort(test3)

	if !testEq([]int32{1, 2, 9, 10}, test3) {
		t.Error("Got: ", test3, "Expected: ", []int32{1, 2, 9, 10})
	}

	test4 := []int32{1, 0, 0, 1}
	RadixSort(test4)

	if !testEq([]int32{0, 0, 1, 1}, test4) {
		t.Error("Got: ", test4, "Expected: ", []int32{0, 0, 1, 1})
	}

	test5 := []int32{100, 20, 90, 80, 10, 0, 30, 70, 50, 60, 11, 23, 34, 45, 56, 67, 78, 99}
	RadixSort(test5)

	if !testEq([]int32{0, 10, 11, 20, 23, 30, 34, 45, 50, 56, 60, 67, 70, 78, 80, 90, 99, 100}, test5) {
		t.Error("Got: ", test5, "Expected: ", []int32{0, 10, 11, 20, 23, 30, 34, 45, 50, 56, 60, 67, 70, 78, 80, 90, 99, 100})
	}
}

func BenchmarkRadixSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RadixSort([]int32{1, 10, 19, 20, 2, 3, 10, 18, 10})
	}
}

func ExampleRadixSort() {
	val := []int32{1, 10, 19, 20, 2, 3, 10, 18, 10}
	RadixSort(val)
	fmt.Println(val)
	// Output:
	// [1 2 3 10 10 10 18 19 20]
}

func testEq(a, b []int32) bool {

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
