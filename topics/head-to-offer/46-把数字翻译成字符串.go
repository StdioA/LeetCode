package main

import (
	"strconv"
)

func translateNum(num int) int {
	var (
		numString = strconv.Itoa(num)
		dp        = make([]int, len(numString)+1)
		charMap   = make(map[string]struct{})
	)
	for i := 0; i < 26; i++ {
		s := strconv.Itoa(i)
		charMap[s] = struct{}{}
	}
	dp[0] = 1

	for i := 0; i < len(numString); i++ {
		for j := 1; j <= 2; j++ {
			if i+1-j < 0 {
				continue
			}
			if _, ok := charMap[numString[i+1-j:i+1]]; ok {
				dp[i+1] += dp[i+1-j]
			}
		}
	}
	return dp[len(numString)]
}
