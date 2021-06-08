package main

func rotateTheBox(box [][]byte) [][]byte {
	// let stones falling right
	for _, row := range box {
		stoneCounter := 0
		for i, obj := range row {
			switch obj {
			case '.': // empty
				continue
			case '#':
				// collect stones
				row[i] = '.'
				stoneCounter++
			case '*':
				// place stones before obstacle
				for stoneCounter > 0 {
					row[i-stoneCounter] = '#'
					stoneCounter--
				}
			}
		}
		// place the left stones
		for stoneCounter > 0 {
			row[len(row)-stoneCounter] = '#'
			stoneCounter--
		}
	}
	// rotate the box
	var (
		rotate = make([][]byte, len(box[0]))
		rows   = len(box)
		cols   = len(box[0])
	)
	for i := 0; i < cols; i++ {
		rotate[i] = make([]byte, rows)
		for j := 0; j < rows; j++ {
			rotate[i][j] = box[rows-1-j][i]
		}
	}
	return rotate
}
