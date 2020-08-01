package algorithm

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	test1, _ := RemoveDuplicates([]int{1, 1, 2})

	if !testEq(test1, []int{1, 2}) {
		t.Error("Got: ", test1, "Expected: ", []int{1, 2})
	}

	test2, _ := RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})

	if !testEq(test2, []int{0, 1, 2, 3, 4}) {
		t.Error("Got: ", test2, "Expected: ", true)
	}

	test3, _ := RemoveDuplicates([]int{})

	if !testEq(test3, []int{}) {
		t.Error("Got: ", test3, "Expected: ", []int{})
	}

	test4, _ := RemoveDuplicates([]int{0, 0, 33, 22, 1, 0, 0, 2, 3, 4, 5, 66, 6, 6, 6, 77, 7, 33})

	if !testEq(test4, []int{0, 1, 2, 3, 4, 5, 6, 7, 22, 33, 66, 77}) {
		t.Error("Got: ", test4, "Expected: ", []int{0, 1, 2, 3, 4, 5, 6, 7, 22, 33, 66, 77})
	}

	test5, _ := RemoveDuplicates([]int{1, 1})

	if !testEq(test5, []int{1}) {
		t.Error("Got: ", test5, "Expected: ", []int{1})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	}
}

func ExampleRemoveDuplicates() {
	fmt.Println(RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4] 5
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
