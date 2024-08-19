package main

func lengthOfLongestSubstring(s string) int {
	// 滑动窗口，用一个缓存来记录每个字符出现的次数
	if len(s) == 0 {
		return 0
	}
	var (
		result  = 1
		p1, p2  = 0, 0
		freqMap = make(map[byte]int)
	)
	freqMap[s[0]]++
	for p2 < len(s) {
		// 向后挪 p2 之前，先判断一下当前字符串长度
		for p2 < len(s) {
			length := p2 - p1 + 1
			if length > result {
				result = length
			}

			p2++
			if p2 >= len(s) {
				break
			}
			// 判断当前字符是否重复，重复了就不能继续移动了
			freqMap[s[p2]]++
			if freqMap[s[p2]] > 1 {
				break
			}
		}

		// 向后挪 p1 直到 s[p2] 不再重复，或 p1 和 p2 重合
		for p1 < p2 && p2 < len(s) && freqMap[s[p2]] > 1 {
			freqMap[s[p1]]--
			p1++
		}
	}
	return result
}
