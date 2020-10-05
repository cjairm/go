package main

func main() {

}

func shuffle(nums []int, n int) []int {
	x := nums[0:n]
	y := nums[n:]
	i := 0
	j := 0
	result := []int{}

	for len(x) > i && len(y) > j {
		result = append(result, x[i])
		result = append(result, y[j])
		i++
		j++
	}

	for len(x) > i {
		result = append(result, x[i])
		i++
	}

	for len(y) > j {
		result = append(result, y[j])
		j++
	}

	return result
}
