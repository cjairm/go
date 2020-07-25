package algorithm

// import (
// 	"fmt"
// 	"testing"
// )

// func TestAveragePair(t *testing.T) {
// 	test1 := AveragePair([]int{19, 3, 3, 10, 6, 7, 5, 12, 1}, 8)

// 	if !test1 {
// 		t.Error("Got: ", test1, "Expected: ", true)
// 	}

// 	test2 := AveragePair([]int{1, 2, 3}, 2.5)

// 	if !test2 {
// 		t.Error("Got: ", test2, "Expected: ", true)
// 	}

// 	test3 := AveragePair([]int{-1, 0, 3, 4, 5, 6}, 4.1)

// 	if test3 {
// 		t.Error("Got: ", test3, "Expected: ", false)
// 	}

// 	test4 := AveragePair([]int{-100, -10, -1, 0, 3, 4, 5, 6, 100, 10, 2, 7, 8}, 100)

// 	if test4 {
// 		t.Error("Got: ", test4, "Expected: ", false)
// 	}

// 	test5 := AveragePair([]int{}, 4)

// 	if test5 {
// 		t.Error("Got: ", test5, "Expected: ", false)
// 	}
// }

// func BenchmarkAveragePair(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		AveragePair([]int{19, 3, 3, 10, 6, 7, 5, 12, 1}, 8)
// 	}
// }

// func ExampleAveragePair() {
// 	fmt.Println(AveragePair([]int{19, 3, 3, 10, 6, 7, 5, 12, 1}, 8))
// 	// Output:
// 	// true
// }
