package algorithm

import (
	"fmt"
	"testing"
)

func TestAreThereDuplicates(t *testing.T) {
	test1 := AreThereDuplicates(1, 10, 19, 20, 2, 3, 10, 18, 10)

	if !test1 {
		t.Error("Got: ", test1, "Expected: ", true)
	}

	test2 := AreThereDuplicates()

	if test2 {
		t.Error("Got: ", test2, "Expected: ", false)
	}

	test3 := AreThereDuplicates(9, 10, 1, 2)

	if test3 {
		t.Error("Got: ", test3, "Expected: ", false)
	}

	test4 := AreThereDuplicates(1, 0, 0, 1)

	if !test4 {
		t.Error("Got: ", test4, "Expected: ", true)
	}

	test5 := AreThereDuplicates(100, 20, 90, 80, 10, 0, 30, 70, 50, 60, 11, 23, 34, 45, 56, 67, 78, 99, 1000, 16, 71, 18, 19)

	if test5 {
		t.Error("Got: ", test5, "Expected: ", false)
	}
}

func BenchmarkAreThereDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreThereDuplicates(1, 10, 19, 20, 2, 3, 10, 18, 10)
	}
}

func ExampleAreThereDuplicates() {
	fmt.Println(AreThereDuplicates(1, 10, 19, 20, 2, 3, 10, 18, 10))
	// Output:
	// true
}
