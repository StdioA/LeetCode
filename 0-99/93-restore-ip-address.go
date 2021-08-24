package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var results []string

	var dfs func(string, []string)
	dfs = func(s string, segments []string) {
		if len(segments) > 4 {
			return
		} else if len(segments) == 4 {
			if len(s) == 0 {
				results = append(results, strings.Join(segments, "."))
			}
			return
		}
		var maxLength = len(s)
		if maxLength > 0 && s[0] == '0' {
			maxLength = 1
		}
		for i := 1; i <= maxLength; i++ {
			// 这里可以继续剪枝
			num, _ := strconv.Atoi(s[:i])
			if num <= 255 {
				dfs(s[i:], append(segments, s[:i]))
			}
		}
	}
	dfs(s, nil)
	return results
}
