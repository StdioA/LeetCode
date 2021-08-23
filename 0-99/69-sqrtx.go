package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var (
		lower = 1
		upper = x/2 + 1
		num   = (lower + upper) / 2
	)
	for {
		product := num * num
		nextProduct := (num + 1) * (num + 1)
		switch {
		case product == x:
			return num
		case nextProduct == x:
			return num + 1
		case product > x:
			upper = num
			num = (num + lower) / 2
		case product < x:
			if nextProduct > x {
				return num
			}
			lower = num + 1
			num = (num + upper) / 2
		}
	}
	return 0
}
