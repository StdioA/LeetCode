package main

import (
	"encoding/json"
	"fmt"
)

func setZeroes(matrix [][]int) {
	var zeroInFirstRow bool
	for i, row := range matrix {
		var flushRow bool
		for j, num := range row {
			if num != 0 {
				continue
			}
			matrix[0][j] = 0
			if i == 0 {
				zeroInFirstRow = true
			} else {
				flushRow = true
			}
		}
		if flushRow {
			for j := range row {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			// Set column to 0
			for i := range matrix {
				matrix[i][j] = 0
			}
		} else if zeroInFirstRow {
			matrix[0][j] = 0
		}
	}
}

func main() {
	var matrix [][]int
	json.Unmarshal([]byte("[[1,1,1],[1,0,1],[1,1,1]]"), &matrix)
	setZeroes(matrix)
	fmt.Println(matrix)
}
