package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int, length)
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	return min(dp[length-1], dp[length-2])
}

func main() {
	l := []int{0, 1, 1, 0}
	fmt.Println(minCostClimbingStairs(l))
}
