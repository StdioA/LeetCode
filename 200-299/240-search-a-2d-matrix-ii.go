package main

import (
	"encoding/json"
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	m, n := 0, 0
	// binary search m
	minM, maxM := 0, len(matrix)-1
	for minM < maxM {
		midM := (minM + maxM) / 2
		if matrix[midM][0] < target {
			minM = midM + 1
		} else if matrix[midM][0] > target {
			maxM = midM
		} else {
			return true
		}
	}
	m = minM
	for m >= 0 && matrix[m][0] > target {
		m--
	}

	// binary search n
	minN, maxN := 0, len(matrix[0])-1
	for minN < maxN {
		midN := (minN + maxN) / 2
		if matrix[midN][0] < target {
			minN = midN + 1
		} else if matrix[midN][0] > target {
			maxN = midN
		} else {
			return true
		}
	}
	n = minN
	for n >= 0 && matrix[0][n] > target {
		n--
	}
	fmt.Println(m, n)
	if m < 0 || n < 0 {
		return false
	}

	// binary search row
	minM, maxM = 0, len(matrix)-1
	if maxM < minM {
		maxM = minM
	}
	for minM < maxM {
		midM := (minM + maxM) / 2
		if matrix[midM][n] < target {
			minM = midM + 1
		} else if matrix[midM][n] > target {
			maxM = midM
		} else {
			return true
		}
	}

	// binary search col
	if maxM < minM {
		maxM = minM
	}
	minM, maxM = 0, len(matrix[0])-1
	for minM < maxM {
		midM := (minM + maxM) / 2
		if matrix[m][midM] < target {
			minM = midM + 1
		} else if matrix[m][midM] > target {
			maxM = midM
		} else {
			return true
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
