package main

func longestPalindrome(s string) string {
	var (
		maxLen    = 1
		maxString = s[:1]
	)
	for i := 0; i < len(s); i++ {
		var curLen = 1
		// s[i] 为中点
		for j := 1; ; j++ {
			if i-j < 0 || i+j >= len(s) {
				break
			}
			if s[i-j] != s[i+j] {
				break
			}
			curLen += 2
		}
		if curLen > maxLen {
			maxLen = curLen
			l := (maxLen - 1) / 2
			maxString = s[i-l : i+l+1]
		}
		// s[i+0.5] 为中点
		curLen = 0
		for j := 1; ; j++ {
			indexA, indexB := i+1-j, i+j
			if indexA < 0 || indexB >= len(s) {
				break
			}
			if s[indexA] != s[indexB] {
				break
			}
			curLen += 2
		}
		if curLen > maxLen {
			maxLen = curLen
			l := maxLen / 2
			maxString = s[i+1-l : i+l+1]
		}
	}
	return maxString
}
