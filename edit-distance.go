package main

import "fmt"

func min(i1, i2, i3 int) int {
	m := i1
	if i2 < m {
		m = i2
	}
	if i3 < m {
		m = i3
	}
	return m
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len(word2); i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			d1, d2 := dp[i-1][j]+1, dp[i][j-1]+1
			d3 := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				d3++
			}
			dp[i][j] = min(d1, d2, d3)
		}
	}
	return dp[len(word1)][len(word2)]
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}
