func maxValue(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	for i, row := range grid {
		for j := range row {
			var m int
			if i > 0 && grid[i-1][j] > m {
				m = grid[i-1][j]
			}
			if j > 0 && grid[i][j-1] > m {
				m = grid[i][j-1]
			}
			grid[i][j] += m
		}
	}
	return grid[m-1][n-1]
}
