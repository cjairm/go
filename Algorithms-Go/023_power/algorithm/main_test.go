package algorithm

import (
	"fmt"
	"testing"
)

func TestPower(t *testing.T) {
	test1 := Power(2, 0)

	if test1 != 1 {
		t.Error("Got: ", test1, "Expected: ", 1)
	}

	test2 := Power(2, 4)

	if test2 != 16 {
		t.Error("Got: ", test2, "Expected: ", 16)
	}

	test3 := Power(2, 2)

	if test3 != 4 {
		t.Error("Got: ", test3, "Expected: ", 4)
	}

	test4 := Power(8, 8)

	if test4 != 16777216 {
		t.Error("Got: ", test4, "Expected: ", 16777216)
	}

	test5 := Power(9, 7)

	if test5 != 4782969 {
		t.Error("Got: ", test5, "Expected: ", 4782969)
	}
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Power(2, 4)
	}
}

func ExamplePower() {
	fmt.Println(Power(2, 4))
	// Output:
	// 16
}
