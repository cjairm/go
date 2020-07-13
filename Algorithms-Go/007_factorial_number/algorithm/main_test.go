package algorithm

import (
	"fmt"
	"testing"
)

func TestFactorialNumber(t *testing.T) {
	test1 := FactorialNumber(0)

	if test1 != 0 {
		t.Error("Got: ", test1, "Expected: ", 0)
	}

	test2 := FactorialNumber(1)

	if test2 != 1 {
		t.Error("Got: ", test2, "Expected: ", 1)
	}

	test3 := FactorialNumber(2)

	if test3 != 2 {
		t.Error("Got: ", test3, "Expected: ", 2)
	}

	test4 := FactorialNumber(3)

	if test4 != 6 {
		t.Error("Got: ", test4, "Expected: ", 6)
	}

	test5 := FactorialNumber(4)

	if test5 != 24 {
		t.Error("Got: ", test5, "Expected: ", 24)
	}
}

func BenchmarkFactorialNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialNumber(3)
	}
}

func ExampleFactorialNumber() {
	fmt.Println(FactorialNumber(3))
	// Output:
	// 6
}
