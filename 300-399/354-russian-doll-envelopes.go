package main

import "sort"

// Slow: O(n^2) time complixity
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		ei, ej := envelopes[i], envelopes[j]
		if ei[0] == ej[0] {
			return ei[1] < ej[1]
		}
		return ei[0] < ej[0]
	})

	var (
		dp     = make([]int, len(envelopes))
		result int
	)
	for i, env := range envelopes {
		dp[i] = 1
		for j, smallEnv := range envelopes[:i] {
			if env[0] > smallEnv[0] && env[1] > smallEnv[1] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}
