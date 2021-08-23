package main

func min(nums ...int) int {
	var m = nums[0]
	for _, n := range nums[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func maximalSquare(matrix [][]byte) int {
	var (
		dp        = make([][]int, 2)
		maxLength int
	)

	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	for _, row := range matrix {
		for j, n := range row {
			if n == '1' {
				dp[1][j+1] = min(dp[0][j], dp[0][j+1], dp[1][j]) + 1
				if dp[1][j+1] > maxLength {
					maxLength = dp[1][j+1]
				}
			}
		}
		dp[0] = dp[1]
		dp[1] = make([]int, len(matrix[0])+1)
	}
	return maxLength * maxLength
}
