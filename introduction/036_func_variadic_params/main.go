package main

import "fmt"

func main() {
	numsToSum := []int{1, 2, 10}
	sum := sumNums(numsToSum...)
	fmt.Println(sum)

	numsToSubs := []int{10, 5, 1}
	subs := subsNums(numsToSubs)
	fmt.Println(subs)
}

func sumNums(xi ...int) int {
	// fmt.Printf("Type: %T, Value: ", xi)
	var totalSum int
	for _, val := range xi {
		totalSum += val
	}
	return totalSum
}

func subsNums(xi []int) int {
	// fmt.Printf("Type: %T, Value: ", xi)
	var totalSub int
	if len(xi) <= 0 {
		return 0
	}
	totalSub = xi[:1][0]
	for _, val := range xi[1:] {
		totalSub -= val
	}
	return totalSub
}
