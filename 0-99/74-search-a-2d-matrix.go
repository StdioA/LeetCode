package main

func searchMatrix(matrix [][]int, target int) bool {
	var index = -1
	for i, row := range matrix {
		if row[0] > target {
			index = i - 1
			break
		}
	}
	if index == -1 {
		index = len(matrix) - 1
	}
	for _, n := range matrix[index] {
		if n == target {
			return true
		}
	}
	return false
}
