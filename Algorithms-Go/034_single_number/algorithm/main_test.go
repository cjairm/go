package algorithm

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	test1 := SingleNumber([]int{2, 2, 1})

	if test1 != 1 {
		t.Error("Got: ", test1, "Expected: ", 1)
	}

	test2 := SingleNumber([]int{4, 1, 2, 1, 2})

	if test2 != 4 {
		t.Error("Got: ", test2, "Expected: ", 4)
	}

	test3 := SingleNumber([]int{1, 1, 3, 5, 3, 4, 2, 4, 2})

	if test3 != 5 {
		t.Error("Got: ", test3, "Expected: ", 5)
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleNumber([]int{1, 1, 3, 5, 3, 4, 2, 4, 2})
	}
}

func ExampleSingleNumber() {
	fmt.Println(SingleNumber([]int{1, 1, 3, 5, 3, 4, 2, 4, 2}))
	// Output:
	// 5
}

func TestSingleNumberTwo(t *testing.T) {
	test1 := SingleNumberTwo([]int{2, 2, 1})

	if test1 != 1 {
		t.Error("Got: ", test1, "Expected: ", 1)
	}

	test2 := SingleNumberTwo([]int{4, 1, 2, 1, 2})

	if test2 != 4 {
		t.Error("Got: ", test2, "Expected: ", 4)
	}

	test3 := SingleNumberTwo([]int{1, 1, 3, 5, 3, 4, 2, 4, 2})

	if test3 != 5 {
		t.Error("Got: ", test3, "Expected: ", 5)
	}
}

func BenchmarkSingleNumberTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleNumberTwo([]int{1, 1, 3, 5, 3, 4, 2, 4, 2})
	}
}

func ExampleSingleNumberTwo() {
	fmt.Println(SingleNumberTwo([]int{1, 1, 3, 5, 3, 4, 2, 4, 2}))
	// Output:
	// 5
}
