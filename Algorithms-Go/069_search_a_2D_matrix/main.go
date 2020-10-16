package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	var i, j int

	if len(matrix) == 0 {
		return false
	}

	for i < len(matrix) {
		if len(matrix[i]) == 0 {
			return false
		}

		if matrix[i][j] == target {
			return true
		}

		j++
		if j >= len(matrix[i]) {
			i++
			j = 0
		}
	}

	return false
}
