package main

import "testing"

func TestIsAnagramValid(t *testing.T) {
	test1 := isAnagramValid("ab", "a")
	if test1 {
		t.Errorf("Validation \"ab\" and \"a\"")
	}

	test2 := isAnagramValid("a", "a")
	if !test2 {
		t.Errorf("Validation \"a\" and \"a\"")
	}

	test3 := isAnagramValid("hello", "llohe")
	if !test3 {
		t.Errorf("Validation \"hello\" and \"llohe\"")
	}

	test4 := isAnagramValid("12ab34cd", "56ef78gh")
	if test4 {
		t.Errorf("Validation \"12ab34cd\" and \"56ef78gh\"")
	}

	test5 := isAnagramValid("aaaaaaaaaaaaaaaauuuuuuuuuuuuuuuuuuuiiiiiiiiiiiiiiiiiiiiii", "iiiiiiiiiiiaaaaaaauuuuiiiiiiiiiiiuuuuuaaaaaaaaauuuuuuuuuu")
	if !test5 {
		t.Errorf("Validation \"aaaaaaaaaaaaaaaauuuuuuuuuuuuuuuuuuuiiiiiiiiiiiiiiiiiiiiii\" and \"iiiiiiiiiiiaaaaaaauuuuiiiiiiiiiiiuuuuuaaaaaaaaauuuuuuuuuu\"")
	}

	test6 := isAnagramValid("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx i", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxi")
	if test6 {
		t.Errorf("Validation \"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx i\" and \"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxi\"")
	}
}
