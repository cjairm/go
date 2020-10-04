package main

func main() {

}

func runningSum(nums []int) []int {
	var results = []int{}
	i := 0
	j := 0
	temp := 0

	for len(nums) > i {
		if i >= j {
			temp += nums[j]
			j++
		} else {
			results = append(results, temp)
			temp = 0
			i++
			j = 0
		}
	}

	return results
}
