func hammingDistance(x int, y int) int {
	result := 0
	for x|y != 0 {
		result += (x % 2) ^ (y % 2)
		x /= 2
		y /= 2
	}
	return result
}
