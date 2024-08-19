package main

func reverse(x int) int {
	var (
		negative bool
		result   int
	)
	if x < 0 {
		negative = true
		x = -x
	}
	// Check the length
	for x > 0 {
		digit := x % 10
		result = result*10 + digit
		x /= 10
		// boundary check
		if result>>31 > 0 {
			return 0
		}
	}
	if negative {
		return -result
	}
	return result
}
