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

// 2021-08-23
func maxSubArray2(nums []int) int {
	var (
		sum    = nums[0]
		maxSum = sum
	)
	for _, num := range nums[1:] {
		if sum < 0 {
			sum = num
		} else {
			sum += num
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{100, 1, -3, 4, -1, 2, 1, -5, 100}))
	fmt.Println(maxSubArray([]int{-2, -1, -1}))
}
