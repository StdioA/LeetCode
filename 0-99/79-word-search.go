package main

func exist(board [][]byte, word string) bool {
	var fill = make([][]bool, len(board))
	for i := range fill {
		fill[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			valid := dfs(board, fill, i, j, word)
			if valid {
				return valid
			}
		}
	}
	return false
}

var directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dfs(board [][]byte, fill [][]bool, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}
	if board[i][j] != word[0] {
		return false
	}
	// fill the current pos
	fill[i][j] = true
	for _, direction := range directions {
		x, y := i+direction[0], j+direction[1]
	}
	fill[i][j] = false
}
