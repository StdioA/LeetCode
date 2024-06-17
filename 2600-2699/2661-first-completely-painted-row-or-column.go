package main

func firstCompleteIndex(arr []int, mat [][]int) int {
	type coordinate struct {
		x, y int
	}
	var (
		row, col       = len(mat), len(mat[0])
		rowSum, colSum = make([]int, len(mat)), make([]int, len(mat[0]))
		coordinateMap  = make(map[int]coordinate)
	)
	for i, row := range mat {
		for j, n := range row {
			coordinateMap[n] = coordinate{x: i, y: j}
		}
	}
	for index, num := range arr {
		coordinate := coordinateMap[num]
		rowSum[coordinate.x]++
		colSum[coordinate.y]++
		if rowSum[coordinate.x] == col || colSum[coordinate.y] == row {
			return index
		}
	}
	return -1
}
