package algorithm

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	test1 := ContainsDuplicate([]int{1, 10, 19, 20, 2, 3, 10, 18, 10})

	if !test1 {
		t.Error("Got: ", test1, "Expected: ", true)
	}

	test2 := ContainsDuplicate([]int{1, 2, 3, 4, 5, 6, 7})

	if test2 {
		t.Error("Got: ", test2, "Expected: ", false)
	}

	test3 := ContainsDuplicate([]int{-1, -100, 3, 99})

	if test3 {
		t.Error("Got: ", test3, "Expected: ", false)
	}

	test4 := ContainsDuplicate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 2, 10, 10})

	if !test4 {
		t.Error("Got: ", test4, "Expected: ", true)
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsDuplicate([]int{1, 2, 3, 4, 5, 6, 7})
	}
}

func ExampleContainsDuplicate() {
	fmt.Println(ContainsDuplicate([]int{1, 2, 3, 4, 5, 6, 7}))
	// Output:
	// false
}
