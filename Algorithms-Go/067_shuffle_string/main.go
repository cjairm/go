package main

import (
	"strings"
)

func main() {

}

func restoreString(s string, indices []int) string {
	newWord := make([]string, len(s))

	for i, v := range s {
		newWord[indices[i]] = string(v)
	}

	return strings.Join(newWord, "")
}
