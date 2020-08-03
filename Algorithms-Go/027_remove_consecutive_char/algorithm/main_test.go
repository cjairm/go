package algorithm

import (
	"fmt"
	"testing"
)

func TestRemoveConsecutiveChar(t *testing.T) {
	test1 := RemoveConsecutiveChar("wwwwwwaabwwwwwww", 3)

	if test1 != "wwaabw" {
		t.Error("Got: ", test1, "Expected: ", "wwaabw")
	}

	test2 := RemoveConsecutiveChar("Eva, can I see bees in a cave?", 3)

	if test2 != "Eva, can I see bees in a cave?" {
		t.Error("Got: ", test2, "Expected: ", "Eva, can I see bees in a cave?")
	}

	test3 := RemoveConsecutiveChar("abbcccb", 3)

	if test3 != "a" {
		t.Error("Got: ", test3, "Expected: ", "a")
	}
}

func BenchmarkRemoveConsecutiveChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveConsecutiveChar("abbcccb", 3)
	}
}

func ExampleRemoveConsecutiveChar() {
	fmt.Println(RemoveConsecutiveChar("abbcccb", 3))
	// Output:
	// a
}
