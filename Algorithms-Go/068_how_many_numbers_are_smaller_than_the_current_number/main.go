package main

func main() {

}

func smallerNumbersThanCurrent(nums []int) []int {
	smallers := []int{}
	var i, j, ni, nj, count int

	for len(nums) > i {
		ni = nums[i]
		nj = nums[j]

		if i != j {
			if ni > nj {
				count++
			}
		}

		j++

		if j >= len(nums) {
			smallers = append(smallers, count)
			i++
			j = 0
			count = 0
		}
	}

	return smallers
}
