package main

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}

	var dp = make([][]int, len(t)+1)
	for i := 0; i <= len(t); i++ {
		dp[i] = make([]int, len(s)+1)
	}
	for i := 0; i < len(s); i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			ct, cs := t[i-1], s[j-1]
			if ct == cs {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
