package main

type retriever func([][]byte, int) []byte

func rowRetriever(board [][]byte, index int) []byte {
	return board[index]
}

func colRetriever(board [][]byte, index int) []byte {
	var column = make([]byte, 9)
	for i := 0; i < 9; i++ {
		column[i] = board[i][index]
	}
	return column
}

func gridRetriever(board [][]byte, index int) []byte {
	var (
		m, n = (index / 3) * 3, (index % 3) * 3
		grid = make([]byte, 0, 9)
	)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid = append(grid, board[m+i][n+j])
		}
	}
	return grid
}

func validate(bytes []byte) bool {
	var count = make([]int, 10)
	for _, b := range bytes {
		if b == '.' {
			continue
		}
		index := int(b - '0')
		count[index]++
		if count[index] > 1 {
			return false
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	var retrievers = []retriever{
		rowRetriever,
		colRetriever,
		gridRetriever,
	}
	for _, ret := range retrievers {
		for i := 0; i < 9; i++ {
			if valid := validate(ret(board, i)); !valid {
				return false
			}
		}
	}
	return true
}
