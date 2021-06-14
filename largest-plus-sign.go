package main

func min(nums ...int) int {
	var minVal = nums[0]
	for _, val := range nums[1:] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func genMatrix(n int) [][]int {
	var matrix = make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	var (
		grid  = genMatrix(n)
		left  = genMatrix(n)
		right = genMatrix(n)
		up    = genMatrix(n)
		down  = genMatrix(n)
	)
	for _, mine := range mines {
		grid[mine[0]][mine[1]] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				if j > 0 {
					left[i][j] = left[i][j-1] + 1
				} else {
					left[i][j] = 1
				}
			}
			if grid[j][i] == 0 {
				if j > 0 {
					down[j][i] = down[j-1][i] + 1
				} else {
					down[j][i] = 1
				}
			}
			if grid[i][n-1-j] == 0 {
				if j > 0 {
					right[i][n-1-j] = right[i][n-1-j+1] + 1
				} else {
					right[i][n-1-j] = 1
				}
			}
			if grid[n-1-j][i] == 0 {
				if j > 0 {
					up[n-1-j][i] = up[n-1-j+1][i] + 1
				} else {
					up[n-1-j][i] = 1
				}
			}
		}
	}
	var maxSize int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			size := min(up[i][j], down[i][j], left[i][j], right[i][j])
			if size > maxSize {
				maxSize = size
			}
		}
	}
	return maxSize
}
