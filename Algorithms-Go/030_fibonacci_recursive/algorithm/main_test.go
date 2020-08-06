package algorithm

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	test1 := Fib(4)

	if test1 != 3 {
		t.Error("Got: ", test1, "Expected: ", 3)
	}

	test2 := Fib(10)

	if test2 != 55 {
		t.Error("Got: ", test2, "Expected: ", 55)
	}

	test3 := Fib(28)

	if test3 != 317811 {
		t.Error("Got: ", test3, "Expected: ", 317811)
	}

	test4 := Fib(35)

	if test4 != 9227465 {
		t.Error("Got: ", test4, "Expected: ", 9227465)
	}

	test5 := Fib(1)

	if test5 != 1 {
		t.Error("Got: ", test5, "Expected: ", 1)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(4)
	}
}

func ExampleFib() {
	fmt.Println(Fib(4))
	// Output:
	// 3
}
