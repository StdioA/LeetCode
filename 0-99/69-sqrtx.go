package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var (
		lower = 1
		upper = (x + 1) / 2
	)
	for lower < upper {
		mid := (lower + upper) / 2
		sqr, nsqr := mid*mid, (mid+1)*(mid+1)
		switch {
		case sqr > x:
			upper = mid
		case nsqr < x:
			lower = mid + 1
		case nsqr == x:
			return mid + 1
		default:
			return mid
		}
	}
	return lower
}
