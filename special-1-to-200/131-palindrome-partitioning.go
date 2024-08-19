package main

func isPalindromeP(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{{}}
	}
	if len(s) == 1 {
		return [][]string{{s}}
	}

	// 从 0 开始往后尝试分割
	var results [][]string
	for i := 1; i <= len(s); i++ {
		left := s[:i]
		if !isPalindromeP(left) {
			continue
		}
		partialResults := partition(s[i:])
		for _, result := range partialResults {
			result = append(result, "")
			for x := len(result) - 1; x > 0; x-- {
				result[x] = result[x-1]
			}
			result[0] = left
			results = append(results, result)
		}
	}
	return results
}
