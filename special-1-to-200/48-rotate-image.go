package main

func rotate(matrix [][]int) {
	var l = len(matrix)

	// 行列处理要单独判断
	// 否则当宽度为奇数时，中间那一条会被旋转两次
	// 建议用实际的例子看看: 3 -> 2 * 1, 4 -> 2 * 2, 5 -> 3 * 2
	for i := 0; i < (l+1)/2; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[j][l-1-i], matrix[l-1-i][l-1-j], matrix[l-1-j][i] = matrix[l-1-j][i], matrix[i][j], matrix[j][l-1-i], matrix[l-1-i][l-1-j]
		}
	}
}
