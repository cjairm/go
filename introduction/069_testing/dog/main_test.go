package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(5)
	if y != 35 {
		t.Error("Got: ", y, "Expected: ", 35)
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(7)
	if y != 49 {
		t.Error("Got: ", y, "Expected: ", 49)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(5)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(5)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output:
	// 35
}
