package main

func longestPrefix(s string) string {
	// 前缀函数的裸题
	// 见 KMP https://oi-wiki.org/string/kmp/
	var prefix = make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := prefix[i-1]
		for j > 0 && s[i] != s[j] {
			j = prefix[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		prefix[i] = j
	}
	return s[:prefix[len(s)-1]]
}
