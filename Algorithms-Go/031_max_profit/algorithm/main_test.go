package algorithm

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	test1 := MaxProfit([]int{7, 1, 5, 3, 6, 4})

	if test1 != 7 {
		t.Error("Got: ", test1, "Expected: ", 7)
	}

	test2 := MaxProfit([]int{1, 2, 3, 4, 5})

	if test2 != 4 {
		t.Error("Got: ", test2, "Expected: ", 4)
	}

	test3 := MaxProfit([]int{7, 6, 4, 3, 1})

	if test3 != 0 {
		t.Error("Got: ", test3, "Expected: ", 0)
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxProfit([]int{7, 1, 5, 3, 6, 4})
	}
}

func ExampleMaxProfit() {
	fmt.Println(MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	// Output:
	// 7
}
