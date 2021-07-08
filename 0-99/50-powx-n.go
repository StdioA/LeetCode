package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	res := myPow(x, n/2)
	res *= res
	if n%2 == 1 {
		res *= x
	}
	if negative {
		res = 1 / res
	}
	return res
}

func main() {
	fmt.Println(myPow(2.0000, -2))
}
