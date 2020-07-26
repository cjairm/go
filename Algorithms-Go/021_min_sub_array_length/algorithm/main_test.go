package algorithm

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLength(t *testing.T) {
	test1 := MinSubArrayLength([]int{2, 3, 1, 2, 4, 3}, 7)

	if test1 != 2 {
		t.Error("Got: ", test1, "Expected: ", 2)
	}

	test2 := MinSubArrayLength([]int{2, 1, 6, 5, 4}, 9)

	if test2 != 2 {
		t.Error("Got: ", test2, "Expected: ", 2)
	}

	test3 := MinSubArrayLength([]int{3, 1, 7, 11, 2, 9, 8, 21, 62, 33, 19}, 52)

	if test3 != 1 {
		t.Error("Got: ", test3, "Expected: ", 1)
	}

	test4 := MinSubArrayLength([]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 39)

	if test4 != 3 {
		t.Error("Got: ", test4, "Expected: ", 3)
	}

	test5 := MinSubArrayLength([]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 55)

	if test5 != 5 {
		t.Error("Got: ", test5, "Expected: ", 5)
	}

	test6 := MinSubArrayLength([]int{4, 3, 3, 8, 1, 2, 3}, 11)

	if test6 != 2 {
		t.Error("Got: ", test6, "Expected: ", 2)
	}

	test7 := MinSubArrayLength([]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 95)

	if test7 != 0 {
		t.Error("Got: ", test7, "Expected: ", 0)
	}
}

func BenchmarkMinSubArrayLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinSubArrayLength([]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 39)
	}
}

func ExampleMinSubArrayLength() {
	fmt.Println(MinSubArrayLength([]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 39))
	// Output:
	// 3
}
