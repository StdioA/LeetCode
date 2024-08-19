package main

func rob(nums []int) int {
	// 线性 dp
	// 状态转移公式 dp[i] = max(dp[i-2], dp[i-3]) + n
	// 不需要再往前看，因为 dp[i-2] 一定大于 dp[i-4]
	// 最后两个结果都有可能是最优结果，所以要做个 max
	var dp = make([]int, len(nums)+3)
	for i, num := range nums {
		index := i + 3
		dp[index] = max(dp[index-2], dp[index-3]) + num
	}
	return max(dp[len(nums)+1], dp[len(nums)+2])
}
