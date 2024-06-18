package main

import "fmt"

func bits(n int) int {
	var bitsNum int
	for n > 0 {
		bitsNum += n & 1
		n >>= 1
	}
	return bitsNum
}

func makeTheIntegerZero(num1 int, num2 int) int {
	// num1 - x * num2 > 0 && bits(num1 - x * num2) < x
	var (
		num = num1
		n   = 0
	)
	for num > 0 {
		num -= num2
		n++
		// Edge case: the num must be devided by at least n `1`s, so there's num > n
		// Like bits(2) < 3, but the 1 cannot devided into three numbers, so there's no solve
		if num > n && bits(num) <= n {
			return n
		}
	}
	return -1
}

func main() {
	// fmt.Println(bits(9))
	fmt.Println(makeTheIntegerZero(85, -15))
}
