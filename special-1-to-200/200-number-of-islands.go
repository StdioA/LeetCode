package main

func floodfillI(grid [][]byte, i, j int) {
	var directions = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	grid[i][j] = '0'
	for _, direction := range directions {
		ni, nj := i+direction[0], j+direction[1]
		if 0 <= ni && ni < len(grid) && 0 <= nj && nj < len(grid[0]) {
			if grid[ni][nj] == '1' {
				floodfillI(grid, ni, nj)
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	// floodfill 裸题
	var result int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				result++
				floodfillI(grid, i, j)
			}
		}
	}
	return result
}
