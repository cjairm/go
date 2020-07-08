package algorithm

import (
	"fmt"
	"testing"
)

func TestSameButSquared(t *testing.T) {
	test1 := SameButSquared([]int{0, 1, 2, 3}, []int{0, 1, 2, 3})

	if test1 {
		t.Error("Got: ", test1, "Expected: ", false)
	}

	test2 := SameButSquared([]int{1, 1, 1, 10, 10, 10, 5, 5, 5, 3, 3, 3, 2, 2, 2}, []int{100, 100, 100, 25, 25, 25, 9, 9, 9, 4, 4, 4, 1, 1, 1})

	if !test2 {
		t.Error("Got: ", test2, "Expected: ", true)
	}

	test3 := SameButSquared([]int{2, 2, 2}, []int{4, 4, 4})

	if !test3 {
		t.Error("Got: ", test3, "Expected: ", true)
	}

	test4 := SameButSquared([]int{}, []int{})

	if !test4 {
		t.Error("Got: ", test4, "Expected: ", true)
	}

	test5 := SameButSquared([]int{}, []int{1, 2, 3})

	if test5 {
		t.Error("Got: ", test5, "Expected: ", false)
	}
}

func BenchmarkSameButSquared(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameButSquared([]int{1, 2, 3}, []int{9, 4, 1})
	}
}

func ExampleSameButSquared() {
	fmt.Println(SameButSquared([]int{1, 2, 3}, []int{9, 4, 1}))
	// Output:
	// true
}
