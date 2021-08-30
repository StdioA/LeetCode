package main

import (
	"encoding/json"
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	var (
		m, n = len(matrix), len(matrix[0])
		i, j = 0, n - 1
	)
	for i < m && j >= 0 {
		num := matrix[i][j]
		switch {
		case num == target:
			return true
		case num < target:
			i++
		default:
			j--
		}
	}
	return false
}

func main() {
	// s := "[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]"
	// s := "[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]"
	s := "[[-5]]"
	var matrix [][]int
	json.Unmarshal([]byte(s), &matrix)
	fmt.Println(searchMatrix(matrix, -5))
}
