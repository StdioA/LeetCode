package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://zh.wikipedia.org/wiki/最大子数列问题
// Kadane 算法
func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	maxEndingHere := nums[0]
	for i, n := range nums[1:] {
		maxEndingHere = max(n, maxEndingHere+n)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{100, 1, -3, 4, -1, 2, 1, -5, 100}))
	fmt.Println(maxSubArray([]int{-2, -1, -1}))
}
