package algorithm

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	test1 := Intersect([]int{1, 2, 2, 1}, []int{2, 2})

	if !testEq(test1, []int{2, 2}) {
		t.Error("Got: ", test1, "Expected: ", []int{2, 2})
	}

	test2 := Intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})

	if !testEq(test2, []int{4, 9}) {
		t.Error("Got: ", test2, "Expected: ", []int{4, 9})
	}

	test3 := Intersect([]int{4}, []int{9, 0, 0, 0, 1, 2, 3, 4, 9, 8, 4})

	if !testEq(test3, []int{4}) {
		t.Error("Got: ", test3, "Expected: ", []int{4})
	}

	test4 := Intersect([]int{4}, []int{})

	if !testEq(test4, []int{}) {
		t.Error("Got: ", test4, "Expected: ", []int{})
	}
}

func BenchmarkIntersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intersect([]int{1, 2, 2, 1}, []int{2, 2})
	}
}

func ExampleIntersect() {
	fmt.Println(Intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	// Output:
	// [2 2]
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
