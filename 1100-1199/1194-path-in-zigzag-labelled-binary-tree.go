package main

func pathInZigZagTree(label int) []int {
	var path []int
	for label > 0 {
		path = append(path, label)
		label >>= 1
	}
	// Reverse array
	l := len(path)
	for i := 0; i < l/2; i++ {
		path[i], path[l-1-i] = path[l-1-i], path[i]
	}

	// Reverse nodes for every two layer
	for i := l % 2; i < l; i += 2 {
		path[i] = 3*(1<<i) - path[i] - 1
	}
	return path
}
