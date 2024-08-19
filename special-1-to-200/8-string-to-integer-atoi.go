package main

func myAtoi(s string) int {
	// 题目描述的还是挺直白的
	const (
		negativeMax = 1 << 31
		positiveMax = (1 << 31) - 1
	)
	var (
		result     int
		index      int
		negative   bool
		upperLimit = positiveMax
	)

	// Step 1: 去除开头的空格
	for ; index < len(s); index++ {
		if s[index] != ' ' {
			break
		}
	}
	// Step 2: 判断符号
	if index >= len(s) {
		return result
	}
	if s[index] == '-' {
		negative = true
		upperLimit = negativeMax
		index++
	} else if s[index] == '+' {
		index++
	}
	// Step 3: 转换数字
	for ; index < len(s); index++ {
		char := s[index]
		if '0' <= char && char <= '9' {
			result = result*10 + int(char-'0')
		} else {
			break
		}
		// Step 4: 截断数字
		if result > upperLimit {
			result = upperLimit
			break
		}
	}
	if negative {
		return -result
	}
	return result
}
