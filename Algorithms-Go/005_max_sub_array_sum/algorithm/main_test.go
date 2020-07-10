package algorithm

import (
	"fmt"
	"testing"
)

func TestMaxSubArraySum(t *testing.T) {
	test1 := MaxSubArraySum([]int{1, 2, 5, 2, 8, 1, 5}, 4)

	if test1 != 17 {
		t.Error("Got: ", test1, "Expected: ", 17)
	}

	test2 := MaxSubArraySum([]int{1, 2, 5, 2, 8, 1, 5}, 2)

	if test2 != 10 {
		t.Error("Got: ", test2, "Expected: ", 10)
	}

	test3 := MaxSubArraySum([]int{}, 2)

	if test3 != 0 {
		t.Error("Got: ", test3, "Expected: ", 0)
	}

	test4 := MaxSubArraySum([]int{4, 2, 1, 6}, 1)

	if test4 != 6 {
		t.Error("Got: ", test4, "Expected: ", 6)
	}

	test5 := MaxSubArraySum([]int{4, 2, 1, 6, 2}, 4)

	if test5 != 13 {
		t.Error("Got: ", test5, "Expected: ", 13)
	}
}

func BenchmarkMaxSubArraySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSubArraySum([]int{1, 2, 5, 2, 8, 1, 5}, 4)
	}
}

func ExampleMaxSubArraySum() {
	fmt.Println(MaxSubArraySum([]int{1, 2, 5, 2, 8, 1, 5}, 4))
	// Output:
	// 17
}
