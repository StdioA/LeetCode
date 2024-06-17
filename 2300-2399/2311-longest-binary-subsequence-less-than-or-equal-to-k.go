package main

func longestSubsequence(s string, k int) int {
	var (
		num    int64
		maxlen int
	)
	// Traverse from lowest bit to highest
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		// If there's an zero, then take it
		if c == '0' {
			maxlen++
		} else if maxlen <= 32 && num+int64(1<<maxlen) <= int64(k) {
			// It an 1 is met, judge if it is overceeded
			// Note that if we already have more than one 0,
			// there's no need to let the subsequence starts with our current 1.
			// e.g. 1[0001]0 is never efficiency then 1[000]1[0]
			num += 1 << maxlen
			maxlen++
		}
	}
	return maxlen
}
