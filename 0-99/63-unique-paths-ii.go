package main

// 二维 DP
func uniquePathsWithObstaclesWithMoreSpace(obstacleGrid [][]int) int {
	var (
		m     = len(obstacleGrid)
		n     = len(obstacleGrid[0])
		paths = make([][]int, m)
	)
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
	}
	paths[0][0] = 1
	for i, row := range obstacleGrid {
		for j, obstacle := range row {
			if obstacle == 1 {
				continue
			}
			if i > 0 {
				paths[i][j] += paths[i-1][j]
			}
			if j > 0 {
				paths[i][j] += paths[i][j-1]
			}
		}
	}
	return paths[m-1][n-1]
}

// 状态压缩，将 O(mn) 的空间复杂度降为 O(n)，同时可以减少一条特判
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var (
		n     = len(obstacleGrid[0])
		paths = make([][]int, 2)
	)
	for i := 0; i < 2; i++ {
		paths[i] = make([]int, n)
	}
	paths[1][0] = 1
	for i, row := range obstacleGrid {
		for j, obstacle := range row {
			if obstacle == 1 {
				paths[1][j] = 0
				continue
			}
			sum := 0
			if i > 0 {
				sum += paths[0][j]
			}
			if j > 0 {
				sum += paths[1][j-1]
			}
			if !(i == 0 && j == 0) {
				paths[1][j] = sum
			}
		}
		paths[0] = paths[1]
	}
	return paths[0][n-1]
}
