package algorithm

import (
	"fmt"
	"testing"
)

func TestProductOfArray(t *testing.T) {
	test1 := ProductOfArray([]int{1, 2, 3})

	if test1 != 6 {
		t.Error("Got: ", test1, "Expected: ", 6)
	}

	test2 := ProductOfArray([]int{1, 2, 3, 10})

	if test2 != 60 {
		t.Error("Got: ", test2, "Expected: ", 60)
	}

	test3 := ProductOfArray([]int{10})

	if test3 != 10 {
		t.Error("Got: ", test3, "Expected: ", 10)
	}

	test4 := ProductOfArray([]int{})

	if test4 != 0 {
		t.Error("Got: ", test4, "Expected: ", 0)
	}

	test5 := ProductOfArray([]int{1, 1, 1, 1, 1, 2, 90})

	if test5 != 180 {
		t.Error("Got: ", test5, "Expected: ", 180)
	}
}

func BenchmarkProductOfArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductOfArray([]int{1, 2, 3})
	}
}

func ExampleProductOfArray() {
	fmt.Println(ProductOfArray([]int{1, 2, 3}))
	// Output:
	// 6
}
