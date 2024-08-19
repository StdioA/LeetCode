package main

func isPalindrome(x int) bool {
	// 小数直接不满足条件
	if x < 0 {
		return false
	}
	// 将数位拆开，然后做回文判断
	var digits = make([]byte, 0, 10)
	for x > 0 {
		digits = append(digits, byte(x%10))
		x /= 10
	}
	length := len(digits) / 2
	for i := 0; i < length; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}
