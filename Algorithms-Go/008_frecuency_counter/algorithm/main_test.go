package algorithm

import (
	"fmt"
	"testing"
)

func TestSameFrecuency(t *testing.T) {
	test1 := SameFrecuency(3589578, 5879385)

	if !test1 {
		t.Error("Got: ", test1, "Expected: ", true)
	}

	test2 := SameFrecuency(182, 281)

	if !test2 {
		t.Error("Got: ", test2, "Expected: ", true)
	}

	test3 := SameFrecuency(34, 14)

	if test3 {
		t.Error("Got: ", test3, "Expected: ", false)
	}

	test4 := SameFrecuency(22, 222)

	if test4 {
		t.Error("Got: ", test4, "Expected: ", false)
	}

	test5 := SameFrecuency(35895783589578351, 57835783518358995)

	if !test5 {
		t.Error("Got: ", test5, "Expected: ", true)
	}
}

func BenchmarkSameFrecuency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SameFrecuency(3589578, 5879385)
	}
}

func ExampleSameFrecuency() {
	fmt.Println(SameFrecuency(3589578, 5879385))
	// Output:
	// true
}
