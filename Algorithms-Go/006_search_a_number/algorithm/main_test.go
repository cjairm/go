package algorithm

import (
	"fmt"
	"testing"
)

func TestSearchNumber(t *testing.T) {
	test1 := SearchNumber([]int{-100, -90, -80, -70, -60, -50, -40, -30, -20, -10, -5, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, 0)

	if test1 != 11 {
		t.Error("Got: ", test1, "Expected: ", 11)
	}

	test2 := SearchNumber([]int{-20, -18, -16, -14, -12, -10, -8, -6, -4, -2, 0, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 19)

	if test2 != 20 {
		t.Error("Got: ", test2, "Expected: ", 20)
	}

	test3 := SearchNumber([]int{}, 10)

	if test3 != -1 {
		t.Error("Got: ", test3, "Expected: ", -1)
	}

	test4 := SearchNumber([]int{1, 2}, 2)

	if test4 != 1 {
		t.Error("Got: ", test4, "Expected: ", 1)
	}

	test5 := SearchNumber([]int{-20, -18, -16, -14, -12, -10, -8, -6, -4, -2, 0, 1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 30)

	if test5 != -1 {
		t.Error("Got: ", test5, "Expected: ", -1)
	}
}

func BenchmarkSearchNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchNumber([]int{-100, -90, -80, -70, -60, -50, -40, -30, -20, -10, -5, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, -90)
	}
}

func ExampleSearchNumber() {
	fmt.Println(SearchNumber([]int{-100, -90, -80, -70, -60, -50, -40, -30, -20, -10, -5, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, -90))
	// Output:
	// 1
}
