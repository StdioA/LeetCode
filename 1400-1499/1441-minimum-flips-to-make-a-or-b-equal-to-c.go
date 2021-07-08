package main

func minFlips(a int, b int, c int) int {
	var count int

	for c > 0 {
		aBit, bBit, cBit := a&1, b&1, c&1
		if aBit|bBit != cBit {
			if cBit == 0 {
				count += aBit + bBit
			} else {
				count++
			}
		}
		a >>= 1
		b >>= 1
		c >>= 1
	}
	// Deal with trailng bits for a & b
	for a > 0 {
		count += a & 1
		a >>= 1
	}
	for b > 0 {
		count += b & 1
		b >>= 1
	}

	return count
}
