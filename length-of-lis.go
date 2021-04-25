package main

import "fmt"

func lengthOfLIS(nums []int) int {
	var (
		l   = len(nums)
		dp  = make([]int, l)
		max int
	)
	for i := range nums {
		dp[i] = 1
		for j := i; j >= 0; j-- {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
