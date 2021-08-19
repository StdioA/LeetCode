package main

import "fmt"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	var dp = make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < len(s); i++ {
		index := i + 1
		cur, prev := s[i], s[i-1]
		if cur != '0' {
			dp[index] = dp[index-1]
			if prev == '1' || prev == '2' && cur <= '6' {
				dp[index] += dp[index-2]
			}
		} else {
			// 10 & 20 可以，其余的非法
			if !(prev == '1' || prev == '2') {
				return 0
			}
			dp[index] = dp[index-2]
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(numDecodings("2101"))
	fmt.Println(numDecodings("10"))
	fmt.Println(numDecodings("1123"))
}
