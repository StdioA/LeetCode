package main

import "fmt"

func singleNumber(nums []int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}

	var (
		mask             int = 1
		result0, result1 int
	)
	for xor&1 == 0 {
		mask <<= 1
		xor >>= 1
	}
	for _, num := range nums {
		if num&mask > 0 {
			result0 ^= num
		} else {
			result1 ^= num
		}
	}
	return []int{result0, result1}
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}
