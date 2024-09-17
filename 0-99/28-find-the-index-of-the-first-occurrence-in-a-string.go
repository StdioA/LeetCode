package main

func strStr(haystack string, needle string) int {
	// KMP 裸题
	var prefix = make([]int, len(needle))
	for i := 1; i < len(needle); i++ {
		j := prefix[i-1]
		for j > 0 && needle[i] != needle[j] {
			j = prefix[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		prefix[i] = j
	}

	var previous int
	for i := 0; i < len(haystack); i++ {
		j := previous
		for j > 0 && needle[j] != haystack[i] {
			j = prefix[j-1]
		}
		if needle[j] == haystack[i] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
		previous = j
	}
	return -1
}
