package main

func main() {

}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	results := []bool{}
	max := 0
	cm := 0

	for _, v := range candies {
		if v >= max {
			max = v
		}
	}

	for _, v := range candies {
		cm = v + extraCandies
		if cm >= max {
			results = append(results, true)
		} else {
			results = append(results, false)
		}
	}

	return results
}
