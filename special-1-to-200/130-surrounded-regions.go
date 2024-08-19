package main

func floodfill(board [][]byte, i, j int) {
	var directions = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	board[i][j] = 'T'
	for _, direction := range directions {
		ni, nj := i+direction[0], j+direction[1]
		if 0 <= ni && ni < len(board) && 0 <= nj && nj < len(board[0]) {
			if board[ni][nj] == 'O' {
				floodfill(board, ni, nj)
			}
		}
	}
}

func solve(board [][]byte) {
	// 从边缘开始 floodfill
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			floodfill(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			floodfill(board, i, len(board[0])-1)
		}
	}
	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 'O' {
			floodfill(board, 0, j)
		}
		if board[len(board)-1][j] == 'O' {
			floodfill(board, len(board)-1, j)
		}
	}
	// 填充没有被 floodfill 到的部分
	// 把 T 还原成 O
	// 把 O 填充成 X
	for i, row := range board {
		for j, val := range row {
			if val == 'O' {
				board[i][j] = 'X'
			} else if val == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}
