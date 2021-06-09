package main

type NumMatrix struct {
	matrixSum [][]int // Prefix sum of square
}

func Constructor(matrix [][]int) NumMatrix {
	var (
		rows      = len(matrix)
		cols      = len(matrix[0])
		numMatrix = NumMatrix{
			matrixSum: make([][]int, rows),
		}
	)
	for i, row := range matrix {
		// Prefix sum of square
		matSum := make([]int, cols)
		numMatrix.matrixSum[i] = matSum
		for j, num := range row {
			matSum[j] = num
			if j > 0 {
				matSum[j] += matSum[j-1]
			}
			if i > 0 {
				matSum[j] += numMatrix.matrixSum[i-1][j]
			}
			if i > 0 && j > 0 {
				matSum[j] -= numMatrix.matrixSum[i-1][j-1]
			}
		}
	}
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum = this.matrixSum[row2][col2]
	if row1 > 0 {
		sum -= this.matrixSum[row1-1][col2]
	}
	if col1 > 0 {
		sum -= this.matrixSum[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		sum += this.matrixSum[row1-1][col1-1]
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
