package algorithm

import (
	"fmt"
	"testing"
)

func TestRotateArray(t *testing.T) {
	test1 := RotateArray([]int{1, 10, 19, 20, 2, 3, 10, 18, 10}, 1)

	if !testEq([]int{10, 1, 10, 19, 20, 2, 3, 10, 18}, test1) {
		t.Error("Got: ", test1, "Expected: ", []int{10, 1, 10, 19, 20, 2, 3, 10, 18})
	}

	test2 := RotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 3)

	if !testEq([]int{5, 6, 7, 1, 2, 3, 4}, test2) {
		t.Error("Got: ", test2, "Expected: ", []int{5, 6, 7, 1, 2, 3, 4})
	}

	test3 := RotateArray([]int{-1, -100, 3, 99}, 2)

	if !testEq([]int{3, 99, -1, -100}, test3) {
		t.Error("Got: ", test3, "Expected: ", []int{3, 99, -1, -100})
	}
}

func BenchmarkRotateArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	}
}

func ExampleRotateArray() {
	fmt.Println(RotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	// Output:
	// [5 6 7 1 2 3 4]
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
