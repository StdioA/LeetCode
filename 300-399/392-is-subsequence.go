package main

func isSubsequence(s string, t string) bool {
	var (
		sIndex, tIndex int
	)
	for sIndex = 0; sIndex < len(s); sIndex++ {
		for ; tIndex < len(t) && s[sIndex] != t[tIndex]; tIndex++ {
		}
		if tIndex >= len(t) {
			return false
		}
		tIndex++
	}
	return true
}
