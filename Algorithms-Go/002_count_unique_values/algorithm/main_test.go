package algorithm

import (
	"fmt"
	"testing"
)

func TestGetCountOfUniqueValues(t *testing.T) {
	test1 := GetCountOfUniqueValues([]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13})

	if test1 != 7 {
		t.Error("Got: ", test1, "Expected: ", 7)
	}

	test2 := GetCountOfUniqueValues([]int{-1, -1, -1, -1, -1, -1, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13})

	if test2 != 9 {
		t.Error("Got: ", test2, "Expected: ", 9)
	}

	test3 := GetCountOfUniqueValues([]int{-1, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 10, 10, 10})

	if test3 != 12 {
		t.Error("Got: ", test3, "Expected: ", 12)
	}

	test4 := GetCountOfUniqueValues([]int{})

	if test4 != 0 {
		t.Error("Got: ", test4, "Expected: ", 0)
	}

	test5 := GetCountOfUniqueValues([]int{1, 1})

	if test5 != 1 {
		t.Error("Got: ", test5, "Expected: ", 1)
	}

	test6 := GetCountOfUniqueValues([]int{1, 2})

	if test6 != 2 {
		t.Error("Got: ", test6, "Expected: ", 2)
	}

	test7 := GetCountOfUniqueValues([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0, 0, 0, 0, 0, 0, 0, 90000})

	if test7 != 3 {
		t.Error("Got: ", test7, "Expected: ", 3)
	}
}

func BenchmarkGetCountOfUniqueValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetCountOfUniqueValues([]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13})
	}
}

func ExampleGetCountOfUniqueValues() {
	fmt.Println(GetCountOfUniqueValues([]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}))
	// Output:
	// 7
}
