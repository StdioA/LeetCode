package main

func matrixScore(grid [][]int) int {
	var (
		rows = len(grid)
		cols = len(grid[0])
	)
	for _, row := range grid {
		// Flip the row to make upper bit set to 1
		if row[0] == 0 {
			for i := range row {
				row[i] = 1 - row[i]
			}
		}
	}
	for j := 1; j < cols; j++ {
		sum := 0
		// Check the amount of ones
		for i := 0; i < rows; i++ {
			sum += grid[i][j]
		}
		if sum < rows-sum {
			// flip the column to gain more ones
			for i := 0; i < rows; i++ {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}
	// Calculate sum (actually it could be merged into last procedure)
	var sum = 0
	for _, row := range grid {
		n := 0
		for _, bit := range row {
			n = n*2 + bit
		}
		sum += n
	}
	return sum
}
