package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	var dp = make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[1][j] = dp[0][j-1] + 1
			} else {
				dp[1][j] = max(dp[1][j-1], dp[0][j])
			}
		}
		copy(dp[0], dp[1])
	}
	return dp[1][len(text2)]
}
