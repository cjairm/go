package algorithm

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	test1 := PlusOne([]int{1, 2, 2, 1})

	if !testEq(test1, []int{1, 2, 2, 2}) {
		t.Error("Got: ", test1, "Expected: ", []int{1, 2, 2, 2})
	}

	test2 := PlusOne([]int{9, 4, 9, 8, 4})

	if !testEq(test2, []int{9, 4, 9, 8, 5}) {
		t.Error("Got: ", test2, "Expected: ", []int{9, 4, 9, 8, 5})
	}

	test3 := PlusOne([]int{8, 9, 9, 9})

	if !testEq(test3, []int{9, 0, 0, 0}) {
		t.Error("Got: ", test3, "Expected: ", []int{9, 0, 0, 0})
	}

	test4 := PlusOne([]int{9, 9, 9})

	if !testEq(test4, []int{1, 0, 0, 0}) {
		t.Error("Got: ", test4, "Expected: ", []int{1, 0, 0, 0})
	}

	test5 := PlusOne([]int{3, 8, 9, 9})

	if !testEq(test5, []int{3, 9, 0, 0}) {
		t.Error("Got: ", test5, "Expected: ", []int{3, 9, 0, 0})
	}

	test6 := PlusOne([]int{9})

	if !testEq(test6, []int{1, 0}) {
		t.Error("Got: ", test6, "Expected: ", []int{1, 0})
	}

	test7 := PlusOne([]int{0})

	if !testEq(test7, []int{1}) {
		t.Error("Got: ", test7, "Expected: ", []int{1})
	}
}

func BenchmarkPlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PlusOne([]int{1, 2, 2, 1})
	}
}

func ExamplePlusOne() {
	fmt.Println(PlusOne([]int{1, 2, 2, 1}))
	// Output:
	// [1 2 2 2]
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
