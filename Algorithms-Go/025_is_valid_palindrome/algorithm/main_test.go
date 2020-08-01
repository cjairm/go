package algorithm

import (
	"fmt"
	"testing"
)

func TestIsValidPalindrome(t *testing.T) {
	test1 := IsValidPalindrome("wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwaabwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww")

	if test1 {
		t.Error("Got: ", test1, "Expected: ", false)
	}

	test2 := IsValidPalindrome("Eva, can I see bees in a cave?")

	if !test2 {
		t.Error("Got: ", test2, "Expected: ", true)
	}

	test3 := IsValidPalindrome("race a car")

	if test3 {
		t.Error("Got: ", test3, "Expected: ", false)
	}

	test4 := IsValidPalindrome("Red rum, sir, is murder")

	if !test4 {
		t.Error("Got: ", test4, "Expected: ", true)
	}

	test5 := IsValidPalindrome("A man, a plan, a canal: Panama")

	if !test5 {
		t.Error("Got: ", test5, "Expected: ", true)
	}
}

func BenchmarkIsValidPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidPalindrome("A man, a plan, a canal: Panama")
	}
}

func ExampleIsValidPalindrome() {
	fmt.Println(IsValidPalindrome("A man, a plan, a canal: Panama"))
	// Output:
	// true
}
