package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	// a[i][j] 代表 s3[:i+j] 是否由 s1[:i]+s2[:j] 构成
	var dp = make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	// 初始值："" + "" == ""
	dp[0][0] = true
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			// 状态转移公式：
			// 如果 s3[i+j] == s1[i]，则可以将上一列迁移过来
			// 如果 s2[i+j] == s1[i]，则可以将上一行迁移过来
			// 注意坐标偏移
			if i > 0 && (s1[i-1] == s3[i+j-1]) {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if j > 0 && (s2[j-1] == s3[i+j-1]) {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
