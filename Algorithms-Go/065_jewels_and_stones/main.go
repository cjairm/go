package main

import "strings"

func main() {

}

func numJewelsInStones(J string, S string) int {
	jewels := strings.Split(J, "")
	stones := strings.Split(S, "")

	var i, j, count int
	var currS, currJ string

	for len(jewels) > i {
		currJ = jewels[i]
		currS = stones[j]

		if currJ == currS {
			count++
		}

		j++
		if j >= len(stones) {
			i++
			j = 0
		}
	}

	return count
}
