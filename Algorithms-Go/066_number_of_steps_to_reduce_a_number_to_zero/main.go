package main

func main() {

}

func numberOfSteps(num int) int {
	counter := 0

	return reduceToZero(num, &counter)
}

func reduceToZero(n int, c *int) int {
	if n <= 0 {
		return *c
	}

	*c = *c + 1

	if n%2 == 0 {
		return reduceToZero(n/2, c)
	} else {
		return reduceToZero(n-1, c)
	}
}
