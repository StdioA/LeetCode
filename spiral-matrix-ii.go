package main

import "fmt"

type Direction struct {
	x, y int
}

var direction = []Direction{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func generateMatrix(length int) [][]int {
	matrix := make([][]int, length)
	for i := 0; i < length; i++ {
		matrix[i] = make([]int, length)
	}

	x, y, cur := 0, 0, 0
	move := func() {
		d := direction[cur]
		nx, ny := x+d.x, y+d.y
		if nx < 0 || ny < 0 || nx >= length || ny >= length {
			cur = (cur + 1) % 4
		} else if matrix[nx][ny] > 0 {
			cur = (cur + 1) % 4
		}
		d = direction[cur]
		x, y = x+d.x, y+d.y
	}

	for i := 1; i <= length*length; i++ {
		matrix[x][y] = i
		move()
	}
	return matrix
}

func main() {
	for _, row := range generateMatrix(3) {
		fmt.Println(row)
	}
}
