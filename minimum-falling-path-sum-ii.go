package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) < 2 {
		return 0
	}

	for i := len(matrix) - 2; i >= 0; i-- {
		// 求每行中最小的两个值
		min1, min2 := int(1e8), int(1e8)
		for _, n := range matrix[i+1] {
			if n < min1 {
				min1, n = n, min1
			}
			if n < min2 {
				min2 = n
			}
		}

		for j := range matrix[i] {
			var m = min1
			if matrix[i+1][j] == m {
				m = min2
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
