package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	numbers := make([]int, 0)
	for _, row := range nums {
		for _, num := range row {
			numbers = append(numbers, num)
		}
	}
	result := make([][]int, r)
	index := 0
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
		for j := 0; j < c; j++ {
			result[i][j] = numbers[index]
			index++
		}
	}
	return result
}

func main() {
	matrix := [][]int{{1, 2}, {3, 4}}

	fmt.Println(matrixReshape(matrix, 1, 4))
}
