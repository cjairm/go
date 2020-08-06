package algorithm

import (
	"fmt"
	"testing"
)

func TestRecursiveRange(t *testing.T) {
	test1 := RecursiveRange(6)

	if test1 != 21 {
		t.Error("Got: ", test1, "Expected: ", 21)
	}

	test2 := RecursiveRange(1)

	if test2 != 1 {
		t.Error("Got: ", test2, "Expected: ", 1)
	}

	test3 := RecursiveRange(0)

	if test3 != 0 {
		t.Error("Got: ", test3, "Expected: ", 0)
	}

	test4 := RecursiveRange(10)

	if test4 != 55 {
		t.Error("Got: ", test4, "Expected: ", 55)
	}

	test5 := RecursiveRange(7)

	if test5 != 28 {
		t.Error("Got: ", test5, "Expected: ", 28)
	}
}

func BenchmarkRecursiveRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveRange(100)
	}
}

func ExampleRecursiveRange() {
	fmt.Println(RecursiveRange(6))
	// Output:
	// 21
}
