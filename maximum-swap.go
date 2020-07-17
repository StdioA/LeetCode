package main

import "fmt"

func maximumSwap(num int) int {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	length := len(digits)
	for i := length - 1; i > 0; i-- {
		max, maxi := digits[i-1], i-1
		for j := i - 1; j >= 0; j-- {
			if digits[j] >= max {
				max, maxi = digits[j], j
			}
		}
		if digits[i] < max {
			digits[i], digits[maxi] = digits[maxi], digits[i]
			break
		}
	}
	result := 0
	for i := length - 1; i >= 0; i-- {
		result = result*10 + digits[i]
	}
	return result
}

func main() {
	fmt.Println(maximumSwap(1993))
}
