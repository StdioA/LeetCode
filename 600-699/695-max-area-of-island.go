package main

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func searchIsland(grid [][]int, i, j int) int {
	if grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 2
	area := 1
	for _, d := range directions {
		ii, jj := i+d[0], j+d[1]
		if ii >= 0 && ii < len(grid) && jj >= 0 && jj < len(grid[0]) && grid[ii][jj] == 1 {
			area += searchIsland(grid, ii, jj)
		}
	}
	return area
}

func maxAreaOfIsland(grid [][]int) int {
	var max int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 1 {
				continue
			}
			area := searchIsland(grid, i, j)
			if area > max {
				max = area
			}
		}
	}
	return max
}
