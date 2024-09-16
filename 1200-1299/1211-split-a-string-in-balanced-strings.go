package main

func balancedStringSplit(s string) int {
	var (
		count      [2]int
		splitCount int
	)
	// 如果 s[:i] 是平衡字符串，那么 s[i:] 一定也是
	// 因此直接从前往后遍历，出现相同出现次数则分割 +1 即可
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			count[0]++
		} else {
			count[1]++
		}
		if count[0] == count[1] {
			splitCount++
		}
	}
	return splitCount
}
