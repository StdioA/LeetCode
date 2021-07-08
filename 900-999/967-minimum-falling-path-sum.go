package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 数字三角形
func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	var width = len(matrix[0])
	for i := len(matrix) - 2; i >= 0; i-- {
		for j := range matrix[i] {
			m := matrix[i+1][j]
			if j > 0 {
				m = min(m, matrix[i+1][j-1])
			}
			if j < width-1 {
				m = min(m, matrix[i+1][j+1])
			}
			matrix[i][j] += m
		}
	}
	m := int(1e8)
	for _, n := range matrix[0] {
		m = min(m, n)
	}
	return m
}
