package algorithm

import (
	"fmt"
	"testing"
)

func TestClockAngle(t *testing.T) {
	test1, error1 := ClockAngle(3, 0)

	if error1 != nil || test1 != 90 {
		t.Error("Got: ", test1, "Expected: ", 90)
	}

	test2, error2 := ClockAngle(20, 40)

	if error2 == nil {
		t.Error("Got: ", test2, "Expected: ", 0)
	}

	test3, error3 := ClockAngle(6, 0)

	if error3 != nil || test3 != 180 {
		t.Error("Got: ", test3, "Expected: ", 180)
	}

	test4, error4 := ClockAngle(0, 30)

	if error4 != nil || test4 != -165 {
		t.Error("Got: ", test4, "Expected: ", -165)
	}
}

func BenchmarkClockAngle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClockAngle(3, 0)
	}
}

func ExampleClockAngle() {
	fmt.Println(ClockAngle(3, 0))
	// Output:
	// 90 <nil>
}
