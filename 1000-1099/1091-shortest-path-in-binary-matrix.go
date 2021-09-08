package main

import (
	"fmt"
)

type bfsElement struct {
	x, y, level int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	var (
		m, n       = len(grid), len(grid[0])
		directions = [][]int{
			{-1, -1}, {-1, 0}, {-1, 1},
			{0, -1}, {0, 1},
			{1, -1}, {1, 0}, {1, 1},
		}
		queue = []bfsElement{{0, 0, 1}}
	)

	for i := 0; i < len(queue); i++ {
		element := queue[i]
		x, y, level := element.x, element.y, element.level
		if x == m-1 && y == n-1 {
			return level
		}

		if grid[x][y] == 1 {
			continue
		}
		grid[x][y] = 1
		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 0 {
				queue = append(queue, bfsElement{nx, ny, level + 1})
			}
		}
	}
	return -1
}

// 优化：启发式的 A* 搜索：https://leetcode.com/problems/shortest-path-in-binary-matrix/discuss/313347/A*-search-in-Python

func main() {
	grid := [][]int{
		{0, 0, 0}, {1, 1, 0}, {1, 1, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(grid))
}
