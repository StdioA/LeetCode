package main

func hammingWeight(n int) int {
	var result int
	for n > 0 {
		result += n & 1
		n = n >> 1
	}
	return result
}
