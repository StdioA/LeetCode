package main

func spiralOrder(matrix [][]int) []int {
	// 模拟题
	var (
		x, y             int
		result           = make([]int, 0, len(matrix)*len(matrix[0]))
		directions       = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		currentDirection = 0
	)
	const visited = -10000
	for i := 0; i < cap(result); i++ {
		result = append(result, matrix[x][y])
		matrix[x][y] = visited
		nx, ny := x+directions[currentDirection][0], y+directions[currentDirection][1]
		if 0 <= nx && nx < len(matrix) && 0 <= ny && ny < len(matrix[0]) && matrix[nx][ny] != visited {
			x, y = nx, ny
		} else {
			currentDirection = (currentDirection + 1) % len(directions)
			x, y = x+directions[currentDirection][0], y+directions[currentDirection][1]
		}
	}
	return result
}
